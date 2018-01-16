package logr

import (
	"github.com/andrewhoff/logr/config"
	"github.com/andrewhoff/logr/reader"
	"github.com/andrewhoff/logr/store"
	"github.com/andrewhoff/logr/writers"
)

const (
	// LowPriority ...
	LowPriority = config.LowPriority

	// MedPriority ...
	MedPriority = config.MedPriority

	// HighPriority ...
	HighPriority = config.HighPriority
)

// InitWithOpts ...
func InitWithOpts(opts *config.Opts) {
	store.Init()
	config.SetOpts(opts)
}

// NewLogReader ...
func NewLogReader() (*reader.LogReader, error) {
	return reader.NewLogReader(), nil
}

// NewGenericLogWriter ...
func NewGenericLogWriter() (*writers.GenericLogWriter, error) {
	return writers.NewGenericLogWriter(), nil
}

// NewLockedPriorityLogWriter ...
func NewLockedPriorityLogWriter(priority int) (*writers.LockedPriorityLogWriter, error) {
	return writers.NewLockedPriorityLogWriter(priority), nil
}

// NewHighPriorityLogWriter ...
func NewHighPriorityLogWriter() (*writers.HighPriorityLogWriter, error) {
	return writers.NewHighPriorityLogWriter(), nil
}

// NewLazyLogWriter ...
func NewLazyLogWriter(priority int) (*writers.LazyLogWriter, error) {
	return writers.NewLazyLogWriter(priority), nil
}
