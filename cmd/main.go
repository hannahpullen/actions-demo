package main

import (
	"fmt"
	"os"

	"github.com/hannahpullen/actions-demo/pkg/adder"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		errorAndExit(fmt.Errorf("unable to create logger: %v", err))
	}
	if len(os.Args) < 3 {
		logger.Error("bad input arguments", zap.Any("args", os.Args))
		errorAndExit(fmt.Errorf("must provide at least two numbers to add"))
	}

	adder := adder.New(logger)
	result, err := adder.Add(os.Args[1:]...)
	if err != nil {
		errorAndExit(err)
	}

	fmt.Println("Answer: ", result)
}

func errorAndExit(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
