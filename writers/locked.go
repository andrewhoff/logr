package writers

import (
	"time"

	"github.com/andrewhoff/logr/store"
	"github.com/fatih/color"
)

// LockedPriorityLogWriter - this type of LogWriter will always try to log messages with the highest priority, automatically and specified priority needed
type LockedPriorityLogWriter struct {
	Priority int
}

// NewLockedPriorityLogWriter ...
func NewLockedPriorityLogWriter(priority int) *LockedPriorityLogWriter {
	return &LockedPriorityLogWriter{
		Priority: priority,
	}
}

// Log - this implementation will always try to log messages with the initialized priority
func (lw *LockedPriorityLogWriter) Log(val string) error {
	_, err := store.InternalDataStore.Enqueue(&store.Item{
		Priority: lw.Priority,
		Value:    val,
		DateTime: time.Now(),
	})

	if err == nil {
		color.Green("[LOCKED PRIORITY [%d] LOG WRITER] - Write Succeeded!", lw.Priority)
	}

	return err
}
