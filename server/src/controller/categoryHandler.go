package controller

import (
	"database/sql"
	"net/http"

	"catch/model"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
)

func SeedCategoryHandler(dbHandler *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		restaurantCategories := []null.String{
			null.NewString("和食", true),
			null.NewString("寿司", true),
			null.NewString("焼肉", true),
			null.NewString("ラーメン", true),
			null.NewString("居酒屋", true),
			null.NewString("フレンチ", true),
			null.NewString("イタリアン", true),
			null.NewString("カフェ", true),
			null.NewString("中華", true),
			null.NewString("ベーカリー", true),
			null.NewString("ファストフード", true),
			null.NewString("ベジタリアン", true),
			null.NewString("韓国料理", true),
			null.NewString("タイ料理", true),
			null.NewString("メキシカン", true),
		}

		// transaction
		tx, err := dbHandler.BeginTx(ctx, nil)
		if err != nil {
			handlerError(ctx, "failed starting transaction", err, http.StatusInternalServerError)
			return
		}
		defer func() {
			rolleback(tx)
		}()

		// create customer
		result, err := model.SeedCategories(ctx, tx, &restaurantCategories)
		if err != nil {
			panic(err)
		}

		ctx.JSON(
			http.StatusOK,
			response("seed category", result, http.StatusOK),
		)
	}
}
