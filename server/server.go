package server

import (
	"database/sql"
	"fmt"
	"net"
	"os"

	"github.com/KobayashiTakaki/sample-webapi-mysql/controller"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	e  *echo.Echo
	db *sql.DB
}

func NewServer() (*Server, error) {
	e := echo.New()
	db, err := newDB()
	if err != nil {
		return nil, err
	}
	e.Use(middleware.Logger())
	postsController := controller.NewPostsController(db)
	e.GET("/posts", postsController.ListPosts)
	e.GET("/posts/:id", postsController.GetPost)
	return &Server{
		e:  e,
		db: db,
	}, nil
}

func (s *Server) Serve(address string) error {
	return s.e.Start(address)
}

func (s *Server) Close() error {
	return s.db.Close()
}

func newDB() (*sql.DB, error) {
	notSet := []string{}
	getEnv := func(name string) string {
		val := os.Getenv(name)
		if val == "" {
			notSet = append(notSet, name)
		}
		return val
	}
	dbUser := getEnv("WEBAPI_DB_USER")
	dbPassword := getEnv("WEBAPI_DB_PASSWORD")
	dbHost := getEnv("WEBAPI_DB_HOST")
	dbPort := getEnv("WEBAPI_DB_PORT")
	dbName := getEnv("WEBAPI_DB_NAME")

	if len(notSet) > 0 {
		return nil, fmt.Errorf("some environment variables are not set: %s", notSet)
	}

	mysqlConfig := mysql.NewConfig()
	mysqlConfig.User = dbUser
	mysqlConfig.Passwd = dbPassword
	mysqlConfig.Net = "tcp"
	mysqlConfig.Addr = net.JoinHostPort(dbHost, dbPort)
	mysqlConfig.DBName = dbName
	mysqlConfig.ParseTime = true

	fmt.Println(mysqlConfig.FormatDSN())
	db, err := sql.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}
