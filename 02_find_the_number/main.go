package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var number int
var counter int = 0

func main() {
	rand.Seed(time.Now().Unix())
	number = rand.Intn(100)
	askNumber()
}

func askNumber() {
	var i int = 0
	fmt.Print("Dévinez le nombre : ")
	_, err := fmt.Scanf("%d", &i)

	if err == nil {
		checkResult(i)
	} else {
		fmt.Println(err)
		os.Exit(2)
	}
}

func checkResult(answer int) {
	if answer == number {
		fmt.Printf("!!!!! Bravo Vous avez trouvé en %v essaie (s) !!!!!!\n", counter)

	} else if answer < number {
		counter++
		fmt.Println("plus grand !")
		askNumber()
	} else if answer > number {
		fmt.Println("plus petit !")
		counter++
		askNumber()
	}
}
