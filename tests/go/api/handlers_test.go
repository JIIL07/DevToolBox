package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/JIIL07/devtoolbox/internal/api"
	"github.com/JIIL07/devtoolbox/internal/core"
)

func TestHandler_Generate(t *testing.T) {
	registry := core.NewGeneratorRegistry()
	handler := api.NewHandler(registry)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/generate", handler.Generate)

	tests := []struct {
		name           string
		request        api.GenerateRequest
		expectedStatus int
		expectError    bool
	}{
		{
			name: "valid request",
			request: api.GenerateRequest{
				Template: "go-struct",
				Input:    `{"name":"string","age":"int"}`,
			},
			expectedStatus: http.StatusOK,
			expectError:    false,
		},
		{
			name: "invalid template",
			request: api.GenerateRequest{
				Template: "invalid-template",
				Input:    `{"name":"string"}`,
			},
			expectedStatus: http.StatusNotFound,
			expectError:    true,
		},
		{
			name: "invalid JSON input",
			request: api.GenerateRequest{
				Template: "go-struct",
				Input:    `{"name":"string",}`,
			},
			expectedStatus: http.StatusInternalServerError,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tt.request)
			req, _ := http.NewRequest("POST", "/generate", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			var response api.GenerateResponse
			json.Unmarshal(w.Body.Bytes(), &response)

			if tt.expectError && response.Error == "" {
				t.Error("expected error in response")
			}
			if !tt.expectError && response.Code == "" {
				t.Error("expected code in response")
			}
		})
	}
}

func TestHandler_ListGenerators(t *testing.T) {
	registry := core.NewGeneratorRegistry()
	handler := api.NewHandler(registry)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/generators", handler.ListGenerators)

	req, _ := http.NewRequest("GET", "/generators", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response api.ListGeneratorsResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	if len(response.Generators) == 0 {
		t.Error("expected at least one generator")
	}

	found := false
	for _, gen := range response.Generators {
		if gen.Name == "go-struct" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected go-struct generator in list")
	}
}

func TestHandler_Health(t *testing.T) {
	registry := core.NewGeneratorRegistry()
	handler := api.NewHandler(registry)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/health", handler.Health)

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
}
