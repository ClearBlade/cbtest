package merge

// OnBeforeMerge is a hook that gets called before we merge a folder into the
// target path.
type OnBeforeMerge interface {
	OnBeforeMerge(targetPath, sourcePath string) error
}

// OnAfterMerge is a hook that gets called after we marge a folder into the
// target path.
type OnAfterMerge interface {
	OnAfterMerge(targetpath, sourcePath string) error
}
