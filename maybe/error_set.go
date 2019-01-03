package maybe

// This file is a part of github.com/thekondor/confita-sugar package.

type errorSet map[error]struct{}

func newErrorSet(errors ...error) errorSet {
	errorSet := make(map[error]struct{})
	for _, err := range errors {
		errorSet[err] = struct{}{}
	}
	return errorSet
}

func (es errorSet) contains(err error) bool {
	_, ok := es[err]
	return ok
}
