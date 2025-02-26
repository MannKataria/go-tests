package example

import (
	"errors"
)

func Add(a, b interface{}) (interface{}, error) {
	aInt, ok1 := a.(int)
	aFloat, ok2 := a.(float64)
	if !ok1 && !ok2 {
		return 0, errors.New("invalid type for a")
	}

	bInt, ok3 := b.(int)
	bFloat, ok4 := b.(float64)
	if !ok3 && !ok4 {
		return 0, errors.New("invalid type for b")
	}

	if ok1 {
		if ok3 {
			return aInt + bInt, nil
		} else {
			return float64(aInt) + bFloat, nil
		}
	} else {
		if ok3 {
			return aFloat + float64(bInt), nil
		} else {
			return aFloat + bFloat, nil
		}
	}
}

func Div(a, b interface{}) (float64, error) {
	aFloat, ok1 := a.(float64)
	aInt, ok2 := a.(int)
	if !ok1 && !ok2 {
		return 0, errors.New("invalid type for a")
	}

	bFloat, ok3 := b.(float64)
	bInt, ok4 := b.(int)
	if !ok3 && !ok4 {
		return 0, errors.New("invalid type for b")
	}
	if bFloat == 0 && bInt == 0 {
		return 0, errors.New("cannot divide by 0")
	}

	if ok1 {
		if ok3 {
			return aFloat / bFloat, nil
		} else {
			return aFloat / float64(bInt), nil
		}
	} else {
		if ok3 {
			return float64(aInt) / bFloat, nil
		} else {
			return float64(aInt) / float64(bInt), nil
		}
	}
}
