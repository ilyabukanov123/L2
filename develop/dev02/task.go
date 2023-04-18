package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func UnpacString(str string) (string, error) {
	// runes := []rune(str)

	// проверка на пустую строку
	if len(str) == 0 {
		// return "", fmt.Errorf("Некорректная строка")
		return "", errors.New("передана пустая строка")
	}

	// Проверяем на число  в качестве 0 символа строки
	// Если строка не число возвращает ошибку. В ином случае возвращает число
	if _, err := strconv.Atoi(string((rune(str[0])))); err == nil {
		return "", errors.New("первый символ в строке является числом")
	}

	var char rune
	result := ""
	for _, v := range str {
		// количество букв
		numberOfLetters := 1
		// Проверяем является ли текущий символ руны десятичным числом
		if !unicode.IsDigit(v) {
			char = v
		} else {
			// Вычисляем каким числом является число в руне
			number, _ := strconv.Atoi(string(v))
			numberOfLetters = number - 1
		}
		// Repeat возвращает новую строку, состоящую из numberOfLetters копий строки char
		result += strings.Repeat(string(char), numberOfLetters)
	}
	return result, nil
}

func main() {
	fmt.Println(UnpacString("He3llo world"))
}
