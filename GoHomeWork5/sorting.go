package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

//Person structure describes a person
type Person struct {
	firstName string
	lastName  string
	birthDay  time.Time
}

//People describes slice of type person
type People []Person

// Len is the number of elements in the collection.
func (p People) Len() int {
	return len(p)
}

// Swap swaps the elements with indexes i and j.
func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (p People) Less(i, j int) bool {
	if p[i].birthDay.Unix() == p[j].birthDay.Unix() {
		if p[i].firstName == p[j].firstName {
			return p[i].lastName < p[j].lastName
		}
		return p[i].firstName < p[j].firstName
	}
	return p[i].birthDay.Unix() > p[j].birthDay.Unix()

}

//GetBirthday takes a string with birthday and returns in time.Time format
func GetBirthday(s string) time.Time {
	res, err := time.Parse("2006-01-02", s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return res
}

func main() {
	p := People{
		{"Ivan", "Ivanov", GetBirthday("2005-08-10")},
		{"Ivan", "Ivanov", GetBirthday("2003-08-05")},
		{"Artiom", "Ivanov", GetBirthday("2005-08-10")},
	}
	sort.Sort(p)

	if p[0].firstName != "Artiom" || p[0].birthDay != GetBirthday("2005-08-10") {
		fmt.Print("result is incorrect")
	} else if p[1].firstName != "Ivan" || p[1].birthDay != GetBirthday("2005-08-10") {
		fmt.Print("result is incorrect")
	} else if p[2].firstName != "Ivan" || p[2].birthDay != GetBirthday("2003-08-05") {
		fmt.Print("result is incorrect")
	} else {
		fmt.Print("sorting was successful")
	}

}
