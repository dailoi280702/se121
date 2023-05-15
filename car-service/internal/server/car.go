package server

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strings"
	"sync"
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
	// Validate and verify ihputs
	err := validateCar(s.db, &req.Name, req.ImageUrl, req.Year, req.HorsePower, req.Torque, req.BrandId, req.SeriesId, req.FuelTypeId, req.TransmissionId)
	if err != nil {
		return nil, err
	}

	// Insert car into the database
	_, err = s.db.Exec(`
    insert into car_models (brand_id, series_id, name, year, horsepower, torque, transmission, fuel_type, review, image_url)
    values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10 )
    `, req.BrandId, req.SeriesId, req.Name, req.Year, req.HorsePower, req.Torque, req.TransmissionId, req.FuelTypeId, req.Review, req.ImageUrl)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "car server got error while inserting car %v to db: %v", req, err)
	}

	return &car.Empty{}, nil
}

func (s *carSerivceServer) UpdateCar(ctx context.Context, req *car.UpdateCarReq) (*car.Empty, error) {
	// Check for car existence
	if err := checkForCarExistence(s.db, int(req.GetId())); err != nil {
		return nil, err
	}

	// Validate car fields
	if err := validateCar(s.db, req.Name, req.ImageUrl, req.Year, req.HorsePower, req.Torque, req.BrandId, req.SeriesId, req.FuelTypeId, req.TransmissionId); err != nil {
		return nil, err
	}

	// Prepare update data
	updateData := map[string]interface{}{}
	if req.BrandId != nil {
		updateData["brand_id"] = *req.BrandId
	}
	if req.SeriesId != nil {
		updateData["series_id"] = *req.SeriesId
	}
	if req.Name != nil {
		updateData["name"] = *req.Name
	}
	if req.Year != nil {
		updateData["year"] = *req.Year
	}
	if req.HorsePower != nil {
		updateData["horsepower"] = *req.HorsePower
	}
	if req.Torque != nil {
		updateData["torque"] = *req.Torque
	}
	if req.TransmissionId != nil {
		updateData["transmission"] = *req.TransmissionId
	}
	if req.FuelTypeId != nil {
		updateData["fuel_type"] = *req.FuelTypeId
	}
	if req.Review != nil {
		updateData["review"] = *req.Review
	}
	if req.ImageUrl != nil {
		updateData["image_url"] = *req.ImageUrl
	}
	updateData["updated_at"] = time.Now()

	// Update car record
	if err := dbUpdateRecord(s.db, "car_models", updateData, int(req.GetId())); err != nil {
		return nil, status.Errorf(codes.Internal, "car service error: %v", err)
	}

	return &car.Empty{}, nil
}

func (s *carSerivceServer) DeleteCar(context context.Context, req *car.DeleteCarReq) (*car.Empty, error) {
	// check for car existence
	if err := checkForCarExistence(s.db, int(req.GetId())); err != nil {
		return nil, err
	}

	// validate and verify ihputs
	err := dbDeleteRecordById(s.db, "car_models", req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "car service error: %v", err)
	}

	return &car.Empty{}, nil
}

func (s *carSerivceServer) SearchForCar(context.Context, *car.SearchReq) (*car.SearchForCarRes, error) {
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

func checkForCarExistence(db *sql.DB, id int) error {
	exists, err := dbIdExists(db, "car_models", id)
	if err != nil {
		return status.Errorf(codes.Internal, "car service error: %v", err)
	}
	if !exists {
		return convertGrpcToJsonError(codes.NotFound, fmt.Sprintf("Car id %v not exists", id))
	}
	return nil
}

func validateCar(db *sql.DB, name, imageUrl *string, year, horsePower, torque, brandId, seriesId, fuelTypeId, transmissionId *int32) error {
	validationErrors := make(map[string]string)

	// Validate inputs
	if name != nil {
		if strings.TrimSpace(*name) == "" {
			validationErrors["name"] = "Name cannot be empty"
		}
	}

	if year != nil {
		if *year < 0 || *year > int32(time.Now().Year()) {
			validationErrors["year"] = "Year is out of range"
		}
	}
	if horsePower != nil {
		if *horsePower <= 0 {
			validationErrors["horsepower"] = "Horsepower is out of range"
		}
	}
	if torque != nil {
		if *torque <= 0 {
			validationErrors["torque"] = "Torque is out of range"
		}
	}
	if imageUrl != nil {
		if !regexp.MustCompile(httpRegex).MatchString(*imageUrl) {
			validationErrors["imageUrl"] = "Image URL is not valid"
		}
	}

	if len(validationErrors) > 0 {
		return convertGrpcToJsonError(codes.InvalidArgument, errorResponse{
			Messages: []string{"Validation error"},
			Details:  validationErrors,
		})
	}

	errCh := make(chan error, 100)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer close(errCh)

		// Verify brand and series existence
		if brandId != nil {
			id := *brandId
			exists, err := dbIdExists(db, "car_brands", id)
			errCh <- err
			if !exists {
				validationErrors["brandId"] = fmt.Sprintf("Brand %d does not exist", id)
			}
		}
		if seriesId != nil {
			id := *seriesId
			exists, err := dbIdExists(db, "car_series", id)
			errCh <- err
			if !exists {
				validationErrors["seriesId"] = fmt.Sprintf("Series %d does not exist", id)
			} else {
				_, ok := validationErrors["brandId"]
				if !ok {
					if brandId != nil {
						brandId := *brandId
						match, err := dbSeriesBrandMatches(db, id, brandId)
						errCh <- err
						if !match {
							validationErrors["seriesId"] = fmt.Sprintf("Series %d does not exist in brand %d", id, brandId)
						}
					} else {
						brand, err := dbGetBrandIdBySeriesId(db, int(id))
						errCh <- err
						brandId32 := int32(brand)
						brandId = &brandId32
					}
				}
			}
		}

		if fuelTypeId != nil {
			id := *fuelTypeId
			exists, err := dbIdExists(db, "fuel_types", id)
			errCh <- err
			if !exists {
				validationErrors["fuel"] = fmt.Sprintf("Fuel type %d does not exist", id)
			}
		}
		if transmissionId != nil {
			id := *transmissionId
			exists, err := dbIdExists(db, "car_transmissions", id)
			errCh <- err
			if !exists {
				validationErrors["transmission"] = fmt.Sprintf("Transmission %d does not exist", id)
			}
		}
	}()

verificationLoop:
	for {
		select {
		case err, ok := <-errCh:
			if err != nil {
				return status.Errorf(codes.Internal, "car service error: %v", err)
			}
			if !ok {
				break verificationLoop
			}
		case <-time.After(5 * time.Second):
			// Timeout duration
			return status.Error(codes.Internal, "Verification timeout")
		}
	}

	if len(validationErrors) > 0 {
		return convertGrpcToJsonError(codes.NotFound, errorResponse{
			Messages: []string{"Validation error"},
			Details:  validationErrors,
		})
	}

	return nil
}
