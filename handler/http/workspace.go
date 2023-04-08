package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/mocha-bot/enjot/core/module"
	"github.com/mocha-bot/enjot/dto"
)

type workspaceHandler struct {
	usecase module.WorkspaceUsecase
}

func NewWorkspaceHandler(r chi.Router, usecase module.WorkspaceUsecase) {
	handler := workspaceHandler{
		usecase: usecase,
	}

	r.Get("/workspaces", handler.GetWorkspace)
	r.Post("/workspace", handler.CreateWorkspace)
	r.Put("/workspace/{workspaceID}", handler.UpdateWorkspace)
	r.Delete("/workspace/{workspaceID}", handler.DeleteWorkspace)
}

func (h *workspaceHandler) GetWorkspace(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	limitStr := values.Get("limit")
	offsetStr := values.Get("offset")

	var limit int
	var offset int

	limit, err := strconv.Atoi(limitStr)

	if err != nil {
		limit = 0
	}

	offset, err = strconv.Atoi(offsetStr)

	if err != nil {
		offset = 0
	}

	workspaceList, err := h.usecase.GetWorkspace(r.Context(), limit, offset)

	if err != nil {
		parseToErrorMsg(w, http.StatusInternalServerError, err)
		return
	}

	response := dto.ParseGetWorkspaceResponse(workspaceList)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func (h *workspaceHandler) CreateWorkspace(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := io.ReadAll(r.Body)

	if err != nil {
		parseToErrorMsg(w, http.StatusBadRequest, err)
		return
	}

	var req dto.CreateWorkspaceRequest

	err = json.Unmarshal(bodyByte, &req)

	if err != nil {
		parseToErrorMsg(w, http.StatusBadRequest, err)
		return
	}

	workspace, err := h.usecase.CreateWorkspace(r.Context(), req.Name)

	if err != nil {
		parseToErrorMsg(w, http.StatusInternalServerError, err)
		return
	}

	response := dto.ParseCreateWorkspaceResponse(workspace)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func (h *workspaceHandler) UpdateWorkspace(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := io.ReadAll(r.Body)

	if err != nil {
		parseToErrorMsg(w, http.StatusBadRequest, err)
		return
	}

	var req dto.UpdateWorkspaceRequest

	err = json.Unmarshal(bodyByte, &req)

	if err != nil {
		parseToErrorMsg(w, http.StatusBadRequest, err)
		return
	}

	workspaceID, err := uuid.Parse(req.WorkspaceID)

	if err != nil {
		parseToErrorMsg(w, http.StatusBadRequest, err)
		return
	}

	if workspaceID == uuid.Nil {
		parseToErrorMsg(w, http.StatusBadRequest, fmt.Errorf("workspace id is not valid"))
		return
	}

	err = h.usecase.UpdateWorkspace(r.Context(), workspaceID, req.Name)

	if err != nil {
		parseToErrorMsg(w, http.StatusInternalServerError, err)
		return
	}

	response := dto.ParseToDefaultReponse()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func (h *workspaceHandler) DeleteWorkspace(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := io.ReadAll(r.Body)

	if err != nil {
		parseToErrorMsg(w, http.StatusBadRequest, err)
		return
	}

	var req dto.DeleteWorkspaceRequest

	err = json.Unmarshal(bodyByte, &req)

	if err != nil {
		parseToErrorMsg(w, http.StatusBadRequest, err)
		return
	}

	workspaceID, err := uuid.Parse(req.WorkspaceID)

	if err != nil {
		parseToErrorMsg(w, http.StatusBadRequest, err)
		return
	}

	if workspaceID == uuid.Nil {
		parseToErrorMsg(w, http.StatusBadRequest, fmt.Errorf("workspace id is not valid"))
		return
	}

	err = h.usecase.DeleteWorkspace(r.Context(), workspaceID)

	if err != nil {
		parseToErrorMsg(w, http.StatusInternalServerError, err)
		return
	}

	response := dto.ParseToDefaultReponse()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
