package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := CreateNewTrie()
	trie.InsertWord("Hello World")

	suggestions := trie.GetWords("Hello", 10)
	if len(suggestions) != 1 {
		t.Fatalf("Wrong lenght")
	}
}
