"use client";

import { useState } from "react";
import { CalculatorService } from "../gen/calculator/v1/calculator_connect";
import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
// import { useEffect } from "react";

// useEffect(() => {
//   // 这里可以放任何需要在客户端执行的初始化逻辑
// }, []);
// export const dynamic = 'force-static';



export default function Calculator() {
  const [num1, setNum1] = useState("");
  const [num2, setNum2] = useState("");
  const [operation, setOperation] = useState("+");
  const [result, setResult] = useState<number | undefined>();
  const [error, setError] = useState("");

  const transport = createConnectTransport({
    baseUrl: "http://localhost:8080",
    useHttpGet: false, // 对简单请求使用GET
    interceptors: [
      (next) => async (req) => {
        // 确保协议版本头
        req.header.set("Connect-Protocol-Version", "1");
        return next(req);
      },
    ],
  });

  const client = createClient(CalculatorService, transport);

  const handleCalculate = async () => {
    try {
      setError("");
      type SafeCalculateResponse = { result: number };
      const response = await client.calculate({
        num1: parseFloat(num1),
        num2: parseFloat(num2),
        operation,
      }) as SafeCalculateResponse;

      setResult(response.result);
      // const response = await client.calculate({
      //   num1: parseFloat(num1),
      //   num2: parseFloat(num2),
      //   operation,
      // });
      // setResult(response.result);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Unknown error");
      setResult(undefined);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white p-8 rounded-lg shadow-md w-96">
        <h1 className="text-2xl font-bold mb-6 text-center">Calculator</h1>

        {error && (
          <div className="mb-4 p-2 bg-red-100 text-red-700 rounded">
            {error}
          </div>
        )}

        <div className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              First Number
            </label>
            <input
              type="number"
              value={num1}
              onChange={(e) => setNum1(e.target.value)}
              className="w-full p-2 border border-gray-300 rounded"
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Operation
            </label>
            <select
              value={operation}
              onChange={(e) => setOperation(e.target.value)}
              className="w-full p-2 border border-gray-300 rounded"
            >
              <option value="+">+</option>
              <option value="-">-</option>
              <option value="*">×</option>
              <option value="/">÷</option>
            </select>
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Second Number
            </label>
            <input
              type="number"
              value={num2}
              onChange={(e) => setNum2(e.target.value)}
              className="w-full p-2 border border-gray-300 rounded"
            />
          </div>

          <button
            onClick={handleCalculate}
            className="w-full bg-blue-600 text-white py-2 px-4 rounded hover:bg-blue-700 transition"
          >
            Calculate
          </button>

          {result !== undefined && (
            <div className="mt-4 p-3 bg-gray-50 rounded text-center">
              <p className="text-sm text-gray-600">Result:</p>
              <p className="text-2xl font-bold">{result}</p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}