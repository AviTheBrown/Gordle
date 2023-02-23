package  gordle

import "fmt"

type Game struct {}

// New -> creates a new intance of Game.
    //Entry point.
func New() *Game {
    g := &Game{}
    return g
}

func (g *Game) Play() {
    fmt.Println("Welcome to Gordle 😎")
    fmt.Printf("Enter a guess: \n")
}