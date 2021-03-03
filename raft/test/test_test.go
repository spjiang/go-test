package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandRange(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var max, min int64
	max = 300
	min = 150
	r := rand.Int63n(max-min) + min
	fmt.Println(r)
}
