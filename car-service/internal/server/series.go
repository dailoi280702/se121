package server

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/dailoi280702/se121/car-service/pkg/car"
	"github.com/dailoi280702/se121/pkg/go/grpc/generated/utils"
	"google.golang.org/grpc/codes"
)

func (s *carSerivceServer) GetSeries(ctx context.Context, req *car.GetSeriesReq) (*car.Series, error) {
	// Check for series existence
	id := req.GetId()
	if err := checkSeriesExistence(s.db, id); err != nil {
		return nil, err
	}

	// Fetch series from database
	series, err := dbGetSeriesById(s.db, int(id))
	if err != nil {
		return nil, serverError(err)
	}

	return series, nil
}

func (s *carSerivceServer) CreateSeries(ctx context.Context, req *car.CreateSeriesReq) (*car.CreateSeriesRes, error) {
	// Validate and verify inputs
	if err := validateSeries(s.db, &req.Name, &req.BrandId); err != nil {
		return nil, err
	}

	// Insert series into database
	var id int32
	err := s.db.QueryRow(`
        INSERT INTO car_series (name, brand_id)
        VALUES ($1, $2)
        `, req.Name, req.BrandId).Scan(id)
	if err != nil {
		return nil, serverError(fmt.Errorf("failed to insert series: %v", err))
	}

	return &car.CreateSeriesRes{Id: id}, nil
}

func (s *carSerivceServer) UpdateSeries(ctx context.Context, req *car.UpdateSeriesReq) (*utils.Empty, error) {
	// Check for series existence
	id := req.GetId()
	if err := checkSeriesExistence(s.db, id); err != nil {
		return nil, err
	}

	// Validate and verify inputs
	if err := validateSeries(s.db, req.Name, req.BrandId); err != nil {
		return nil, err
	}

	// Prepare update data
	updateData := map[string]interface{}{"updated_at": time.Now()}
	if req.Name != nil {
		updateData["name"] = *req.Name
	}
	if req.BrandId != nil {
		updateData["brand_id"] = *req.BrandId
	}

	// Update series record
	if err := dbUpdateRecord(s.db, "car_series", updateData, int(id)); err != nil {
		return nil, serverError(err)
	}

	return &utils.Empty{}, nil
}

func (s *carSerivceServer) SearchForSeries(ctx context.Context, req *utils.SearchReq) (*car.SearchForSeriesRes, error) {
	res := car.SearchForSeriesRes{}
	errCh := make(chan error)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		var err error
		res.Series, err = fetchSeriesDetails(s.db, req)
		errCh <- err
	}()

	go func() {
		defer wg.Done()
		var err error
		res.Total, err = fetchNumSeries(s.db, req)
		errCh <- err
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

	// verify name
	if name != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			nameExists, err := dbExists(db, `
            SELECT EXISTS(SELECT 1 FROM car_series WHERE name = $1)
            `, *name)
			errCh <- err
			if nameExists {
				validateErrors["name"] = fmt.Sprintf("Series %s already exists", *name)
			}
		}()
	}

	// verify brand
	if brandID != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			brandExists, err := dbIdExists(db, "car_brands", *brandID)
			errCh <- err
			if !brandExists {
				validateErrors["brandId"] = fmt.Sprintf("Brand %d does not exist", *brandID)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	// return interal error if exist
	for err := range errCh {
		if err != nil {
			return serverError(err)
		}
	}

	// return an error if verification failed
	if len(validateErrors) > 0 {
		return convertGrpcToJsonError(codes.InvalidArgument, errorResponse{Details: validateErrors})
	}

	return nil
}

// return SQL query string for get series from search request
func generateSearchSeriesQuery(req *utils.SearchReq) string {
	query := `
    SELECT car_series.id, car_series.name, car_series.brand_id
    FROM car_series
    LEFT JOIN car_brands on car_series.brand_id = car_brands.id
    WHERE 1=1`

	// Add search conditions if a query is provided
	if req.GetQuery() != "" {
		query += fmt.Sprintf(` 
            AND (car_series.name ILIKE '%%%s%%'
            OR car_brands.name ILIKE '%%%s%%')`,
			req.GetQuery(), req.GetQuery())
	}

	// Add ordering if orderby field is provided
	if req.GetOrderby() != "" {
		var orderBy string
		switch req.GetOrderby() {
		case "name":
			orderBy = "car_series.name"
		default:
			orderBy = "car_series.created_at"
		}
		query += fmt.Sprintf(" ORDER BY %s", orderBy)
		if req.GetIsAscending() {
			query += " ASC"
		} else {
			query += " DESC"
		}
	}

	// Add pagination if startAt field is provided
	if req.GetStartAt() > 0 {
		query += fmt.Sprintf(" OFFSET %d", req.GetStartAt())
	}

	// Add limit if limit field is provided
	if req.GetLimit() > 0 {
		query += fmt.Sprintf(" LIMIT %d", req.GetLimit())
	}

	return query
}

// fetch series from database by SQL query
func fetchSeriesDetails(db *sql.DB, req *utils.SearchReq) ([]*car.SeriesDetail, error) {
	query := generateSearchSeriesQuery(req)
	brands := map[int]*car.Brand{}
	brandIDs := []int{}
	seriesList := []*car.Series{}
	seriesDetails := []*car.SeriesDetail{}

	rows, err := db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return seriesDetails, nil
		}
		return nil, err
	}
	defer rows.Close()

	// fetch list of series
	for rows.Next() {
		var series car.Series

		err = rows.Scan(&series.Id, &series.Name, &series.BrandId)
		if err != nil {
			return nil, fmt.Errorf("failed to get record: %v", err)
		}

		seriesList = append(seriesList, &series)
		brandIDs = append(brandIDs, int(series.BrandId))
	}

	// fetch existed brands in list of series
	brandList, err := fetchBrandsByIDs(db, brandIDs...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch brands %v", err)
	}
	// convet list of brands of map of brand id and brand
	for _, brand := range brandList {
		if brand != nil {
			brands[int((*brand).GetId())] = brand
		}
	}

	// Combine brands and list of series
	for _, series := range seriesList {
		seriesDetails = append(seriesDetails, &car.SeriesDetail{
			Id:    series.Id,
			Name:  series.Name,
			Brand: brands[int(series.BrandId)],
		})
	}

	return seriesDetails, nil
}

// fetch numbser of series from database by SQL query
func fetchNumSeries(db *sql.DB, req *utils.SearchReq) (int32, error) {
	query := `
    SELECT COUNT(*)
    FROM car_series
    LEFT JOIN car_brands on car_series.brand_id = car_brands.id
    WHERE 1=1`

	// Add search conditions if a query is provided
	if req.GetQuery() != "" {
		query += fmt.Sprintf(` 
            AND (car_series.name ILIKE '%%%s%%'
            OR car_brands.name ILIKE '%%%s%%')`,
			req.GetQuery(), req.GetQuery())
	}

	var nums int
	if err := db.QueryRow(query).Scan(&nums); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return 0, nil
		default:
			return 0, err
		}
	}

	return int32(nums), nil
}
