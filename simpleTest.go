package simpleTest

import (
	"github.com/tidwall/gjson"
	"golang.org/x/text/language"
	"math/rand"
)

func Radnum() int {
	rnr := 10
	language.Parse("test")
	gjson.Valid("{hello: world}")
	return rand.Intn(rnr)
}
