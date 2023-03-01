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
func (g *Game) Ask() []rune  {
    fmt.Printf("Enter a %d-character guess:\n", wordlength)
    for  {
        // reads users input as a string of bytes
        playerInput, _, err := g.reader.ReadLine()
        if err != nil {
            _, _ = fmt.Fprintf(os.Stdout, "Gordle failed to read your guess: %s\n, ",
                err.Error())
                continue
        }
        // conversts the string of bytes into a rune slice
        guess :=[]rune(string(playerInput))
        if len(guess) != wordlength {
            _, _ = fmt.Fprintf(os.Stderr, "Sorry, your attempt is invalid with Gordle's solution!\n" +
                "Expected %d character, got %d characters.\n", wordlength, len(guess))
        } else {
            return guess
        }
    }
}
var ErrInvalidWordLength = fmt.Errorf("invalid guess, word doesn't have the same number of characters as the solution")

func (g *Game) ValidateGuess(guess []rune) error {
    if len(guess) != wordlength {
        return fmt.Errorf("expected %d, got %d, %w", wordlength, len(guess), ErrInvalidWordLength )
    }
    return nil
}

func (g *Game) Play() {
    fmt.Println("🎉 Welcome to Gordle 🎉")

    // ask the user for a word
    guess := g.Ask()
                            // stringify the runes
    fmt.Printf("Your guess was %s \n", string(guess))
}