package main

import "fmt"

func main() {
	num := 6
	if num%2 == 0 {
		fmt.Println("by 2")
	} else {
		fmt.Println("else")
	}

	for i := 0; i < 6; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}
		if i > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}

	l := []string{"python", "java", "go"}

	for _, v := range l {
		fmt.Print(v + ` `)
	}

	os := "window"

	switch os {
	case "mac":
		fmt.Println("Mac!")
	case "window":
		fmt.Println("Window!")
	default:
		fmt.Println("default...")
	}

}
