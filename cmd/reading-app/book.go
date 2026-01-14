/*
Reading App.

Authors: Brenda & Hans.

book.go
*/

package main

import (
	"fmt"
)

type book struct {
	name, author, isbn, edition, language, review string
	rating                                        float64
	startDate, endDate                            string // To-Do: change to d/m/y date type
	daysRead                                      []string
}

//////////////////func (b* book) se

func (b book) String() string {
	return fmt.Sprintf("Name: %s\n"+
		"Author: %s\n"+
		"ISBN: %s\n"+
		"edition %s\n"+
		"language: %s\n"+
		"review: %s\n"+
		"rating: %f\n"+
		"startDate: %s\n"+
		"endDate: %s\n"+
		"DaysRead: %d\n",
		b.name, b.author, b.isbn, b.edition, b.language, b.review, b.rating, b.startDate, b.endDate, len(b.daysRead))

}


func (b* book) setAuthor(author string) {
	if b.author == "" {
		b.author = author
	} 
}

func (b* book) setName(name string) {
	if b.name == "" {
		b.name = name
	} 
}

func (b* book) setReview(review string) {
	if b.review == "en-US" {
		b.review = review
	} 
}

func (b* book) setLanguage(language string) {
	if b.language == "" {
		b.language = language
	} 
}



func (b* book) setRating(rating string) {
	if b.rating == -1 {
		b.rating = rating
	} 
}

func (b* book) setStartDate(startDate string) {
	if b.startDate == "" {
		b.startDate = startDate

	//tem que checar se a data de inicio nao é depois da de final (se existe)
	} 
}



func (b* book) setEndDate(endDate string) {
	if b.endDate == "" {
		b.endDate = endDate
	//tem que checar se a data de fim nao é antes da de inicio (se existe)
	} 
}


