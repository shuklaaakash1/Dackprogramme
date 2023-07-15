package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("expacted the len of deck is 52 but we got %v ", len(d))
	}
	if d[0] != "aceofSpades" {
		t.Errorf("epacted the card of aceofSpades but we got %v", d[0])
	}
	if d[len(d)-1] != "jokerofclub" {
		t.Errorf("epacted the card of aceofSpades but we got %v", d[len(d)-1])
	}
}
func TestSaveToDeckAndNewDackFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckfromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("expacted the 52 no of cade from and we got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
