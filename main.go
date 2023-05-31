package main

import (
	"fizzbuzztest/processor"
	"fizzbuzztest/processor/linear"
	"fmt"
)

type UnexpectedInputArgumentError struct {
	HandlerName string
	Err         error
}

func (e UnexpectedInputArgumentError) Error() string {
	return fmt.Sprintf("error %s, happened in handler: %s", e.Err, e.HandlerName)
}

func main() {
	lc := linear.CreateLinearCore()
	var fizzbuzzHandler processor.Handler = func(args ...any) (any, error) {
		inputArg, ok := args[0].(int)
		if !ok {
			return nil, UnexpectedInputArgumentError{Err: fmt.Errorf("type conversion failed, argument %v is not a int", args[0]), HandlerName: "FizzbuzzHandler"}
		}
		if inputArg%15 == 0 {
			return "FizzBuzz", nil
		}
		return inputArg, nil
	}
	var fizzHandler processor.Handler = func(args ...any) (any, error) {
		switch v := args[0].(type) {
		case int:
			if v%3 == 0 {
				return "Fizz", nil
			} else {
				return v, nil
			}
		case string:
			return v, nil
		default:
			return nil, UnexpectedInputArgumentError{Err: fmt.Errorf("type conversion failed, argument %v is not a int or string", args[0]), HandlerName: "FizzHandler"}
		}
	}
	var buzzHandler processor.Handler = func(args ...any) (any, error) {
		switch v := args[0].(type) {
		case int:
			if v%5 == 0 {
				return "Buzz", nil
			} else {
				return v, nil
			}
		case string:
			return v, nil
		default:
			return nil, UnexpectedInputArgumentError{Err: fmt.Errorf("type conversion failed, argument %v is not a int or string", args[0]), HandlerName: "BuzzHandler"}
		}
	}
	handlers := []processor.Handler{fizzbuzzHandler, fizzHandler, buzzHandler}
	p, err := processor.CreateNewProcessor(lc, handlers)
	if err != nil {
		fmt.Println("processor init error: ", err)
	}
	for i := 1; i < 106; i++ {
		res, err := p.Run(i)
		if err != nil {
			fmt.Println("processor run error: ", err)
		}
		fmt.Println(res)
	}
}
