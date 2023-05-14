package server

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"

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
        SELECT true
        FROM car_brands
        JOIN car_series ON car_brands.id = car_series.brand_id
        WHERE car_brands.id = $1 AND car_series.brand_id = $2;
        `, series_id, brand_id)
}

func getCarFromBd(db *sql.DB, id int) (*car.Car, error) {
	// :TODO
	return nil, nil
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
        select id, name, Description
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
        select id, name, Description
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

func dbIdExists(db *sql.DB, table, id any) (bool, error) {
	query := `
        SELECT true FROM %s WHERE id = $1
        `
	query = fmt.Sprintf(query, table)
	return dbExists(db, query, id)
}

func convertGrpcToJsonError(e any) error {
	if e == nil {
		return nil
	}

	data, err := json.Marshal(e)
	if err != nil {
		return status.Error(codes.Internal, fmt.Sprintf("car service err: %v", err))
	}
	return status.Error(codes.InvalidArgument, string(data))
}
