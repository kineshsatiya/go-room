package bst_test

import (
	"go-room/bst"
	"testing"
)

func Test_Contains(t *testing.T) {
	b := bst.New()
	found := b.Contains(1)
	if found {
		t.Error("Found value in empty bst")
	}
	b.Insert(1)
	found = b.Contains(1)
	if !found {
		t.Errorf("expecting the root value %d", 1)
	}
}

func Test_InOrder(t *testing.T) {
	b := bst.New()
	inOrder := b.InOrder()
	if inOrder != nil {
		t.Error("In order array should be nil for an empty tree")
	}

	b.Insert(1)
	inOrder = b.InOrder()
	if len(inOrder) != 1 || inOrder[0] != 1 {
		t.Error("Incorrect inorder array")
	}

	b.Insert(2)
	inOrder = b.InOrder()
	if len(inOrder) != 2 || inOrder[0] != 1 || inOrder[1] != 2 {
		t.Error("Incorrect inorder array")
	}

	b.Insert(0)
	inOrder = b.InOrder()
	if len(inOrder) != 3 || inOrder[0] != 0 || inOrder[1] != 1 || inOrder[2] != 2 {
		t.Error("Incorrect inorder array")
	}
}
