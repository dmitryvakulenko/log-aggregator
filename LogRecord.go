package main

import "ErrorCollector/simhash"

type LogRecord struct {
	Type int `json:"type"`
	Date string `json:"date"`
	Message string `json:"message"`
	Trace map[string]string `json:"trace,omitempty"`
	Uri string `json:"uri"`
	Referer string `json:"referer"`
	UserAgent string `json:"userAgent"`
	Tags string `json:"tags"`
	User string `json:"user"`
	OrderId string `json:"orderId"`
	IpAddress string `json:"ipAddress"`
	RequestParams string `json:"requestParams"`
}

func (r *LogRecord) SimHash() simhash.SimHash {
	resText := r.Message + "\n" + r.Uri
	for _, v := range r.Trace {
		resText += "\n" + v
	}
	return simhash.Calculate(resText)
}