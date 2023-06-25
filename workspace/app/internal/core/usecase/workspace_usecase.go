package usecase

import (
	"errors"
	"fmt"
	"time"
	"workspace/internal/core/domain"
	"workspace/internal/core/dto"
	"workspace/internal/core/port/driven"
	"workspace/pkg/auth/authenticator"

	"github.com/google/uuid"

)

type WorkspaceUsecase interface {
	Create(workspace dto.CreateWorkspaceRequest) (dto.CreateWorkspaceResponse, error)
	Update(workspace dto.UpdateWorkspaceRequest) (dto.UpdateWorkspaceResponse, error)
	CustomerWorkspaces(token string) ([]domain.Workspace, error)
	Index(token string) ([]domain.Workspace, error)
	WorkspaceMembers(workspace dto.WorkspaceMembersRequest) ([]dto.WorkspaceMembersResponse, error)
}

type workspaceUsecase struct {
	workspaceDao driven.WorkspacePortDriven
}

func NewWorkspaceUsecase(workspaceDao driven.WorkspacePortDriven) *workspaceUsecase {
	return &workspaceUsecase{
		workspaceDao: workspaceDao,
	}
}

func (w *workspaceUsecase) Create(workspace dto.CreateWorkspaceRequest) (dto.CreateWorkspaceResponse, error) {

	//TODO: domain

	customerId, err := authenticator.IsCustomerToken(workspace.AccessToken)

	if err != nil {
		return dto.CreateWorkspaceResponse{}, err
	}

	id := uuid.New()

	newWorkspace, err := w.workspaceDao.Create(domain.Workspace{
		ID:        id,
		Title:     workspace.Title,
		CreatorId: customerId,
		CreatedAt: time.Now(),
	})

	fmt.Println(newWorkspace)

	if err != nil {
		return dto.CreateWorkspaceResponse{}, err
	}

	fmt.Println("f")
	return dto.CreateWorkspaceResponse{
		ID:    id,
		Title: "newWorkspace.Title",
	}, nil

}

func (w *workspaceUsecase) Update(workspace dto.UpdateWorkspaceRequest) (dto.UpdateWorkspaceResponse, error) {

	customerId, err := authenticator.IsCustomerToken(workspace.AccessToken)

	if err != nil {
		return dto.UpdateWorkspaceResponse{}, err
	}

	isAdmin := w.workspaceDao.IsAdmin(customerId, workspace.ID)
	if !isAdmin {
		return dto.UpdateWorkspaceResponse{}, errors.New("not valid")
	}

	workspaceInfo, err := w.workspaceDao.Update(workspace.ID, workspace.Title)

	if err != nil {
		return dto.UpdateWorkspaceResponse{}, err
	}

	return dto.UpdateWorkspaceResponse{
		ID:    workspaceInfo.ID,
		Title: workspaceInfo.Title,
	}, nil

}

// TODO: saga
func (w *workspaceUsecase) Delete(workspace dto.UpdateWorkspaceRequest) (dto.UpdateWorkspaceResponse, error) {

	//TODO: delete Transactional

	//userId := uuid.New()

	//TODO:check customer id equle with creatorid

	newWorkspace, err := w.workspaceDao.Update(workspace.ID, workspace.Title)

	if err != nil {
		return dto.UpdateWorkspaceResponse{}, err
	}

	return dto.UpdateWorkspaceResponse{
		ID:    newWorkspace.ID,
		Title: newWorkspace.Title,
	}, nil

}

func (w *workspaceUsecase) AddMember(workspace dto.UpdateWorkspaceRequest) (dto.UpdateWorkspaceResponse, error) {

	//TODO: get user id from token

	//TODO: check user is valid for do this action

	//TODO: add member

	//userId := uuid.New()

	newWorkspace, err := w.workspaceDao.Update(workspace.ID, workspace.Title)

	if err != nil {
		return dto.UpdateWorkspaceResponse{}, err
	}

	return dto.UpdateWorkspaceResponse{
		ID:    newWorkspace.ID,
		Title: newWorkspace.Title,
	}, nil

}

func (w *workspaceUsecase) InviteUser() {

	//Use cache

	//TODO: get user id

	//TODO: check access

	//TODO: check user has account

	//TODO: check email
}

func (w *workspaceUsecase) WorkspaceMembers(workspace dto.WorkspaceMembersRequest) ([]dto.WorkspaceMembersResponse, error) {

	//TODO: Check user token

	//userId := uuid.New()

	//check user is a member of workspace or not

	customers, err := w.workspaceDao.WorkspaceMembers(workspace.ID)

	fmt.Println(customers)

	if err != nil {
		return nil, err
	}

	return nil, nil

	//TODO: Check user is memeber of workspace
}

func (w *workspaceUsecase) Index(token string) ([]domain.Workspace, error) {

	//TODO: check user token and get id

	//userID := uuid.New()

	workspaces, err := w.workspaceDao.Index()

	if err != nil {
		return nil, err
	}
	fmt.Println(workspaces)
	return nil, nil
}

func (w *workspaceUsecase) CustomerWorkspaces(token string) ([]domain.Workspace, error) {

	customerId, err := authenticator.IsCustomerToken(token)

	if err != nil {
		return nil, err
	}

	workspaces, err := w.workspaceDao.CustomerWorkspacesIndex(customerId)

	if err != nil {
		return nil, err
	}

	//TODO: wrap response into response
	fmt.Println(workspaces)
	return nil, nil

}
