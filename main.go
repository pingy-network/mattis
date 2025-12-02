package main

import (
	"github.com/pingy-network/mattis/cmd"
	"github.com/xh3b4sd/tracer"
)

func main() {
	err := cmd.New().Execute()
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}
}
