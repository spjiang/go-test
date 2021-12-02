package Builder

import (
	"fmt"
	"testing"
)

func TestConcreteBuilder_GetResult(t *testing.T) {
	builder := NewConcreteBuilder()
	director := NewBuilder(&builder)
	director.Construct()
	result := builder.GetResult()
	fmt.Println(result.Built)

}

