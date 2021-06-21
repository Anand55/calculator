package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Anand55/calculator/proto"
	"google.golang.org/grpc"
)

type Server struct{
	proto.UnimplementedCalculatorServer
}

// Addition method return sum of a and b
func (*Server) Addition(ctx context.Context, req *proto.CalculationRequest) (*proto.CalculationResponse, error) {
	fmt.Printf("Addition function invoked %v\n", req)
	a := req.GetA()
	b := req.GetB()
	result := a + b
	res := proto.CalculationResponse{
		Result: result,
	}
	return &res, nil
}

// Substaction method return difference of a and b
func (*Server) Subtraction(ctx context.Context, req *proto.CalculationRequest) (*proto.CalculationResponse, error) {
	fmt.Printf("Subtraction function invoked %v\n", req)
	a := req.GetA()
	b := req.GetB()
	result := a - b
	res := proto.CalculationResponse{
		Result: result,
	}
	return &res, nil
}

// Multiplication method return product of a and b
func (*Server) Multiplication(ctx context.Context, req *proto.CalculationRequest) (*proto.CalculationResponse, error) {
	fmt.Printf("Multiplication function invoked %v\n", req)
	a := req.GetA()
	b := req.GetB()
	result := a * b
	res := proto.CalculationResponse{
		Result: result,
	}
	return &res, nil
}

// Division method return division of a and b
func (*Server) Division(ctx context.Context, req *proto.CalculationRequest) (*proto.CalculationResponse, error) {
	fmt.Printf("Division function invoked %v\n", req)
	a := req.GetA()
	b := req.GetB()
	if b != 0{
		result := a / b
		res := proto.CalculationResponse{
			Result: result,
		}
		
		return &res, nil
	}else{
		return nil, fmt.Errorf("Division by 0 %v",nil)
	}
}

func main(){
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Failed to listen", err)
	}
	fmt.Println("Server Started..")
	s := grpc.NewServer()
	proto.RegisterCalculatorServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: ", err)
	}
}