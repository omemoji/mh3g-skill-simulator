package controllers_test

import (
	"encoding/json"
	"mh3g-skill-simulator/internal/controllers"
	"mh3g-skill-simulator/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type MockSimulatorService struct {
	ExecuteFunc func(searchQuery models.SearchQuery) ([]models.Hunter, error)
}

func (m *MockSimulatorService) Execute(searchQuery models.SearchQuery) ([]models.Hunter, error) {
	return m.ExecuteFunc(searchQuery)
}

func TestSimulatorController_GetHunters(t *testing.T) {
	// モックサービスの準備
	mockService := new(MockSimulatorService)
	expectedHunters := []models.Hunter{
		{
			Head:  &models.Equipment{Name: "Test Head"},
			Body:  &models.Equipment{Name: "Test Body"},
			Arms:  &models.Equipment{Name: "Test Arms"},
			Waist: &models.Equipment{Name: "Test Waist"},
			Legs:  &models.Equipment{Name: "Test Legs"},
		},
	}
	mockService.ExecuteFunc = func(searchQuery models.SearchQuery) ([]models.Hunter, error) {
		return expectedHunters, nil
	}

	// コントローラーの初期化
	ctrl := controllers.NewSimulatorController(mockService)

	// Ginルーターの作成
	router := gin.New()
	router.GET("/api/v1", ctrl.GetHunters)

	// HTTPリクエストの作成
	req, _ := http.NewRequest("GET", "/api/v1?max_results=10", nil)
	w := httptest.NewRecorder()

	// リクエストの実行
	router.ServeHTTP(w, req)

	// ステータスコードの検証
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// レスポンスボディの検証
	var response controllers.SimulatorResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// レスポンスの内容検証
	if len(response.Hunters) != len(expectedHunters) {
		t.Errorf("Expected %d hunters, got %d", len(expectedHunters), len(response.Hunters))
	}

	if len(response.Hunters) > 0 {
		hunter := response.Hunters[0]
		if hunter.Head != expectedHunters[0].Head.Name {
			t.Errorf("Expected head to be %s, got %s", expectedHunters[0].Head.Name, hunter.Head)
		}
		if hunter.Body != expectedHunters[0].Body.Name {
			t.Errorf("Expected body to be %s, got %s", expectedHunters[0].Body.Name, hunter.Body)
		}
		if hunter.Arms != expectedHunters[0].Arms.Name {
			t.Errorf("Expected arms to be %s, got %s", expectedHunters[0].Arms.Name, hunter.Arms)
		}
		if hunter.Waist != expectedHunters[0].Waist.Name {
			t.Errorf("Expected waist to be %s, got %s", expectedHunters[0].Waist.Name, hunter.Waist)
		}
		if hunter.Legs != expectedHunters[0].Legs.Name {
			t.Errorf("Expected legs to be %s, got %s", expectedHunters[0].Legs.Name, hunter.Legs)
		}
	}

}
