package main

import (
   "fmt"
   "bufio"
   "os"
   "io/ioutil"
   "encoding/json"
	"math/rand"
	"strings"
)

// main
func main() {
   masterDeck := readIn()
   reader := bufio.NewReader(os.Stdin)
   fmt.Println("Welcome to blackjack. What is your name please?")
   name, _ := reader.ReadString('\n')
   // different replaces for Unix or Windows
   name = strings.Replace(name, "\r\n", "", -1)
   name = strings.Replace(name, "\n", "", -1)
   for {
      //fmt.Println("Would you like to play a hand, %v? [y/n]", name)
      fmt.Println(fmt.Sprintf("Would you like to play a hand, %v? [y/n]", name))
      switch getKey() {
         case 'y':
           gameLoop(masterDeck)
           break
         case 'n':
           fmt.Println("Goodbye")
           return
         default:
           fmt.Println("Command not understood")
      }
   }
}

func gameLoop(masterDeck Deck) {
   deck := shuffle(masterDeck)
   deckIndex := 4
   playerHandSize := 2
   dealerHandSize := 2
   playerHandArr := [5]Card{deck[0], deck[0], deck[0], deck[0], deck[0]}
   dealerHandArr := [5]Card{deck[0], deck[0], deck[0], deck[0], deck[0]}
   
   playerHandArr[0] = deck[0]
   playerHandArr[1] = deck[1]
   dealerHandArr[0] = deck[2]
   dealerHandArr[1] = deck[3]
   
   // player draw loop
   keepLooping := true
   for ; keepLooping; {
      hitOrStand(playerHandArr, playerHandSize, dealerHandArr)
      switch getKey() {
         case 'h':
            playerHandArr[playerHandSize] = deck[deckIndex]
            playerHandSize++;
            deckIndex++
            if playerHandSize == 5 {
               keepLooping = false
            }
            if getVal(playerHandArr, playerHandSize) > 21 {
               keepLooping = false
            }
            break
         case 's':
            keepLooping = false
            break
         default:
            fmt.Println("Command not understood")
      }
   }
   
   // dealer draw loop
   keepLooping = true
   // player has already won or lost
   if getVal(playerHandArr, playerHandSize) > 21 {
      keepLooping = false
   } else if playerHandSize == 5 {
      keepLooping = false
   }
   for ; keepLooping; {
      if getVal(dealerHandArr, dealerHandSize) < getVal(playerHandArr, playerHandSize) {
         dealerHandArr[playerHandSize] = deck[deckIndex]
         dealerHandSize++;
         deckIndex++
      }
      if getVal(dealerHandArr, dealerHandSize) >= getVal(playerHandArr, playerHandSize) {
         keepLooping = false
      }
      if dealerHandSize == 5 {
         keepLooping = false
      }
   }
   
   // show results
   fmt.Println(fmt.Sprintf("You have %v. ", handString(playerHandArr, playerHandSize)))
   fmt.Println(fmt.Sprintf("Dealer has %v. ", handString(dealerHandArr, dealerHandSize)))
   if getVal(playerHandArr, playerHandSize) > 21 {
      fmt.Println("Busted! You lose.")
   } else  if getVal(dealerHandArr, dealerHandSize) > 21{
      fmt.Println("Dealer busted! You win!")
   } else {
      resultVal := getVal(playerHandArr, playerHandSize) - getVal(dealerHandArr, dealerHandSize)
      if resultVal > 0 {
         fmt.Println("You win!")
      } else if resultVal < 0 {
         fmt.Println("You lose!")
      } else {
         fmt.Println("Push! It's a tie!")
      }
   }
}

// ask hit or stand 
func hitOrStand(pCards [5]Card, size int, dCards [5]Card) {
   queryStr := fmt.Sprintf("You have %v. ", handString(pCards, size))
   queryStr += fmt.Sprintf("Dealer is showing %v.", dCards[0].String())
   fmt.Println(queryStr)
   fmt.Println("[h]it or [s]tand?")
}

// get a hand string 
func handString(cards [5]Card, size int) string {
   outStr := ""
	for i := 0; i < size; i++ {
      outStr += cards[i].String() + " "
	}
   outStr += fmt.Sprintf("(%v)", getVal(cards, size))
   return outStr
}

// read key input
func getKey() rune {
   reader := bufio.NewReader(os.Stdin)
   char, _, err := reader.ReadRune()
   
   if err != nil {
      fmt.Println(err)
   }
   
   return char
}


// cards
//////////////////////////////////////////////////////////

type Card struct {
	Value int `json:"value"`
	Face string `json:"face"`
   Suit string `json:"suit"`
}

type Deck struct {
    Cards []Card `json:"cards"`
}

type Hand struct {
    Cards [5]Card
    Size int
}

func (c Card) isAce() bool {
	return c.Value == 1
}

func (c Card) String() string {
	return c.Face + c.Suit
}

// add a card to a hand
func (h Hand) add(c Card) {
   h.Cards[h.Size] = c
   h.Size = h.Size + 1
   fmt.Println(h.Size)
}

// return the value of a hand of cards
func getVal(cards [5]Card, size int) int{
	sum := 0
   hasAce := false
	for i := 0; i < size; i++ {
      sum += cards[i].Value
      if cards[i].isAce() {
         hasAce = true
      }
	}
   if hasAce && sum + 10 <= 21 {
      return sum + 10
   }
   return sum
}

// return a shuffled copy of the deck
func shuffle(deck Deck) [52]Card {
   // make a copy
   var newDeck [52]Card
	for i := 0; i < len(newDeck); i++ {
      newDeck[i] = deck.Cards[i]
	}
   // shuffle seven times
   var temp = newDeck[0]
	for j := 0; j < 7; j++ {
   	for k := 0; k < len(newDeck); k++ {
         if rand.Intn(2) == 0 {
            temp = newDeck[0]
            newDeck[0] = newDeck[k]
            newDeck[k] = temp
         }
   	}
	}
   return newDeck
}

// json
//////////////////////////////////////////////////////////

func readIn() Deck {
   jsonFile, err := os.Open("cards.json")
   if err != nil {
       fmt.Println(err)
   }
   // close when we're done
   defer jsonFile.Close()
   byteValue, _ := ioutil.ReadAll(jsonFile)
   var deck Deck
   json.Unmarshal(byteValue, &deck)
   
   return deck
}