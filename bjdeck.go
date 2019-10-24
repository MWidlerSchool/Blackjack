package main

import (
   "math/rand"
   "time"
)

type Deck struct {
    Cards []Card `json:"cards"`
    Index int
}

// pop the top card
func draw(d* Deck) Card {
   var c = d.Cards[d.Index]
   d.Index += 1
   return c
}

// return a shuffled copy of the deck
func shuffle(deck Deck) Deck {
   deck.Index = 0
   rand.Seed(time.Now().UTC().UnixNano())
   // shuffle seven times
   var temp = deck.Cards[0]
   for i := 0; i < 7; i++ {
      for j := 0; j < 52; j++ {
         if rand.Intn(2) == 0 {
            temp = deck.Cards[0]
            deck.Cards[0] = deck.Cards[j]
            deck.Cards[j] = temp
         }
      }
   }
   return deck
}