package hdu_cas_helper

type LoginStatus struct {
	method  string
	tgc     string
	expired bool
	service string
	err     error
}

func (s *LoginStatus) Error() error {
	return s.err
}
