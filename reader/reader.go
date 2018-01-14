package reader

import (
	"log"

	"github.com/andrewhoff/logr/store"
)

// LogReader ...
type LogReader struct {
}

// NewLogReader ...
func NewLogReader() *LogReader {
	return &LogReader{}
}

// Get ...
func (reader *LogReader) Get() string {
	if store.InternalDataStore == nil {
		log.Println("[WARNING] logr hasn't finished initializing")
		return ""
	}

	if store.InternalDataStore.Empty() {
		log.Println("[WARNING] log is empty")
		return ""
	}

	dequeued := store.InternalDataStore.Dequeue()
	if dequeued == nil {
		return ""
	}

	return dequeued.Value
}
