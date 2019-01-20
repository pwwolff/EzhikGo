package store

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct{}

const SERVER = "mongodb://localhost:27017"

const DBNAME = "Russian"

func (r Repository) GetStressed(unstressed string) AccentPairs {
	session, err := mgo.Dial(SERVER)

	const COLLECTION = "AccentPairs"

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	results := AccentPairs{}

	if err := c.Find(bson.M{"unstressed": unstressed}).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

func (r Repository) GetWordByID(id int) Word {
	session, err := mgo.Dial(SERVER)

	const COLLECTION = "Words"

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	var result Word

	if err := c.Find(bson.M{"word_id": id}).One(&result); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return result
}

func (r Repository) GetText(textName string) RussianTexts {
	session, err := mgo.Dial(SERVER)

	const COLLECTION = "Texts"
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	results := RussianTexts{}

	if err := c.Find(bson.M{"urlTitle": textName}).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}
