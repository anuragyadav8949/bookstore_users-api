package services

import (
	"fmt"
	"strings"

	"github.com/anuragyadav8949/bookstore_users-api/domain/users"
	"github.com/anuragyadav8949/bookstore_users-api/utils/errors"
)

func CheckForWord(textValue users.BookText) (map[string]int, *errors.RestError) {
	fmt.Println("Check for word")
	stringData := textValue.BookTitle

	//Split the word into array
	splitString := strings.Split(stringData, " ")

	//Create a map to store word in a key and count as value
	var WordMap = make(map[string]int)

	for _, word := range splitString {
		_, matched := WordMap[word]
		if matched {
			WordMap[word] += 1
		} else {
			WordMap[word] = 1
		}
	}
	fmt.Println(WordMap)
	return WordMap, nil
}
