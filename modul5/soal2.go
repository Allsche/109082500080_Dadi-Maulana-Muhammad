package main

import "fmt"

func printBintang(n int, nilai int) {
	if nilai > n {
		return
	}
	for i := 0; i < nilai; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	printBintang(n, nilai+1)
}

func main() {
	var n int
	fmt.Scan(&n)
	printBintang(n, 1)
}
