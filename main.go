package main

import (
	"fmt"

	app "github.com/MidhunRajeevan/Employee/app"
)

func main() {
	// cards := cr.NewDeck()
	// cards.Print()
	// selectCard := cards.PickOne()
	// fmt.Println("selected Card", selectCard)

	midhun := app.Employee{
		FirstName: "midhun",
		LastName:  "r",
		EmpId:     2406,
		Contact: app.ContactInfo{
			Email: "midhun@gmail.com",
			Po:    "2406",
		},
	}
	fmt.Println(midhun)
}
