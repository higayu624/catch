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
)

type Customer struct {
	Name              string `boil:"name" json:"name"`
	Email             string `boil:"email" json:"email"`
	Gender            string `boil:"gender" json:"gender"`
	Age               int64  `boil:"age" json:"age"`
	GoogleAccessToken string `boil:"google_access_token" json:"google_access_token"`
	Store             *Store `boil:"store" json:"store,omitempty"`
}

func ReadCustomerStoreCategorization(ctx context.Context, db *sql.DB, request Customer) (*models.Customer, error) {
	customer, err := models.Customers(
		Active,
		qm.Where("email = ?", request.Email),
		qm.Load(models.CustomerRels.Stores, Active),
		qm.Load(
			qm.Rels(
				models.CustomerRels.Stores,
				models.StoreRels.Categorizations,
				models.CategorizationRels.Category,
			), Active),
	).One(ctx, db)
	log.Print("customer", customer)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func CreateCustomerStoreCategorization(ctx context.Context, tx *sql.Tx, request *CreateCustomer) (*CreateCustomer, error) {
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
	for _, requestCategory := range request.Store.Categories {
		var categorization models.Categorization
		category, err := models.Categories(
			Active,
			qm.Where("name = ?", requestCategory.Name),
		).One(ctx, tx)
		if err != nil {
			return nil, err
		}
		categorization.CategoryID = null.Int64From(category.ID)
		categorization.StoreID = null.Int64From(store.ID)
		err = store.AddCategorizations(ctx, tx, false, &categorization)
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
		qm.Load(qm.Rels(models.CustomerRels.Stores, models.StoreRels.Categorizations), Active),
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
	categorizations := store.R.Categorizations
	if categorizations != nil {
		for _, categorization := range categorizations {
			err = store.RemoveCategorizations(ctx, tx, categorization)
			if err != nil {
				return err
			}
		}
	} else {
		makeResult(customer, response)
		return nil
	}
	log.Print("categories", categorizations)

	makeResult(customer, response)

	return err
}
