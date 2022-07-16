package main

import "fmt"

func main() {
	var a, p1, p2, z1, z2, y1, y2 byte
	var massive [10]byte
	for i := 0; i < len(massive); i++ {
		fmt.Scan(&a)
		massive[i] = a
	}
	fmt.Println(massive)
	fmt.Scan(&p1)
	fmt.Scan(&p2)
	fmt.Scan(&z1)
	fmt.Scan(&z2)
	fmt.Scan(&y1)
	fmt.Scan(&y2)
	massive[p1], massive[p2] = massive[p2], massive[p1]
	massive[z1], massive[z2] = massive[z2], massive[z1]
	massive[y1], massive[y2] = massive[y2], massive[y1]
	for j := 0; j < len(massive); j++ {
		fmt.Print(massive[j], " ")
	}
	fmt.Println("type ok")
}
