package trie_test

import (
	"go-room/trie"
	"testing"
)

func Test_Find(t *testing.T) {
	trie := trie.New()
	_, found := trie.Find([]byte("kinesh"))
	if found {
		t.Error("Found element in empty trie")
	}
	trie.Insert([]byte("kinesh"), "Kinesh Satiya")
	_, found = trie.Find([]byte("kinesh"))
	if !found {
		t.Error("Could not find inserted element")
	}
}
