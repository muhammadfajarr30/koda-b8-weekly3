package utils

import (
	"bufio"
	"fmt"
	"os"
)

func Input(opt string) string {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(opt)

	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}
	return scanner.Text()
}
