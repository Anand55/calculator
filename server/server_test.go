package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	pb "github.com/Anand55/calculator/proto"
)


func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)
 
	server := grpc.NewServer()
 
	pb.RegisterCalculatorServer(server, &Server{})
 
	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()
 
	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestAddition(t *testing.T) {
	tests := []struct {
		name    string
		a  float32
		b  float32
		res     *pb.CalculationResponse
		errCode codes.Code
		errMsg  string
	}{
		{
			"Add1",
			5,
			3,
			&pb.CalculationResponse{Result: 8},
			codes.OK,
			"",
		},
		{
			"Add2",
			-3,
			2,
			&pb.CalculationResponse{Result: -1},
			codes.OK,
			"",
		},
		{
			"Add3",
			-3,
			-2,
			&pb.CalculationResponse{Result: -5},
			codes.OK,
			"",
		},
	}
 
	ctx := context.Background()
 
	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
 
	client := pb.NewCalculatorClient(conn)
 
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := &pb.CalculationRequest{A: tt.a,B: tt.b}
 
			response, err := client.Addition(ctx, request)
 
			if response != nil {
				if response.GetResult() != tt.res.GetResult() {
					t.Error("response: expected", tt.res.GetResult(), "received", response.GetResult())
				}
			}
 
			if err != nil {
				if er, ok := status.FromError(err); ok {
					if er.Code() != tt.errCode {
						t.Error("error code: expected", codes.InvalidArgument, "received", er.Code())
					}
					if er.Message() != tt.errMsg {
						t.Error("error message: expected", tt.errMsg, "received", er.Message())
					}
				}
			}
		})
	}
}


func TestSubtraction(t *testing.T) {
	tests := []struct {
		name    string
		a  float32
		b  float32
		res     *pb.CalculationResponse
		errCode codes.Code
		errMsg  string
	}{
		{
			"Sub1",
			5,
			3,
			&pb.CalculationResponse{Result: 2},
			codes.OK,
			"",
		},
		{
			"Sub2",
			-3,
			2,
			&pb.CalculationResponse{Result: -5},
			codes.OK,
			"",
		},
		{
			"Sub3",
			-3,
			-2,
			&pb.CalculationResponse{Result: -1},
			codes.OK,
			"",
		},
	}
 
	ctx := context.Background()
 
	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
 
	client := pb.NewCalculatorClient(conn)
 
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := &pb.CalculationRequest{A: tt.a,B: tt.b}
 
			response, err := client.Subtraction(ctx, request)
 
			if response != nil {
				if response.GetResult() != tt.res.GetResult() {
					t.Error("response: expected", tt.res.GetResult(), "received", response.GetResult())
				}
			}
 
			if err != nil {
				if er, ok := status.FromError(err); ok {
					if er.Code() != tt.errCode {
						t.Error("error code: expected", codes.InvalidArgument, "received", er.Code())
					}
					if er.Message() != tt.errMsg {
						t.Error("error message: expected", tt.errMsg, "received", er.Message())
					}
				}
			}
		})
	}
}


func TestMultiplication(t *testing.T) {
	tests := []struct {
		name    string
		a  float32
		b  float32
		res     *pb.CalculationResponse
		errCode codes.Code
		errMsg  string
	}{
		{
			"Mul1",
			5,
			3,
			&pb.CalculationResponse{Result: 15},
			codes.OK,
			"",
		},
		{
			"Mul2",
			-3,
			2,
			&pb.CalculationResponse{Result: -6},
			codes.OK,
			"",
		},
		{
			"Mul3",
			-3,
			-2,
			&pb.CalculationResponse{Result: 6},
			codes.OK,
			"",
		},
		{
			"Mul4",
			-3,
			0,
			&pb.CalculationResponse{Result: 0},
			codes.OK,
			"",
		},
	}
 
	ctx := context.Background()
 
	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
 
	client := pb.NewCalculatorClient(conn)
 
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := &pb.CalculationRequest{A: tt.a,B: tt.b}
 
			response, err := client.Multiplication(ctx, request)
 
			if response != nil {
				if response.GetResult() != tt.res.GetResult() {
					t.Error("response: expected", tt.res.GetResult(), "received", response.GetResult())
				}
			}
 
			if err != nil {
				if er, ok := status.FromError(err); ok {
					if er.Code() != tt.errCode {
						t.Error("error code: expected", codes.InvalidArgument, "received", er.Code())
					}
					if er.Message() != tt.errMsg {
						t.Error("error message: expected", tt.errMsg, "received", er.Message())
					}
				}
			}
		})
	}
}



func TestDivision(t *testing.T) {
	tests := []struct {
		name    string
		a  float32
		b  float32
		res     *pb.CalculationResponse
		errCode codes.Code
		errMsg  string
	}{
		{
			"Div1",
			6,
			3,
			&pb.CalculationResponse{Result: 2},
			codes.OK,
			"",
		},
		{
			"Div2",
			-3,
			2,
			&pb.CalculationResponse{Result: -1.5},
			codes.OK,
			"",
		},
		{
			"Div3",
			-3,
			-2,
			&pb.CalculationResponse{Result: 1.5},
			codes.OK,
			"",
		},
		{
			"Div4",
			-3,
			0,
			nil,
			codes.Unknown,
			fmt.Sprintf("Division by 0 %v",nil),
		},
	}
 
	ctx := context.Background()
 
	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
 
	client := pb.NewCalculatorClient(conn)
 
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := &pb.CalculationRequest{A: tt.a,B: tt.b}
 
			response, err := client.Division(ctx, request)
 
			if response != nil {
				if response.GetResult() != tt.res.GetResult() {
					t.Error("response: expected", tt.res.GetResult(), "received", response.GetResult())
				}
			}
 
			if err != nil {
				if er, ok := status.FromError(err); ok {
					if er.Code() != tt.errCode {
						t.Error("error code: expected", codes.Unknown, "received", er.Code())
					}
					if er.Message() != tt.errMsg {
						t.Error("error message: expected", tt.errMsg, "received", er.Message())
					}
				}
			}
		})
	}
}
