package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {

	for {

		fmt.Println(uuid.New)
		time.Sleep(1000 * time.Millisecond)
	}
}
