package server

import (
	"database/sql"

	"github.com/dailoi280702/se121/car-service/pkg/car"
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

func getAllBrandFromDb(db *sql.DB) ([]*car.Brand, error) {
	brands := []*car.Brand{}
	rows, err := db.Query(`
        select id, name, country_of_origin, founded_year, website_ur, logo_url
        from brands
        `)
	if err != nil {
		if err == sql.ErrNoRows {
			return brands, nil
		}
		return nil, err
	}

	for rows.Next() {
		var brand *car.Brand
		err = rows.Scan(brand.Id, brand.Name, brand.CountryOfOrigin, brand.FoundedYear, brand.WebsiteUrl, brand.LogoUrl)
		if err != nil {
			return nil, err
		}
		brands = append(brands, brand)
	}

	return brands, nil
}

func getAllSeriesFromDb(db *sql.DB) ([]*car.Series, error) {
	series := []*car.Series{}
	rows, err := db.Query(`
        select id, brand_id, name
        from brands
        `)
	if err != nil {
		if err == sql.ErrNoRows {
			return series, nil
		}
		return nil, err
	}

	for rows.Next() {
		var s *car.Series
		err = rows.Scan()
		if err != nil {
			return nil, err
		}
		series = append(series, s)
	}

	return series, nil
}
