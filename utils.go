package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// refer https://gist.github.com/r0l1/3dcbb0c8f6cfe9c66ab8008f55f8f28b
func confirm(s string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s[y/n]: ", s)
		res, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		res = strings.ToLower(strings.TrimSpace(res))

		if res == "y" || res == "yes" {
			return true
		}
		// if res == no
		return false
	}
}

// this can't use for multibyte string
func getLastChar(s string) byte {
	return s[len(s)-1]
}

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
