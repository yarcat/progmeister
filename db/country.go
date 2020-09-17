package main

type CountryID int

type Country struct {
	ID   CountryID // Autoincremented.
	Code string
	Name string
}

type CountryDB struct {
	records []Country
}

func (db *CountryDB) Insert(record Country) {
	record.ID = CountryID(len(db.records) + 1)
	db.records = append(db.records, record)
}

func (db *CountryDB) InsertMany(records []Country) {
	for _, record := range records {
		db.Insert(record)
	}
}

func (db *CountryDB) Select(pred CountryPredFunc) *CountryCursor {
	return &CountryCursor{db: db, pred: pred, index: -1}
}

type CountryCursor struct {
	db    *CountryDB
	pred  CountryPredFunc
	index int
}

func (it *CountryCursor) Next() bool {
	for {
		it.index++
		if it.index >= len(it.db.records) {
			return false
		}
		if it.pred(it.Value()) {
			return true
		}
	}
}

func (it CountryCursor) Value() Country {
	return it.db.records[it.index]
}

func (it *CountryCursor) Must() Country {
	if !it.Next() {
		panic("Not found")
	}
	return it.Value()
}

type CountryPredFunc func(Country) bool

func AnyCountry(Country) bool { return true }