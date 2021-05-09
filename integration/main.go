package main

import (
	"github.com/noguchidaisuke/go-table-testing/integration/repository"
	"log"
	"strconv"
)

// 技術と実装を分離 entityには依存。
// databaseを一番外, handlerを中, entityを中心に置けば疎結合にできる
type Repository interface {
	Save(title, body string) error
	Get(title string) ([]repository.SomeData, error)
}

func calcurate(repo Repository, keys []string) (int, error) {
	sum := 0
	var values []repository.SomeData
	for _, key := range keys {
		data, err := repo.Get(key)
		if err != nil {
			return 0, err
		}
		values = append(values, data...)
	}

	for _, v := range values {
		i, err := strconv.Atoi(v.Body)
		if err != nil {
			return 0, err
		}
		log.Printf("[INFO] sum = sum:%v + i:%v", sum, i)
		sum = sum + i
	}
	return sum, nil
}

func main() {
	repo, err := repository.NewDatabase("root", "mysql", "127.0.0.1", "3306", "todo", "?parseTime=true&loc=Japan")
	if err != nil {
		log.Fatal(err)
	}

	err = repo.Save("someTitle", "sample body")
	if err != nil {
		log.Fatal(err)
	}

	var data []repository.SomeData
	data, err = repo.Get("someTitle")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[INFO] num:%v, result = %v", len(data), data)

	err = repo.Save("t1", "100")
	if err != nil {
		log.Fatal(err)
	}
	err = repo.Save("t2", "200")
	if err != nil {
		log.Fatal(err)
	}
	sum, err := calcurate(repo, []string{"t1", "t2"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[INFO] sum:%v", sum)
}