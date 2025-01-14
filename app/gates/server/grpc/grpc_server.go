package server

import (
	"golang.org/x/sync/semaphore"
	"log/slog"
	"os"
	"sync"
	"tages.local/app/internal/config"
	pb "tages.local/grpc/images/pb"
)

type server struct {
	pb.UnimplementedImageServiceServer
	uploadSem   *semaphore.Weighted // Семафор для ограничения загрузок
	downloadSem *semaphore.Weighted // Семафор для ограничения скачиваний
	mu          sync.RWMutex        // Мьютекс для безопасного доступа к мапе файлов
	files       map[string]string   // Мапа для хранения соответствия id -> filename
	cfg         *config.Config
	log         *slog.Logger
}

func NewServer(cfg *config.Config, log *slog.Logger) *server {
	// Создаем директорию для загрузок, если её нет
	if err := os.MkdirAll(cfg.Pictures.Path, 0755); err != nil {
		panic(err)
	}

	return &server{
		uploadSem:   semaphore.NewWeighted(int64(cfg.Grpc.UploadLimit)),
		downloadSem: semaphore.NewWeighted(int64(cfg.Grpc.DownloadLimit)),
		files:       make(map[string]string),
		cfg:         cfg,
		log:         log,
	}
}
