package main

import (
	"flag"
	"fmt"
)

var colorMap = map[string]string{
	"white": "FFFFFF",
	"red":   "FF0000",
	"green": "00FF00",
	"blue":  "0000FF",
}

type Value interface {
	String() string
	Set(string) error
}

type color struct {
	rgb string
}

func (c *color) String() string {
	return fmt.Sprintf("color: {rgb: #%s}", c.rgb)
}

type colorFlag struct {
	color
}

func (c *colorFlag) Set(s string) error {
	if code, ok := colorMap[s]; ok {
		c.rgb = code
		return nil
	}
	return fmt.Errorf("unknown color %s", s)
}

func ColorFlag(name string, value string, usage string) *color {
	c := &colorFlag{}

	if code, ok := colorMap[value]; !ok {
		code = colorMap["white"]
	} else {
		c.rgb = code
	}

	flag.CommandLine.Var(c, name, usage)
	return &c.color
}

func main() {
	color := ColorFlag("color", "red", "color")
	flag.Parse()
	fmt.Println(color)
}
