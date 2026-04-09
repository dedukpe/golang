package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	// Manually compute length of deck
	deckLen := 0
	for range deck {
		deckLen++
	}

	cardsPerPlayer := deckLen / 4
	playerNum := 1
	index := 0

	for playerNum <= 4 {
		fmt.Printf("Player %d: ", playerNum)
		cardCount := 0
		for cardCount < cardsPerPlayer {
			fmt.Printf("%d", deck[index])
			cardCount++
			index++
			if cardCount < cardsPerPlayer {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("\n")
		playerNum++
	}
}
