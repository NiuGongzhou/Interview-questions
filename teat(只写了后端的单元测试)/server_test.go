// backend/server_test.go
package main

import (
	"calculator/proto/calculator/v1/calculatorv1connect"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPRoutes(t *testing.T) {
	t.Parallel()

	// 初始化服务
	calculator := &CalculatorServer{}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	path, handler := calculatorv1connect.NewCalculatorServiceHandler(calculator)
	mux.Handle(path, handler)

	// 测试健康检查端点

	t.Run("health check", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	// 测试Connect协议端点
	t.Run("connect endpoint", func(t *testing.T) {
		req := httptest.NewRequest(
			"POST",
			"/calculator.v1.CalculatorService/Calculate",
			strings.NewReader(`{"num1":1,"num2":2,"operation":"+"}`),
		)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Connect-Protocol-Version", "1")
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var response struct{ Result float64 }
		json.NewDecoder(w.Body).Decode(&response)
		assert.Equal(t, 3.0, response.Result)
	})

}
