package main

import (
    "fmt"
    _"os" //remove _

    _ tea "github.com/charmbracelet/bubbletea" //remove _
)

func main(){
	fmt.Printf("This is a test %v\n", "Tyler")
}

type model struct {
	choices []string
	cursor int
	selected map[int]struct{}
}