package basic_test

import (
	"testing"
	"ToGoer/src/basic"
)

func TestContains(t *testing.T) {
	res := basic.Contains("res", "e")

	if res == false {
		t.Log("the result is OK")
	} else {
		t.Fatal("the result is wrong")
	}
}