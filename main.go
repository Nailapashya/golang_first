package main

import (
	"context"
	"fmt"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	controllers "microdata/kemendagri/sipd/service/boilerplate_go/controller"
	hdl "microdata/kemendagri/sipd/service/boilerplate_go/handler"
	"microdata/kemendagri/sipd/service/boilerplate_go/handler/configs"
	"microdata/kemendagri/sipd/service/boilerplate_go/handler/http_util"
	_deliveryMiddleware "microdata/kemendagri/sipd/service/boilerplate_go/handler/middleware"
	"microdata/kemendagri/sipd/service/boilerplate_go/utils"
	"os"
	"strconv"
	"time"

	// docs are generated by Swag CLI, you have to import them.
	_ "microdata/kemendagri/sipd/service/boilerplate_go/docs" // load API Docs files (Swagger)
)

var serverName, serverUrl, serverReadTimeout, alwOrg, dbServerUrl string

var pgxConn *pgxpool.Pool
var vld *validator.Validate
var err error

func init() {
	// Server Env-
	serverName = os.Getenv("SERVER_NAME")
	if serverName == "" {
		exitf("SERVER_NAME env is required")
	}
	serverUrl = os.Getenv("SERVER_URL")
	if serverUrl == "" {
		exitf("SERVER_URL env is required")
	}
	serverReadTimeout = os.Getenv("SERVER_READ_TIMEOUT")
	if serverReadTimeout == "" {
		exitf("SERVER_READ_TIMEOUT env is required")
	}

	// CORS
	alwOrg = os.Getenv("SIPD_CORS_WHITELISTS")
	if alwOrg == "" {
		exitf("SIPD_CORS_WHITELISTS config is required")
	}

	// Databse Env
	dbServerUrl = os.Getenv("DB_SERVER_URL")
	if dbServerUrl == "" {
		exitf("DB_SERVER_URL config is required")
	}
}

func dbConnection() {
	var maxConnLifetime, maxConnIdleTime time.Duration
	maxConnLifetime = 5 * time.Minute
	maxConnIdleTime = 2 * time.Minute

	var cfg *pgxpool.Config

	// sipd_master_data
	cfg, err = pgxpool.ParseConfig(dbServerUrl + " application_name=" + serverName)
	if err != nil {
		exitf("Unable to create db pool config sipd_master_data %v\n", err)
	}
	cfg.MaxConns = 1000                   // Maximum total connections in the pool
	cfg.MaxConnLifetime = maxConnLifetime // Maximum lifetime of a connection
	cfg.MaxConnIdleTime = maxConnIdleTime // Maximum time a connection can be idle
	pgxConn, err = pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		exitf("Unable to connect to database sipd_master_data %v\n", err)
	}
}

//	@title						SIPD Service Boilerpate
//	@version					1.0
//	@description				SIPD Service Boilerpate Rest API.
//	@termsOfService				http://swagger.io/terms/
//	@contact.name				API Support
//	@contact.email				lifelinejar@mail.com
//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@BasePath					/
func main() {
	// Pgx Pool Connection
	dbConnection()

	// close all open component
	defer func() {
		pgxConn.Close()
	}()

	serverReadTimeoutInt, err := strconv.Atoi(serverReadTimeout)
	if err != nil {
		exitf("failed casting timeout context: ", err)
	}
	timeoutContext := time.Duration(serverReadTimeoutInt) * time.Second

	// Define a validator
	vld = utils.NewValidator()

	// Define Fiber config.
	config := configs.FiberConfig()
	app := fiber.New(config)
	middL := _deliveryMiddleware.InitMiddleware(app)
	app.Use(middL.CORS())
	app.Use(middL.LOGGER())

	// Swagger handler
	app.Get("/swagger/*", swagger.HandlerDefault)

	// public router
	siteCtl := controllers.NewSiteController(pgxConn, timeoutContext)
	hdl.NewSiteHandler(app, vld, siteCtl)
	// end public router

	// strict router
	rStrict := app.Group("/strict", middL.JWT()) // router for api private access

	hdl.NewUrusanHandler(rStrict, vld, controllers.NewUrusanController(pgxConn, timeoutContext))
	// end strict router

	http_util.StartServer(app)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}