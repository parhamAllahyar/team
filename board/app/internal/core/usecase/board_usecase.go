package usecase

import (
	"board/internal/core/domain"
	"board/internal/core/dto"
	"board/internal/core/port/driven"
	"board/pkg/auth"
	"board/pkg/message"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

)

type BoardUsecase interface {
	CreateBoard(dto.CreateBoardRequest) (*dto.CreateBoardResponse, error)
	DeleteBoard(dto.DeleteBoardRequest) error
	UpdateBoard(dto.UpdateBoardRequest) (*dto.UpdateBoardResponse, error)
}

type boardUsecase struct {
	boardDao driven.BoardPortDriven
}

func NewBoardUsecase(boardDao driven.BoardPortDriven) BoardUsecase {
	return &boardUsecase{
		boardDao: boardDao,
	}
}

func (b *boardUsecase) CreateBoard(board dto.CreateBoardRequest) (*dto.CreateBoardResponse, error) {

	customerId, err := auth.IsCustomer(board.AccessToken)

	if err != nil {
		return nil, err
	}

	//TODO:check permission
	fmt.Println("customer id", customerId)

	//TODO: verifying  workspace id

	//TODO: validation

	boardId := uuid.New()

	//TODO: validation

	newBoard, err := b.boardDao.Create(domain.Board{
		ID:          boardId,
		WorkspaceID: board.WorkspaceID,
		Title:       board.Title,
		CreatedAt:   time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &dto.CreateBoardResponse{
		ID:          newBoard.ID,
		Title:       newBoard.Title,
		WorkspaceID: newBoard.WorkspaceID,
	}, nil

}

func (b *boardUsecase) DeleteBoard(board dto.DeleteBoardRequest) error {

	boardData, err := b.boardDao.GetById(board.ID)

	if err != nil {
		return err
	}

	if boardData == nil {
		return errors.New(message.BoardNotFound)
	}

	id, err := auth.IsCustomer(board.AccessToken)

	if err != nil {
		return err
	}

	//TODO:check permission
	fmt.Println(id)

	//TODO: validation

	err = b.boardDao.Delete(boardData.ID)

	if err != nil {
		return nil
	}

	return nil
}

func (b *boardUsecase) UpdateBoard(board dto.UpdateBoardRequest) (*dto.UpdateBoardResponse, error) {

	boardData, err := b.boardDao.GetById(board.ID)

	if err != nil {
		return nil, err
	}

	if boardData == nil {
		return nil, errors.New(message.BoardNotFound)
	}

	id, err := auth.IsCustomer(board.AccessToken)

	if err != nil {
		return nil, err
	}

	//TODO:check permission
	fmt.Println(id)

	//TODO: validation

	updatedBoard, err := b.boardDao.Update(domain.Board{
		ID:    board.ID,
		Title: board.Title,
	})

	if err != nil {
		return nil, nil
	}

	return &dto.UpdateBoardResponse{
		ID:          updatedBoard.ID,
		Title:       updatedBoard.Title,
		WorkspaceId: updatedBoard.WorkspaceID,
	}, nil

}
