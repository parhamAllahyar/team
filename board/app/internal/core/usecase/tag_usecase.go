package usecase

import (
	"board/internal/core/domain"
	"board/internal/core/dto"
	"board/internal/core/port/driven"
	"board/pkg/auth"
	"fmt"

	"github.com/google/uuid"
)

type TagUsecase interface {
	CreateTag(dto.CreateTagRequest) (*dto.CreateTagResponse, error)
	DeleteTag(dto.DeleteTagRequest) error
	UpdateTag(dto.UpdateTagRequest) (*dto.UpdateTagResponse, error)
}

type tagUsecase struct {
	tagDao driven.TagPortDriven
}

func NewTagUsecase(tagDao driven.TagPortDriven) TagUsecase {
	return &tagUsecase{
		tagDao: tagDao,
	}
}

func (t *tagUsecase) CreateTag(tag dto.CreateTagRequest) (*dto.CreateTagResponse, error) {

	id, err := auth.IsCustomer(tag.AccessToken)

	if err != nil {
		return nil, err
	}

	//check permission
	fmt.Println(id)

	//TODO: check board id is valid

	//TODO: Validation

	newTag, err := t.tagDao.Create(domain.Tag{
		ID:      uuid.New(),
		Title:   tag.Title,
		Pattern: tag.Pattern,
		Order:   tag.Order,
		BoardID: tag.BoardID,
	})

	if err != nil {
		return nil, err
	}

	return &dto.CreateTagResponse{
		ID:      newTag.ID,
		BoardID: newTag.BoardID,
		Title:   newTag.Title,
		Pattern: newTag.Pattern,
		Order:   newTag.Order,
	}, nil
}
func (t *tagUsecase) DeleteTag(tag dto.DeleteTagRequest) error {

	id, err := auth.IsCustomer(tag.AccessToken)

	if err != nil {
		return err
	}

	//TODO: check permission
	fmt.Println(id)

	err = t.tagDao.Delete(tag.ID)

	if err != nil {
		return err
	}

	return nil
}
func (t *tagUsecase) UpdateTag(tag dto.UpdateTagRequest) (*dto.UpdateTagResponse, error) {

	id, err := auth.IsCustomer(tag.AccessToken)

	if err != nil {
		return nil, err
	}

	//check permission
	fmt.Println(id)

	tagData, err := t.tagDao.Update(domain.Tag{
		ID:      tag.ID,
		Title:   tag.Title,
		BoardID: tag.BoardID,
		Order:   tag.Order,
	})

	if err != nil {
		return nil, err
	}

	return &dto.UpdateTagResponse{
		ID:      tagData.ID,
		BoardID: tagData.BoardID,
		Title:   tagData.Title,
		Pattern: tagData.Pattern,
		Order:   tagData.Order,
	}, nil

}
