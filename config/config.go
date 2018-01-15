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
	// Capacity - How many items in total the logging system can hold
	Capacity int

	// Overwrite - set this to true to, in the event of the logging system becoming completely full, remove the oldest, lowest priority item, in order to make space for the new item
	Overwrite bool
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

// ShouldOverwrite ...
func ShouldOverwrite() bool {
	return opts.Overwrite
}
