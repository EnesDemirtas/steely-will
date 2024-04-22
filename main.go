package main

import (
	"fmt"

	"github.com/EnesDemirtas/steely-will/ui"
)

func main() {
	if err := ui.Run(); err != nil {
		fmt.Println(err)
	}
}