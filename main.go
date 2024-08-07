package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// Функція для перевірки, чи слово містить дві однакові букви, розділені будь-яким символом
func hasTwoSameLettersSeparated(word string) bool {
	for i := 0; i < len(word)-2; i++ {
		if word[i] == word[i+2] {
			return true
		}
	}
	return false
}

func main() {
	// Відкриваємо текстовий файл
	file, err := os.Open("text2.txt")
	if err != nil {
		fmt.Println("Помилка при відкритті файлу:", err)
		return
	}
	defer file.Close()

	// Регулярний вираз для знаходження слів
	wordRegExp := regexp.MustCompile(`[а-яА-Яa-zA-Z]+`)

	// Змінна для зберігання результатів
	var wordsWithTwoSameLetters []string

	// Читаємо файл рядок за рядком
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Знаходимо всі слова в рядку
		words := wordRegExp.FindAllString(line, -1)
		for _, word := range words {
			if hasTwoSameLettersSeparated(word) {
				wordsWithTwoSameLetters = append(wordsWithTwoSameLetters, word)
			}
		}
	}

	// Виводимо результати
	fmt.Println("Слова з двома однаковими буквами, розділеними будь-яким символом:")
	for _, word := range wordsWithTwoSameLetters {
		fmt.Println(word)
	}

	// Перевіряємо на помилки під час читання
	if err := scanner.Err(); err != nil {
		fmt.Println("Помилка при читанні файлу:", err)
	}
}
