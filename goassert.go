package goassert

import (
	"log"
	"reflect"
)

var (
	debug_mode = true
)

// AssertEq: check equality
func AssertEq(val1, val2 interface{}) {
	// checker
	if !reflect.DeepEqual(val1, val2) {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		panic("Eq assertion fails.")
	}
}

// AssertNEq: check non-equality
func AssertNEq(val1, val2 interface{}) {
	// checker
	if reflect.DeepEqual(val1, val2) {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		panic("NEq assertion fails.")
	}
}

// AssertGt: check greater than
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
		log.Printf("loss=%f\n", loss)
	}
	return loss
}

// AssertGte: check greater or equal
func AssertGte(val1, val2 interface{}) float64 {
	// checker and loss
	loss := allToFloat64(val1) - allToFloat64(val2)
	if loss < 0 {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		panic("Gte assertion fails.")
	}
	if debug_mode {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		log.Printf("loss=%f\n", loss)
	}
	return loss
}

// AssertLt: check less than
func AssertLt(val1, val2 interface{}) float64 {
	// checker and loss
	loss := allToFloat64(val1) - allToFloat64(val2)
	if loss >= 0 {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		panic("Lt assertion fails.")
	}
	if debug_mode {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		log.Printf("loss=%f\n", loss)
	}
	return loss
}

// AssertLte: check less or equal
func AssertLte(val1, val2 interface{}) float64 {
	// checker and loss
	loss := allToFloat64(val1) - allToFloat64(val2)
	if loss > 0 {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		panic("Lte assertion fails.")
	}
	if debug_mode {
		log.Printf("val1=%v\n", val1)
		log.Printf("val2=%v\n", val2)
		log.Printf("loss=%f\n", loss)
	}
	return loss
}

// AssertEmpty: check empty
func AssertEmpty(val1 []interface{}) int {
	length := len(val1)
	if length > 0 {
		log.Printf("val1=%v\n", val1)
		panic("Empty assertion fails.")
	}
	return length
}

// AssertNEmpty: check none-empty
func AssertNEmpty(val1 []interface{}) int {
	length := len(val1)
	if length == 0 {
		log.Printf("val1=%v\n", val1)
		panic("NEmpty assertion fails.")
	}
	return length
}
