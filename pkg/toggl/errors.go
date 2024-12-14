package toggl

func (s *APIErrorString) Error() string {
	return string(*s)
}
