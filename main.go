package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	b := []string{"Name1", "Name2", "Name3", "Name4"}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(b), func(i, j int) { b[i], b[j] = b[j], b[i] })
	fmt.Println("Prize 1 goes to", b[0])
	fmt.Println("Prize 2 goes to", b[1])
	fmt.Println("Prize 3 goes to", b[2])
	fmt.Println(`  _______
 |        |
(| Winner |)
 | Winner |
 | Chiken |
 | Dinner |
  \      /
   _|__|_`)

}
