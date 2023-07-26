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

func (e *Employee) UpdateEmploye(firstName string) {
	e.FirstName = firstName
}
