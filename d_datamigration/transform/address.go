package transform

import "syreclabs.com/go/faker"

func GetAddress(ID string) string {
	return faker.Address().StreetAddress()
}

func GetName(firstName, lastName string) string {
	return firstName + " " + lastName
}
