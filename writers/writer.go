package writers

import (
	"github.com/andrewhoff/logr/store"
)

// LogWriter ...
type LogWriter struct {
}

// NewLogWriter ...
func NewLogWriter() *LogWriter {
	return &LogWriter{}
}

// Log ...
func (lw *LogWriter) Log(priority int, val string) error {
	store.InternalDataStore.Enqueue(&store.Item{
		Priority: priority,
		Value:    val,
	})

	return nil
}
