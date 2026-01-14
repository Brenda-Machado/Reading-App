package main

import "fmt"

func main() {

	b := book.book{"the stand", "stephen king", "9780385121682", "1", "en-us", "", 9.5, "1-1-2026", "30-1-2026", string{"1-1-2026", "2-1-2026"}}
	fmt.Println(b.String()) 
}

