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

func TestOrString(t *testing.T){
	res := basic.OrString()

	if res != "Elpoep olleh" {
		t.Log(res)
	}
}

func TestIsSameSignal(t *testing.T) {
	res := basic.IsSameSignal(12, -2)

	if res {
		t.Fatal("Error occur")
	} else {
		t.Log("Nice job")
	}
}

func TestReverseInt(t *testing.T) {
	res := basic.ReverseInt(0x11110000)

	if res == 0x00001111 {
		t.Log("success")
	} else {
		t.Log(res)
		t.Fatal("Error occer")
	}
}