package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"catch/model"

	"github.com/gin-gonic/gin"
)

type postCustomerResponse struct {
	Name              *string      `json:"name"`
	Email             *string      `json:"email"`
	Gender            *string      `json:"gender"`
	Age               *int64       `json:"age"`
	GoogleAccessToken *string      `json:"google_access_token"`
	Store             *model.Store `json:"store,omitempty"`
}

func makeResult(responseModel interface{}, result interface{}) error {
	responseModelJson, err := json.Marshal(responseModel)
	if err != nil {
		return err
	}
	err = json.Unmarshal(responseModelJson, result)
	if err != nil {
		return err
	}
	return err
}

func response(message string, result any, status int) map[string]any {
	response := make(map[string]any)
	response["message"] = message
	response["result"] = result
	response["status"] = status

	return response
}

func responseError(message string, err error, status int) map[string]any {
	response := make(map[string]any)
	response["message"] = message
	response["result"] = err.Error()
	response["status"] = status

	return response
}

func handlerError(ctx *gin.Context, message string, err error, httpStatus int) {
	log.Print("err ", err)
	ctx.JSON(
		http.StatusInternalServerError,
		responseError(message, err, httpStatus),
	)
	ctx.Abort()
}
