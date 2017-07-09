package trie

import (
	"goroom/trie"
	"testing"
)

func Test_Find(t *testing.T) {
	trie := trie.New()
	values, found := trie.Find([]byte("kinesh"))
	if found {
		t.Error("Found element in empty trie")
	}
	trie.Insert([]byte("kinesh"), "Kinesh Satiya")
	values, found = trie.Find([]byte("kinesh"))
	if !found {
		t.Error("Could not find inserted element")
	}
}
