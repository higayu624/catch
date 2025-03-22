package appconfig

import (
	"embed"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type PostgresInfo struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	User     string `yaml:"user"`
	SslMode  string `yaml:"sslmode"`
}

var staticYamlDir embed.FS

// DatabaseInfo returns struct for connecting to postgres
func DatabaseInfo() *PostgresInfo {
	fileName := "config/yaml/config.local.yaml"
	b, err := os.ReadFile(fileName)
	if err != nil {
		panic("failed to read. file=" + fileName + " err=" + err.Error())
	}
	gin.SetMode(gin.DebugMode)
	config := PostgresInfo{}
	yaml.Unmarshal(b, &config)
	return &config
}

// DatabaseInfoSqlboiler returns struct for connecting to postgres for sqlboiler
func DatabaseInfoSqlboiler() *PostgresInfo {
	fileName := "config/yaml/config.local.yaml"
	b, err := os.ReadFile(fileName)
	if err != nil {
		panic("failed to read. file=" + fileName + " err=" + err.Error())
	}
	config := PostgresInfo{}
	yaml.Unmarshal(b, config)
	return &config
}
