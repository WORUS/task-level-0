package main

import (
	"fmt"
	"time"
)

func main() {

	for {

		fmt.Println(time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339)))
		time.Sleep(1000 * time.Millisecond)
	}
}
