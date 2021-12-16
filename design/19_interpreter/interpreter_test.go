package interpreter

import (
	"fmt"
	"testing"
)

func TestInterpreter(t *testing.T) {
	p := &Parser{}
	p.Parse("1 + 2 + 3")
	res := p.Result().Interpret()
	expect := 6
	if res != expect {
		t.Fatalf("expect %d got %d", expect, res)
	}
	fmt.Println(res)
}
