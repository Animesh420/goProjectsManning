package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

var errInvalidWordLength = fmt.Errorf("invalid guess, word doesn't have the same number of characters as the solution.")

type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

func New(playerInput io.Reader, corpus []string, maxAttempts int) (*Game, error) {

	if len(corpus) == 0 {
		return nil, ErrCorpusIsEmpty
	}

	g := &Game{
		reader:      bufio.NewReader(playerInput),
		solution:    []rune(strings.ToUpper(pickWord(corpus))),
		maxAttempts: maxAttempts,
	}
	return g, nil
}

func NewRuneOnly(playerInput io.Reader, solution string, maxAttempts int) *Game {

	g := &Game{
		reader:      bufio.NewReader(playerInput),
		solution:    splitToUppercaseCharacters(solution),
		maxAttempts: maxAttempts,
	}
	return g
}

func (g *Game) Play() {
	fmt.Println("Welcome to Gordle !")

	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		// ask for a valid word
		guess := g.ask()
		fb := computeFeedback(guess, g.solution)
		fmt.Println(fb.String())

		if slices.Equal(guess, g.solution) {
			fmt.Printf("🎉 You won! You found it in %d guess(es)!  The word was: %s.\n", currentAttempt, string(g.solution))
			return
		} else {
			fmt.Printf("Incorrect solution %s in attempt %d\n", string(guess), currentAttempt)
		}

	}

	fmt.Printf("Your have lost, the solution was: %s\n", string(g.solution))

}

// split to uppercase characters is a naive implementation to turn a string into list of characters
func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}

func (g *Game) validateGuess(guess []rune) error {
	fmt.Println("Validating", len(guess), len(g.solution))
	if len(guess) != len(g.solution) {
		return fmt.Errorf("expected %d, got %d, %w", len(g.solution), len(guess), errInvalidWordLength)
	}

	return nil

}

func computeFeedback(guess, solution []rune) feedback {
	result := make(feedback, len(guess))
	used := make([]bool, len(solution))

	if len(guess) != len(solution) {
		_, _ = fmt.Fprintf(os.Stderr, "Internal error!, Guess and Solution"+" have different lengths: %d vs %d", len(guess), len(solution))
	}

	// check for correct letters
	for posInGuess, character := range guess {
		if character == solution[posInGuess] {
			result[posInGuess] = correctPosition
			used[posInGuess] = true
		}
	}

	// look for letters in the wrong position
	for posInGuess, character := range guess {
		if result[posInGuess] != absentCharacter {
			continue
		}

		for posInSolution, target := range solution {
			if used[posInSolution] {
				continue
			}

			if character == target {
				result[posInGuess] = wrongPosition
				used[posInSolution] = true
				break
			}

		}
	}

	return result
}

func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", len(g.solution))
	for {
		playerInput, _, err := g.reader.ReadLine()

		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}

		guess := splitToUppercaseCharacters(string(playerInput))

		if len(guess) != len(g.solution) {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's soltion! Expected %d characters, got %d.\n", len(g.solution), len(guess))
		} else {
			err := g.validateGuess(guess)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution %s.\n", err.Error())
			} else {
				return guess
			}

		}

	}
}
