package controller

import (
	"catch/model"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostCustomer(dbHandler *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request model.Customer
		var responseBody *model.Customer

		// Recieved request's body and validation
		err := ctx.BindJSON(&request)
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				response("request body is bad", err, http.StatusBadRequest),
			)
			return
		}

		// transaction
		tx, err := dbHandler.BeginTx(ctx, nil)
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				response("failed starting transaction", err, http.StatusInternalServerError),
			)
			return
		}
		defer func() {
			if p := recover(); p != nil { // panicに渡された値をrecover()でキャッチ
				ctx.JSON(
					http.StatusInternalServerError,
					response("panic!!", p, http.StatusInternalServerError),
				)
				return
			} else { // transactionの実行
				err = tx.Commit()
				if err != nil {
					tx.Rollback()
					ctx.JSON(
						http.StatusInternalServerError,
						response("failed execution of transaction", err, http.StatusInternalServerError),
					)
					return
				}
			}
		}()

		// create customer
		responseBody, err = model.CreateCustomerStoreCategorys(ctx, tx, &request)
		if err != nil {
			tx.Rollback()
			ctx.JSON(
				http.StatusBadRequest,
				response("failed Insert customerStoreCategorys", err, http.StatusInternalServerError),
			)

			return
		}

		// response
		ctx.JSON(
			http.StatusOK,
			response("create customer store cateogrys", responseBody, http.StatusOK),
		)
		return
	}
}
