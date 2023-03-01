package test

import (
    "errors"
    "gordle/gordle"
    "strings"
    "testing"
    "golang.org/x/exp/slices"

)

func TestGameAsk(t *testing.T) {

    testCase := map[string]struct{
        input string
        want []rune
    }{
        "5 cases in english": {
            input: "Hello",
            want: []rune("Hello"),
        },
        "5 characters in arabic": {
            input: "مرحبا",
            want: []rune("مرحبا"),
        },
        "5 characters in japanese" : {
            input: "こんにちは",
            want: []rune("こんにちは",),
        },
        "3 characters in japanese" : {
            input: "こんに",
            want: []rune("こんに"),
        },
    }

    for name, tc := range testCase {
        t.Run(name, func(t *testing.T) {
            g := gordle.New(strings.NewReader(tc.input), string(tc.want), 0)
            got := g.Ask()

            if !slices.Equal(got, tc.want) {
                    t.Errorf("readRunes() got = %v, want %v", string(got),
                        string(tc.want))
            }
        })
    }

}

func TestValidateGuess(t *testing.T) {
    // needs to test the input of the parameters
        // determin if the parameters are runes or not.
    // test the word length
        // correct and incorrect lenghts
    // test the erroe message and if it displays correctly

    // prep phase
    tt := map[string]struct{
        word []rune
        expected error
    }{
        "nominal" : {
            word: []rune("GUESS"),
            expected: nil,
        },
        "too long" : {
            word:     []rune("Lunchbox"),
            expected: gordle.ErrInvalidWordLength,
        },
        "too short" : {
            word: []rune("egg"),
            expected: gordle.ErrInvalidWordLength,
        },
    }

    // execution phase
    for name, tc := range tt {
        t.Run(name, func(t *testing.T) {
            g := gordle.New(nil)

            err := g.ValidateGuess(tc.word)
    //validation phase
            if !errors.Is(err,tc.expected) {
                t.Errorf("%c, expected %q, got %q", tc.word, tc.expected, err)
            }
        })
    }




}