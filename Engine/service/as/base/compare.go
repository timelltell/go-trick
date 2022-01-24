package basic_condition

import "GolangTrick/Engine/constant"

//return (number1 expr number2)
func compareInt(number1, number2 int, expr string) bool {
	result := false
	switch expr {
	case constant.COND_EQUAL:
		result = (number1 == number2)
	case constant.COND_LARGER:
		result = (number1 > number2)
	case constant.COND_LESS:
		result = (number1 < number2)
	case constant.COND_LARGER_EQUAL:
		result = (number1 >= number2)
	case constant.COND_LESS_EQUAL:
		result = (number1 <= number2)
	case constant.COND_NOT_EQUAL:
		result = (number1 != number2)
	}
	return result
}
