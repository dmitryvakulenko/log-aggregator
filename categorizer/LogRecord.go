package categorizer

import (
	"github.com/globalsign/mgo/bson"
	"log-aggregator/simhash"
)

type Category struct {
	Id   bson.ObjectId `bson:"_id"`
	Hash simhash.SimHash
}

type LogRecord struct {
	Id            bson.ObjectId `bson:"_id"`
	Type          int           `json:"type"`
	Date          string        `json:"date"`
	Message       string        `json:"message"`
	Trace         []string      `json:"trace,omitempty"`
	Uri           string        `json:"uri"`
	Referer       string        `json:"referer"`
	UserAgent     string        `json:"userAgent"`
	Tags          string        `json:"tags"`
	User          string        `json:"user"`
	OrderId       string        `json:"orderId"`
	IpAddress     string        `json:"ipAddress"`
	RequestParams string        `json:"requestParams"`
	CategoryId    bson.ObjectId `bson:"category_id"`
}

func (r *LogRecord) SimHash() simhash.SimHash {
	resText := r.Message + "\n" + r.Uri
	for _, v := range r.Trace {
		resText += "\n" + v
	}
	return simhash.Calculate(resText)
}
