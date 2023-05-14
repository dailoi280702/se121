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
	car, err := getCarFromBd(s.db, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while get car data from db: %v", err)
	}
	if car == nil {
		return nil, status.Errorf(codes.NotFound, "car %d not exists", id)
	}
	return nil, nil
}

func (s *carSerivceServer) CreateCar(ctx context.Context, req *car.CreateCarReq) (*car.Empty, error) {
	errorsRes := struct {
		Messages []string          `json:"messages"`
		Details  map[string]string `json:"details"`
	}{[]string{}, map[string]string{}}

	name := req.GetName()

	// validate
	if strings.TrimSpace(name) == "" {
		errorsRes.Details["name"] = "Name can not be empty"
	}
	if req.Year != nil {
		if req.GetYear() < 0 || req.GetYear() > int32(time.Now().Year()) {
			errorsRes.Details["year"] = "Year out of range"
		}
	}
	if req.HorsePower != nil {
		if req.GetHorsePower() < 0 {
			errorsRes.Details["horsepower"] = "Horse power out of range"
		}
	}
	if req.Torque != nil {
		if req.GetTorque() < 0 {
			errorsRes.Details["torque"] = "Torque out of range"
		}
	}
	if req.ImageUrl != nil {
		if !regexp.MustCompile(httpRegex).MatchString(req.GetImageUrl()) {
			errorsRes.Details["imageUrl"] = "Image url is not a valid url"
		}
	}

	if len(errorsRes.Details) != 0 {
		return nil, convertGrpcToJsonError(codes.InvalidArgument, errorsRes)
	}

	// verify
	if req.BrandId != nil {
		id := req.GetBrandId()
		exists, err := dbIdExists(s.db, "car_brands", id)
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("car service err: %v", err))
		}
		if !exists {
			errorsRes.Details["brandId"] = fmt.Sprintf("Brand %d not exists", id)
		}
	}
	if req.SeriesId != nil {
		id := req.GetSeriesId()
		exists, err := dbIdExists(s.db, "car_series", id)
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("car service err: %v", err))
		}
		if !exists {
			errorsRes.Details["seriesId"] = fmt.Sprintf("Series %d not exists", id)
		} else {
			_, ok := errorsRes.Details["brandId"]
			if !ok {
				if req.BrandId != nil {
					brandId := req.GetBrandId()
					match, err := dbSeriesBrandMatches(s.db, id, brandId)
					if err != nil {
						return nil, status.Error(codes.Internal, fmt.Sprintf("car service err: %v", err))
					}
					if !match {
						errorsRes.Details["seriesId"] = fmt.Sprintf("Series %d not exists in brand %d", id, brandId)
					}
				} else {
					brandId, err := dbGetBrandIdBySeriesId(s.db, int(id))
					if err != nil {
						return nil, status.Error(codes.Internal, fmt.Sprintf("car service err: %v", err))
					}
					brandId32 := int32(brandId)
					req.BrandId = &brandId32
				}
			}
		}
		if req.FuelTypeId != nil {
			id := req.GetFuelTypeId()
			exists, err := dbIdExists(s.db, "fuel_types", id)
			if err != nil {
				return nil, status.Error(codes.Internal, fmt.Sprintf("car service err: %v", err))
			}
			if !exists {
				errorsRes.Details["fuel"] = fmt.Sprintf("Fuel type %d not exists", id)
			}
		}
		if req.TransmissionId != nil {
			id := req.GetTransmissionId()
			exists, err := dbIdExists(s.db, "car_transmissions", id)
			if err != nil {
				return nil, status.Error(codes.Internal, fmt.Sprintf("car service err: %v", err))
			}
			if !exists {
				errorsRes.Details["transmission"] = fmt.Sprintf("Transmission %d not exists", id)
			}
		}
	}

	if len(errorsRes.Details) != 0 {
		return nil, convertGrpcToJsonError(codes.NotFound, errorsRes)
	}

	// insert to db
	_, err := s.db.Exec(`
        insert into car_models (brand_id, series_id, name, year, horsepower, torque, transmission, fuel_type, review, image_url)
        values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10 )
        `, req.BrandId, req.SeriesId, req.Name, req.Year, req.HorsePower, req.Torque, req.TransmissionId, req.FuelTypeId, req.Review, req.ImageUrl)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while inserting car %v to db: %v", req, err)
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
