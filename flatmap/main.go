package main

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/numaproj/numaflow-go/pkg/mapper"
)

func mapFn(_ context.Context, keys []string, d mapper.Datum) mapper.Messages {
	msg := d.Value()
	_ = d.EventTime() // Event time is available
	_ = d.Watermark() // Watermark is available
	// Split the msg into an array with comma.
	wordMap := make(map[string]int)
	results := mapper.MessagesBuilder()
	for _, s := range strings.Split(string(msg), " ") {
		wordMap[s]++
	}
	for k, v := range wordMap {
		results = results.Append(mapper.NewMessage([]byte(strconv.Itoa(v))).WithKeys([]string{k}))
	}
	return results
}

func main() {
	err := mapper.NewServer(mapper.MapperFunc(mapFn)).Start(context.Background())
	if err != nil {
		log.Panic("Failed to start map function server: ", err)
	}
}
