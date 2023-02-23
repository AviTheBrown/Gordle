package gordle

type Game struct {}

// New -> creates a new intance of Game.
    //Entry point.
func New() *Game {
    g := &Game{}
    return g
}