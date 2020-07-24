package basic_test

import (
	"testing"
	"ToGoer/src/basic"
)

func TestContains(t *testing.T) {
	res := basic.Contains("res", "e")

	if res == true {
		t.Log("the result is OK")
	} else {
		t.Fatal("the result is wrong")
	}
}

func TestAdd(t *testing.T) {
	
	// define table
	var addtests = []struct {
		a int
		b int
		res int    // expected result
	}{
		{1,1,2},
		{2,3,5},
		{3,4,7},
		{5,6,11},
		{6,5,11},
		{1,22,23},
	}

	for _, sum := range addtests {
		actual := basic.Add(sum.a, sum.b)

		if actual != sum.res {
			t.Errorf("Add(%d + %d ) = %d, expected result is %d", sum.a, sum.b, actual ,sum.res)
		}
	}
}

func TestEqualFold(t *testing.T) {
	res := basic.EqualFold("Go", "go")

	if res == true {
		t.Log("EqualFold Success")
	} else {
		t.Fatal("EqualFold failed")
	}
}

func TestHasPrefix(t *testing.T) {
	res := basic.HasPrefix("Gopher", "Go")
	
	if res == true {
		t.Log("EqualFold Success")
	} else {
		t.Fatal("EqualFold failed")
	}
}

func TestHasSuffix(t *testing.T) {
	res := basic.HasSuffix("Gopher", "re")

	if res == true {
		t.Fatal("Has Suffix failed")
	}
}

func TestIndex(t *testing.T) {
	var tablestring = []struct{
		s string
		substr string
		index int
	}{
		{"chicken", "ken", 4},
		{"chicken", "dmr", -1},
	}

	for _, res := range tablestring {
		actual := basic.Index(res.s, res.substr)

		if actual != res.index {
			t.Fatal("Index failed")
		}
	}
}