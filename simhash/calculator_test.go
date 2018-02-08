package simhash

import (
	"testing"
	"crypto/md5"
)

func TestCalculateEmpty(t *testing.T) {
	res := Calculate("")
	if res != md5.Sum([]byte("")) {
		t.Error()
	}
}


func TestCalculateOneWord(t *testing.T) {
	word := "Hello"
	res := Calculate(word)
	if res != md5.Sum([]byte(word)) {
		t.Error()
	}
}

func TestCalculateToSameWords(t *testing.T) {
	word := "Hello"
	res := Calculate(word + " " + word)
	// это потому, что биты в хеше одинаковые
	if res != md5.Sum([]byte(word)) {
		t.Error()
	}
}

func TestDifferenceSameHashes(t *testing.T) {
	res := Calculate("Hello")
	diff := Difference(res, res)
	if diff != 0 {
		t.Errorf("Difference should be 0, %f instead", diff)
	}
}

func TestSmallDifference(t *testing.T) {
	h1 := Calculate("This is a reference manual for the Go programming language")
	h2 := Calculate("This is a reference manual for the Go programming languag")
	diff := Difference(h1, h2)
	if diff > 0.2 {
		t.Errorf("Difference should be < 0.5, %f instead", diff)
	}
}


func TestBigDifference(t *testing.T) {
	h1 := Calculate("This is a reference manual for the Go programming language")
	h2 := Calculate("The grammar is compact and regular")
	diff := Difference(h1, h2)
	if diff < 0.5 {
		t.Errorf("Difference should be < 0.5, %f instead", diff)
	}
}