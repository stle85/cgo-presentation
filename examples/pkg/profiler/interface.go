package profiler

// Profile is profile interface.
type Profile interface {
	Stop()
}

// ProfileFunc type of profile function.
type ProfileFunc func() func(*profile)
