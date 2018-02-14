package categorizer

import (
	"github.com/globalsign/mgo/bson"
	"log-aggregator/simhash"
	"strconv"
	"time"
)

type Category struct {
	Id      bson.ObjectId `bson:"_id"`
	Updated time.Time
	Hash    simhash.SimHash
}

type LogRecord struct {
	Id            bson.ObjectId `bson:"_id"`
	Type          int           `json:"type"`
	Timestamp     time.Time     `json:"timestamp"`
	Message       string        `json:"message"`
	File          string        `json:"file"`
	Line          int           `json:"line"`
	Trace         []string      `json:"trace,omitempty"`
	Uri           string        `json:"uri"`
	Referer       string        `json:"referer"`
	UserAgent     string        `json:"userAgent"`
	Tags          string        `json:"tags"`
	User          string        `json:"user"`
	OrderId       string        `json:"order"`
	IpAddress     string        `json:"address"`
	Context       string        `json:"context"`
	RequestParams string        `json:"requestParams"`
	CategoryId    bson.ObjectId `bson:"category_id"`
}

func (r *LogRecord) SimHash() simhash.SimHash {
	resText := r.Message + "\n" + r.File + "\n" + strconv.Itoa(r.Line)
	for _, v := range r.Trace {
		resText += "\n" + v
	}
	return simhash.Calculate(resText)
}
