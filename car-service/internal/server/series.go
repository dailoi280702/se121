package server

import (
	"context"
	"database/sql"
	"fmt"

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

func (s *carSerivceServer) CreateSeries(context.Context, *car.CreateSeriesReq) (*car.Empty, error) {
	// Validate and verify inputs

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

	// Prepare update data

	// Update series record

	return nil, status.Errorf(codes.Unimplemented, "method UpdateSeries not implemented")
}

func (s *carSerivceServer) SearchForSeries(context.Context, *car.SearchReq) (*car.SearchForSeriesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchForCar not implemented")
}

// check for series existence in database
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
