package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Three, Suit: Diamond})
	fmt.Println(Card{Rank: Four, Suit: Club})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Three of Diamonds
	// Four of Clubs
	// Joker

}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 13*4 {
		t.Error("Wrong many cards a new deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card, Received:", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card, Received:", cards[0])
	}
}

func TestShuffle(t *testing.T) {
	// make shuffle rand deterministic
	// First call to shuffleRand.Perm(52) sshould be
	//[40, 35 ...]
	shuffleRand = rand.New(rand.NewSource(0)) // will establish a seed so that permutation of indexes will always be the same

	origDeck := New()
	first := origDeck[40]
	second := origDeck[35]
	cards := New(Shuffle)

	if cards[0] != first {
		t.Errorf("Expected the first card to be %s, received %s", first, cards[0])
	}

	if cards[1] != second {
		t.Errorf("Expected the first card to be %s, received %s", second, cards[1])
	}

}

func TestJokers(t *testing.T) {

	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}

	if count != 3 {
		t.Error("Expected 3 Jokers, received:", count)
	}

}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}

	cards := New(Filter(filter))

	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected all twos and threes to be filtered out")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	// 13 ranks * 4 suits * 3 decks
	if len(cards) != 13*4*3 {
		t.Errorf("Expected %d cards, received %d cards.", 13*4*3, len(cards))
	}
}
