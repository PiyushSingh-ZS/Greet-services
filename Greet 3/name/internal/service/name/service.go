package name

import (
	"fmt"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) Greet(name string) (string, error) {

	return fmt.Sprint("hello, ", name), nil
}
