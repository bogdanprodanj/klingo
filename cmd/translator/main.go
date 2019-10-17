package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/gookit/color"
	"github.com/klingo/app"
)

var (
	errInvalidInput = errors.New("input can not be empty")
	space           = regexp.MustCompile(`\s+`)
)

func main() {
	input := strings.Join(os.Args[1:], " ")
	if input == "" {
		exit(errInvalidInput)
	}
	input = space.ReplaceAllString(input, " ")
	translator := app.New()
	err := translator.Validate(input)
	if err != nil {
		exit(err)
	}
	species, err := translator.FindSpecies(input)
	if err != nil {
		exit(err)
	}
	if len(species) == 0 {
		color.Blue.Println(translator.Translate(input))
		color.Red.Println(app.UnknownSpecies)
	} else {
		for fullName, speciesName := range species {
			if err = translator.Validate(input); err != nil {
				// some of the full names from stapi.co are not translatable, ignoring
				continue
			}
			color.Blue.Println(translator.Translate(fullName))
			color.Green.Println(fullName)
			for k := range speciesName {
				if k == app.UnknownSpecies {
					color.Red.Printf("%s ", k)
				} else {
					color.Cyan.Printf("%s ", k)
				}
			}
			fmt.Println()
		}
	}
}

func exit(err error) {
	fmt.Printf("%v", err)
	os.Exit(1)
}
