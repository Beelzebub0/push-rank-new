package main

import (
	"fmt"
	"time"
)

func main() {
	in := "11:06 PM"
	lay := "03:04 PM"

	// x, _ := parserP(lay, in)

	fmt.Println(parserP(lay, in))

}

func parserP(layout, input string) (time.Time, error) {
	p, err := time.Parse(layout, input)
	if err != nil {
		return p, err
	}
	return p, nil
}
