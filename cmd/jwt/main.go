package main

type Algorism int

const (
	HS256 Algorism = iota
)

type TokenType int

const (
	JWT TokenType = iota
)

type Header struct {
   Algorism `json:"alg"`
   TokenType `json:"typ"`
}

func main() {
}
