package controller

import (
	"catch/model"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
)

func SeedCategoryHandler(dbHandler *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var restaurantCategories = []null.String{
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
		var seedResponse *[]null.String

		// transaction
		tx, err := dbHandler.BeginTx(ctx, nil)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				response("failed starting transaction", err, http.StatusInternalServerError),
			)
		}
		defer func() {
			if p := recover(); p != nil { // panicに渡された値をrecover()でキャッチ
				ctx.JSON(
					http.StatusInternalServerError,
					response("panic!!", p, http.StatusInternalServerError),
				)
			} else { // transactionの実行
				err = tx.Commit()
				if err != nil {
					tx.Rollback()
					ctx.JSON(
						http.StatusInternalServerError,
						response("failed execution of transaction", err, http.StatusInternalServerError),
					)
				}
				log.Print("seed category Successed!", seedResponse)
			}
		}()

		// create customer
		seedResponse, err = model.SeedCategorys(ctx, tx, &restaurantCategories)
		if err != nil {
			panic(err)
		}
	}
}
