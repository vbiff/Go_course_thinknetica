package main

import (
	"crawler/pkg/spider"
	"fmt"
	"log"
	"os"
	"strings"
)

var word string
var str []string

func main() {
	const url1 = "https://www.opennet.ru"
	const url2 = "https://habr.com/ru/"

	if len(os.Args) == 1 {
		word = "" // если ничего не ввели, то выводим все
	} else {
		word = os.Args[1]
	}
	sets := []spider.Settings{}
	se := spider.Settings{
		URL:   url1,
		Depth: 2,
	}
	sets = append(sets, se)
	se = spider.Settings{
		URL:   url2,
		Depth: 2,
	}
	sets = append(sets, se)
	for _, set := range sets {
		saveScan(&set)
	}
	for {
		for i := range str {
			if strings.Contains(strings.ToLower(str[i]), strings.ToLower(word)) { // подходит ли это для решения?
				fmt.Printf("Результат сканирования :%s\n", str[i])
			}
		}

		fmt.Print("Введите слово для поиска: ")
		fmt.Fscan(os.Stdin, &word)

	}
}

func saveScan(sc spider.Scanner) {
	data, err := sc.Scan()
	if err != nil {
		log.Printf("ошибка при сканировании сайта %v\n", err) //как теперь сюда подставить страницу URL?
	}
	for k, v := range data {
		str = append(str, k+v) // сохраняем в slice

	}

}
