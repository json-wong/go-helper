package helper

import "testing"

func TestIsNumeric(t *testing.T) {
	t.Log(IsNumeric("2342342.45.234"))
}

func TestNumberFormat(t *testing.T) {
	t.Log(NumberFormat(34234.232434323, 2))
}