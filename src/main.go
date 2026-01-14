/*
Reading App.

Authors: Brenda & Hans.

main.go
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	inputData := bufio.NewReader(os.Stdin)
	fmt.Println("New book record: ")
	input, err := inputData.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	fileName := "lib.json"
	writer, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	_, err = writer.WriteString(input)
	if err != nil {
		log.Fatal(err)
	}

	// Read contents of the file
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Contents:", string(data))

	b := book{"the stand", "stephen king", "9780385121682", "1", "en-us", "", 9.5, "1-1-2026", "30-1-2026", []string{"1-1-2026", "2-1-2026"}}
	fmt.Println(b.String())
}
