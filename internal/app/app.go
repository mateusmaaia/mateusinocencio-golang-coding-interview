package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	SetDB sync.Once

	db *gorm.DB
)

func Start(database *gorm.DB) {
	err := initHtmlTemplates()
	if err != nil {
		panic(err)
	}

	SetDB.Do(func() {
		db = database
	})

	port := fmt.Sprintf("%d", viper.GetInt("port"))
	fmt.Println("Listening in port: " + port)
	listenAndServe(port)
}

func listenAndServe(port string) {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// HTML endpoints
	e.GET("/", home)
	e.GET("/home", home)
	e.GET("/states", getStates)

	// REST endpoints
	e.GET("api/states", getStatesJson)
	e.GET("api/categories", getCategoriesJson)

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}

func initHtmlTemplates() error {
	fmt.Println("Loading html templates...")
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	path = filepath.Join(path, "./internal/web/templates")

	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	templates := make(map[string]*template.Template)
	for _, file := range files {
		filename := file.Name()
		templateName := strings.TrimSuffix(filename, filepath.Ext(filename))

		fmt.Println("Processing template: ", templateName)
		fileBytes, err := ioutil.ReadFile(filepath.Join(path, filename))
		if err != nil {
			return err
		}

		templates[templateName], err = template.New(templateName).Parse(string(fileBytes))
		if err != nil {
			return err
		}
	}

	htmlTemplates = templates
	return nil
}
