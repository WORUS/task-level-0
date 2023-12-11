package main

import "fmt"

func main() { mps := make(map[string]string); mps["1"] = "1"; mps["1"] = "2"; fmt.Print(mps["1"]) }
