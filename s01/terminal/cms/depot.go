package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path"
	"strconv"

	"github.com/pkg/errors"
)

type item struct {
	id       uint64
	name     string
	quantity uint32
}

func dbFilename() string {
	const fn = "depot.csv"
	if _, err := os.Stat(fn); err == nil {
		return fn
	}
	relative := path.Join("s01", "terminal", "cms", fn)
	if _, err := os.Stat(relative); err == nil {
		return relative
	}
	log.Fatalln("dbFilename:", fn, "not found")
	return ""
}

func itemsFromCSV(r io.Reader) ([]item, error) {
	rows, err := csv.NewReader(r).ReadAll()
	if err != nil {
		return nil, err
	}
	var items []item
	for _, row := range rows {
		item, err := itemFromCSVRow(row)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func itemFromCSVRow(row []string) (item, error) {
	if len(row) != 3 {
		return item{}, errors.Errorf("wrong column count: len(row)=%d, want=3", len(row))
	}
	id, err := strconv.ParseUint(row[0] /* base */, 10 /* bits */, 64)
	if err != nil {
		return item{}, err
	}
	name := row[1]
	quantity, err := strconv.ParseUint(row[2] /* base */, 10 /* bits */, 32)
	if err != nil {
		return item{}, err
	}
	return item{id: id, name: name, quantity: uint32(quantity)}, nil
}

func loadDepotItems() []item {
	f, err := os.Open(dbFilename())
	if err != nil {
		log.Fatalln("Unable to open db:", err)
	}

	items, err := itemsFromCSV(f)
	if err != nil {
		log.Fatalln("Failed to load CSV:", err)
	}

	return items
}
