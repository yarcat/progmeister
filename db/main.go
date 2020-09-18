package main

import (
	"fmt"
	"strings"
)

func main() {
	countries := CountryDB{}
	countries.InsertMany([]Country{
		{Code: "LV", Name: "Latvia"},
		{Code: "CH", Name: "Switzerland"},
		{Code: "RU", Name: "Russia"},
	})

	var (
		code = func(code string) CountryPredFunc {
			return func(c Country) bool { return c.Code == code }
		}
		switzerland = countries.Select(code("CH")).Must()
		latvia      = countries.Select(code("LV")).Must()
		russia      = countries.Select(code("RU")).Must()
	)
	people := PersonDB{}
	people.InsertMany([]Person{
		{FirstName: "Jaroslavs", LastName: "Samcuks", CountryID: switzerland.ID},
		{FirstName: "Pavels", LastName: "Zaicenkovs", CountryID: switzerland.ID},
		{FirstName: "Sergejs", LastName: "Melniks", CountryID: latvia.ID},
		{FirstName: "Julija", LastName: "Pecerska", CountryID: switzerland.ID},
		{FirstName: "Aleksejs", LastName: "Lucko", CountryID: russia.ID},
	})
	fmt.Println()
	fmt.Println("### SELECT * FROM people")
	fmt.Println()
	for c := people.Select(AnyPerson); c.Next(); {
		fmt.Println(c.Value())
	}
	fmt.Println()
	fmt.Println(`### SELECT * FROM people WHERE FirstName LIKE "J%"`)
	fmt.Println()
	startsWithJ := func(p Person) bool { return strings.HasPrefix(p.FirstName, "J") }
	for c := people.Select(startsWithJ); c.Next(); {
		fmt.Println(c.Value())
	}
	fmt.Println()
	fmt.Println("### SELECT p.FirstName, p.LastName, c.Name")
	fmt.Println("### FROM people AS p JOIN countries AS c")
	fmt.Println("### ON (p.CountryID = c.ID)")
	fmt.Println()
	countryID := func(p Person, c Country) bool { return p.CountryID == c.ID }
	joined := PersonCountryJoin{P: &people, C: &countries}.On(countryID)
	for c := joined.Select(AnyPersonCountry); c.Next(); {
		row := c.Value()
		fmt.Println(row.P.FirstName, row.P.LastName, row.C.Name)
	}
}

type PersonCountry struct {
	P Person
	C Country
}

type PersonCountryJoin struct {
	P *PersonDB
	C *CountryDB
}

type PersonCountryJoinPred func(Person, Country) bool

func (join PersonCountryJoin) On(pred PersonCountryJoinPred) *PersonCountryJoinDB {
	return &PersonCountryJoinDB{cross: func() func() *PersonCountry {
		var country *CountryCursor
		person := join.P.Select(AnyPerson)
		nextPerson := true
		return func() *PersonCountry {
			for {
				if nextPerson && !person.Next() {
					return nil
				}
				nextPerson = false
				if country == nil {
					country = join.C.Select(AnyCountry)
				}
				for country.Next() {
					if pred(person.Value(), country.Value()) {
						return &PersonCountry{P: person.Value(), C: country.Value()}
					}
				}
				nextPerson = true
				country = nil
			}
		}
	}}
}

type PersonCountryJoinDB struct {
	cross func() func() *PersonCountry
}

func (db *PersonCountryJoinDB) Select(pred PersonCountryPred) *PersonCountryCursor {
	return &PersonCountryCursor{next: db.cross()}
}

type PersonCountryCursor struct {
	next func() *PersonCountry
	val  *PersonCountry
}

func (c *PersonCountryCursor) Next() bool {
	c.val = c.next()
	return c.val != nil
}

func (c *PersonCountryCursor) Value() PersonCountry { return *c.val }

type PersonCountryPred func(Person, Country) bool

func AnyPersonCountry(Person, Country) bool { return true }
