package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type grpc struct {
	Port          string `yaml:"port" env-required:"true"`
	UploadLimit   int    `yaml:"upload_limit" env-required:"true"`
	DownloadLimit int    `yaml:"download_limit" env-required:"true"`
	ChunkSize     int    `yaml:"chunk_size" env-required:"true"`
}

type Log struct {
	FilePath string `yaml:"logger_file_path"`
}

type Pictures struct {
	Path string `yaml:"path" env-required:"true"`
}

type Config struct {
	Env      string   `yaml:"env"`
	Grpc     grpc     `yaml:"grpc"`
	Log      Log      `yaml:"logger"`
	Pictures Pictures `yaml:"pictures"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./config.yaml"
	}

	//проверка существует ли файл
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("cannot read config file")
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &cfg
}
