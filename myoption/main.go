package main

import (
	"fmt"
	"github.com/spjiang/go-test/myoption/entity"
)

func main() {
	message2 := entity.NewByOption(entity.WithID(2), entity.WithName("message2"), entity.WithAddress("cache2"), entity.WithPhone(456))
	message3 := entity.NewRequireIDByOption(3, entity.WithAddress("cache3"), entity.WithPhone(789), entity.WithName("message3"))
	fmt.Println(message2)
	fmt.Println(message3)
}
