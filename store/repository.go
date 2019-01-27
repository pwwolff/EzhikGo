package store

import (
	"fmt"
	"log"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pwwolff/EzhikGo/config"
)

var session *mgo.Session

type Repository struct{}

func GetConnection() *mgo.Database {

	conf := config.GetConfig()
	if session == nil {
		SERVER := "mongodb://" + conf.DataBaseAddr
		CRED := &mgo.Credential{Username: conf.Username, Password: conf.Password}
		session, err := mgo.Dial(SERVER)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		session.Login(CRED)
		return session.DB(conf.DataBaseName)
	}

	newSession := session.Copy()
	return newSession.DB(conf.DataBaseName)
}

func (r Repository) GetStressed(unstressed string) AccentPairs {
	const COLLECTION = "AccentPairs"
	db := GetConnection()
	defer db.Session.Close()
	c := db.C(COLLECTION)

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
	db := GetConnection()
	defer db.Session.Close()
	c := db.C(COLLECTION)

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
	db := GetConnection()
	defer db.Session.Close()

	c := db.C(COLLECTION)
	results := RussianTexts{}

	if err := c.Find(bson.M{"urlTitle": textName}).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}
