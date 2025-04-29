package main

import (
	"calculator/proto/calculator/v1"
	"calculator/proto/calculator/v1/calculatorv1connect"
	"context"
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"

	"github.com/rs/cors"
)

type CalculatorServer struct{}

func (s *CalculatorServer) Calculate(
	ctx context.Context,
	req *connect.Request[calculatorv1.CalculateRequest],
) (*connect.Response[calculatorv1.CalculateResponse], error) {
	log.Println("Request headers: ", req.Header())

	num1 := req.Msg.Num1
	num2 := req.Msg.Num2
	var result float64

	switch req.Msg.Operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("division by zero"))
		}
		result = num1 / num2
	default:
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid operation"))
	}

	res := connect.NewResponse(&calculatorv1.CalculateResponse{
		Result: result,
	})
	res.Header().Set("Calculator-Version", "v1")
	return res, nil
}

func main() {
	calculator := &CalculatorServer{}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
			return
		}
		http.NotFound(w, r)
	})
	path, handler := calculatorv1connect.NewCalculatorServiceHandler(calculator)
	mux.Handle(path, handler)
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // 前端地址
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Connect-Protocol-Version"},
		AllowCredentials: true,
		Debug:            true, // 开发环境开启调试
	}).Handler(mux)
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe("localhost:8080", corsHandler)
}
