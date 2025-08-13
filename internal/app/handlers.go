package app

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"

	"github.com/BEON-Tech-Studio/golang-live-coding-challenge/internal/models"
)

var htmlTemplates map[string]*template.Template

type params map[string]interface{}

func home(c echo.Context) error {
	return c.HTML(http.StatusOK, execTemplateFromBase("Home", "home", params{}))
}

func getStates(c echo.Context) error {
	var states []models.State

	result := db.Table("states").Find(&states)
	if result.RowsAffected == 0 {
		statesData, err := FetchStates()
		if err != nil {
			// %w will pass the error value to the caller which means if you use error.Is/error.As you can check the error type
			return fmt.Errorf("failed to fetch states: %w", err)
		}

		// Save the fetched states to the database
		if len(statesData.States) > 0 {
			result := db.Create(&statesData.States)
			if result.Error != nil {
				return fmt.Errorf("failed to save states to database: %w", result.Error)
			}
		}

		states = statesData.States
	}

	statesHtml := execTemplateFromBase("States", "states", params{
		"states":       states,
		"row_template": htmlTemplates["states-row"],
	})
	return c.HTML(http.StatusOK, statesHtml)
}

func getStatesJson(c echo.Context) error {
	var states []models.State
	result := db.Table("states").Find(&states)
	if result.RowsAffected == 0 {
		statesData, err := FetchStates()
		if err != nil {
			// %w will pass the error value to the caller which means if you use error.Is/error.As you can check the error type
			return fmt.Errorf("failed to fetch states: %w", err)
		}

		// Save the fetched states to the database
		if len(statesData.States) > 0 {
			result := db.Create(&statesData.States)
			if result.Error != nil {
				return fmt.Errorf("failed to save states to database: %w", result.Error)
			}
		}

		states = statesData.States
	}

	return c.JSON(http.StatusOK, states)
}

func getStateByCode(c echo.Context) error {
	codeID := c.Param("code")

	var state models.State
	result := db.First(&state, "code = ?", codeID)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "State not found"})
	}

	return c.JSON(http.StatusOK, state)
}

func getReports(c echo.Context) error {
	reports, err := FetchReports()
	if err != nil {
		return fmt.Errorf("failed to fetch reports") // api should not know internal error // add a error wrapper
	}

	return c.JSON(http.StatusOK, reports)
}

func getCategoriesJson(c echo.Context) error {
	var categories []models.Category

	return c.JSON(http.StatusOK, categories)
}

// Helpers

func execTemplateFromBase(title, templateName string, p params) string {
	return execTemplate(
		"base-template",
		params{
			"title": title,
			"body":  execTemplate(templateName, p),
		},
	)
}

func execTemplate(templateName string, p params) string {
	t := htmlTemplates[templateName]
	if t == nil {
		return ""
	}
	var res bytes.Buffer
	err := t.Execute(&res, p)
	if err != nil {
		fmt.Println("ERROR in execTemplate:", err.Error())
		return ""
	}
	return res.String()
}
