package main

import (
    "fmt"
)

type Hand struct {
    Cards [12]Card
    Size int
}

func handDefaultConstructor() Hand {
    var c = Card{0, "?", "?"}
    return Hand{[12]Card{c, c, c, c, c, c, c, c, c, c, c, c}, 0}
}

// add a copy of a card to a hand
func add(h *Hand, c Card) {
    h.Cards[h.Size] = c
    h.Size += 1
}

// get the value of a hand
func (h Hand) getValue() int {
    sum := 0
    hasAce := false
    for i := 0; i < h.Size; i++ {
        sum += h.Cards[i].Value
        if h.Cards[i].isAce() {
            hasAce = true
        }
    }
    if hasAce && sum + 10 <= 21 {
        return sum + 10
    }
    return sum
}

// get a hand string 
func (h Hand) String() string {
    outStr := ""
    for i := 0; i < h.Size; i++ {
        outStr += h.Cards[i].String() + " "
    }
    outStr += fmt.Sprintf("(%v)", h.getValue())
    return outStr
}