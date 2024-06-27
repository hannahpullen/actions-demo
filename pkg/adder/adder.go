// package adder defines the Adder struct that can add two numbers
package adder

import (
	"strconv"

	"go.uber.org/zap"
)

type Adder struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *Adder {
	return &Adder{
		logger: logger,
	}
}

// Add takes a list of strings, converts them to floats, and adds them.
// Returns an error if any string cannot be converted to a float.
func (a *Adder) Add(numsToAdd ...string) (float64, error) {
	toAdd := []float64{}
	for _, numStr := range numsToAdd {
		num, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			a.logger.Error("error parsing number", zap.String("arg", numStr), zap.Error(err))
			return 0, err
		}
		toAdd = append(toAdd, num)
	}
	a.logger.Info("Adding numbers", zap.Any("numbers", toAdd))

	total := 0.0
	for _, n := range toAdd {
		total += n
	}
	a.logger.Info("Got total", zap.Float64("total", total))
	return total, nil
}
