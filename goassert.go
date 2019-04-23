package goassert

import (
	"log"
	"reflect"
)

var (
	debug_mode = true
)

/*
	AssertEq: check equality of variables
	val1 and val2: number, string, array, pointer and struct
*/
func AssertEq(val1, val2 interface{}) {
	// checker
	if !reflect.DeepEqual(val1, val2) {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		panic("Eq assertion fails.")
	}
}

func AssertGt(val1, val2 interface{}) float64 {
	// checker and loss
	loss := allToFloat64(val1) - allToFloat64(val2)
	if loss <= 0 {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		panic("Gt assertion fails.")
	}
	if debug_mode {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		log.Printf("loss=%v\n", loss)
	}
	return loss
}
