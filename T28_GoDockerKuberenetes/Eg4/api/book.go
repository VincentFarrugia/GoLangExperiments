package api

import (
	"encoding/json"
)

// Book type with Name, Author and ISBN
type Book struct {
	Name   string
	Author string
	ISBN   string
}

/* Examples of specifying that we want custom names
for our JSON elements. Here we want lowercase names.
type Book struct {
	Name   string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}*/

// ToJSON to be used for marshalling of Book type
func (b Book) ToJSON() []byte {
	jsonByteArr, err := json.Marshal(b)
	panicIfErr(err)
	return jsonByteArr
}

// FromJSON to be used for unmarshalling of Book type
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	panicIfErr(err)
	return book
}

// HELPER FUNCTIONS:

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
