package helper

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// NumberFormat — Format a number with grouped thousands, opts params:decimals, decPoint, thousandsSep
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

// Rand Rand()
// Range: [0, 2147483647]
func Rand(min, max int) int {
	if min > max {
		panic("min: min cannot be greater than max")
	}
	// PHP: getRandMax()
	if int31 := 1<<31 - 1; max > int31 {
		panic("max: max can not be greater than " + strconv.Itoa(int31))
	}
	if min == max {
		return min
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max+1-min) + min
}

// Round php round function
func Round(value float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Trunc((value+0.5/p)*p) / p
}
