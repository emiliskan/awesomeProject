package main

import (
	"fmt"
	"strconv"
	"time"
)

func countAge(birthYear int) int {
	return time.Now().Year() - birthYear
}

type dog struct {
	name      string
	birthYear int
}

type person struct {
	lastname  string
	firstname string
	birthYear int
	dogs      []dog
}

func (d dog) updateName(newName string) {
	d.name = newName
}

func (p person) toString() string {
	age := countAge(p.birthYear)

	// dogs := strings.Join(p.dogs, "\n")
	dogs := ""
	return p.lastname + " " + p.firstname + " " + strconv.Itoa(age) + " " + dogs
}

func (d dog) toString() string {
	age := countAge(d.birthYear)
	return d.name + " at age " + strconv.Itoa(age)
}

func main() {

	tanj := dog{
		name:      "Tanjirs",
		birthYear: 2014,
	}
	tanj.updateName("Tanj")

	emil := person{
		firstname: "Iska",
		lastname:  "Emil",
		birthYear: 1996,
		dogs:      []dog{tanj},
	}

	fmt.Println(emil.toString())
	fmt.Printf("%+v", emil)

}
