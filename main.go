package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inFile, err := os.Open("/tmp/dm.prof")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer inFile.Close()

		scanner := bufio.NewScanner(inFile)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			msg := parseLine(scanner.Text())
			fmt.Printf("%#v", msg)
		}
	}
}
