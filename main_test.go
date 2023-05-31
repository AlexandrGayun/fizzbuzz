package main

import (
	"fizzbuzztest/processor"
	"fizzbuzztest/processor/linear"
	"fmt"
	"strconv"
	"testing"
)

func TestFizzbuzz(t *testing.T) {
	testCases := []struct {
		input  int
		output string
	}{
		{0, "FizzBuzz"},
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "7"},
		{8, "8"},
		{9, "Fizz"},
		{10, "Buzz"},
		{11, "11"},
		{12, "Fizz"},
		{13, "13"},
		{14, "14"},
		{15, "FizzBuzz"},
		{16, "16"},
		{17, "17"},
		{18, "Fizz"},
	}
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
	fizzbuzz, err := processor.CreateNewProcessor(lc, handlers)
	if err != nil {
		t.Error("processor init failed")
	}
	for _, test := range testCases {
		output, err := fizzbuzz.Run(test.input)
		if err != nil {
			t.Error("processor internal error ", err)
		}
		var outputString string
		switch v := output.(type) {
		case string:
			outputString = v
		case int:
			outputString = strconv.Itoa(v)
		}
		if test.output != outputString {
			t.Errorf("test failed, input %s, output %s", output, test.output)
		}
	}
}
