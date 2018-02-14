package categorizer

import (
	"log-aggregator/simhash"
	"time"
)

type Categorized interface {
	SaveCategory(*Category)
	GetCategories() []Category
	GetLogRecords(*Category) []LogRecord
	AddLogRecord(*LogRecord)
}

type Categorizer struct {
	Storage Categorized
}


func (c *Categorizer) AddRecord(rec *LogRecord) {
	hash := rec.SimHash()
	var minCat *Category = nil
	var diff float32 = -1
	for _, cat := range c.Storage.GetCategories() {
		newDiff := simhash.Difference(cat.Hash, hash)
		if diff < newDiff {
			continue
		}

		diff = newDiff
		minCat = &cat
	}

	var resCat *Category = nil
	if diff == -1 || diff > 0.1 {
		resCat = &Category{"", time.Now(), hash}
	} else {
		resCat = minCat
		resCat.Updated = time.Now()
	}

	c.Storage.SaveCategory(resCat)
	rec.CategoryId = resCat.Id
	c.Storage.AddLogRecord(rec)
}
