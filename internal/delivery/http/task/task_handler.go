package task

import (
	"encoding/json"
	"go-task-manager/internal/delivery/http/middleware"
	"go-task-manager/internal/domain"
	"net/http"
	"strconv"
)

type taskHandler struct {
	useCase domain.TaskUseCase
}

func NewTaskHandler(mux *http.ServeMux, useCase domain.TaskUseCase) {
	handler := &taskHandler{
		useCase: useCase,
	}

	// Helper function for protected routes
	handleFunc := func(pattern string, handlerFunc http.HandlerFunc) {
		mux.Handle(pattern, middleware.AuthMiddleware(http.HandlerFunc(handlerFunc)))
	}

	handleFunc("/tasks", handler.HandleTasks)          // GET (List) & POST (Create)
	handleFunc("/tasks/update", handler.UpdateStatus)  // PUT (Update) - Basitlik için query param veya body
	handleFunc("/tasks/delete", handler.Delete)        // DELETE
}

func (h *taskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(int)

	if r.Method == http.MethodGet {
		tasks, err := h.useCase.GetUserTasks(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(tasks)
		return
	}

	if r.Method == http.MethodPost {
		var req struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		if err := h.useCase.CreateTask(userID, req.Title, req.Description); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Task created"})
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func (h *taskHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.Context().Value(middleware.UserIDKey).(int)
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var req struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.useCase.UpdateTaskStatus(id, userID, req.Status); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // Unauthorized da olabilir ama basitleştirdik
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Task updated"})
}

func (h *taskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.Context().Value(middleware.UserIDKey).(int)
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	if err := h.useCase.DeleteTask(id, userID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
