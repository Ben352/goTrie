package trie

import (
	"fmt"
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
	assert.Equal(t, suggestions[0].Depth, 6, "check Depth")
	assert.Equal(t, suggestions[0].Suggestion, "Hello World")
	fmt.Println(suggestions)
}
