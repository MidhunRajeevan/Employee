package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// cards := cr.NewDeck()
	// cards.Print()
	// selectCard := cards.PickOne()
	// fmt.Println("selected Card", selectCard)

	// midhun := app.Employee{
	// 	FirstName: "midhun",
	// 	LastName:  "r",
	// 	EmpId:     2406,
	// 	Contact: app.ContactInfo{
	// 		Email: "midhun@gmail.com",
	// 		Po:    "2406",
	// 	},
	// }
	// fmt.Println(midhun)
	// midhun.UpdateEmploye("abcd")
	// fmt.Println(midhun)

	// circle := Circle{Radius: 5}
	// rectagle := Rectangle{Hieght: 5, Width: 5}

	// shape := []Shape{circle, rectagle}

	// for _, shape := range shape {
	// 	fmt.Println(shape.Area())
	// }

	// GetGoogle()
	http.HandleFunc("/", Index)
	fmt.Println("application Runs at 9090")
	http.ListenAndServe(":9090", nil)

}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Hieght float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Hieght
}

type Shape interface {
	Area() float64
}

func GetGoogle() {
	res, err := http.Get("http://google.com")
	if err != nil {
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// res.Body.Read(bs)
	// fmt.Println(string(bs))
	body, err := io.ReadAll(res.Body)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(body))
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"app": "My First Service on 21-07-2023"}`))
}
