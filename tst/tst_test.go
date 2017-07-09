package tst

import (
	"goroom/tst"
	"testing"
)

func Test_Find(t *testing.T) {
	tst := tst.New()
	values := tst.Query([]byte("g"))
	if len(values) > 0 {
		t.Error("Found values from empty tst")
	}

	tst.Insert([]byte("go"), "Google")
	tst.Insert([]byte("golang"), "Google language")

	values = tst.Query([]byte("g"))
	if len(values) != 2 || values[0] != "Google" || values[1] != "Google language" {
		t.Errorf("Could not find correct values for query %v, values found : %v\n", "g", values)
	}

	values = tst.Query([]byte("gol"))
	if len(values) != 1 || values[0] != "Google language" {
		t.Errorf("Could not find correct values for query %v, values found %v\n", "gol", values)
	}

	values = tst.Query([]byte("z"))
	if len(values) > 0 {
		t.Errorf("Found values for key that was not inserted, query %v\n", "z")
	}
}
