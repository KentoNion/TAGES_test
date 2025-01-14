package server

import (
	"context"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"tages.local/app/internal/config"
	"time"

	"github.com/google/uuid"
	"golang.org/x/sync/semaphore"
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

// UploadImage обрабатывает загрузку изображения
func (s *server) UploadImage(stream pb.ImageService_UploadImageServer) error {
	const op = "app.gates.server.UploadImage"
	// Получаем разрешение на загрузку
	if err := s.uploadSem.Acquire(stream.Context(), 1); err != nil {
		s.log.Error(op, "failed to acquire upload semaphore", err)
		return err
	}
	defer s.uploadSem.Release(1)

	var fileID string
	var filename string
	var fileWriter *os.File

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Завершаем загрузку и отправляем ответ
			s.log.Info(op, "upload finished")
			fileInfo, err := fileWriter.Stat() //точка выхода ---------------------------
			if err != nil {
				s.log.Error(op, "failed to stat file", err)
				return err
			}

			return stream.SendAndClose(&pb.ImageUploadResponse{
				Id:       fileID,
				Filename: filename,
				Size:     fileInfo.Size(),
			})
		}
		if err != nil {
			s.log.Error(op, "failed to receive file", err)
			return err
		}

		// Обрабатываем первый чанк с метаданными
		if md := req.GetMetadata(); md != nil {
			fileID = uuid.New().String()
			filename = md.GetFilename()
			s.log.Info(op, "starting upload file: ", filename)

			// Создаем файл
			filePath := filepath.Join(s.cfg.Pictures.Path, fileID)
			var err error
			fileWriter, err = os.Create(filePath)
			if err != nil {
				s.log.Error(op, "failed to create file", err)
				return err
			}
			defer fileWriter.Close()

			// Сохраняем информацию о файле
			s.mu.Lock()
			s.files[fileID] = filename
			s.mu.Unlock()

			continue
		}

		// Записываем чанк в файл
		if chunk := req.GetChunk(); chunk != nil {
			if _, err := fileWriter.Write(chunk); err != nil {
				s.log.Error(op, "failed to write chunk", err)
				return err
			}
		}
	}
}

// DownloadImage обрабатывает скачивание изображения
func (s *server) DownloadImage(req *pb.ImageDownloadRequest, stream pb.ImageService_DownloadImageServer) error {
	const op = "app.gates.server.DownloadImage"
	// Получаем разрешение на скачивание
	if err := s.downloadSem.Acquire(stream.Context(), 1); err != nil {
		s.log.Error(op, "failed to acquire download semaphore", err)
		return err
	}
	defer s.downloadSem.Release(1)

	// Проверяем существование файла
	id := req.GetId()
	filePath := filepath.Join(s.cfg.Pictures.Path, id)
	s.log.Info(op, "starting download file id: ", id)
	file, err := os.Open(filePath)
	if err != nil {
		s.log.Error(op, "failed to open file", err)
		return err
	}
	defer file.Close()

	buffer := make([]byte, s.cfg.Grpc.ChunkSize)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			s.log.Error(op, "failed to read file", err)
			return err
		}

		if err := stream.Send(&pb.ImageDownloadResponse{
			Chunk: buffer[:n],
		}); err != nil {
			s.log.Error(op, "failed to send chunk", err)
			return err
		}
	}
	s.log.Info(op, "finished download file")
	return nil
}

// ListImages возвращает список доступных изображений
func (s *server) ListImages(ctx context.Context, req *pb.ListImagesRequest) (*pb.ListImagesResponse, error) {
	const op = "app.gates.server.ListImages"
	s.mu.RLock()
	defer s.mu.RUnlock()

	start := int(req.GetPage() * req.GetPerPage())
	end := int((req.GetPage() + 1) * req.GetPerPage())

	var images []*pb.ImageInfo
	i := 0

	for id, filename := range s.files {
		if i >= start && i < end {
			filePath := filepath.Join(s.cfg.Pictures.Path, id)
			fileInfo, err := os.Stat(filePath)
			if err != nil {
				continue
			}

			images = append(images, &pb.ImageInfo{
				Id:        id,
				Filename:  filename,
				Size:      fileInfo.Size(),
				CreatedAt: time.Unix(0, fileInfo.Sys().(*syscall.Win32FileAttributeData).CreationTime.Nanoseconds()).Format(time.RFC3339),
				UpdatedAt: fileInfo.ModTime().Format(time.RFC3339),
			})
		}
		i++
		if i >= end {
			break
		}
	}

	return &pb.ListImagesResponse{
		Images: images,
		Total:  int32(len(s.files)),
	}, nil
}
