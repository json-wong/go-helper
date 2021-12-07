package helper

import "time"

func GetCSTLoc() *time.Location {
	return time.FixedZone("CST", 8*3600)
}
