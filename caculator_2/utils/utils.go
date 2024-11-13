package utils

import (
	"errors"
	"fmt"
	"golang.org/x/exp/slices"
	"sockets/entity"
	"strconv"
	"strings"
)

func validateInput(input string) bool {
	result := strings.Split(input, " ")
	if len(result) != 3 {
		return false
	}
	operation := result[1]
	symbols := []string{"+", "-", "*", "/", "**"}
	contains := slices.Contains(symbols, operation)
	if !contains {
		return false
	}
	return true
}

func InputToRequest(input string) (*entity.Request, error) {
	isValid := validateInput(input)
	if !isValid {
		return nil, errors.New("invalid input")
	}
	result := strings.Split(input, " ")
	num1, err := strconv.ParseFloat(result[0], 64)
	if err != nil {
		return nil, err
	}
	num2, err := strconv.ParseFloat(result[2], 64)
	if err != nil {
		return nil, err
	}
	return &entity.Request{
		Num1:      num1,
		Num2:      num2,
		Operation: result[1],
	}, nil
}

func GetReponse(calculator entity.Calculator, req entity.Request) entity.Response {
	value, err := calculator.Calculate(req.Num1, req.Num2, req.Operation)
	if err != nil {
		return entity.Response{
			Result: err.Error(),
		}
	}
	return entity.Response{
		Result: fmt.Sprintf("%s = %f", req.GetInput(), value),
	}
}
