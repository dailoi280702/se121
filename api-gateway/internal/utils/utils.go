package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

const JSON = "application/json"

var UnplementedError = errors.New("method unimplemented yet")

type MalformedRequest struct {
	Status int
	Msg    string
}

func (mr *MalformedRequest) Error() string {
	return mr.Msg
}

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	if contentType := r.Header.Get("Content-Type"); contentType != "" && contentType != JSON {
		msg := "Content-Type header is not application/json"
		return &MalformedRequest{Status: http.StatusUnsupportedMediaType, Msg: msg}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := "Request body contains badly-formed JSON"
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return &MalformedRequest{Status: http.StatusRequestEntityTooLarge, Msg: msg}

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		msg := "Request body must only contain a single JSON object"
		return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
	}

	return nil
}

func ConvertQueryParams(queryParams url.Values, data interface{}) error {
	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() != reflect.Ptr || dataValue.IsNil() {
		return fmt.Errorf("data must be a non-nil pointer")
	}

	dataValue = dataValue.Elem()
	dataType := dataValue.Type()

	var wg sync.WaitGroup
	errCh := make(chan error, dataType.NumField())

	for i := 0; i < dataType.NumField(); i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			field := dataType.Field(index)
			fieldValue := dataValue.Field(index)

			// Get JSON tag value of the field
			jsonTag := field.Tag.Get("json")

			// Use field name if JSON tag is empty
			if jsonTag == "" {
				jsonTag = field.Name
			}

			// Skip processing if the JSON tag is set to "-"
			if jsonTag == "-" {
				return
			}

			// Split JSON tag to separate options
			jsonParts := strings.Split(jsonTag, ",")

			// Get query parameter name from JSON tag (first part)
			paramName := jsonParts[0]

			// Skip processing if the parameter name is empty
			if paramName == "" {
				return
			}

			// Get query parameter value by parameter name
			paramValue := queryParams.Get(paramName)
			// Skip processing if the parameter value is empty
			if paramValue == "" {
				return
			}

			// Handle pointer types
			if fieldValue.Kind() == reflect.Ptr {
				// Create a new instance of the field type if it's nil
				if fieldValue.IsNil() {
					fieldValue.Set(reflect.New(fieldValue.Type().Elem()))
				}

				// Set the value of the pointer type
				fieldValue = fieldValue.Elem()
			}

			// Convert and set field value based on its type
			switch fieldValue.Kind() {
			case reflect.Int32:
				intValue, err := strconv.ParseInt(paramValue, 10, 32)
				if err != nil {
					// msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
					// return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
					errCh <- fmt.Errorf("failed to convert '%s' to int32: %w", paramValue, err)
					return
				}
				fieldValue.SetInt(intValue)
			case reflect.Bool:
				boolValue, err := strconv.ParseBool(paramValue)
				if err != nil {
					errCh <- fmt.Errorf("failed to convert '%s' to bool: %w", paramValue, err)
					return
				}
				fieldValue.SetBool(boolValue)
			case reflect.String:
				fieldValue.SetString(paramValue)
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}

func ValidateField(field string, value string, required bool, pattern *regexp.Regexp) error {
	isEmpty := strings.TrimSpace(value) == ""
	if required && isEmpty {
		return errors.New(field + " cannot be empty")
	}

	if pattern != nil && !isEmpty && !pattern.MatchString(value) {
		return errors.New(field + " is not valid")
	}

	return nil
}
