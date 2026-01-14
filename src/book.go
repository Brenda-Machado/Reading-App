/*
Reading App.

Authors: Brenda & Hans.

book.go
*/

package book

type book struct {
	name, author, isbn, edition, language, review string
	rating                                        float64
	start_date, end_date                          string // To-Do: change to d/m/y date type
	days_read                                     []string
}
