package helper

import (
	"fmt"
	"strconv"
)

// IsNumeric - Finds whether a variable is a number or a numeric string
func IsNumeric(x interface{}) (result bool) {
	//Figure out result
	switch x.(type) {

	case int, uint:
		result = true
	case int8, uint8:
		result = true
	case int16, uint16:
		result = true
	case int32, uint32:
		result = true
	case int64, uint64:
		result = true

	case float32, float64:
		result = true

	case complex64, complex128:
		result = true

	case string:
		if xAsString, ok := x.(string); ok {
			result = isStringNumeric(xAsString)
		} else {
			result = false
		}

	default:
		result = false

	}

	return result
}

func isStringNumeric(x string) bool {

	hasPeriod := false
	for i, c := range x {
		switch c {

		case '-':
			if i != 0 {
				return false
			}

		case '.':
			if hasPeriod {
				return false
			}
			hasPeriod = true

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			//Nothing here.

		default:
			return false

		}
	}

	return true
}

// NumberFormat — Format a number with grouped thousands
func NumberFormat(number float64, opts ...interface{}) string {
	decimals := 2
	decPoint := "."
	thousandsSep := ","

	switch len(opts) {
	case 1:
		decimals = opts[0].(int)
	case 2:
		decPoint = opts[1].(string)
	case 3:
		thousandsSep = opts[2].(string)
	}

	neg := false
	if number < 0 {
		number = -number
		neg = true
	}
	dec := int(decimals)
	// Will round off
	str := fmt.Sprintf("%."+strconv.Itoa(dec)+"F", number)
	prefix, suffix := "", ""
	if dec > 0 {
		prefix = str[:len(str)-(dec+1)]
		suffix = str[len(str)-dec:]
	} else {
		prefix = str
	}
	sep := []byte(thousandsSep)
	n, l1, l2 := 0, len(prefix), len(sep)
	// thousands sep num
	c := (l1 - 1) / 3
	tmp := make([]byte, l2*c+l1)
	pos := len(tmp) - 1
	for i := l1 - 1; i >= 0; i, n, pos = i-1, n+1, pos-1 {
		if l2 > 0 && n > 0 && n%3 == 0 {
			for j := range sep {
				tmp[pos] = sep[l2-j-1]
				pos--
			}
		}
		tmp[pos] = prefix[i]
	}
	s := string(tmp)
	if dec > 0 {
		s += decPoint + suffix
	}
	if neg {
		s = "-" + s
	}
	return s
}
