package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"math/rand"
)

func main() {
	// arraySign([]int{2, 1})                    // 1
	// arraySign([]int{-2, 1})                   // -1
	// arraySign([]int{-1, -2, -3, -4, 3, 2, 1}) // 1

	// fmt.Println(a, b, c)
	// isAnagram("anak", "kana") // true
	// isAnagram("anak", "mana") // false
	// isAnagram("anagram", "managra") // true

	// findTheDifference("abcd", "abcde") // 'e'
	// findTheDifference("abcd", "abced") // 'e'
	// findTheDifference("", "y")         // 'y'

	// canMakeArithmeticProgression([]int{1, 5, 3})    // true; 1, 3, 5 adalah baris aritmatik +2
	// canMakeArithmeticProgression([]int{5, 1, 9})    // true; 9, 5, 1 adalah baris aritmatik -4
	// canMakeArithmeticProgression([]int{1, 2, 4, 8}) // false; 1, 2, 4, 8 bukan baris aritmatik, melainkan geometrik x2

	tesDeck()
}

// https://leetcode.com/problems/sign-of-the-product-of-an-array
func arraySign(nums []int) int {
	var countNegativeNumbers int = 0
	for i := 0; i < len(nums); i++ {
		// saya tambah untuk handle kasus perkaliannya sama dengan 0, karena 0 bukan positif atau negatif
		if nums[i] == 0 {
			return 0
		}
		if nums[i] < 0 {
			countNegativeNumbers += 1
		}
	}

	if countNegativeNumbers%2 == 0 {
		return 1
	} else {
		return -1
	}
}

// https://leetcode.com/problems/valid-anagram
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	firstLetters := strings.Split(s, "")
	secondLetters := strings.Split(t, "")
	length := len(firstLetters)

	for i := 0; i < length; i++ {
		if firstLetters[i] != secondLetters[length-1-i] {
			return false
		}
	}

	return true
}

// https://leetcode.com/problems/find-the-difference
func findTheDifference(s string, t string) byte {
	if len(s) == 0 {
		return byte(t[0])
	}

	var letterCountMap = map[byte]int{}
	var letterCountMap2 = map[byte]int{}

	for i := 0; i < len(s); i++ {
		letterCountMap[s[i]] += 1
	}
	for i := 0; i < len(t); i++ {
		letterCountMap2[t[i]] += 1
	}

	for k, v := range letterCountMap2 {
		if v != letterCountMap[k] {
			return k
		}
	}

	return byte('x')
}

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence
func canMakeArithmeticProgression(arr []int) bool {
	slices.Sort(arr)

	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] != arr[1]-arr[0] {
			return false
		}
	}
	return true
}

// Deck represent "standard" deck consist of 52 cards
type Deck struct {
	cards []Card
}

// Card represent a card in "standard" deck
type Card struct {
	symbol int // 0: spade, 1: heart, 2: club, 3: diamond
	number int // Ace: 1, Jack: 11, Queen: 12, King: 13
}

// New insert 52 cards into deck d, sorted by symbol & then number.
// [A Spade, 2 Spade,  ..., A Heart, 2 Heart, ..., J Diamond, Q Diamond, K Diamond ]
// assume Ace-Spade on top of deck.
func (d *Deck) New() {
	for i := 0; i < 4; i++ {
		for j := 1; j <= 13; j++ {
			d.cards = append(d.cards, Card{i, j})
		}
	}

}

// PeekTop return n cards from the top
func (d Deck) PeekTop(n int) []Card {
	return d.cards[:n]
}

// PeekTop return n cards from the bottom
func (d Deck) PeekBottom(n int) []Card {
	return d.cards[len(d.cards)-n:]
}

// PeekCardAtIndex return a card at specified index
func (d Deck) PeekCardAtIndex(idx int) Card {
	return d.cards[idx]
}

// Shuffle randomly shuffle the deck
func (d *Deck) Shuffle() {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	for i,_ := range d.cards {
		newPos := r.Intn(len(d.cards)-1)
		d.cards[i], d.cards[newPos] = d.cards[newPos], d.cards[i]
	}
}

// Cut perform single "Cut" technique. Move n top cards to bottom
// e.g. Deck: [1, 2, 3, 4, 5]. Cut(3) resulting Deck: [4, 5, 1, 2, 3]
func (d *Deck) Cut(n int) {
	// write code here
	
}

func (c Card) ToString() string {
	textNum := ""
	switch c.number {
	case 1:
		textNum = "Ace"
	case 11:
		textNum = "Jack"
	case 12:
		textNum = "Queen"
	case 13:
		textNum = "King"
	default:
		textNum = fmt.Sprintf("%d", c.number)
	}
	texts := []string{"Spade", "Heart", "Club", "Diamond"}
	return fmt.Sprintf("%s %s", textNum, texts[c.symbol])
}

func tesDeck() {
	deck := Deck{}
	deck.New()

	top5Cards := deck.PeekTop(5)
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}
	fmt.Println("---")

	fmt.Println(deck.PeekCardAtIndex(12).ToString()) // King Spade
	fmt.Println(deck.PeekCardAtIndex(13).ToString()) // Ace Heart
	fmt.Println(deck.PeekCardAtIndex(14).ToString()) // 2 Heart
	fmt.Println(deck.PeekCardAtIndex(15).ToString()) // 3 Heart
	fmt.Println("---")

	deck.Shuffle()
	top5Cards = deck.PeekTop(10)
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}

	fmt.Println("---")
	deck.New()
	deck.Cut(5)
	bottomCards := deck.PeekBottom(10)
	for _, c := range bottomCards {
		fmt.Println(c.ToString())
	}
}
