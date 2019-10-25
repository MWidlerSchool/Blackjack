package main

type Card struct {
	Value int `json:"value"`
	Face string `json:"face"`
	Suit string `json:"suit"`
}

func (c Card) isAce() bool {
	return c.Value == 1
}

func (c Card) String() string {
	return c.Face + c.Suit
}
