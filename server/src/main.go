package main

import (
	"log"
	"os"
	"os/exec"

	appconfig "catch/config"
	appdb "catch/database"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	dbConfig := appconfig.DatabaseInfo()
	gorm, err := appdb.DatabaseConnector(dbConfig)
	sqlBoiler, err := appdb.DatabaseConnectorSqlboiler(dbConfig)
	defer sqlBoiler.Close()
	boil.DebugMode = true
	if err != nil {
		log.Print("接続失敗", err)
	}
	log.Printf("db connect successed!")

	// migrate db
	appdb.Migrate(gorm)
	log.Printf("migrate successed!")

	// create sqlboiler models
	cmd := exec.Command(
		"sqlboiler",
		"psql",
		"-c", "config/yaml/sqlboiler.yaml",
		"-o", "model/models",
		"--no-tests",
		"--wipe",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// model作成実行
	err = cmd.Run()
	if err != nil {
		log.Print("can't create sqlboiler models", err)
	}
	log.Print("create sqlboiler models")

	// seed category
	appdb.SeedCategorys(sqlBoiler)
	log.Print("seed Categorys successed")

	// seed db
	// appdb.Seed(gorm)
	// log.Printf("seed successed!")

	log.Printf("server activation")

	router := initRouter(sqlBoiler)
	router.Run(":8080")
}
