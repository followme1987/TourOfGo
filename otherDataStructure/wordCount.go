package otherDataStructure

import (
	"fmt"
	"strings"
)

func WordCounter(str string) {
	words := strings.Fields(str)
	wordMap := make(map[string]int)

	for _, val := range words {
		if _, ok := wordMap[val]; ok {
			wordMap[val]++
		} else {
			wordMap[val] = 1
		}
	}

	for k, v := range wordMap {
		fmt.Printf("%v has happened %d times \n", k, v)
	}

}
