package model

import (
	"catch/model/models"
	"context"
	"database/sql"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Category struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Categorization Categorization
}

type Categorys []Category

func SeedCategorys(ctx context.Context, tx *sql.Tx, restaurantCategories *[]null.String) (*[]null.String, error) {
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

	return restaurantCategories, nil
}
