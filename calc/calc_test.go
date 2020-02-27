package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	expressions := []string {
		"2+2+6",
		"2*2*3",
		"6/3/2",
		"6-4-2",
		"55+44+1",
		"124-24",
		"2+2*2",
		"2+6/3",
		"3*(5+3)/8",
		"3+4*5-(1*2)",
	}
	answers := []string {
		"10",
		"12",
		"1",
		"0",
		"100",
		"100",
		"6",
		"4",
		"3",
		"21",
	}
	for i,_ := range expressions {
		received := calculate(expressions[i])
		if received != answers[i] {
			t.Error("Expected ", answers[i], "got ", received[i])
		}
	}


}
