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
	email_send()
	assignment()

}

func age_of_person(){
	var agePaul, ageJohn int = rand.Intn(110), rand.Intn(110)
	fmt.Println("John is", ageJohn , "years old")
	fmt.Println("Paul is", agePaul , "years old")

	if agePaul > ageJohn {
		fmt.Println("Paul is older than John")}
	fmt.Println("End of the Program")	
	// 	else {
	// 	// ANOTHER IF/ELSE STRUCTURE
	// 	if agePaul == ageJohn {
	// 		fmt.Println("Paul and John have the same age.")
	// 	} else {
	// 		fmt.Println("Paul is younger than John")
	// 	}
	// }
}

func email_send(){
	const emailToSend = 4
	emailSent := 0
	for emailSent < emailToSend {
		fmt.Println("Sending email...")
		emailSent++
	}
	fmt.Println("end of program")
}

func assignment(){
	const no_of_rooms = 134
	var roomsAvailable, roomsBooked int
	roomsBooked = rand.Intn(110) 
	roomsAvailable = no_of_rooms - roomsBooked

	if roomsAvailable == 0{
		fmt.Println("No rooms available for tonight")
	} else if roomsAvailable == 3 {
		fmt.Println("- 110 : 2 people / 3 nights \n, - 111 : 4 people / 1 nights \n, - 112 : 6 people / 12 nights \n")
	}

	var occupancy_rate = (roomsBooked / no_of_rooms) * 100
	if occupancy_rate <= 30 {
		fmt.Println("Low")
	} else if occupancy_rate > 30 && occupancy_rate < 60 {
		fmt.Println("Medium")
	} else {
		fmt.Println("High")
	}
}