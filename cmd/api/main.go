package main

import (
	"fmt"
	"go-task-manager/internal/config"
	userHandler "go-task-manager/internal/delivery/http/user"
	taskHandler "go-task-manager/internal/delivery/http/task"
	userRepoPkg "go-task-manager/internal/repository/postgres"
	taskRepo "go-task-manager/internal/repository/postgres/task"
	userUseCase "go-task-manager/internal/usecase/user"
	taskUseCase "go-task-manager/internal/usecase/task"
	"go-task-manager/pkg/database"
	"log"
	"net/http"
)

// CORS Middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

func corsHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		handler(w, r)
	}
}

func main() {
	fmt.Println("Starting Task Manager API...")

	// 1. Config Yükle
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	// 2. Veritabanına Bağlan
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	defer db.Close()

	// 3. Katmanları Bağla (Dependency Injection)
	// User
	userRepo := userRepoPkg.NewUserRepository(db)
	userUC := userUseCase.NewUserUseCase(userRepo)

	// Task
	taskRepo := taskRepo.NewTaskRepository(db)
	taskUC := taskUseCase.NewTaskUseCase(taskRepo)

	// 4. Router ve Handler'ları Ayarla
	mux := http.NewServeMux()
	
	// Static files (Frontend)
	fs := http.FileServer(http.Dir("./web"))
	mux.Handle("/", corsMiddleware(fs))
	mux.Handle("/css/", corsMiddleware(fs))
	mux.Handle("/js/", corsMiddleware(fs))
	
	// API endpoints (with CORS)
	userHandler.NewUserHandler(mux, userUC)
	taskHandler.NewTaskHandler(mux, taskUC)

	// 5. Server'ı Başlat
	address := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server running on port %s", cfg.Port)
	if err := http.ListenAndServe(address, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
