package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {

	startA := MeasureResponseTime(a)

	startB := MeasureResponseTime(b)

	if startA > startB {
		return b
	}

	return a
}

func MeasureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
