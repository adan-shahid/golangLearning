package main

import (
	"fmt"
	"math/rand"
)

// func main() {
// 	fmt.Println("Hi There!")
// }

// HEXADECIMAL REPRESENTATION OF A NUMBER

func main(){
	// number := 2586
	// fmt.Printf("%x\n", number)
	// fmt.Printf("Decimal : %d", number)
	hotel_name := "The Gopher Hotel"
	fmt.Println("Hotel Name: "+hotel_name)
	var roomsAvailable int
	var rooms, roomsOccupied = 100, rand.Intn(100)
	roomsAvailable = rooms - roomsOccupied
	fmt.Println(roomsAvailable, "Rooms Available")
	age_of_person()

}

func age_of_person(){
	var agePaul, ageJohn int = rand.Intn(110), rand.Intn(110)
	fmt.Println("John is", ageJohn , "years old")
	fmt.Println("Paul is", agePaul , "years old")

	if agePaul > ageJohn {
		fmt.Println("Paul is older than John")
	} else {
		fmt.Println("Paul is younger than John, or both have the same age")
	}
}

// TO PRINT A NUMBER IN BASE 10