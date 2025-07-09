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
	hotel_name := "The Gopher Hotel"
	
	const no_of_rooms = 134
	const firstRoomNumber = 110
	var roomsAvailable, roomsBooked int
	roomsBooked = rand.Intn(110) 
	roomsAvailable = no_of_rooms - roomsBooked

	occupancyRate := occupancyRate(roomsBooked, no_of_rooms)
	occupancyLevel := occupancyLevel(occupancyRate)

	fmt.Println("Hotel Name: ", hotel_name)
	fmt.Println("Total Rooms ",no_of_rooms)
	fmt.Println("Rooms Available", roomsAvailable)
	fmt.Println("                   Occupancy Level", occupancyLevel)
	fmt.Println("                   Occupancy Rate", occupancyRate)

	if roomsAvailable > 0 {
		fmt.Println("Rooms:")
		for i := 0; roomsAvailable > i; i++ {
			roomNumber := firstRoomNumber + i
			size := rand.Intn(6) + 1
			nights := rand.Intn(10) + 1
			printRoomDetails(roomNumber, size, nights)
		}
	} else {
		fmt.Println("No rooms available for tonight")
	}

	// number := 2586
	// fmt.Printf("%x\n", number)
	// fmt.Printf("Decimal : %d", number)
	
	// var roomsAvailable int
	// var rooms, roomsOccupied = 100, rand.Intn(100)
	// roomsAvailable = rooms - roomsOccupied
	// fmt.Println(roomsAvailable, "Rooms Available")
	// age_of_person()
	// email_send()
	// assignment()
	// johnPrice := computePrice(148.0, 2)
	// fmt.Println("John:", johnPrice)
}

func printRoomDetails(roomNumber, size, nights int){
	fmt.Println(roomNumber, ":", size, "People /", nights, "nights")
}

func occupancyLevel (occupancyRate float32) string {
	if occupancyRate > 70 {
		return "High"
	} else if occupancyRate > 20 {
		return "Medium"
	} else {
		return "Low"
	}

}

func occupancyRate (roomsOccupied int, no_of_rooms int) float32{
	return (float32(roomsOccupied) / float32(no_of_rooms)) * 100
}

// func age_of_person(){
// 	var agePaul, ageJohn int = rand.Intn(110), rand.Intn(110)
// 	fmt.Println("John is", ageJohn , "years old")
// 	fmt.Println("Paul is", agePaul , "years old")

// 	if agePaul > ageJohn {
// 		fmt.Println("Paul is older than John")}
// 	fmt.Println("End of the Program")	
	// 	else {
	// 	// ANOTHER IF/ELSE STRUCTURE
	// 	if agePaul == ageJohn {
	// 		fmt.Println("Paul and John have the same age.")
	// 	} else {
	// 		fmt.Println("Paul is younger than John")
	// 	}
	// }
// }

// func email_send(){
// 	const emailToSend = 4
// 	emailSent := 0
// 	for emailSent < emailToSend {
// 		fmt.Println("Sending email...")
// 		emailSent++
// 	}
// 	fmt.Println("end of program")
// }



// func computePrice(rate float32, nights int) float32 {
// 	return rate * float32(nights)
// }