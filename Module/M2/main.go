package main

import "fmt"

func main() {

	// var age int

	// fmt.Print("Masukkan Umur = ")
	// fmt.Scanf("%d", &age)

	// if age < 17 {
	// 	fmt.Println("You are under age!")
	// } else {
	// 	fmt.Println("You are eligible!")
	// }

	// Switch

	var trafficLamp string

	fmt.Print("Masukkan Kondisi Lampu Lalu Lintas = ")
	fmt.Scanf("%s", &trafficLamp)

	switch trafficLamp {
	case "Hijau":
		fmt.Print("Jalan")
	case "Kuning":
		fmt.Print("Hati-Hati")
	case "Merah":
		fmt.Print("Berhenti")
	default:
		fmt.Print("Error")
	}

}
