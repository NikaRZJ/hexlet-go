package main

import (
	"github.com/NikaRZJ/hexlet-go/greeting"
	"github.com/fatih/color"
)

func main() {
	color.Cyan(greeting.Hello())
	color.Blue(greeting.Hello())
	color.Red(greeting.Hello())
	color.Magenta(greeting.Hello())
}
