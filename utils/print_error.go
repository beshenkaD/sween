package utils

import (
	"fmt"
	. "github.com/logrusorgru/aurora"
)

func PrintError(err error) {
	if err != nil {
		fmt.Println(Red("ERROR:"), err)
	}
}
