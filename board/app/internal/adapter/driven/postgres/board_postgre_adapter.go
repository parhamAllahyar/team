package postgres

import (
	"board/internal/core/domain"
	"board/internal/core/port/driven"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoardGorm struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;"`
	WorkspaceID uuid.UUID
	Title       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type boardRepo struct {
	boardPort driven.BoardPortDriven
	db        *sql.DB
}

func NewBoardRepo(db *sql.DB) *boardRepo {
	return &boardRepo{
		db: db,
	}
}

func (b *boardRepo) Create(board domain.Board) (*domain.Board, error) {
	query := callGorm(b.db).Table("boards").Create(&board).Find(&board)
	if query.Error != nil {
		return nil, query.Error
	}
	return &board, nil
}

func (b *boardRepo) Update(board domain.Board) (*domain.Board, error) {
	query := callGorm(b.db).Table("boards").Where("id = ?", board.ID).Updates(board).Find(&board)
	if query.Error != nil {
		return nil, query.Error
	}
	return &board, nil
}

func (b *boardRepo) GetById(id uuid.UUID) (*domain.Board, error) {
	var board domain.Board
	query := callGorm(b.db).Table("boards").Where("id = ?", id).Find(&board)
	if query.Error != nil {
		return nil, query.Error
	}
	return &board, nil
}
func (b *boardRepo) Index() []domain.Board {
	return []domain.Board{}
}

func (b *boardRepo) Delete(id uuid.UUID) error {
	query := callGorm(b.db).Table("boards").Where("id = ?", id).Delete(&domain.Board{})
	return query.Error
}
