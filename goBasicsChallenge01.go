package main

import "fmt"

var name, city, state string = "Bruno", "Fishers", "Indiana"

func main() {
    fmt.Println("Name:", name)
    fmt.Println("Location:", city + ",", state)
}
