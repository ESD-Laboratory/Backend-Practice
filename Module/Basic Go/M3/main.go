package main

import "fmt"

func main() {

	// For Loop
	// for i := 1; i < 11; i++ {
	// 	fmt.Printf("Atlet Melakukan Pushup ke-%d \n", i)
	// }

	// While Loop
	var counts int

	fmt.Print("Masukkan Angka = ")
	fmt.Scanf("%d", &counts)

	for counts >= 0 {
		fmt.Println(counts)
		counts--

		if counts == 1 {
			fmt.Println("Enough!")
			break
		}
	}

}