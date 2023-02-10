package main

import "fmt"

func add(n1 int, n2 int) {
	var hasil int  = n1 + n2
	fmt.Printf("Hasil Tambah %d dengan %d adalah = %d ", n1,n2,hasil)
}

func main() {
	add(1,2)
}