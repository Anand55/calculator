package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Anand55/calculator/proto"
	"google.golang.org/grpc"
)

// Creating InputObj which stores method name, first number and second number
type InputObj struct{
	Method 		string
	FirstInput  float32
	SecondInput float32
}

// Getting command line args and constructing InputObj struct
func GetInputObject(argsArr []string) (InputObj,error) {
	inputObj := InputObj{}
	if argsArr[0] != "-method"{
		fmt.Printf("Please provide proper command\n")
		return InputObj{}, fmt.Errorf("Incorrect command format\n")
	}
	inputObj.Method = argsArr[1]
	firstInput, err := strconv.ParseFloat(argsArr[3],32) 
	if err != nil {
		log.Fatalf("Error converting args to integer: %s", err)
		return InputObj{},err
	}
	inputObj.FirstInput = float32(firstInput)
	secondInput, err := strconv.ParseFloat(argsArr[5],32) 
	if err != nil {
		log.Fatalf("Error converting args to integer: %s", err)
		return InputObj{},err
	}
	inputObj.SecondInput = float32(secondInput)
	return inputObj, nil
}

func main() {
	fmt.Println("Client..")
	inputObj, err := GetInputObject(os.Args[1:])
	if err != nil{
		log.Fatal("Error getting input object\n",err)
	}
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect %v", err)
	}
	defer conn.Close()
	c := proto.NewCalculatorClient(conn)
	
	doCalculation(c,inputObj)
	
}

// Checking for given calculation method and calling rpc functions accordingly 
func doCalculation(c proto.CalculatorClient, inputVal InputObj) {
	req := &proto.CalculationRequest{
		A : inputVal.FirstInput,
		B : inputVal.SecondInput,
	}
	
	if inputVal.Method == "add" {
		res, err := c.Addition(context.Background(), req)
		if err != nil {
			log.Fatal("error while calling Addition rpc %v", err)
		}
		fmt.Printf("Response from server: %v\n", res.Result)
	}else if inputVal.Method == "sub"{
		res, err := c.Subtraction(context.Background(), req)
		if err != nil {
			log.Fatal("error while calling Substraction rpc %v", err)
		}
		fmt.Printf("Response from server: %v\n", res.Result)
	}else if inputVal.Method == "mul"{
		res, err := c.Multiplication(context.Background(), req)
		if err != nil {
			log.Fatal("error while calling Multiplication rpc %v", err)
		}
		fmt.Printf("Response from server: %v\n", res.Result)
	}else if inputVal.Method == "div"{
		if req.B != 0{
			res, err := c.Division(context.Background(), req)
			if err != nil {
				log.Fatal("error while calling Division rpc %v", err)
			}
			fmt.Printf("Response from server: %v\n", res.Result)
		}else{
			fmt.Printf("Invalid Input: Division by 0\n")
		}
	}else{
		fmt.Printf("Invalid method [%s], Please give valid method\n",inputVal.Method)
		fmt.Printf("Available valid methods: add, sub, mul, div\n")
	}
}