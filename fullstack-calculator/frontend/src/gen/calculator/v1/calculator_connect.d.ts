// @generated by protoc-gen-connect-es v1.6.1
// @generated from file calculator/v1/calculator.proto (package calculator.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CalculateRequest, CalculateResponse } from "./calculator_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service calculator.v1.CalculatorService
 */
export declare const CalculatorService: {
  readonly typeName: "calculator.v1.CalculatorService",
  readonly methods: {
    /**
     * @generated from rpc calculator.v1.CalculatorService.Calculate
     */
    readonly calculate: {
      readonly name: "Calculate",
      readonly I: typeof CalculateRequest,
      readonly O: typeof CalculateResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

