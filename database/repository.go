package database

import (
	"github.com/globalsign/mgo"
	"log-aggregator/categorizer"
	"github.com/globalsign/mgo/bson"
)

type Repository struct {
	session       *mgo.Session
	catCollection *mgo.Collection
}

func (l *Repository) Connect() {
	var err error
	l.session, err = mgo.Dial("localhost:27017/uts24")
	if err != nil {
		panic(err)
	}
	l.catCollection = l.session.DB("").C("log_categories")
}

func (l *Repository) Disconnect() {
	l.session.Close()
}

func (l *Repository) CreateCategory(category *categorizer.Category) {
	category.Id = bson.NewObjectId()
	err := l.catCollection.Insert(category)
	if err != nil {
		panic(err)
	}
}

func (l *Repository) GetCategories() []categorizer.Category {
	var tmpRes []categorizer.Category
	err := l.catCollection.Find(nil).All(&tmpRes)
	if  err != nil {
		panic(err)
	}

	return tmpRes
}


func (l *Repository) GetLogRecords(category *categorizer.Category) []categorizer.LogRecord {
	return nil
}

func (l *Repository) AddLogRecord(rec *categorizer.LogRecord) {

}
