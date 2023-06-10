package server

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/dailoi280702/se121/car-service/pkg/car"
	"github.com/dailoi280702/se121/pkg/go/grpc/generated/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// const httpRegex = `/^https?:\/\/(?:www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b(?:[-a-zA-Z0-9()@:%_\+.~#?&\/=]*)$/`
const httpRegex = `/(((ftp|http|https):\/\/)|(\/)|(..\/))(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?`

func (s *carSerivceServer) GetCar(ctx context.Context, req *car.GetCarReq) (*car.Car, error) {
	id := int(req.GetId())
	exists, err := dbIdExists(s.db, "car_models", id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while checking for car existence: %v", err)
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

func (s *carSerivceServer) CreateCar(ctx context.Context, req *car.CreateCarReq) (*car.CreateCarRes, error) {
	// Validate and verify inputs
	err := validateCar(s.db, &req.Name, req.ImageUrl, req.Year, req.HorsePower, req.Torque, req.BrandId, req.SeriesId, req.FuelTypeId, req.TransmissionId)
	if err != nil {
		return nil, err
	}

	// Insert car into the database
	var id int32
	err = s.db.QueryRow(`
    insert into car_models (brand_id, series_id, name, year, horsepower, torque, transmission, fuel_type, review, image_url)
    values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10 )
    RETURNING id
    `, req.BrandId, req.SeriesId, req.Name, req.Year, req.HorsePower, req.Torque, req.TransmissionId, req.FuelTypeId, req.Review, req.ImageUrl).Scan(&id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "car server got error while inserting car %v to db: %v", req, err)
	}

	return &car.CreateCarRes{Id: id}, nil
}

func (s *carSerivceServer) UpdateCar(ctx context.Context, req *car.UpdateCarReq) (*utils.Empty, error) {
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

	return &utils.Empty{}, nil
}

func (s *carSerivceServer) DeleteCar(context context.Context, req *car.DeleteCarReq) (*utils.Empty, error) {
	// check for car existence
	if err := checkForCarExistence(s.db, int(req.GetId())); err != nil {
		return nil, err
	}

	// validate and verify ihputs
	err := dbDeleteRecordById(s.db, "car_models", req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "car service error: %v", err)
	}

	return &utils.Empty{}, nil
}

func (s *carSerivceServer) SearchForCar(ctx context.Context, req *utils.SearchReq) (*car.SearchForCarRes, error) {
	query := generateSearchForCarQuery(req)
	idList := []int{}

	rows, err := s.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return &car.SearchForCarRes{Cars: []*car.Car{}, Total: 0}, nil
		}
		return nil, status.Errorf(codes.Internal, "car service error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		idList = append(idList, id)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "car service error: %v", err)
		}
	}

	errCh := make(chan error, 2)
	var wg sync.WaitGroup
	var res car.SearchForCarRes
	wg.Add(2)

	go func() {
		res.Cars, err = getCarsFromIds(s.db, int(math.Ceil(math.Sqrt(float64(len(idList))))), idList...)
		errCh <- err
		wg.Done()
	}()

	go func() {
		// total, err := dbCountRecords(s.db, "car_models")
		total, err := countCarsFromQuery(s.db, req)
		res.Total = int32(total)
		errCh <- err
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			return nil, serverError(err)
		}
	}

	return &res, nil
}

func (s *carSerivceServer) GetCarMetadata(context.Context, *utils.Empty) (*car.GetCarMetadataRes, error) {
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

func (s *carSerivceServer) GetCars(context context.Context, req *car.GetCarsReq) (*car.GetCarsRes, error) {
	query := generateGetCarIDs(req)
	idList := []int{}

	rows, err := s.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return &car.GetCarsRes{Cars: []*car.Car{}}, nil
		}
		return nil, status.Errorf(codes.Internal, "car service error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		idList = append(idList, id)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "car service error: %v", err)
		}
	}

	cars, err := getCarsFromIds(s.db, int(math.Ceil(math.Sqrt(float64(len(idList))))), idList...)
	if err != nil {
		return nil, serverError(err)
	}

	return &car.GetCarsRes{Cars: cars}, nil
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

// generateSearchForCarQuery generates a SQL query for searching cars based on the provided search request.
// The generated query performs a left join on various tables and applies search conditions, ordering, and pagination.
// The query is returned as a string.
func generateSearchForCarQuery(req *utils.SearchReq) string {
	query := `
    SELECT car_models.id 
    FROM  car_models
    LEFT JOIN car_brands on car_models.brand_id = car_brands.id
    LEFT JOIN car_series on car_models.series_id = car_series.id
    LEFT JOIN fuel_types on car_models.fuel_type = fuel_types.id
    LEFT JOIN car_transmissions on car_models.transmission = car_transmissions.id
    WHERE 1=1`

	// Add search conditions if a query is provided
	if req.GetQuery() != "" {
		query += fmt.Sprintf(` 
            AND (car_models.name ILIKE '%%%s%%'
            OR car_brands.name ILIKE '%%%s%%'
            OR car_series.name ILIKE '%%%s%%'
            OR fuel_types.name ILIKE '%%%s%%'
            OR car_transmissions.name ILIKE '%%%s%%')`,
			req.GetQuery(), req.GetQuery(), req.GetQuery(), req.GetQuery(), req.GetQuery())
	}

	// Add ordering if orderby field is provided
	if req.GetOrderby() != "" {
		orderBy := "car_models.create_at"
		switch req.GetOrderby() {
		case "date":
			orderBy = "car_models.create_at"
		case "torque":
			orderBy = "car_models.torque"
		case "horsePower":
			orderBy = "car_models.horsepower"
		case "year":
			orderBy = "car_models.year"
		case "name":
			orderBy = "car_models.name"
		}
		query += fmt.Sprintf(" ORDER BY %s", orderBy)
		if req.GetIsAscending() {
			query += " ASC"
		} else {
			query += " DESC"
		}
	}

	// Add limit if limit field is provided
	if req.GetLimit() > 0 {
		query += fmt.Sprintf(" LIMIT %d", req.GetLimit())
	}

	// Add pagination if startAt field is provided
	if req.GetStartAt() > 0 {
		query += fmt.Sprintf(" OFFSET %d", req.GetStartAt())
	}

	return query
}

// generate a SQL query for get cars from request
func generateGetCarIDs(req *car.GetCarsReq) string {
	query := `
    SELECT id 
    FROM car_models
    WHERE 1=1`

	updateData := map[string]any{
		"id":           req.Id,
		"name":         req.Name,
		"year":         req.Year,
		"torque":       req.Torque,
		"transmission": req.TransmissionID,
		"fueltype":     req.FuelTypeID,
		"brand_id":     req.BrandID,
		"series_id":    req.SeriesID,
	}

	for k, v := range updateData {
		var data any
		value := reflect.ValueOf(v)
		if !value.IsNil() {
			if value.Kind() == reflect.Pointer {
				data = reflect.ValueOf(v).Elem().Interface()
			} else {
				data = v
			}
			query += fmt.Sprintf(" AND %s = %v", k, data)
		}
	}

	return query
}

// getCarsFromIds retrieves cars from the database based on the provided IDs.
// It uses a worker pool approach to parallelize the database queries and
// returns the cars in the same order as the provided IDs.
func getCarsFromIds(db *sql.DB, numWorkers int, ids ...int) ([]*car.Car, error) {
	numTasks := len(ids)

	cars := make([]*car.Car, len(ids))
	carsCh := make(chan struct {
		idex int
		car  *car.Car
	}, len(ids))

	type task struct {
		idex  int
		carId int
	}
	taskCh := make(chan task, numTasks)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// worker performs the actual retrieval of cars from the database
	worker := func() {
		for task := range taskCh {
			c, err := dbGetCarById(db, task.carId)
			if err != nil {
				// Send a nil car to indicate an error occurred during retrieval
				carsCh <- struct {
					idex int
					car  *car.Car
				}{task.idex, nil}
				continue
			}

			// Send the retrieved car and its index to the cars channel
			carsCh <- struct {
				idex int
				car  *car.Car
			}{task.idex, c}
		}
		wg.Done()
	}

	// Spawn worker goroutines
	for i := 0; i < numWorkers; i++ {
		go worker()
	}

	// Launch tasks
	for i, id := range ids {
		taskCh <- task{i, id}
	}
	close(taskCh)

	go func() {
		wg.Wait()
		close(carsCh)
	}()

	// Collect the retrieved cars from the cars channel
	for result := range carsCh {
		cars[result.idex] = result.car
	}

	return cars, nil
}

func countCarsFromQuery(db *sql.DB, req *utils.SearchReq) (int, error) {
	query := `
    SELECT COUNT(*)
    FROM  car_models
    LEFT JOIN car_brands on car_models.brand_id = car_brands.id
    LEFT JOIN car_series on car_models.series_id = car_series.id
    LEFT JOIN fuel_types on car_models.fuel_type = fuel_types.id
    LEFT JOIN car_transmissions on car_models.transmission = car_transmissions.id
    WHERE 1=1`

	// Add search conditions if a query is provided
	if req.GetQuery() != "" {
		query += fmt.Sprintf(` 
            AND (car_models.name ILIKE '%%%s%%'
            OR car_brands.name ILIKE '%%%s%%'
            OR car_series.name ILIKE '%%%s%%'
            OR fuel_types.name ILIKE '%%%s%%'
            OR car_transmissions.name ILIKE '%%%s%%')`,
			req.GetQuery(), req.GetQuery(), req.GetQuery(), req.GetQuery(), req.GetQuery())
	}

	var count int
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count records: %v", err)
	}

	return count, nil
}
