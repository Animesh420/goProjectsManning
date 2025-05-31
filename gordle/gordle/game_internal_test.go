package gordle

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestGameAsk(t *testing.T) {
	tt := map[string]struct {
		input string
		want  []rune
	}{
		"5 characters in english": {
			input: "HELLO",
			want:  []rune("HELLO"),
		},
		"5 characters in arabic": {

			input: "مرحبا",

			want: []rune("مرحبا"),
		},

		"5 characters in japanese": {

			input: "こんにちは",

			want: []rune("こんにちは"),
		},

		"3 characters in japanese": {

			input: "こんに\nこんにちは",

			want: []rune("こんに"),
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := NewRuneOnly(strings.NewReader(tc.input), string(tc.want), 0)

			got := g.ask()

			if !slices.Equal(got, tc.want) {
				t.Errorf("got = %v, want %v", string(got), string(tc.want))
			}
		})
	}
}

func TestGameValidateGuess(t *testing.T) {
	tt := map[string]struct {
		word     []rune
		expected error
	}{
		"matches length": {
			word:     []rune("guess"),
			expected: nil,
		},
		"too long": {
			word:     []rune("Pocket"),
			expected: errInvalidWordLength,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g, err1 := New(strings.NewReader("apple"), []string{"apple"}, 0)
			if err1 != nil {
				t.Errorf("Setup error in test validate guess")
			}

			err := g.validateGuess(tc.word)
			if !errors.Is(err, tc.expected) {
				t.Errorf("%c, expected %q, got %q", tc.word, tc.expected, err)
			}
		})
	}
}

func TestComputeFeedback(t *testing.T) {
	tt := map[string]struct {
		guess            string
		solution         string
		expectedFeedback feedback
	}{
		"nominal": {guess: "hello", solution: "hello", expectedFeedback: feedback{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition}},
		"two identical, but not in the right position": {guess: "hlleo", solution: "hello", expectedFeedback: feedback{correctPosition, wrongPosition, correctPosition, wrongPosition, correctPosition}},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			fb := computeFeedback([]rune(tc.guess), []rune(tc.solution))
			if !tc.expectedFeedback.Equal(fb) {
				t.Errorf(
					"guess: %q, got the wrong feedback, wanted %v, got %v", tc.guess, tc.expectedFeedback, fb)

			}
		})
	}
}
