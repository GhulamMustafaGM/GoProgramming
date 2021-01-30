package main

import (
	"fmt"
	"time"
)

func main() {
	
	for j:=10; j>0; j-- {
		fmt.Println(j)
		time.Sleep(time.Second)
	}
	fmt.Println("Happy New Year!")

}
