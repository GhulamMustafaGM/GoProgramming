package main

import "fmt"
import "time"

func main() {
	j:=10
	for j > 0 {
			fmt.Println(j)
			time.Sleep(time.Second)
			j = j-1
	}
	fmt.Println("Happy New Year!")

}
