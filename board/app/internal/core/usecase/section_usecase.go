package usecase

import (
	"board/internal/core/domain"
	"board/internal/core/dto"
	"board/internal/core/port/driven"
	"board/pkg/auth"
	"fmt"

	"github.com/google/uuid"
)

type SectionUsecase interface {
	CreateSection(dto.CreateSectionRequest) (*dto.CreateSectionResponse, error)
	DeleteSection(dto.DeleteSectionRequest) error
	UpdateSection(dto.UpdateSectionRequest) (*dto.UpdateSectionResponse, error)
}

type sectionUsecase struct {
	sectionDao driven.SectionPortDriven
}

func NewSectionUsecase(sectionDao driven.SectionPortDriven) SectionUsecase {
	return &sectionUsecase{
		sectionDao: sectionDao,
	}
}

func (s *sectionUsecase) CreateSection(section dto.CreateSectionRequest) (*dto.CreateSectionResponse, error) {

	id, err := auth.IsCustomer(section.AccessToken)

	if err != nil {
		return nil, err
	}

	//check permission
	fmt.Println(id)

	return &dto.CreateSectionResponse{
		ID:          uuid.New(),
		AccessToken: "token",
		BoardID:     uuid.New(),
		Title:       "title",
		Order:       34,
	}, nil
}

func (s *sectionUsecase) DeleteSection(section dto.DeleteSectionRequest) error {

	id, err := auth.IsCustomer(section.AccessToken)

	if err != nil {
		return err
	}

	//check permission
	fmt.Println(id)

	return nil
}

func (s *sectionUsecase) UpdateSection(section dto.UpdateSectionRequest) (*dto.UpdateSectionResponse, error) {

	id, err := auth.IsCustomer(section.AccessToken)

	if err != nil {
		return nil, err
	}

	//check permission
	fmt.Println(id)

	sectionData, err := s.sectionDao.Create(domain.Section{
		ID:      section.ID,
		BoardID: section.BoardID,
		Title:   section.Title,
		Order:   section.Order,
	})

	if err != nil {
		return nil, err
	}

	return &dto.UpdateSectionResponse{
		ID:      sectionData.ID,
		BoardID: sectionData.BoardID,
		Title:   section.Title,
		Order:   section.Order,
	}, nil

}
