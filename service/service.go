package service

import (
	"errors"
	"strconv"
)

// Service implements the business logic.
type Service interface {
	Single(number int) (string, error)
}

type service struct{}

// New returns an empty service, ready to be used.
func New() Service {
	return &service{}
}

func (s *service) Single(number int) (string, error) {
	if number < 1 {
		return "", errors.New("number must be greater than 0")
	}

	var res string
	if number%15 == 0 {
		res = "fizzbuzz"
	} else if number%5 == 0 {
		res = "buzz"
	} else if number%3 == 0 {
		res = "fizz"
	} else {
		res = strconv.Itoa(number)
	}

	return res, nil
}
