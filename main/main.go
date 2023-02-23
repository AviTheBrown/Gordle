package main

import (
    "gordle/gordle"
    "os"
)

func main() {
    g := gordle.New(os.Stdin)
    g.Play()
}