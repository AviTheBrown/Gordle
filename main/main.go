package main

import (

    "gordle/gordle"
    "os"
)

func main() {
    g := gordle.New(os.Stdin, "Words", 5 )
    g.Play()
}