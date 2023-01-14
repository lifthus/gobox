package tableTest

import (
	"errors"
	"fmt"
)

func DoMath(num1, num2 int, op string) (int, error) {
	// to test this kind of things, the test code get too complicated. so in this case
	switch op {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 + num2, nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("division by zero")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("unknown operator %s", op)
	}
}
