package main

import "fmt"

func main() {
	var sum int = 0

	for i:=0; i < 10; i++ {
		if i%2==0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

	// while loop alternative in golang
	i:=0
	for i < 10 {
		if i%2==0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
		i++
	}

	// infinity loop
	j:=0
	for {
		if j >= 10 {
			break
		}

		if j%2==0 {
			fmt.Println(j, "even")
		} else {
			fmt.Println(j, "odd")
		}
		j++
	}
}
