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

func init() {
	store.Init()
}

// InitWithOpts ...
func InitWithOpts(opts *config.Opts) {
	config.SetOpts(opts)
}

// NewLogReader ...
func NewLogReader() (*reader.LogReader, error) {
	return reader.NewLogReader(), nil
}

// NewLogWriter ...
func NewLogWriter() (*writers.LogWriter, error) {
	return writers.NewLogWriter(), nil
}
