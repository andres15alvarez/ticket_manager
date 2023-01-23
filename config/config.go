package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var db *bun.DB
var server *gin.Engine

type Config struct {
	Environment string
	DBUser      string
	DBPassword  string
	DBHost      string
	DBDatabase  string
	DBPort      string
}

func InitializeConfig() *Config {
	return &Config{
		Environment: os.Getenv("ENVIRONMENT"),
		DBUser:      os.Getenv("DATABASE_USER"),
		DBPassword:  os.Getenv("DATABASE_PASSWORD"),
		DBHost:      os.Getenv("DATABASE_HOST"),
		DBDatabase:  os.Getenv("DATABASE_NAME"),
		DBPort:      os.Getenv("DATABASE_PORT"),
	}
}

func (c *Config) InitializeDB() {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBDatabase,
	)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db = bun.NewDB(sqldb, pgdialect.New())
}

func GetDB() *bun.DB {
	return db
}
