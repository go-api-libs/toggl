package toggl

var ErrTimeEntryAlreadyStopped = new(APIErrorString("Time entry already stopped"))

func (s *APIErrorString) Error() string {
	return string(*s)
}

func (s *APIErrorString) Is(err error) bool {
	e, ok := err.(*APIErrorString)
	if !ok {
		return false
	}

	if s == nil || e == nil {
		return s == e // pointer identity: nil==nil → true, else false
	}

	return *s == *e
}
