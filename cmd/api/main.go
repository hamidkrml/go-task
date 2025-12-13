
import (
	"fmt"
	"go-task-manager/internal/config"
	userHandler "go-task-manager/internal/delivery/http/user"
	"go-task-manager/internal/repository/postgres"
	userUseCase "go-task-manager/internal/usecase/user"
	"go-task-manager/pkg/database"
	"log"
	"net/http"
)

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
	userRepo := postgres.NewUserRepository(db)
	userUC := userUseCase.NewUserUseCase(userRepo)

	// 4. Router ve Handler'ları Ayarla
	mux := http.NewServeMux()
	userHandler.NewUserHandler(mux, userUC)

	// 5. Server'ı Başlat
	address := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server running on port %s", cfg.Port)
	if err := http.ListenAndServe(address, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
