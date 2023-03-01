package gordle

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/slices"
	"io"
	"os"
	"strings"
)

type Game struct {
	reader   *bufio.Reader
	solution []rune
	// max before losing
	maxAttempts int
}


// New -> creates a new intance of Game.
// Entry point.
func New(playerInfo io.Reader, solution string, maxAttempts int) *Game {
	g := &Game{
		reader:      bufio.NewReader(playerInfo),
		solution:    SplitToUppercaseCharacters(solution),
		maxAttempts: maxAttempts,
	}
	return g
}



func (g *Game) Ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", len(g.solution))
	for {
		// reads users input as a string of bytes
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "Gordle failed to read your guess: %s\n, ",
				err.Error())
			continue
		}
		// conversts the string of bytes into a rune slice
		guess := []rune(string(playerInput))
		if len(guess) != len(g.solution) {
			_, _ = fmt.Fprintf(os.Stderr, "Sorry, your attempt is invalid with Gordle's solution!\n"+
				"Expected %d character, got %d characters.\n", len(g.solution), len(guess))
		} else {
			return guess
		}
	}
}

var ErrInvalidWordLength = fmt.Errorf("invalid guess, word doesn't have the same number of characters as the solution")

func (g *Game) ValidateGuess(guess []rune) error {
	if len(guess) != len(g.solution) {
		return fmt.Errorf("expected %d, got %d, %w", len(g.solution), len(guess), ErrInvalidWordLength)
	}
	return nil
}

func SplitToUppercaseCharacters(intput string) []rune {
	return []rune(strings.ToUpper(intput))
}

func (g *Game) Play() {
	fmt.Println("🔠 Welcome to Gordle 🔠")

	// ask the user for a word with loop
	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		guess := g.Ask()
		// stringify the runes
		fmt.Printf("Your guess was %s.\n You have %d more attempts.\n", string(guess), (g.maxAttempts - currentAttempt))
        // conspare the two slices  of runes
		if slices.Equal(guess, g.solution) {
            fmt.Printf("🎉 You won! 🎉 , you guessed it in %d tries: the word was %s.\n",
                currentAttempt, string(g.solution))
		}
	}
    // after exhausting the number of allowed attempts.
    fmt.Printf("You lost! The solution was %s .\n", string(g.solution))
}
