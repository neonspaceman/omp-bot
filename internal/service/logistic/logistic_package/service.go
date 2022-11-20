package logistic_package

import (
	"errors"
)

type Service struct {
	CountOfRow int
}

func NewService(countOfRow int) *Service {
	if countOfRow <= 0 {
		panic("Count of row must be greater than 0")
	}

	return &Service{
		CountOfRow: countOfRow,
	}
}

var ErrIndexOutOfRange = errors.New("Logistic.Package index out of range")

func (s *Service) New(p *Package) {
	allEntities = append(allEntities, p)
}

func (s *Service) Edit(idx int, p *Package) error {
	if err := s.checkCorrectIndex(idx); err != nil {
		return err
	}

	allEntities[idx] = p

	return nil
}

func (s *Service) List(from int) ([]*Package, bool) {
	if from < 0 {
		from = 0
	}

	// start index out of range
	if from >= len(allEntities) {
		return []*Package{}, false
	}

	to := from + s.CountOfRow
	hasNext := true

	if to >= len(allEntities) {
		to = len(allEntities)
		hasNext = false
	}

	return allEntities[from:to], hasNext
}

func (s *Service) Get(idx int) (*Package, error) {
	if err := s.checkCorrectIndex(idx); err != nil {
		return nil, err
	}

	return allEntities[idx], nil
}

func (s *Service) Delete(idx int) error {
	if err := s.checkCorrectIndex(idx); err != nil {
		return err
	}

	copy(allEntities[idx:], allEntities[idx+1:])
	allEntities[len(allEntities)-1] = nil
	allEntities = allEntities[:len(allEntities)-1]

	return nil
}

func (s *Service) checkCorrectIndex(idx int) error {
	if idx < 0 || idx >= len(allEntities) {
		return ErrIndexOutOfRange
	}

	return nil
}
