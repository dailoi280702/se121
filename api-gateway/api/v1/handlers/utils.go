package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/dailoi280702/se121/api-gateway/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SendJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Panic(err)
	}
}

func MustSendError(err error, w http.ResponseWriter) {
	log.Printf("\nBitch, you got an err: %s\n", err.Error())
	http.Error(w, fmt.Sprintf("server errror: %s", err.Error()), http.StatusInternalServerError)
}

func SendJsonFromGrpcError(w http.ResponseWriter, err error, actions map[codes.Code]func()) {
	code := status.Code(err)
	s, _ := status.FromError(err)

	action, ok := actions[code]

	switch {
	case ok:
		action()
	case code == codes.InvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	case code == codes.AlreadyExists:
		w.WriteHeader(http.StatusConflict)
	case code == codes.NotFound:
		w.WriteHeader(http.StatusNotFound)
	case code == codes.Unavailable:
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	default:
		MustSendError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(s.Message()))
	if err != nil {
		MustSendError(err, w)
	}
}

type convertOpt struct {
	before   *func()
	callback func() error
	after    *func()
	actions  map[codes.Code]func()
	reqData  interface{}
}

type convertOptFunc func(*convertOpt)

func defaultConvertOpt(callback func() error) *convertOpt {
	return &convertOpt{
		before:   nil,
		callback: callback,
		after:    nil,
		actions:  map[codes.Code]func(){},
		reqData:  nil,
	}
}

func convertWithPreFunc(fn func()) convertOptFunc {
	return func(opt *convertOpt) {
		opt.before = &fn
	}
}

func convertWithPostFunc(fn func()) convertOptFunc {
	return func(opt *convertOpt) {
		opt.after = &fn
	}
}

func convertWithCustomCodes(actions map[codes.Code]func()) convertOptFunc {
	return func(opt *convertOpt) {
		opt.actions = actions
	}
}

func convertWithJsonReqData(data interface{}) convertOptFunc {
	return func(opt *convertOpt) {
		opt.reqData = data
	}
}

func convertJsonApiToGrpc(w http.ResponseWriter, r *http.Request, callback func() error, opts ...convertOptFunc) {
	opt := defaultConvertOpt(callback)
	for _, fn := range opts {
		fn(opt)
	}

	w.Header().Set("Content-Type", "application/json")
	if opt.reqData != nil {
		err := utils.DecodeJSONBody(w, r, opt.reqData)
		if err != nil {
			var mr *utils.MalformedRequest
			if errors.As(err, &mr) {
				http.Error(w, mr.Msg, mr.Status)
			} else {
				MustSendError(err, w)
			}
			return
		}
	}

	if opt.before != nil {
		(*opt.before)()
	}

	err := callback()
	if err != nil {
		SendJsonFromGrpcError(w, err, opt.actions)
		return
	}

	if opt.after != nil {
		(*opt.after)()
	}
}
