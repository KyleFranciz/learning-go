package forloops

import "fmt"

func doWhile(name string) {
	for name == "kyle" {
		fmt.Println("The user's name is kyle")
		break
	}
}
