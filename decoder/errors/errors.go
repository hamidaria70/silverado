package errors

import (
	"fmt"
	"os"
)

func ConnectionError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
