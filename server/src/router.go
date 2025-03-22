package main

import (
	"catch/controller"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var productName = "catch"
var version = "v1"
var APIRoot = version + "/" + productName + "/"

func endpoint(path string) string {
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

	router.POST(endpoint("customer"), controller.PostCustomer(dbHandler))

	return router
}
