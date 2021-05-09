package main

import (
	"github.com/noguchidaisuke/go-table-testing/integration/repository"
	"strings"
	"testing"
)

func TestCalculate(t *testing.T) {
	trepo := repoMock{}
	want := 300

	got, err := calcurate(trepo, []string{"t_100", "t_200"})
	if err != nil {
		t.Fatal("calculate() fails err:", err)
	}
	if got != want {
		t.Fatalf("want %v, but %v", want, got)
	}

}

type repoMock struct{}

func (r repoMock) Get(title string) ([]repository.SomeData, error) {
	ary := strings.Split(title, "_")
	if len(ary) == 0 {
		return []repository.SomeData{}, nil
	}
	return []repository.SomeData{
		{Body:      ary[1]},
	}, nil
}

func (r repoMock) Save(title, body string) error {
	return nil
}