package  gordle

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

type Game struct {
    reader *bufio.Reader
}

// New -> creates a new intance of Game.
    //Entry point.
func New(playerInfo io.Reader) *Game {
    g := &Game{
        reader: bufio.NewReader(playerInfo),
    }
    return g
}

const wordlength = 5
func (g *Game) ask() []rune  {
    fmt.Printf("Enter a %d-character guess:\n", wordlength)

    for  {
        playerInput, _, err := g.reader.ReadLine()
        if err != nil {
            _, _ = fmt.Fprintf(os.Stdout, "Gordle failed to read your guess: %s\n, ",
                err.Error())
        }
        guess :=[]rune(string(playerInput))
    }
}

func (g *Game) Play() {
    fmt.Println("Welcome to Gordle 😎")
    fmt.Printf("Enter a guess: \n")
}