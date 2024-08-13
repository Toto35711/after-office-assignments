package main

import "testing"

func TestArraySign(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 1}, 1},
		{[]int{-2, 1}, -1},
		{[]int{-1, -2, -3, -4, 3, 2, 1}, 1},
		{[]int{1, 2, 0, 2, 1}, 0},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			res := arraySign(c.input)
			if res != c.expected {
				t.Errorf("arraySign(%v) = %v; expected %v", c.input, res, c.expected)
			}
		})
	}
}

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		firstWord  string
		secondWord string
		expected   bool
	}{
		{"anak", "kana", true},
		{"anak", "mana", false},
		{"anagram", "managra", false},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			res := isAnagram(c.firstWord, c.secondWord)
			if res != c.expected {
				t.Errorf("isAnagram(%v, %v) = %v; expected %v", c.firstWord, c.secondWord, res, c.expected)
			}
		})
	}
}

func TestFindTheDifference(t *testing.T) {
	cases := []struct {
		s        string
		t        string
		expected byte
	}{
		{"abcd", "abcde", 'e'},
		{"abcd", "abced", 'e'},
		{"", "y", 'y'},
		{"abceed", "abced", 'e'},
		{"afteroffice", "afterxoffice", 'x'},
		{"sayasukagolang", "sayatsukagolang", 't'},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			res := findTheDifference(c.s, c.t)
			if res != c.expected {
				t.Errorf("isAnagram(%v, %v) = %v; expected %v", c.s, c.t, res, c.expected)
			}
		})
	}
}

func TestCanMakeArithmeticProgressio(t *testing.T) {
	cases := []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 5, 3}, true},
		{[]int{5, 1, 9}, true},
		{[]int{1, 2, 4, 8}, false},
		{[]int{1, 1, 1, 1}, true},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			res := canMakeArithmeticProgression(c.input)
			if res != c.expected {
				t.Errorf("canMakeArithmeticProgression(%v) = %v; expected %v", c.input, res, c.expected)
			}
		})
	}
}
