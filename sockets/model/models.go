package model

import "fmt"

type Request struct {
	Num1   float64 `json:"num1"`
	Num2   float64 `json:"num2"`
	Action string  `json:"action"`
}

type Response struct {
	Result float64
}

func (r Response) Show() {
	fmt.Printf("Result: %f\n", r.Result)
}
