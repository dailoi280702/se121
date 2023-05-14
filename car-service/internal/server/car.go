package server

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/dailoi280702/se121/car-service/pkg/car"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const httpRegex = `/^https?:\/\/(?:www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b(?:[-a-zA-Z0-9()@:%_\+.~#?&\/=]*)$/`

func (s *carSerivceServer) GetCar(ctx context.Context, req *car.GetCarReq) (*car.Car, error) {
	id := int(req.GetId())
	exists, err := dbIdExists(s.db, "car_models", id)
	if err != nil {
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error while checking for car existence: %v", err)
		}
	}
	if !exists {
		return nil, convertGrpcToJsonError(codes.NotFound, fmt.Sprintf("Car %d not exists", id))
	}

	car, err := dbGetCarById(s.db, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while get car data from db: %v", err)
	}
	if car == nil {
		return nil, status.Errorf(codes.NotFound, "car %d not exists", id)
	}
	return car, nil
}

func (s *carSerivceServer) CreateCar(ctx context.Context, req *car.CreateCarReq) (*car.Empty, error) {
	validationErrors := make(map[string]string)

	// Validate inputs
	if strings.TrimSpace(req.GetName()) == "" {
		validationErrors["name"] = "Name cannot be empty"
	}

	if req.Year != nil {
		if req.GetYear() < 0 || req.GetYear() > int32(time.Now().Year()) {
			validationErrors["year"] = "Year is out of range"
		}
	}
	if req.HorsePower != nil {
		if req.GetHorsePower() <= 0 {
			validationErrors["horsepower"] = "Horsepower is out of range"
		}
	}
	if req.Torque != nil {
		if req.GetTorque() <= 0 {
			validationErrors["torque"] = "Torque is out of range"
		}
	}
	if req.ImageUrl != nil {
		if !regexp.MustCompile(httpRegex).MatchString(req.GetImageUrl()) {
			validationErrors["imageUrl"] = "Image URL is not valid"
		}
	}

	if len(validationErrors) > 0 {
		return nil, convertGrpcToJsonError(codes.InvalidArgument, errorResponse{
			Messages: []string{"Validation error"},
			Details:  validationErrors,
		})
	}

	errCh := make(chan error, 1)
	go func() {
		// Verify brand and series existence
		if req.BrandId != nil {
			id := req.GetBrandId()
			exists, err := dbIdExists(s.db, "car_brands", id)
			errCh <- err
			if !exists {
				validationErrors["brandId"] = fmt.Sprintf("Brand %d does not exist", id)
			}
		}
		if req.SeriesId != nil {
			id := req.GetSeriesId()
			exists, err := dbIdExists(s.db, "car_series", id)
			errCh <- err
			if !exists {
				validationErrors["seriesId"] = fmt.Sprintf("Series %d does not exist", id)
			} else {
				_, ok := validationErrors["brandId"]
				if !ok {
					if req.BrandId != nil {
						brandId := req.GetBrandId()
						match, err := dbSeriesBrandMatches(s.db, id, brandId)
						errCh <- err
						if !match {
							validationErrors["seriesId"] = fmt.Sprintf("Series %d does not exist in brand %d", id, brandId)
						}
					} else {
						brandId, err := dbGetBrandIdBySeriesId(s.db, int(id))
						errCh <- err
						brandId32 := int32(brandId)
						req.BrandId = &brandId32
					}
				}
			}

			if req.FuelTypeId != nil {
				id := req.GetFuelTypeId()
				exists, err := dbIdExists(s.db, "fuel_types", id)
				errCh <- err
				if !exists {
					validationErrors["fuel"] = fmt.Sprintf("Fuel type %d does not exist", id)
				}
			}
			if req.TransmissionId != nil {
				id := req.GetTransmissionId()
				exists, err := dbIdExists(s.db, "car_transmissions", id)
				errCh <- err
				if !exists {
					validationErrors["transmission"] = fmt.Sprintf("Transmission %d does not exist", id)
				}
			}
		}
		close(errCh)
	}()

	for range errCh {
		err := <-errCh
		if err != nil {
			return nil, status.Errorf(codes.Internal, "car service error: %v", err)
		}
	}

	if len(validationErrors) > 0 {
		return nil, convertGrpcToJsonError(codes.NotFound, errorResponse{
			Messages: []string{"Validation error"},
			Details:  validationErrors,
		})
	}

	// Insert car into the database
	_, err := s.db.Exec(`
    insert into car_models (brand_id, series_id, name, year, horsepower, torque, transmission, fuel_type, review, image_url)
    values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10 )
    `, req.BrandId, req.SeriesId, req.Name, req.Year, req.HorsePower, req.Torque, req.TransmissionId, req.FuelTypeId, req.Review, req.ImageUrl)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "car server got error while inserting car %v to db: %v", req, err)
	}

	return &car.Empty{}, nil
}

func (s *carSerivceServer) UpdateCar(context.Context, *car.UpdateCarReq) (*car.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCar not implemented")
}

func (s *carSerivceServer) DeleteCar(context context.Context, req *car.DeleteCarReq) (*car.Empty, error) {
	id := req.GetId()

	exists, err := dbIdExists(s.db, "car_models", id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while deleting car %v to db: %v", req, err)
	}
	if !exists {
		return nil, convertGrpcToJsonError(codes.NotFound, fmt.Sprintf("Car id %v not exists", id))
	}

	err = dbDeleteRecordById(s.db, "car_models", id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while deleting car %v to db: %v", req, err)
	}

	return &car.Empty{}, nil
}

func (s *carSerivceServer) SearchForCar(context.Context, *car.SearchForCarReq) (*car.SearchForCarRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchForCar not implemented")
}

func (s *carSerivceServer) GetCarMetadata(context.Context, *car.Empty) (*car.GetCarMetadataRes, error) {
	brands, err := getAllBrandFromDb(s.db)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "car service err: %v", err)
	}
	series, err := getAllSeriesFromDb(s.db)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "car service err: %v", err)
	}
	fuelTypes, err := getAllFuelTypesFromDb(s.db)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "car service err: %v", err)
	}
	transmissions, err := getAllTransmissionFromDb(s.db)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "car service err: %v", err)
	}
	res := car.GetCarMetadataRes{
		Brands:       brands,
		Series:       series,
		FuelType:     fuelTypes,
		Transmission: transmissions,
	}
	return &res, nil
}
