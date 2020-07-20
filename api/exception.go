package api

// DeleteExcepction to represent excaptions caused by delete action.
type DeleteExcepction interface {
	ErrorOf(int) string
}
