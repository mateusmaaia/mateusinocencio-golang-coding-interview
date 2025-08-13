package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"

	"github.com/BEON-Tech-Studio/golang-live-coding-challenge/internal/config"
	"github.com/BEON-Tech-Studio/golang-live-coding-challenge/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetStatesJson_Success(t *testing.T) {
	config.LoadConfig([]string{"../config"}, "config")

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/states", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := getStatesJson(c)

	if err != nil {
		// If external API is not available, skip this test
		t.Skip("External API not available:", err.Error())
		return
	}

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "application/json; charset=UTF-8", rec.Header().Get("Content-Type"))

	var states []models.State
	err = json.Unmarshal(rec.Body.Bytes(), &states)
	assert.NoError(t, err)

	// Validate array is populated with actual states
	assert.NotEmpty(t, states, "States array should contain data")
	assert.Greater(t, len(states), 0, "Should have at least one state")
}

func TestHome_ReturnsHTML(t *testing.T) {
	// Load actual templates for testing
	htmlTemplates = make(map[string]*template.Template)

	// Load base template
	baseTemplate, err := template.ParseFiles("../web/templates/base-template.html")
	assert.NoError(t, err)
	htmlTemplates["base-template"] = baseTemplate

	// Load home template
	homeTemplate, err := template.ParseFiles("../web/templates/home.html")
	assert.NoError(t, err)
	htmlTemplates["home"] = homeTemplate

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = home(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "text/html; charset=UTF-8", rec.Header().Get("Content-Type"))

	// Verify response contains expected HTML content
	responseBody := rec.Body.String()
	assert.Contains(t, responseBody, "<html>")
	assert.Contains(t, responseBody, "<title>Home</title>")
	assert.Contains(t, responseBody, "<h1>Home</h1>")
	assert.Contains(t, responseBody, "Welcome!")
	assert.Contains(t, responseBody, "</html>")
}

func TestHome_WithoutTemplates(t *testing.T) {
	// Test behavior when templates are not loaded
	htmlTemplates = make(map[string]*template.Template)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := home(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "text/html; charset=UTF-8", rec.Header().Get("Content-Type"))

	// Response should be empty when templates are missing
	responseBody := rec.Body.String()
	assert.Empty(t, responseBody)
}
