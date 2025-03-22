package model

import (
	"catch/model/models"
	"context"
	"database/sql"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model

	Name              null.String `json:"name"`
	Email             null.String `json:"email"`
	Gender            null.String `json:"gender"`
	Age               null.Int64  `json:"age"`
	GoogleAccessToken null.String `json:"google_access_token"`
	Store             Store
}

func CreateCustomerStoreCategorys(ctx context.Context, tx *sql.Tx, request *Customer) (*Customer, error) {
	// Insert customer table
	customer := models.Customer{
		Name:              request.Name,
		Email:             request.Email,
		Gender:            request.Gender,
		Age:               request.Age,
		GoogleAccessToken: request.GoogleAccessToken,
	}
	err := customer.Insert(ctx, tx, boil.Infer())
	if err != nil {
		return nil, err
	}

	// Insert store table
	store := models.Store{
		CustomerID:  customer.ID,
		Name:        request.Store.Name,
		Description: request.Store.Description,
		Address:     request.Store.Address,
		Latitude:    request.Store.Latitude,
		Longitude:   request.Store.Longitude,
	}
	err = store.Insert(ctx, tx, boil.Infer())
	if err != nil {
		return nil, err
	}

	// Insert categorys table
	for _, categorization := range request.Store.Categorizations {
		categorizations := models.Categorization{
			StoreID:    null.Int64From(store.ID),
			CategoryID: null.Int64From(int64(categorization.CategoryId)),
		}
		err = categorizations.Insert(ctx, tx, boil.Infer())
	}

	return request, nil
}

// func PhysicalDeleteCustomer(ctx context.Context, tx *sql.Tx, request *models.Customer) (*models.Customer, error) {

// }

// func LogicalDeleteCustomer(ctx context.Context, tx *sql.Tx, request *models.Customer) (*models.Customer, error) {

// }
