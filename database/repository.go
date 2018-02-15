package database

import (
	"github.com/globalsign/mgo"
	"log-aggregator/categorizer"
	"github.com/globalsign/mgo/bson"
	"time"
	"log"
)

const subDayDuration = -24 * time.Hour

type Repository struct {
	session       *mgo.Session
	catCollection *mgo.Collection
	recCollection *mgo.Collection
}

func (l *Repository) Connect() {
	var err error
	l.session, err = mgo.Dial("localhost:27017/uts24")
	if err != nil {
		log.Fatalf("Error database connection %v", err)
	}
	l.catCollection = l.session.DB("").C("log_categories")
	l.recCollection = l.session.DB("").C("log_records")
}

func (l *Repository) Disconnect() {
	l.session.Close()
}

func (l *Repository) SaveCategory(category *categorizer.Category) {
	var err error
	if len(category.Id) == 0 {
		category.Id = bson.NewObjectId()
		err = l.catCollection.Insert(category)
	} else {
		err = l.catCollection.UpdateId(category.Id, category)
	}

	if err != nil {
		log.Printf("Error saving category %v", err)
	}
}

func (l *Repository) GetCategories() []categorizer.Category {
	var tmpRes []categorizer.Category
	startDate := time.Now().Add(subDayDuration)
	err := l.catCollection.Find(bson.M{"updated": bson.M{"$gt": startDate}}).All(&tmpRes)
	if  err != nil {
		log.Fatalf("Error getting categories list %v", err)
	}

	return tmpRes
}


func (l *Repository) AddLogRecord(rec *categorizer.LogRecord) {
	rec.Id = bson.NewObjectId()
	err := l.recCollection.Insert(rec)
	if err != nil {
		log.Printf("Error saving log record %v", err)
	}
}
