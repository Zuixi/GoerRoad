package basic_test

import (
	"testing"
	"ToGoer/src/basic"
)

func TestJudgeNumber(t *testing.T) {
	res := basic.JudgeNumber(13)

	if res == true {
		t.Fatal("Error occur")
	}
}