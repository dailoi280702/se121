package server

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/dailoi280702/se121/car-service/pkg/car"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type carSerivceServer struct {
	db *sql.DB
	car.UnimplementedCarServiceServer
}

func NewServer(db *sql.DB) *carSerivceServer {
	return &carSerivceServer{
		db: db,
	}
}

type errorResponse struct {
	Messages []string          `json:"messages,omitempty"`
	Details  map[string]string `json:"details,omitempty"`
}

func dbGetBrandIdBySeriesId(db *sql.DB, id int) (int, error) {
	brand_id := 0
	if err := db.QueryRow(`
        SELECT brand_id
        FROM car_series
        WHERE id = $1;
        `, id).Scan(&brand_id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return 0, nil
		default:
			return brand_id, err
		}
	}
	return brand_id, nil
}

func dbSeriesBrandMatches(db *sql.DB, series_id, brand_id any) (bool, error) {
	return dbExists(db, `
        SELECT EXISTS(SELECT 1 FROM car_series WHERE id = $1 AND brand_id = $2)
        `, series_id, brand_id)
}

// Fetch car details and related entities from the database
func dbGetCarById(db *sql.DB, id int) (*car.Car, error) {
	var car car.Car
	var brandId *int
	var seriesId *int
	var transmissionId *int
	var fuelTypeId *int

	// Fetch car details from the database
	err := dbScanRecordById(db, "car_models", id,
		"id", &car.Id,
		"brand_id", &brandId,
		"series_id", &seriesId,
		"name", &car.Name,
		"year", &car.Year,
		"horsepower", &car.HorsePower,
		"torque", &car.Torque,
		"transmission", &transmissionId,
		"fuel_type", &fuelTypeId,
		"review", &car.Review,
		"image_url", &car.ImageUrl,
	)
	if err != nil {
		return nil, err
	}

	// Fetch related entities concurrently
	var wg sync.WaitGroup
	errCh := make(chan error, 3)

	// Fetch brand details concurrently
	if brandId != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			car.Brand, err = dbGetBrandById(db, *brandId)
			errCh <- err
		}()
	}

	// Fetch series details concurrently
	if seriesId != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			car.Series, err = dbGetSeriesById(db, *seriesId)
			errCh <- err
		}()
	}

	// Fetch transmission details concurrently
	if transmissionId != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			car.Transmission, err = dbGetTransmissionById(db, *transmissionId)
			errCh <- err
		}()
	}

	// Fetch fuel type details concurrently
	if fuelTypeId != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			car.FuelType, err = dbGetFuelTypeById(db, *fuelTypeId)
			errCh <- err
		}()
	}

	// Wait for all related entities to be fetched
	go func() {
		wg.Wait()
		close(errCh)
	}()

	// Check for errors in fetching related entities
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	return &car, nil
}

func dbGetBrandById(db *sql.DB, id int) (*car.Brand, error) {
	var brand car.Brand
	if err := dbScanRecordById(db, "car_brands", id,
		"id", &brand.Id,
		"name", &brand.Name,
		"founded_year", &brand.FoundedYear,
		"country_of_origin", &brand.CountryOfOrigin,
		"website_url", &brand.WebsiteUrl,
		"logo_url", &brand.LogoUrl,
	); err != nil {
		return nil, err
	}
	return &brand, nil
}

func dbGetSeriesById(db *sql.DB, id int) (*car.Series, error) {
	var series car.Series
	if err := dbScanRecordById(db, "car_series", id,
		"id", &series.Id,
		"name", &series.Name,
		"brand_id", &series.BrandId,
	); err != nil {
		return nil, err
	}
	return &series, nil
}

func dbGetFuelTypeById(db *sql.DB, id int) (*car.FuelType, error) {
	var fuelType car.FuelType
	if err := dbScanRecordById(db, "fuel_types", id,
		"id", &fuelType.Id,
		"name", &fuelType.Name,
		"description", &fuelType.Description,
	); err != nil {
		return nil, err
	}
	return &fuelType, nil
}

func dbGetTransmissionById(db *sql.DB, id int) (*car.Transmission, error) {
	var transmission car.Transmission
	if err := dbScanRecordById(db, "car_transmissions", id,
		"id", &transmission.Id,
		"name", &transmission.Name,
		"description", &transmission.Description,
	); err != nil {
		return nil, err
	}
	return &transmission, nil
}

func dbDeleteRecordById(db *sql.DB, tableName string, id any) error {
	query := "DELETE FROM " + tableName + " WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func getAllBrandFromDb(db *sql.DB) ([]*car.Brand, error) {
	brands := []*car.Brand{}
	rows, err := db.Query(`
        select id, name, country_of_origin, founded_year, website_url, logo_url
        from car_brands
        `)
	if err != nil {
		if err == sql.ErrNoRows {
			return brands, nil
		}
		return nil, err
	}

	for rows.Next() {
		var brand car.Brand
		err = rows.Scan(&brand.Id, &brand.Name, &brand.CountryOfOrigin, &brand.FoundedYear, &brand.WebsiteUrl, &brand.LogoUrl)
		if err != nil {
			return nil, err
		}
		brands = append(brands, &brand)
	}

	return brands, nil
}

func getAllSeriesFromDb(db *sql.DB) ([]*car.Series, error) {
	series := []*car.Series{}
	rows, err := db.Query(`
        select id, brand_id, name
        from car_series
        `)
	if err != nil {
		if err == sql.ErrNoRows {
			return series, nil
		}
		return nil, err
	}

	for rows.Next() {
		var s car.Series
		err = rows.Scan(&s.Id, &s.BrandId, &s.Name)
		if err != nil {
			return nil, err
		}
		series = append(series, &s)
	}

	return series, nil
}

func getAllFuelTypesFromDb(db *sql.DB) ([]*car.FuelType, error) {
	fuelTypes := []*car.FuelType{}
	rows, err := db.Query(`
        select id, name, description
        from fuel_types
        `)
	if err != nil {
		if err == sql.ErrNoRows {
			return fuelTypes, nil
		}
		return nil, err
	}

	for rows.Next() {
		var f car.FuelType
		err = rows.Scan(&f.Id, &f.Name, &f.Description)
		if err != nil {
			return nil, err
		}
		fuelTypes = append(fuelTypes, &f)
	}

	return fuelTypes, nil
}

func getAllTransmissionFromDb(db *sql.DB) ([]*car.Transmission, error) {
	transmissions := []*car.Transmission{}
	rows, err := db.Query(`
        select id, name, description
        from car_transmissions
        `)
	if err != nil {
		if err == sql.ErrNoRows {
			return transmissions, nil
		}
		return nil, err
	}

	for rows.Next() {
		var s car.Transmission
		err = rows.Scan(&s.Id, &s.Name, &s.Description)
		if err != nil {
			return nil, err
		}
		transmissions = append(transmissions, &s)
	}

	return transmissions, nil
}

func countNumberOFRows(db *sql.DB, query string, args ...any) (int, error) {
	c := 0
	err := db.QueryRow(query, args...).Scan(&c)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return c, nil
}

func dbExists(db *sql.DB, query string, args ...any) (bool, error) {
	exists := false
	if err := db.QueryRow(query, args...).Scan(&exists); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return false, nil
		default:
			log.Println(args...)
			return false, err
		}
	}
	return exists, nil
}

// arguments is slice of key-value, for example: k1 v1 k2 v2 k3 v3 k4 v4

func dbScanRecordById(db *sql.DB, table string, id any, args ...any) error {
	if len(args)%2 != 0 {
		return errors.New("number of arguments must be even")
	}

	var keys []string
	var vals []any
	for i := 0; i < len(args); i += 1 {
		if i%2 == 0 {
			keys = append(keys, args[i].(string))
			continue
		}
		vals = append(vals, args[i])
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = $1", strings.Join(keys, ", "), table)
	if err := db.QueryRow(query, id).Scan(vals...); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil
		default:
			return err
		}
	}
	return nil
}

func dbIdExists(db *sql.DB, table, id any) (bool, error) {
	query := `
        SELECT exists( SELECT 1 FROM %s WHERE id = $1)
        `
	query = fmt.Sprintf(query, table)
	return dbExists(db, query, id)
}

func convertGrpcToJsonError(c codes.Code, e any) error {
	if e == nil {
		return nil
	}

	data, err := json.Marshal(e)
	if err != nil {
		return status.Error(codes.Internal, fmt.Sprintf("car service err: %v", err))
	}
	return status.Error(c, string(data))
}
