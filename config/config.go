package config

const (
	// LowPriority - Not that important
	LowPriority = iota + 1

	// MedPriority - Kind of important
	MedPriority

	// HighPriority - Really important
	HighPriority
)

// NumPriorities - the number of configured logging message priority levels
const NumPriorities = 3

// DefaultCapacity - // #items, it would be cool to set this based on item size..so like setting the capacity to the amount of memory it can use then determining the numerical item capacity that way (1 MB capacity/ITEM_SIZE) = MAX_LENGTH
const DefaultCapacity = 256

// opts ...
var opts *Opts

// Opts ...
type Opts struct {
	Capacity int
}

// SetOpts ...
func SetOpts(in *Opts) {
	opts = in
}

// Capacity ...
func Capacity() int {
	if opts.Capacity == 0 {
		return DefaultCapacity
	}

	return opts.Capacity
}
