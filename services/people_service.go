package services

import (
	"errors"
	"golang_mvc_starterpack/dto"
	"golang_mvc_starterpack/models"
	"golang_mvc_starterpack/repositories"
	"strconv"
	"time"
)

type PeopleService interface {
	Store(p *dto.PeopleCreate) error
	Index() ([]models.People, error)
	Show(id string) (models.People, error)
	Update(id uint, patch dto.PeoplePatch) error
}

type peopleService struct {
	repo repositories.PeopleRepository
}

func NewPeopleService(repo repositories.PeopleRepository) PeopleService {
	return &peopleService{repo}
}

func (s *peopleService) Store(p *dto.PeopleCreate) error {
	date, err := time.Parse("2006-01-02", p.DateOfBirth)
	if err != nil {
		return errors.New("format tanggal salah, gunakan YYYY-MM-DD")
	}

	if date.After(time.Now()) {
		return errors.New("tanggal lahir tidak boleh lebih dari hari ini")
	}

	people := &models.People{
		FirstName:    p.FirstName,
		LastName:     p.LastName,
		PlaceOfBirth: p.PlaceOfBirth,
		DateOfBirth:  date,
		Address:      p.Address,
	}

	return s.repo.Store(people)
}

func (s *peopleService) Index() ([]models.People, error) {
	return s.repo.Index()
}

func (s *peopleService) Show(id string) (models.People, error) {
	return s.repo.Show(id)
}

func (s *peopleService) Update(id uint, patch dto.PeoplePatch) error {
	// cek apakah data ada
	existing, err := s.repo.Show(strconv.FormatUint(uint64(id), 10))
	if err != nil {
		return err
	}

	// validasi tanggal lahir hanya kalau dikirim
	if patch.DateOfBirth != nil {
		date, err := time.Parse("2006-01-02", *patch.DateOfBirth)
		if err != nil {
			return errors.New("format tanggal salah, gunakan YYYY-MM-DD")
		}
		if date.After(time.Now()) {
			return errors.New("tanggal lahir tidak boleh lebih dari hari ini")
		}
		existing.DateOfBirth = date
	}

	if patch.FirstName != nil {
		existing.FirstName = *patch.FirstName
	}
	if patch.LastName != nil {
		existing.LastName = *patch.LastName
	}
	if patch.PlaceOfBirth != nil {
		existing.PlaceOfBirth = *patch.PlaceOfBirth
	}
	if patch.Address != nil {
		existing.Address = *patch.Address
	}

	existing.UpdatedAt = time.Now()

	return s.repo.Update(existing)
}
