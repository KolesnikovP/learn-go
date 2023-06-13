// package main - объявили новый пакет, в любом проекте должен быть обязательно пакет main. Запуск программы начинается именно с этого пакета
package main

// package for output and input
import (
	"fmt"
	"strconv"
	"strings"
	"sort"
)

// func main(){} - объявили функцию c названием "main".
// Имя main является особенным, эта функция будет вызываться сама при запуске
// программы. Эта функция тоже обязательная в программе, с неё начинает работать ваш код.
func main() {
	fmt.Println("hello fucking world again")
	fmt.Println(strconv.Itoa(100))
	fmt.Println(strings.Split(strconv.Itoa(100), ""))
	var arr []int
	arr = append(arr, 1)
	fmt.Println(strconv.Atoi("100"))
	num, _ := strconv.Atoi("100")
	isEqual := num == 100
	fmt.Print(isEqual)
	fmt.Println(len(strings.Split(strconv.Itoa(100), "")))

	fmt.Println(UniqueSortedUserIDs([]int64{22, 22, 22}))
	fmt.Println(MostPopularWord([]string{"22", "22", "33", "pop", "pop"}))
}

func UniqueSortedUserIDs(userIDs []int64) []int64 {
  sort.Slice(userIDs, func(i, j int) bool {
    return userIDs[i] < userIDs[j]
  })
  
  j := 0
	for i := 1; i < len(userIDs); i++ {
		if userIDs[j] != userIDs[i] {
			j++
			userIDs[j] = userIDs[i]
		}
	}

	return userIDs[:j+1]
}

type Element struct {
	num  int
	str string
}

func MostPopularWord(words []string) string {
  dict := make(map[string]int, 0)
  for i:=0; i<len(words); i++{
    dict[words[i]] += 1 
  }

	var mostPop [2]Element

	for key, value := range dict{
		if value > mostPop[0].num {
			mostPop[0].num = value
			mostPop[1].str = key
		}
	}

  fmt.Println(dict)
	return mostPop[1].str
}

// Тоже самое но проще и в одну итерацию
func MostPopularWord2(words []string) string {
	wordsCount := make(map[string]int, 0)
	mostPopWord := ""
	max := 0

	for i := len(words) - 1; i >= 0; i -= 1 {
		word := words[i]
		wordsCount[word] += 1
		if wordsCount[word] >= max {
			max = wordsCount[word]
			mostPopWord = word
		}
	}

	return mostPopWord
}
