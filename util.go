package util

import (
	"fmt"

	_ "github.com/lib/pq" //Importing needed postgres driver
)

func version() {
	fmt.Println("Util package, v0.2")
	fmt.Println("Author: Simon Lindgren, email: simon.g.lindgren@gmail.com")
}
