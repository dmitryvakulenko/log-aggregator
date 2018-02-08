package comparator

import "testing"

func TestCompareEmptyStrings(t *testing.T) {
	a, b := "", ""
	if Levenshtein(a, b) != 0 {
		t.Error()
	}
}


func TestCompareEqualStrings(t *testing.T) {
	a, b := "A", "A"
	if Levenshtein(a, b) != 0 {
		t.Error()
	}
}

func TestInserting(t *testing.T) {
	a, b := "A", "Ab"
	res := Levenshtein(a, b)
	if res != 1 {
		t.Error()
	}
}

func TestRemovingUTF8(t *testing.T) {
	a, b := "Я", ""
	res := Levenshtein(a, b)
	if res != 1 {
		t.Errorf("Should be 1 deletion, but got %d", res)
	}
}


func TestComplexUTF8(t *testing.T) {
	a, b := "Привет", "Пока"
	res := Levenshtein(a, b)
	if res != 5 {
		t.Errorf("Complexcity should be 5, %d instead", res)
	}
}