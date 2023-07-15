package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new type of dack a kind a slic

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuit := []string{"Spades", "Diamond", "heart", "club"}
	cardValues := []string{"ace", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "king", "quene", "joker"}
	for _, suit := range cardSuit {
		for _, value := range cardValues {
			cards = append(cards, value+"of"+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}
func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}
func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}
func newDeckfromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}
func (d deck) shuffle() {
	Source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(Source)

	for i := range d {
		newPostion := r.Intn(len(d) - 1)
		d[i], d[newPostion] = d[newPostion], d[i]
	}

}
