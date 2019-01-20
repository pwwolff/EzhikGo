package store

import "gopkg.in/mgo.v2/bson"

type Word struct {
	ID             bson.ObjectId `bson:"_id"`
	Word_Id        uint64        `json:"word_id"`
	Query_Count    uint32        `json:"query_count"`
	Word_Name      string        `json:"word_name"`
	Type           string        `json:"type"`
	Translation_En string        `json:"translation_en"`
	Translation_De string        `json:"translation_de"`
	Forms          bson.M        `json:"forms"`
}

type AccentPair struct {
	ID             bson.ObjectId `bson:"_id"`
	Word_Id        uint64        `json:"word_id"`
	Query_Count    uint32        `json:"query_count"`
	Unstressed     string        `json:"unstressed"`
	Stressed       string        `json:"stressed"`
	Word_Name      string        `json:"word_name"`
	Form           string        `json:"form"`
	Type           string        `json:"type"`
	Translation_En string        `json:"translation_en"`
}

type RussianText struct {
	ID        bson.ObjectId `bson:"_id"`
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
}

type TextWord struct {
	Index    int32  `json:"wordIndex"`
	WordText string `json:"wordText"`
	WordId   int32  `json:"wordId"`
	Tail     string `json:"tail"`
}

type RussianTexts []RussianText
type AccentPairs []AccentPair
type Words []Word
