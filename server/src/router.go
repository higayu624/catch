package main

import (
	"database/sql"
	"net/http"
	"time"

	"catch/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	productName = "catch"
	version     = "v1"
	APIRoot     = version + "/" + productName + "/"
)

func endpointGroups(path string) string {
	return APIRoot + path
}

func initRouter(dbHandler *sql.DB) *gin.Engine {
	router := gin.Default()
	router.ContextWithFallback = true

	// エンドポイントの存在しないアクセス処理
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "404 page not found",
		})
	})

	// ルーター
	router.Use(cors.New(cors.Config{
		AllowMethods: []string{
			"GET",
			"UPDATE",
			"PUT",
			"POST",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
			"Set-Cookie",
			"Cookies",
		},
		// 外部からのアクセス許可
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowCredentials: false,
		MaxAge:           24 * time.Hour,
	}))

	// v1/catch/customer
	customerGroup := router.Group(endpointGroups("customer"))
	{
		route := ""
		customerGroup.GET(route, controller.GetCustomer(dbHandler))
		customerGroup.POST(route, controller.PostCustomer(dbHandler))
		customerGroup.DELETE(route, controller.DeleteCustomer(dbHandler))
		// route = "/force"
		// customerGroup.DELETE(route, controller.DeleteCustomerForce(dbHandler))
	}

	// v1/catch/category
	categoryGroup := router.Group(endpointGroups("category"))
	{
		route := ""
		categoryGroup.POST(route, controller.SeedCategoryHandler(dbHandler))
	}

	return router
}
