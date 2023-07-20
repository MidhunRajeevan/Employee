package app

type Employee struct {
	FirstName string
	LastName  string
	EmpId     int
	Contact   ContactInfo
}

type ContactInfo struct {
	Email string
	Po    string
}

type person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

// func emp() {
// 	midhun := Employee{
// 		firstName: "midhun",
// 		lastName:  "r",
// 		empId:     2406,
// 		contact: contactInfo{
// 			"midhun@gmail.com",
// 			"2406",
// 		},
// 	}
// }
