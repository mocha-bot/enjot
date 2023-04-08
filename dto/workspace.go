package dto

import "github.com/mocha-bot/enjot/core/entity"

type Workspace struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateWorkspaceRequest struct {
	Name string `json:"name"`
}

type CreateWorkspaceResponse struct {
	Workspace Workspace `json:"workspace"`
}

func ParseCreateWorkspaceResponse(workspace entity.Workspace) CreateWorkspaceResponse {
	return CreateWorkspaceResponse{
		Workspace: Workspace{
			ID:   workspace.ID.String(),
			Name: workspace.Name,
		},
	}
}

type GetWorkspaceResponse struct {
	Workspaces []Workspace `json:"workspaces"`
	Page       PageDTO     `json:"page"`
}

func ParseGetWorkspaceResponse(workspace entity.WorkspaceList) GetWorkspaceResponse {
	workspaces := make([]Workspace, len(workspace.Workspaces))

	for idx, entity := range workspace.Workspaces {
		workspace := Workspace{
			ID:   entity.ID.String(),
			Name: entity.Name,
		}

		workspaces[idx] = workspace
	}

	return GetWorkspaceResponse{
		Workspaces: workspaces,
		Page: PageDTO{
			Limit:  workspace.Page.Limit,
			Offset: workspace.Page.Offset,
		},
	}
}

type UpdateWorkspaceRequest struct {
	WorkspaceID string `json:"workspaceId"`
	Name        string `json:"name"`
}

type DeleteWorkspaceRequest struct {
	WorkspaceID string `json:"workspaceId"`
}
