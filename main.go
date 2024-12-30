package main

import (
	"fullstack-beego-app/middlewares"
	_ "fullstack-beego-app/middlewares"
	_ "fullstack-beego-app/models"
	_ "fullstack-beego-app/routers"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/beego/beego/v2/server/web/swagger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func init() {

	err := godotenv.Load()

	err = orm.RegisterDriver(
		"postgres",
		orm.DRPostgres,
	)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Database credentials retrieval
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" || dbPort == "" || dbSSLMode == "" {
		log.Fatalln("One or more required environment variables are missing")
	}

	connectionString := "user=" + dbUser +
		" password=" + dbPassword +
		" dbname=" + dbName +
		" host=" + dbHost +
		" port=" + dbPort +
		" sslmode=" + dbSSLMode

	log.Println("Connecting to database with connection string: " + connectionString)

	err = orm.RegisterDataBase(
		"default",
		"postgres",
		connectionString,
	)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func main() {

	// Debugging

	orm.Debug = true

	err := orm.RunSyncdb(
		"default",
		false,
		true,
	)

	if err != nil {
		log.Fatalln(err)
	}

	// Register the middleware
	beego.InsertFilter("*", beego.BeforeRouter, middlewares.LoggingMiddleware)
	beego.InsertFilter("*", beego.BeforeRouter, middlewares.AuthMiddleware)
	beego.InsertFilter("*", beego.BeforeRouter, middlewares.CORSMiddleware)
	beego.InsertFilter("*", beego.BeforeRouter, middlewares.CacheControlMiddleware)
	beego.InsertFilter("*", beego.BeforeRouter, middlewares.MethodOverride)
	beego.InsertFilter("*", beego.BeforeRouter, middlewares.AdminOnlyMiddleware)

	beego.Run()
}
