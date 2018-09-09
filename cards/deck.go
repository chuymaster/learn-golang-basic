package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
type deck []string // this extends []string behavior, not subclass

func (d deck) print() { // (d deck) defines deck struct as a 'receiver'. 'd' = value of 'deck' type as a variable.
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{} // empty deck

	// card suits
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Club"}

	// card values
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// add a combination to the deck
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	d.shuffle()
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) toByteSlice() []byte {
	return []byte(d.toString())
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, d.toByteSlice(), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Log the error and entirely quit the program, not return nil
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		// newPosition := rand.Intn(len(d) - 1)
		// get random int number. this one has constant seed so not really random...

		newPosition := r.Intn(len(d) - 1)           // get random int number
		d[i], d[newPosition] = d[newPosition], d[i] // swap the elements
	}
}
