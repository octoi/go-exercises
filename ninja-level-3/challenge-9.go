package main

import "fmt"

func main() {
	favSport := "FootBall"

	switch favSport {
		case "FootBall":
			fmt.Println("So you like balls don't you ?")
		case "Cricket":
			fmt.Println("Hit it hard")
		default:
			fmt.Println(favSport + "!,", "i never played that before, if you don't mind can you teach me")
	}
}
