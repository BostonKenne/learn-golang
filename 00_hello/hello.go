// this code will just say hello and print a random number
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	x := rand.Intn(100)
	fmt.Println("Hello my favorite number is : ", x)
}
