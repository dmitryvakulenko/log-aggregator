package categorizer

import (
	"log-aggregator/simhash"
	"time"
)

type Categorized interface {
	SaveCategory(*Category)
	GetCategories() []Category
	AddLogRecord(*LogRecord)
}

type Categorizer struct {
	Storage Categorized
}


func (c *Categorizer) AddRecord(rec *LogRecord) {
	hash := rec.SimHash()
	var minCat *Category = nil
	var diff float32 = 2 // 2 оно не может быть по определению
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
		resCat = &Category{"", rec.Message, 0, time.Now(), hash}
	} else {
		resCat = minCat
	}

	resCat.MarkUpdated()
	c.Storage.SaveCategory(resCat)
	rec.CategoryId = resCat.Id
	c.Storage.AddLogRecord(rec)
}
