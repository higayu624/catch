package controller

import (
	"database/sql"
	"net/http"

	"catch/model"

	"github.com/gin-gonic/gin"
)

func PostCustomer(dbHandler *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request model.Customer
		var customer *model.Customer
		// Recieved request's body and validation
		err := ctx.BindJSON(&request)
		if err != nil {
			handlerError(ctx, "request body is bad", err, http.StatusBadRequest)
			return
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
		customer, err = model.CreateCustomerStoreCategorization(ctx, tx, &request)
		if err != nil {
			tx.Rollback()
			handlerError(ctx, "failed Insert customerStoreCategorys", err, http.StatusInternalServerError)
			return
		}
		// response
		var result model.Customer
		err = makeResult(*customer, &result)
		if err != nil {
			tx.Rollback()
			handlerError(ctx, "failed replacement structure", err, http.StatusInternalServerError)
			return
		}
		ctx.JSON(
			http.StatusOK,
			response("create customer store cateogrys", result, http.StatusOK),
		)
	}
}

func DeleteCustomer(dbHandler *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request model.Customer
		var responseBody model.Customer
		// Recieved request's body and validation
		err := ctx.BindJSON(&request)
		if err != nil {
			handlerError(ctx, "request body is bad", err, http.StatusBadRequest)
			return
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
		// delete customer
		err = model.DeleteCustomerStoreCategorization(ctx, tx, &request, &responseBody)
		if err != nil {
			tx.Rollback()
			handlerError(ctx, "failed Delete customerStoreCategorys", err, http.StatusInternalServerError)
			return
		}
		// response
		ctx.JSON(
			http.StatusOK,
			response("delete customer store cateogrys", responseBody, http.StatusOK),
		)
	}
}
