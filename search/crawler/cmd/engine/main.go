package main

import (
	"crawler/pkg/spider"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	const url1 = "https://www.opennet.ru"
	const url2 = "https://habr.com/ru/"
	var word string
	var u []string
	var data map[string]string
	u = append(u, url2) // очень долго с habr
	u = append(u, url1) // как это сделать нормально?
	if len(os.Args) == 1 {
		word = "" // если ничего не ввели, то выводим все
	} else {
		word = os.Args[1]
	}
	for _, url := range u {
		var err error
		data, err = spider.Scan(url, 2)
		if err != nil {
			log.Printf("ошибка при сканировании сайта %s: %v\n", url, err)
			continue
		}
	}
	for {
		fmt.Print("Введите слово для поиска: ")
		fmt.Fscan(os.Stdin, &word)
		for url, article := range data {
			if strings.Contains(strings.ToLower(url+article), strings.ToLower(word)) { // подходит ли это для решения?
				fmt.Printf("Результат сканирования :%s\n", url+data[url]) // поиск по URL
			}
		}

	}
}
