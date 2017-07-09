package util

import (
	"fmt"
	"os"

	_ "github.com/lib/pq" //Importing needed postgres driver
)

func version() {
	fmt.Println("Util package, v0.2")
	fmt.Println("Author: Simon Lindgren, email: simon.g.lindgren@gmail.com")
}

//FileExists checks if a file identified by a supplied path exist
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
