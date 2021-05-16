package algorithm

import (
	"fmt"
)

type Trie struct {
	word     string
	children [26]*Trie
	index    int
}

func NewTrie(word string) *Trie {
	return &Trie{
		word:  word,
		index: -1,
	}
}

func (t *Trie) add(word string, index int) {
	cur := t
	for i, e := range word {
		loc := e - 'a'
		if cur.children[loc] == nil {
			cur.children[loc] = NewTrie(string(word[i]))
		}
		cur = cur.children[loc]
	}
	cur.index = index
}

func (t *Trie) check_bfs() int {

	cur := t
	que := make([]*Trie, 0)
	que = append(que, cur)
	var leafNode *Trie
	for len(que) > 0 {
		top := que[0]
		que = que[1:]
		for i := len(top.children) - 1; i >= 0; i-- {
			if top.children[i] != nil && top.children[i].index != top.index && top.children[i].index != -1 {
				fmt.Println(top.word, top.index)
				que = append(que, top.children[i])
			}
			leafNode = top
		}
	}
	return leafNode.index

}

func LongestWord(words []string) string {
	root := NewTrie("root")
	for i := range words {
		root.add(words[i], i)
	}
	index := root.check_bfs()
	if index == -1 {
		return ""
	}
	return words[index]
}
