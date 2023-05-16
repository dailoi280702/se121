package server

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/dailoi280702/se121/car-service/pkg/car"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *carSerivceServer) GetBrand(ctx context.Context, req *car.GetBrandReq) (*car.Brand, error) {
	id := int(req.GetId())
	brand, err := dbGetBrandById(s.db, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while get car data from db: %v", err)
	}
	if brand == nil {
		return nil, status.Errorf(codes.NotFound, "car %d not exists", id)
	}
	return brand, nil
}

func (s *carSerivceServer) CreateBrand(context.Context, *car.CreateBrandReq) (*car.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBrand not implemented")
}

func (s *carSerivceServer) UpdateBrand(context.Context, *car.UpdateBrandReq) (*car.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBrand not implemented")
}

func (s *carSerivceServer) SearchForBrand(ctx context.Context, req *car.SearchReq) (*car.SearchForBrandRes, error) {
	res := car.SearchForBrandRes{}
	errCh := make(chan error)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		query := generateSearchForBrandsQuery(req)
		brands, err := fetchBrands(s.db, query)
		errCh <- err
		res.Brands = brands
		defer wg.Done()
	}()

	go func() {
		total, err := dbCountRecords(s.db, "car_brands")
		errCh <- err
		res.Total = int32(total)
		defer wg.Done()
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

func generateSearchForBrandsQuery(req *car.SearchReq) string {
	query := `
    SELECT id, name, country_of_origin, founded_year ,website_url, logo_url
    FROM car_brands
    WHERE 1=1`

	// Add search conditions if a query is provided
	if req.GetQuery() != "" {
		query += fmt.Sprintf(` 
            AND (name ILIKE '%%%s%%'
            OR country_of_origin ILIKE '%%%s%%')`,
			req.GetQuery(), req.GetQuery())
	}

	// Add ordering if orderby field is provided
	if req.GetOrderby() != "" {
		orderBy := "car_brands.created_at"
		switch req.GetOrderby() {
		case "date":
			orderBy = "car_brands.created_at"
		case "country":
			orderBy = "car_brands.country_of_origin"
		case "year":
			orderBy = "car_brands.founded_year"
		case "name":
			orderBy = "car_brands.name"
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

func fetchBrands(db *sql.DB, query string) ([]*car.Brand, error) {
	brands := []*car.Brand{}
	rows, err := db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return brands, nil
		}
		return nil, fmt.Errorf("failed to get records: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var brand car.Brand
		err := rows.Scan(&brand.Id, &brand.Name, &brand.CountryOfOrigin, &brand.FoundedYear, &brand.LogoUrl, &brand.WebsiteUrl)
		if err != nil {
			return nil, fmt.Errorf("failed to get record: %v", err)
		}
		brands = append(brands, &brand)
	}
	return brands, nil
}
