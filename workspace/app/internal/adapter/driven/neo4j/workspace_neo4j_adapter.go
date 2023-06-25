package neo4j

import (
	//"workspace/internal/core/domain"
	"context"
	"fmt"
	"workspace/internal/core/domain"
	"workspace/internal/core/port/driven"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"

)

type workspaceRepo struct {
	workspacePort driven.WorkspacePortDriven
	Driver        neo4j.DriverWithContext
	ctx           context.Context
}

func NewWorkspaceRepo(Driver neo4j.DriverWithContext, ctx context.Context) workspaceRepo {
	return workspaceRepo{
		Driver: Driver,
		ctx:    ctx,
	}
}

func (w workspaceRepo) Get(id uuid.UUID) (*domain.Workspace, error) {
	return nil, nil
}

// func (w workspaceRepo) Index(ctx context.Context) ([]domain.Workspace, error) {
// 	return []domain.Workspace{}, nil
// }

func (w workspaceRepo) Create(workspace domain.Workspace) (*domain.Workspace, error) {

	result, err := neo4j.ExecuteQuery(w.ctx, w.Driver,
		"CREATE (workspace:Workspace {workspace_id: $workspace_id, title: $title })-[:creator]->(customer:Customer { id: $customer_id}) RETURN workspace",
		map[string]any{
			"workspace_id": workspace.ID,
			"title":        workspace.Title,
			"customer_id":  workspace.CustomerID,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	itemNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "workspace")
	if err != nil {
		return nil, fmt.Errorf("could not find node workspace")
	}
	workspace_id, err := neo4j.GetProperty[int64](itemNode, "workspace_id")
	if err != nil {
		return nil, err
	}
	fmt.Println(workspace_id)
	title, err := neo4j.GetProperty[string](itemNode, "title")
	if err != nil {
		return nil, err
	}
	fmt.Println(title)
	return nil, nil

}

// MATCH (al:Person{name:"Al Pacino"})
// SET al.born=1960,al.Salary=10000

func (w workspaceRepo) Update(id uuid.UUID, title string) (*domain.Workspace, error) {

	result, err := neo4j.ExecuteQuery(w.ctx, w.Driver,
		"MATCH (workspace:Workspace {workspace_id: $workspace_id }) SET workspace.title = $title  RETURN workspace",
		map[string]any{
			"workspace_id": id,
			"title":        title,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	itemNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "workspace")
	if err != nil {
		return nil, fmt.Errorf("could not find node workspace")
	}
	workspace_id, err := neo4j.GetProperty[int64](itemNode, "workspace_id")
	if err != nil {
		return nil, err
	}
	fmt.Println(workspace_id)
	title, err = neo4j.GetProperty[string](itemNode, "title")
	if err != nil {
		return nil, err
	}
	fmt.Println(title)
	return nil, nil

}

func (w workspaceRepo) CustomerWorkspacesIndex(id uuid.UUID) ([]domain.Workspace, error) {

	result, err := neo4j.ExecuteQuery(w.ctx, w.Driver,
		"MATCH (workspace:Workspace{workspace_id: $workspace_id })->MATCH (customer:Customer) RETURN customer, RETURN workspace",
		map[string]any{
			"workspace_id": id,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	itemNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "workspace")
	if err != nil {
		return nil, fmt.Errorf("could not find node workspace")
	}
	workspace_id, err := neo4j.GetProperty[int64](itemNode, "workspace_id")
	if err != nil {
		return nil, err
	}
	fmt.Println(workspace_id)
	title, err := neo4j.GetProperty[string](itemNode, "title")
	if err != nil {
		return nil, err
	}
	fmt.Println(title)
	return nil, nil

}

func (w workspaceRepo) WorkspaceMembers(id uuid.UUID) ([]domain.Customer, error) {
	result, err := neo4j.ExecuteQuery(w.ctx, w.Driver,
		"MATCH (customer:Customer)->(workspace:Workspace{workspace_id: $workspace_id }) RETURN customer",
		map[string]any{
			"workspace_id": id,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	itemNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "customer")
	if err != nil {
		return nil, fmt.Errorf("could not find node customer")
	}
	customer_id, err := neo4j.GetProperty[int64](itemNode, "customer_id")
	if err != nil {
		return nil, err
	}
	fmt.Println(customer_id)
	name, err := neo4j.GetProperty[string](itemNode, "name")
	if err != nil {
		return nil, err
	}
	fmt.Println(name)
	return nil, nil
}


func (w workspaceRepo)Index() ([]domain.Workspace, error){
	return nil , nil
}


// TODO:
func (w workspaceRepo) IsAdmin(customerId uuid.UUID, workspacId uuid.UUID) bool {
	return false
}
