package main

import (
	"crawler/pkg/spider"
	"fmt"
	"log"
	"os"
	"strings"
)

var word string
var s []string

func main() {
	const url1 = "https://www.opennet.ru"
	const url2 = "https://habr.com/ru/"
	var u []string
	u = append(u, url2) // очень долго с habr
	u = append(u, url1) // как это сделать нормально?
	if len(os.Args) == 1 {
		word = "" // если ничего не ввели, то выводим все
	} else {
		word = os.Args[1]
	}
	for _, url := range u {
		data, err := spider.Scan(url, 2)
		if err != nil {
			log.Printf("ошибка при сканировании сайта %s: %v\n", url, err)
		}
		for k, v := range data {
			s = append(s, k+v) // сохраняем в slice

		}

	}
	for {
		for i := range s {
			if strings.Contains(strings.ToLower(s[i]), strings.ToLower(word)) { // подходит ли это для решения?
				fmt.Printf("Результат сканирования :%s\n", s[i])
			}
		}

		fmt.Print("Введите слово для поиска: ")
		fmt.Fscan(os.Stdin, &word)

	}
}
