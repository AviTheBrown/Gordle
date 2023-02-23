package test

import (
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
            g := gordle.New(strings.NewReader(tc.input))
            got := g.Ask()

            if !slices.Equal(got, tc.want) {
                    t.Errorf("readRunes() got = %v, want %v", string(got),
                        string(tc.want))
            }
        })
    }

}