package categorizer

import "log-aggregator/simhash"

type Categorized interface {
	CreateCategory(*Category)
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
		if diff == -1 {
			diff = newDiff
		} else if newDiff < diff {
			diff = newDiff
			minCat = &cat
		}
	}

	if diff > 0.1 {
		newCategory := &Category{"", hash}
		c.Storage.CreateCategory(newCategory)
		rec.CategoryId = newCategory.Id
	} else {
		rec.CategoryId = minCat.Id
	}

	c.Storage.AddLogRecord(rec)
}
