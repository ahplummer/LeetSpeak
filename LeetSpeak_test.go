package main

import "testing"

func TestGenerateMap(t *testing.T) {
	ls := new(LeetSpeaker)
	ls.Init()
	expected := "|D"
	actual := ls.leetmap["P"]
	if expected != actual {
		t.Errorf("Failed in map translation, expected %s, got %s", expected, actual)
	}
}

func TestTranslateToLeet(t *testing.T){
	ls := new(LeetSpeaker)
	ls.Init()
	english := "&*668hAcKor990"
	expected := "&*668#4(|<0|2990"
	actual := ls.translateToLeet(english)
	if expected != actual {
		t.Errorf("Failed in map translation, expected %s, got %s", expected, actual)
	}
}