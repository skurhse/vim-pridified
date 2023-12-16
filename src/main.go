package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/beevik/etree"
)

const (
	LightBlue = "#5BCEFA"
	LightPink = "#F5A9B8"
	White     = "#FFFFFF"
)

var transColors = [3]string{LightBlue, LightPink, White}

const numArgs = 3

func main() {
	inp, out := parseArgs()

	doc := etree.NewDocument()

	if err := doc.ReadFromFile(inp); err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}

	var colors map[string]string

	paths := doc.FindElements("//path")

	for _, path := range paths {
		style := path.SelectAttr("style")

		if style != nil {
			props := strings.Split(style.Value, ";")

			for _, prop := range props {
				entry := strings.Split(strings.TrimSpace(prop), ":")
				key, val := entry[0], entry[1]

				if key == "fill" {
					fmt.Println(val)
				}
			}
		}
		originalColor := path.SelectAttrValue("fill", "")

		var newColor string
		var ok bool
		if newColor, ok = colors[originalColor]; ok {
		} else {
			newColor = transColors[rand.Intn(3)]
		}
		path.CreateAttr("fill", newColor)
	}

	if err := doc.WriteToFile(out); err != nil {
		fmt.Fprintf(os.Stderr, "error writing output: %v\n", err)
		os.Exit(1)
	}
}

func parseArgs() (string, string) {
	if lenArgs := len(os.Args); lenArgs != numArgs {
		fmt.Fprintf(os.Stderr, "expected %d args, got %d\n", numArgs, lenArgs)
		os.Exit(1)
	}

	return os.Args[1], os.Args[2]
}
