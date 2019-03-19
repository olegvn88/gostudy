package main

import (
	"fmt"
	"github.com/olegvn88/gostudy/app/main/app/database/properties"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	sess *mgo.Session
	prop = properties.JsonProp()
)

type student struct {
	ID    bson.ObjectId `json:"id" bson:"_id"`
	Fio   string        `json:"fio" bson:"_fio"`
	Work  string        `json:"work" bson:"_work"`
	Score int           `json:"score" bson:"_score"`
}

func main() {
	var err error
	sess, err = mgo.Dial("mongodb://" + prop.MongodbUser + ":" + prop.MongodbPassword + "@" + prop.MongoHost)
	PanicForErr(err)

	//если коллекции не будет, то она создастся автоматичечски
	collection := sess.DB("tesstbase").C("students")

	index := mgo.Index{
		Key: []string{"fio"},
	}

	err = collection.EnsureIndex(index)
	PanicForErr(err)

	//для mongo нет такого красивого дампа SQL, так тчо я вставляю демо запись, если коллекция пустая
	if n, _ := collection.Count(); n == 0 {
		firstStudent := &student{bson.NewObjectId(), "Oleg Nesterov", "work: softserve", 10}
		err = collection.Insert(firstStudent)
		PanicForErr(err)
	}

	var allStudents []student
	err = collection.Find(bson.M{}).All(&allStudents) // M{} - message .
	PanicForErr(err)
	for i, v := range allStudents {
		fmt.Printf("student[%d]: %+v\n", i, v)
	}

	//Генерим какой-то ИДшник
	id := bson.NewObjectId()
	var nonExistentStudent student
	err = collection.Find(bson.M{"_id": id}).One(&nonExistentStudent)
	PanicForErr(err)

	secondStudent := &student{id, "Dmytro Oleksiyovych", "", 0}
	err = collection.Insert(secondStudent)
	PanicForErr(err)

	err = collection.Find(bson.M{"_id": id}).One(&nonExistentStudent)
	PanicForErr(err)
	fmt.Printf("Seconf student: %+v\n", nonExistentStudent)

	collection.UpdateAll(
		bson.M{"fio": "Dmytro Oleksiyovych"},
		bson.M{
			"$set": bson.M{"info": "all Dmytro's info"},
		},
	)
}

func PanicForErr(e error) {
	if e != nil {
		panic(e)
	}
}
