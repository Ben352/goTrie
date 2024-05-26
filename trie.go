package trie

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Trie struct {
	Root         *Node
	Words        []string
	OrigianWords map[string]map[string]struct{}
}

type Node struct {
	Letter   rune
	Children map[rune]*Node
	IsLeaf   bool
}

type SearchResult struct {
	Depth      int    `json:"Depth"`
	Suggestion string `json:"Suggestions"`
}

func (t *Trie) AddWord(word string) {
	t.Words = append(t.Words, word)
}

func CreateNewNode(letter rune) *Node {
	node := &Node{Letter: letter, Children: make(map[rune]*Node), IsLeaf: false}

	return node
}

func (t *Trie) InsertWord(word string) {
	lowercaseWord := strings.ToLower(word)
	currentRoot := t.Root
	for _, letter := range lowercaseWord {

		_, ok := currentRoot.Children[letter]
		if ok {

		} else {
			currentRoot.Children[letter] = CreateNewNode(letter)

		}
		currentRoot = currentRoot.Children[letter]
	}
	currentRoot.IsLeaf = true
	t.AddWord(word)

	_, ok := t.OrigianWords[lowercaseWord]
	// If the key exists
	if !ok {
		t.OrigianWords[lowercaseWord] = make(map[string]struct{})
		t.OrigianWords[lowercaseWord][word] = struct{}{}
	}

}

func (t *Trie) GetWords(word string, depth int) []SearchResult {
	options := []SearchResult{}
	var currentWord []rune
	currentRoot := t.Root
	lowercaseWord := strings.ToLower(word)
	for _, letter := range lowercaseWord {

		_, ok := currentRoot.Children[letter]
		if !ok {
			return options
		}
		currentWord = append(currentWord, letter)
		currentRoot = currentRoot.Children[letter]
	}

	type option struct {
		CurrentNode *Node
		CurrentWord []rune
		Depth       int
	}
	var stack []option
	stack = append(stack, option{currentRoot, currentWord, 0})
	for len(stack) != 0 {

		currentOption := stack[0]
		stack = stack[1:]
		if currentOption.CurrentNode.IsLeaf {
			options = append(options, SearchResult{Depth: currentOption.Depth, Suggestion: string(currentOption.CurrentWord)})
		}
		if currentOption.Depth < depth {
			for k, v := range currentOption.CurrentNode.Children {
				newWord := make([]rune, len(currentOption.CurrentWord))
				_ = copy(newWord, currentOption.CurrentWord)

				newWord = append(newWord, k)

				newOption := option{v, newWord, currentOption.Depth + 1}

				stack = append(stack, newOption)
			}
		}
	}
	caseInsensetiveResults := []SearchResult{}
	for _, element := range options {
		for k, _ := range t.OrigianWords[element.Suggestion] {
			caseInsensetiveResults = append(caseInsensetiveResults, SearchResult{Suggestion: k, Depth: element.Depth})
		}
	}
	return caseInsensetiveResults

	//  findalloptions

}

func (t *Trie) WordExsists(word string) bool {
	currentRoot := t.Root

	for _, letter := range word {

		_, ok := currentRoot.Children[letter]
		if ok {

		} else {

			fmt.Println("Not found, but added")
			return false
		}
		currentRoot = currentRoot.Children[letter]
	}
	return currentRoot.IsLeaf

}

func (t *Trie) Erase() {
	t.Root.Children = make(map[rune]*Node)
}

func CreateNewTrie() *Trie {
	emptyNode := &Node{Letter: 'h', Children: make(map[rune]*Node), IsLeaf: false}
	return &Trie{Root: emptyNode, OrigianWords: make(map[string]map[string]struct{})}
}

func (t *Trie) LoadTrie(fileName string) {
	t.Erase()
	file, _ := os.ReadFile(fileName)

	laodedResults := []string{}
	err := json.Unmarshal(file, &laodedResults)
	if err != nil {
		fmt.Println(err)
	}
	intCounte := 0
	for intCounte < len(laodedResults) {
		t.InsertWord(laodedResults[intCounte])
		intCounte += 1
	}
	fmt.Println(laodedResults)
}

func (t *Trie) SerializeTrie(fileName string) {
	// results := t.GetWords("Hello Wor", 10)
	jsonData, err := json.MarshalIndent(t.Words, "", "    ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	os.WriteFile(fileName, jsonData, os.ModePerm)
	fmt.Println(string(jsonData))
}

// func main() {

// 	newTrie := Trie{Root: CreateNewNode('a')}
// 	newTrie.LoadTrie("data.json")
// 	// newTrie.InsertWord("Hello World")
// 	// newTrie.InsertWord("Hello World2")
// 	// newTrie.InsertWord("Hello Wor467891")
// 	// newTrie.InsertWord("Hello World3")
// 	// newTrie.InsertWord("Hello World4")
// 	// newTrie.InsertWord("Hello World5")
// 	fmt.Println(newTrie.GetWords("Hel", 10))
// 	fmt.Println(newTrie.GetWords("H", 10))
// 	// newTrie.SerializeTrie()
// }
