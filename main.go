package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"os"

	"html/template"

	glog "github.com/labstack/gommon/log"

	"database/sql"
	"github.com/labstack/echo/engine/standard"
	"github.com/rumyantseva/vk-authorization-example/config"
	"github.com/rumyantseva/vk-authorization-example/handler"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	"io"
)

type Template struct {
	templates *template.Template
}

var (
	FDebug = flag.Bool("debug", true, "Run server in debug mode")
)

func main() {
	log.SetOutput(os.Stdout)

	e := echo.New()

	flag.Parse()
	if *FDebug {
		e.SetLogLevel(glog.DEBUG)
	} else {
		e.SetLogLevel(glog.INFO)
	}

	e.SetLogOutput(os.Stdout)

	var conf config.Config
	if _, err := toml.DecodeFile("config-dev.toml", &conf); err != nil {
		log.Fatalf("Couldn't open config file: %+v", err)
	}

	conn, err := sql.Open(
		"postgres",
		"postgres://postgres:postgres@localhost:5432/auth?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}

	db := reform.NewDB(
		conn,
		postgresql.Dialect,
		reform.NewPrintfLogger(log.Printf),
	)

	e.Use(middleware.Static("/static"))

	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.SetRenderer(t)

	h := handler.NewHandler(db, conf.VK)

	ui := e.Group("/ui")
	{
		e.Static("/static", "./static")
		ui.GET("/hello", h.Hello)
		ui.GET("/verify", h.Verify)
	}

	// ToDo: API for auth service
	//api := e.Group("/api")
	//{
	//	...
	//}

	e.Run(standard.New(conf.Server.Address))
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
