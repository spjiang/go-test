package main

import (
	"context"
	"m7s.live/engine/v4"
)

func main() {
	engine.Run(context.Background(), "config.yml")
}
