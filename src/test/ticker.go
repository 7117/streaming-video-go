package main

import (
	"fmt"
	"time"
)

func foo()  {
	fmt.Println("wilson...")
}

func main() {
	ticker := time.NewTicker(1 * time.Second)

	sum:=6;

	for i:=0;i<sum; i++ {
		select {
		case <-ticker.C:
			go foo()
		}
	}
}
