package trie

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := CreateNewTrie()
	trie.InsertWord("Hello World")

	suggestions := trie.GetWords("Hello", 10)
	if len(suggestions) != 1 {
		t.Fatalf("Wrong lenght")
	}
	assert.Equal(t, 6, suggestions[0].Depth, "check Depth")
	assert.Equal(t, "Hello World", suggestions[0].Suggestion)
	fmt.Println(suggestions)
	fmt.Println(trie.OrigianWords[strings.ToLower("Hello")])

	trie.Erase()
	suggestionsAfterErase := trie.GetWords("Hello", 10)
	if len(suggestionsAfterErase) != 0 {
		t.Fatalf("Did not reset")
	}

}

func TestCaseInsensitiveTrie(t *testing.T) {
	fmt.Println("Check Case insensitive querries")
	trie := CreateNewTrie()
	trie.InsertWord("Hello World")
	trie.InsertWord("whats up")

	suggestionsHello := trie.GetWords("hello", 10)
	suggestionsWhats := trie.GetWords("What", 10)

	assert.Equal(t, "Hello World", suggestionsHello[0].Suggestion)
	assert.Equal(t, "whats up", suggestionsWhats[0].Suggestion)

}
