package main

import (
	"fmt"
	"reflect"
)

// a struct gathers information toguether that is related.
// the fields of a struct can have different types of data
// copiyng structs variables creates independent data from the original,
// unless used with &structVariable

// naming conventions:
// upper case first letter -> exported from package
// lower case first letter -> scoped to package;
// this includes the fields of the struct itself

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

// there is no inheritance, but structs can be embedded in other structs
// ChiefDoctor has Doctor
type ChiefDoctor struct {
	Doctor
	yearsOfService int
}

// struct with tags
type Animal struct {
	Name string `required:"true" max:"100"`
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Salchichon",
		companions: []string{
			"Rambo",
			"Po",
		},
	}

	fmt.Println(aDoctor)

	// pulling field from a struct
	fmt.Println(aDoctor.actorName)

	// anonymous struct
	otherDoc := struct{ name string }{name: "Dam Daniel"}
	fmt.Println(otherDoc)

	// embedded structs
	chief := ChiefDoctor{}
	chief.number = 5
	chief.actorName = "Maneman"
	chief.yearsOfService = 10

	fmt.Println(chief)

	// struct with tag
	t := reflect.TypeOf((Animal{}))
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
