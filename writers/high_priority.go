package writers

import (
	"github.com/andrewhoff/logr/config"
)

// HighPriorityLogWriter - this type of LogWriter will always try to log messages with the highest priority, automatically and specified priority needed
type HighPriorityLogWriter struct {
	*LockedPriorityLogWriter
}

// NewHighPriorityLogWriter really just a special LockedPriorityWriter
func NewHighPriorityLogWriter() *HighPriorityLogWriter {
	return &HighPriorityLogWriter{
		&LockedPriorityLogWriter{
			Priority: config.HighPriority,
		},
	}
}

// Log - this implementation will always try to log messages with the initialized priority
func (lw *HighPriorityLogWriter) Log(val string) error {
	return lw.LockedPriorityLogWriter.Log(val)
}
