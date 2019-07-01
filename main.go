package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	b := []string{"Gabriel reis", "Caroline Santos", "Sergio Panaka", "Anderson", "Claudio Furini", "Gabriel Casteglioni", "Bruno Casteluci", "Katarine Leal", "Edson Junior", "Rafael Dias", "Lucio Charallo", "Flavio Muniz", "Rodrigo Moraes", "Yuri Levenhagen"}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(b), func(i, j int) { b[i], b[j] = b[j], b[i] })
	fmt.Println("The t-shirt goes to", b[1])
	fmt.Println("The agenda 1 goes to", b[2])
	fmt.Println("The agenda 2 goes to", b[3])
	fmt.Println("The webcam cover goes to", b[4])
	fmt.Println(`  _______
 |        |
(| Pipefy |)
 |        |
  \      /
   _|__|_`)

}
