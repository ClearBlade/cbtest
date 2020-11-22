package merge

// defaultIgnore contains all the folders that are gonna be ignored during
// the merging process.
var defaultIgnore = []string{
	"node_modules",
}

// Options struct specifies different merge options that can be passed to the
// MergeWithOptions function.
type Options struct {

	// Skip function should return true for files that should be skipped.
	Skip func(string) (bool, error)

	// OnBeforeMerge is a hook that gets called before a folder is merged.
	OnBeforeMerge OnBeforeMerge

	// OnAfterMerge is a hook that gets called after a folder is merged.
	OnAfterMerge OnAfterMerge
}

// DefaultOptions returns the default merge options.
func DefaultOptions() *Options {
	return &Options{
		Skip:          MakeSkip(defaultIgnore),
		OnBeforeMerge: nil,
		OnAfterMerge:  nil,
	}
}
