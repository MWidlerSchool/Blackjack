package main

import (
	"fmt"
	"bufio"
	"os"
	"io/ioutil"
	"encoding/json"
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
		fmt.Println(fmt.Sprintf("Would you like to play a hand, %v? [y/n]", name))
		switch getKey() {
		case 'y':
			gameLoop(masterDeck)
		case 'n':
			fmt.Println("Goodbye")
			return
		default:
			fmt.Println("Command not understood")
		}
	}
}

func gameLoop(deck Deck) {
	deck = shuffle(deck)
	playerHand := handDefaultConstructor()
	dealerHand := handDefaultConstructor()

	add(&playerHand, draw(&deck))
	add(&playerHand, draw(&deck))
	add(&dealerHand, draw(&deck))
	add(&dealerHand, draw(&deck))
	
	// blank line for readability
	fmt.Println("")
	
	// player draw loop
	keepLooping := true
	for ; keepLooping; {
		hitOrStand(playerHand, dealerHand)
		switch getKey() {
			case 'h':
				add(&playerHand, draw(&deck))
				fmt.Println(playerHand.String())
				if playerHand.getValue() > 21 {
					keepLooping = false
				}
			case 's':
				keepLooping = false
			default:
				fmt.Println("Command not understood")
		}
	}
   
	// dealer draw loop
	keepLooping = true
   
	// player has already won or lost
	if playerHand.getValue() > 21 {
		keepLooping = false
	}
	for ; keepLooping; {
		if dealerHand.getValue() < playerHand.getValue() {
			add(&dealerHand, draw(&deck))
		}
		if dealerHand.getValue() >= playerHand.getValue() {
			keepLooping = false
		}
	}
   
	// show values
	fmt.Println("Your hand: " + playerHand.String())
	fmt.Println("Dealer hand: " + dealerHand.String())
   
	// show game results
	if playerHand.getValue() > 21 {
		fmt.Println("Busted! You lose.")
	} else  if dealerHand.getValue() > 21{
		fmt.Println("Dealer busted! You win!")
	} else {
		resultVal := playerHand.getValue() - dealerHand.getValue()
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
func hitOrStand(playerHand Hand, dealerHand Hand) {
	queryStr := fmt.Sprintf("You have %v. ", playerHand.String())
	queryStr += fmt.Sprintf("Dealer is showing %v.", dealerHand.Cards[0].String())
	fmt.Println(queryStr)
	fmt.Println("[h]it or [s]tand?")
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