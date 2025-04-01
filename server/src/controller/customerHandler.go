package controller

import (
	"database/sql"
	"log"
	"net/http"

	"catch/model"
	"catch/model/models"

	"github.com/gin-gonic/gin"
)

func GetCustomer(dbHandler *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request model.Customer
		var customer *models.Customer
		// Recieved request's body and validation
		err := ctx.BindJSON(&request)
		if err != nil {
			handlerError(ctx, "request body is bad", err, http.StatusBadRequest)
			return
		}
		// read customer
		customer, err = model.ReadCustomerStoreCategorization(ctx, dbHandler, request)
		if err != nil {
			handlerError(ctx, "failed Read customerStoreCategorization", err, http.StatusInternalServerError)
			return
		}
		// response
		var result model.Customer
		log.Print("custoer", customer)
		err = makeResult(customer, &result)
		_ = makeResult(customer.R.Stores, &result.Store)
		// var categories []model.Category
		// for _, categorization := range customer.R.Stores[0].R.Categorizations {
		// 	categories = append(categories, categorization.R.Category)
		// }
		log.Print("result", result)
		if err != nil {
			handlerError(ctx, "failed conversion struct", err, http.StatusInternalServerError)
			return
		}
		ctx.JSON(
			http.StatusOK,
			response("get customer store cateogrys", customer, http.StatusOK),
		)
	}
}

func PostCustomer(dbHandler *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request model.CreateCustomer
		var customer *model.CreateCustomer
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
		var result model.CreateCustomer
		err = makeResult(customer, &result)
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
