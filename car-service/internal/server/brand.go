package server

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

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

func (s *carSerivceServer) CreateBrand(ctx context.Context, req *car.CreateBrandReq) (*car.Empty, error) {
	// Verify inputs
	if err := validateBrand(&req.Name, req.CountryOfOrigin, req.WebsiteUrl, req.LogoUrl, req.FoundedYear); err != nil {
		return nil, err
	}

	// Insert brand into database
	if _, err := s.db.Exec(`
        INSERT INTO car_brands (name, country_of_origin, founded_year, website_url, logo_url)
        VALUES ($1, $2, $3, $4, $5)
        `, req.Name, req.CountryOfOrigin, req.FoundedYear, req.WebsiteUrl, req.LogoUrl); err != nil {
		return nil, serverError(err)
	}

	return &car.Empty{}, nil
}

func (s *carSerivceServer) UpdateBrand(ctx context.Context, req *car.UpdateBrandReq) (*car.Empty, error) {
	// Verify brand existence
	if err := checkForBrandExistence(s.db, req.GetId()); err != nil {
		return nil, err
	}

	// Verify inputs
	if err := validateBrand(req.Name, req.CountryOfOrigin, req.WebsiteUrl, req.LogoUrl, req.FoundedYear); err != nil {
		return nil, err
	}

	// Prepare update data
	updateData := map[string]interface{}{"updated_at": time.Now()}
	if req.Name != nil {
		updateData["name"] = *req.Name
	}
	if req.CountryOfOrigin != nil {
		updateData["country_of_origin"] = *req.CountryOfOrigin
	}
	if req.FoundedYear != nil {
		updateData["founded_year"] = *req.FoundedYear
	}
	if req.WebsiteUrl != nil {
		updateData["website_url"] = *req.WebsiteUrl
	}
	if req.LogoUrl != nil {
		updateData["logo_url"] = *req.LogoUrl
	}

	// Update brand record
	if err := dbUpdateRecord(s.db, "car_brands", updateData, int(req.GetId())); err != nil {
		return nil, status.Errorf(codes.Internal, "car service error: %v", err)
	}

	return &car.Empty{}, nil
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

// validate brand's inputs before modifying database
// return an json encoded errorResponse as error if inputs are not in correct formats
func validateBrand(name, countryOfOrigin, webSiteUrl, logoUrl *string, foundedYear *int32) error {
	validationErrors := map[string]string{}

	if name != nil {
		if strings.TrimSpace(*name) == "" {
			validationErrors["name"] = "Name can not be empty"
		}
	}
	if foundedYear != nil {
		if *foundedYear < 0 || *foundedYear > int32(time.Now().Year()) {
			validationErrors["foundedYear"] = "Founded year is out of range"
		}
	}
	if countryOfOrigin != nil {
		if strings.TrimSpace(*countryOfOrigin) == "" {
			validationErrors["countryOfOrigin"] = "Country of origin can not be empty"
		}
	}
	if webSiteUrl != nil {
		if strings.TrimSpace(*webSiteUrl) == "" || !regexp.MustCompile(httpRegex).MatchString(*webSiteUrl) {
			validationErrors["webSiteUrl"] = "Websiate URL is not valid"
		}
	}
	if logoUrl != nil {
		if strings.TrimSpace(*logoUrl) == "" || !regexp.MustCompile(httpRegex).MatchString(*logoUrl) {
			validationErrors["logoUrl"] = "Logo URL is not valid"
		}
	}

	if len(validationErrors) > 0 {
		return convertGrpcToJsonError(codes.InvalidArgument, errorResponse{Details: validationErrors})
	}
	return nil
}

// check for existence of a brand in car_brands table using its id
// return an error if brand does not exist or got an internal error
func checkForBrandExistence(db *sql.DB, id int32) error {
	exists, err := dbIdExists(db, "car_brands", id)
	if err != nil {
		return serverError(err)
	}
	if !exists {
		return convertGrpcToJsonError(codes.NotFound, fmt.Sprintf("Brand id %v not exists", id))
	}
	return nil
}

// genreate sql query for searching brands from grpc request as string
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

func fetchBrandsByIDs(db *sql.DB, ids ...int) ([]*car.Brand, error) {
	// Guard condition
	if len(ids) == 0 {
		return nil, nil
	}

	// Convert list of ids to list of string
	idStrings := make([]string, len(ids))
	for i, num := range ids {
		idStrings[i] = strconv.Itoa(num)
	}

	// Prepare the SQL query to retrieve records based on the list of IDs
	query := fmt.Sprintf(`
    SELECT id, name, country_of_origin, founded_year ,website_url, logo_url
    FROM car_brands
    WHERE id IN (%s)`, strings.Join(idStrings, ", "))

	// Fetch brands
	return fetchBrands(db, query)
}
