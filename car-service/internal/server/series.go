package server

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"sync"

	"github.com/dailoi280702/se121/car-service/pkg/car"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *carSerivceServer) GetSeries(ctx context.Context, req *car.GetSeriesReq) (*car.Series, error) {
	// Check for series existence
	id := req.GetId()
	if err := checkSeriesExistence(s.db, id); err != nil {
		return nil, err
	}

	// Fetch series from database

	return nil, status.Errorf(codes.Unimplemented, "method GetSeries not implemented")
}

func (s *carSerivceServer) CreateSeries(ctx context.Context, req *car.CreateSeriesReq) (*car.Empty, error) {
	// Validate and verify inputs
	if err := validateSeries(s.db, &req.Name, &req.BrandId); err != nil {
		return nil, err
	}

	// Insert series into database

	return nil, status.Errorf(codes.Unimplemented, "method CreateSeries not implemented")
}

func (s *carSerivceServer) UpdateSeries(ctx context.Context, req *car.UpdateSeriesReq) (*car.Empty, error) {
	// Check for series existence
	id := req.GetId()
	if err := checkSeriesExistence(s.db, id); err != nil {
		return nil, err
	}

	// Validate and verify inputs
	if err := validateSeries(s.db, &req.Name, &req.BrandId); err != nil {
		return nil, err
	}

	// Prepare update data

	// Update series record

	return nil, status.Errorf(codes.Unimplemented, "method UpdateSeries not implemented")
}

func (s *carSerivceServer) SearchForSeries(context.Context, *car.SearchReq) (*car.SearchForSeriesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchForCar not implemented")
}

// check for series existence in database by its id
// if series does not exist return error
func checkSeriesExistence(db *sql.DB, id int32) error {
	exists, err := dbIdExists(db, "car_series", id)
	if err != nil {
		return serverError(err)
	}
	if !exists {
		return convertGrpcToJsonError(codes.NotFound, fmt.Sprintf("Series id %v not exists", id))
	}
	return nil
}

// validate and verify series inputs before modify database
// return an json encoded errorResponse as error if inputs are not in correct formats
func validateSeries(db *sql.DB, name *string, brandID *int32) error {
	validateErrors := map[string]string{}

	// validate name and brandID
	if name != nil {
		if strings.TrimSpace(*name) == "" {
			validateErrors["name"] = "Name can not be empty"
		}
	}
	if brandID != nil {
		if *brandID == 0 {
			validateErrors["brandId"] = "Brand can not be empty"
		}
	}

	// return an error if validation failed
	if len(validateErrors) > 0 {
		return convertGrpcToJsonError(codes.InvalidArgument, errorResponse{Details: validateErrors})
	}

	errCh := make(chan error)
	var wg sync.WaitGroup
	go func() {
		wg.Done()
		close(errCh)
	}()

	// verify name
	go func() {
		nameExists, err := dbExists(db, `
            SELECT EXISTS(SELECT 1 FROM car_series WHERE name = $1
            `, *name)
		errCh <- err
		if nameExists {
			validateErrors["name"] = fmt.Sprintf("Series %s already exists", *name)
		}
	}()

	// verify brand
	go func() {
		brandExists, err := dbIdExists(db, "car_brands", *brandID)
		errCh <- err
		if !brandExists {
			validateErrors["brandId"] = fmt.Sprintf("Brand %d does not exist", *brandID)
		}
	}()

	// return an error if verification failed
	if len(validateErrors) > 0 {
		return convertGrpcToJsonError(codes.InvalidArgument, errorResponse{Details: validateErrors})
	}

	// return interal error if exist
	for err := range errCh {
		if err != nil {
			return serverError(err)
		}
	}

	return nil
}

// return SQL query string for series from search request
func generateSQLSearchQueryForSeries(req *car.SearchReq) string {
	return ""
}

// fetch series from database by SQL query
func fetchSeries(db *sql.DB, query string) ([]*car.Series, error) {
	return nil, nil
}
