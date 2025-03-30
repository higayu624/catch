package model

import (
	"context"
	"database/sql"

	"catch/model/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Category struct {
	Name string `boil:"name" json:"name"`
}

type Categorys []Category

func SeedCategories(ctx context.Context, tx *sql.Tx, restaurantCategories *[]null.String) ([]null.String, error) {
	// Insert customer table
	for _, value := range *restaurantCategories {
		category := models.Category{
			Name: value,
		}
		err := category.Insert(ctx, tx, boil.Infer())
		if err != nil {
			return nil, err
		}
	}

	return *restaurantCategories, nil
}
