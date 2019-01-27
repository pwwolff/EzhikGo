package store

import (
	"fmt"

	"github.com/pwwolff/EzhikGo/config"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct{}

func GetConnection() *mgo.Session {
	conf := config.GetConfig()
	SERVER := "mongodb://" + conf.DataBaseAddr
	CRED := &mgo.Credential{Username: conf.Username, Password: conf.Password}
	session, err := mgo.Dial(SERVER)
	session.Login(CRED)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	return session
}

func (r Repository) GetStressed(unstressed string) AccentPairs {
	const COLLECTION = "AccentPairs"
	DBNAME := "Russian"
	session := GetConnection()
	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)

	results := AccentPairs{}

	if err := c.Find(bson.M{"unstressed": unstressed}).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	} else {
		c.UpdateAll(
			bson.M{"unstressed": unstressed},
			bson.M{"$inc": bson.M{"query_count": 1}},
		)
	}

	return results
}

func (r Repository) GetWordByID(id int) Word {
	const COLLECTION = "Words"
	DBNAME := "Russian"
	session := GetConnection()

	c := session.DB(DBNAME).C(COLLECTION)
	var result Word

	if err := c.Find(bson.M{"word_id": id}).One(&result); err != nil {
		fmt.Println("Failed to write results:", err)
	} else {
		c.Upsert(
			bson.M{"word_id": id},
			bson.M{"$inc": bson.M{"query_count": 1}},
		)
	}

	return result
}

func (r Repository) GetText(textName string) RussianTexts {
	const COLLECTION = "Texts"
	DBNAME := "Russian"
	session := GetConnection()

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	results := RussianTexts{}

	if err := c.Find(bson.M{"urlTitle": textName}).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}
