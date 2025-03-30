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
	if err != nil {
		panic(err)
	}
	sqlBoiler, err := appdb.DatabaseConnectorSqlboiler(dbConfig)
	if err != nil {
		panic(err)
	}
	defer sqlBoiler.Close()
	boil.DebugMode = true
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

	// seed db
	// appdb.Seed(gorm)
	// log.Printf("seed successed!")

	log.Printf("server activation")

	router := initRouter(sqlBoiler)
	router.Run(":8080")
}
