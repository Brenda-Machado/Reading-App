/*
Reading App.

Authors: Brenda & Hans.

book.go
*/

package main //lalal =)

import (
	"fmt"
)

type book struct {
	name, author, isbn, edition, language, review string
	rating                                        float64
	startDate, endDate                          string // To-Do: change to d/m/y date type
	daysRead                                     []string
}

//////////////////func (b* book) se

func (b book) String() string {
	return fmt.Sprintf("Name: %s\n"+
			    Author: %s\n
		    	    ISBN: %s\n
		    	    edition %s\n
		    	    language: %s\n
		    	    review: %s\n
		    	    rating: %f\n
		    	    startDate: %s\n
		    	    endDate: %s\n
		    	    DaysRead: %d\n",
			    b.name, b.author, b.isbn, b.edition, b.language, b.review, b.rating, b.startDate, b.endDate, len(b.daysRead))
			    
}
