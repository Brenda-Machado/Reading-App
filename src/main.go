/*
Reading App.

Authors: Brenda & Hans.

main.go
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	new_book := readBookFromStdin()

	if err := saveBook(new_book, "lib.json"); err != nil {
		log.Fatalf("Failed to save book: %v", err)
	}

	fmt.Println("\nBook record added:")
}

func readBookFromStdin() *book {
	reader := bufio.NewReader(os.Stdin)

	b := &book{}

	fmt.Print("Name: ")
	b.name = readLine(reader)

	fmt.Print("Author: ")
	b.author = readLine(reader)

	fmt.Print("ISBN: ")
	b.isbn = readLine(reader)

	fmt.Print("Edition: ")
	b.edition = readLine(reader)

	fmt.Print("Language: ")
	b.language = readLine(reader)

	fmt.Print("Review (optional): ")
	b.review = readLine(reader)

	fmt.Print("Rating (0-10): ")
	ratingStr := readLine(reader)
	if ratingStr != "" {
		fmt.Sscanf(ratingStr, "%f", &b.rating)
	}

	fmt.Print("Start Date (DD-MM-YYYY): ")
	b.startDate = readLine(reader)

	fmt.Print("End Date (DD-MM-YYYY): ")
	b.endDate = readLine(reader)

	fmt.Print("Days Read (DD-MM-YYYY): ")
	b.daysRead = []string{}

	return b
}

func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')

	return strings.TrimSpace(line)
}

func saveBook(book *book, fileName string) error {

	fmt.Println(book.String())
	jsonData, err := json.Marshal(book)

	if err != nil {
		return fmt.Errorf("failed to marshal book: %w", err)
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	_, err = file.WriteString(string(jsonData) + "\n")

	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
