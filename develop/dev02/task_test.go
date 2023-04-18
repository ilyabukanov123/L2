package main

import "testing"

// пустая строка
// первый символ число

func TestStringEmpty(t *testing.T) {
	str, err := UnpacString("")
	if err != nil {
		return
	}
	t.Fatalf("Должна была произойти ошибка, но вернулось строка: %v", str)
}

func TestForThePresenceOfCharactersIsAString(t *testing.T) {
	str, err := UnpacString("3Hello World")
	if err != nil {
		return
	}
	t.Fatalf("Должна была произойти ошибка, но вернулось строка: %v", str)
}

func TestStringForTheCorrect(t *testing.T) {
	str, err := UnpacString("He3llo Wo2rld")
	if err != nil {
		t.Fatalf("Должно была вернутся строка, но вернулась ошибка: %v", err)
	}
	t.Logf("Полученная строка: %v", str)
}
