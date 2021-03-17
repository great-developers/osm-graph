package store

import (
	"os"

	"github.com/gocarina/gocsv"
)

type Store struct {
	StoreID int     `csv:"store_id"`
	Lat     float64 `csv:"lat"`
	Lng     float64 `csv:"lng"`
}

func FromCSV(path string) []Store {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var result []Store
	if err := gocsv.UnmarshalFile(f, &result); err != nil {
		panic(err)
	}
	for _, store := range result {
		result = append(result, store)
	}
	return result
}
