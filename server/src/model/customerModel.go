package model

import (
	"context"
	"database/sql"
	"log"
	"time"

	"catch/model/models"

	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/types"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model

	Name              string `boil:"name" json:"name"`
	Email             string `boil:"email" json:"email"`
	Gender            string `boil:"gender" json:"gender"`
	Age               int64  `boil:"age" json:"age"`
	GoogleAccessToken string `boil:"google_access_token" json:"google_access_token"`
	Store             *Store `boil:"store" json:"store,omitempty"`
}

func CreateCustomerStoreCategorization(ctx context.Context, tx *sql.Tx, request *Customer) (*Customer, error) {
	// Insert customer table
	customer := models.Customer{
		Name:              null.NewString(request.Name, true),
		Email:             null.NewString(request.Email, true),
		Gender:            null.NewString(request.Gender, true),
		Age:               null.NewInt64(request.Age, true),
		GoogleAccessToken: null.NewString(request.GoogleAccessToken, true),
	}
	err := customer.Insert(ctx, tx, boil.Infer())
	if err != nil {
		return nil, err
	}

	// Insert store table
	latitudeDecimal := new(decimal.Big)
	store := models.Store{
		CustomerID:  null.Int64From(customer.ID),
		Name:        null.NewString(request.Store.Name, true),
		Description: null.NewString(request.Store.Description, true),
		Address:     null.NewString(request.Store.Address, true),
		Latitude:    types.NewNullDecimal(latitudeDecimal.SetFloat64(request.Store.Latitude)),
		Longitude:   types.NewNullDecimal(latitudeDecimal.SetFloat64(request.Store.Longitude)),
	}
	err = store.Insert(ctx, tx, boil.Infer())
	if err != nil {
		return nil, err
	}

	// Insert categorization table
	for _, requestCategory := range *request.Store.Categories {
		category, err := models.Categories(
			Active,
			qm.Where("name = ?", requestCategory.Name),
		).One(ctx, tx)
		if err != nil {
			return nil, err
		}
		log.Print(category)
		err = store.AddCategories(ctx, tx, false, category)
		if err != nil {
			return nil, err
		}
	}

	return request, nil
}

func DeleteCustomerStoreCategorization(ctx context.Context, tx *sql.Tx, request *Customer, response *Customer) error {
	// delete customer table
	customer, err := models.Customers(
		Active,
		qm.Where("email = ?", request.Email),
		qm.Load(models.CustomerRels.Stores, Active),
	).One(ctx, tx)
	if err != nil {
		return err
	}
	now := time.Now()
	customer.DeletedAt = null.TimeFrom(now)
	_, err = customer.Update(ctx, tx, boil.Whitelist("deleted_at"))
	if err != nil {
		return err
	}

	// delete store table
	store := customer.R.Stores[0]
	if store != nil {
		store.DeletedAt = null.TimeFrom(now)
	} else {
		makeResult(customer, response)
		return nil
	}
	_, err = store.Update(ctx, tx, boil.Whitelist("deleted_at"))
	if err != nil {
		return err
	}

	// delete categorization table
	categories := store.R.Categories
	if categories != nil {
		for _, category := range store.R.Categories {
			err = store.RemoveCategories(ctx, tx, category)
			if err != nil {
				return err
			}
		}
	} else {
		makeResult(store, response.Store)
	}
	log.Print("categories", categories)

	makeResult(customer, response)

	return err
}
