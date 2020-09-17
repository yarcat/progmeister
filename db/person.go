package main

type PersonID int

type Person struct {
	ID        PersonID // Autoincremented.
	FirstName string
	LastName  string
	CountryID CountryID
}

type PersonDB struct{ records []Person }

func (db *PersonDB) Insert(record Person) {
	record.ID = PersonID(len(db.records) + 1)
	db.records = append(db.records, record)
}

func (db *PersonDB) InsertMany(records []Person) {
	for _, record := range records {
		db.Insert(record)
	}
}

func (db *PersonDB) Select(pred PersonPredFunc) *PersonCursor {
	return &PersonCursor{db: db, pred: pred, index: -1}
}

type PersonCursor struct {
	db    *PersonDB
	pred  PersonPredFunc
	index int
}

func (it *PersonCursor) Next() bool {
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

func (it PersonCursor) Value() Person { return it.db.records[it.index] }

type PersonPredFunc func(Person) bool

func AnyPerson(Person) bool { return true }
