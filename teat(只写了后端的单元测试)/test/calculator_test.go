// backend/calculator_test.go
package main

import (
	calculatorv1 "calculator/proto/calculator/v1"
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/require" 
)

func TestCalculatorService(t *testing.T) {
	t.Parallel()

	srv := &CalculatorServer{}

	tests := []struct {
		name     string
		num1     float64
		num2     float64
		op       string
		expected float64
		wantErr  bool
	}{
		{"addition", 2, 3, "+", 5, false},
		{"subtraction", 5, 3, "-", 2, false},
		{"multiplication", 2, 4, "*", 8, false},
		{"division", 6, 2, "/", 3, false},
		{"divide by zero", 5, 0, "/", 0, true},
		{"invalid op", 1, 1, "?", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := connect.NewRequest(&calculatorv1.CalculateRequest{
				Num1:      tt.num1,
				Num2:      tt.num2,
				Operation: tt.op,
			})

			res, err := srv.Calculate(context.Background(), req)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, res.Msg.Result)
		})
	}
}

func TestEdgeCases(t *testing.T) {
    srv := &CalculatorServer{}
    ctx := context.Background()

    // 测试大数计算
    t.Run("large number multiplication", func(t *testing.T) {
        req := connect.NewRequest(&calculatorv1.CalculateRequest{
            Num1:      1e20,
            Num2:      1e20,
            Operation: "*",
        })
        res, err := srv.Calculate(ctx, req)
        assert.NoError(t, err)
        assert.Equal(t, 1e40, res.Msg.Result)
    })

    // 测试协议头传递
    t.Run("headers propagation", func(t *testing.T) {
        req := connect.NewRequest(&calculatorv1.CalculateRequest{
            Num1:      1,
            Num2:      1,
            Operation: "+",
        })
        req.Header().Set("X-Trace-ID", "test123")

        res, err := srv.Calculate(ctx, req)
        assert.NoError(t, err)
        assert.Equal(t, "v1", res.Header().Get("Calculator-Version"))
    })
}
