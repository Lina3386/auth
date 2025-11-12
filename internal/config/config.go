package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type GRPCConfig interface {
	Addres() string
}

type PGConfig interface {
	DSN() string
}

type grpcConfig struct {
	host string
	port string
}

func (g *grpcConfig) Address() string {
	return fmt.Sprintf("%s:%s", g.host, g.port)
}

type pgConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func (p *pgConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		p.host, p.port, p.user, p.password, p.dbname)
}

func Load(path string) (*grpcConfig, PGConfig, error) {
	if err := godotenv.Load(path); err != nil {
		return nil, nil, fmt.Errorf("error loading .env file", err)
	}

	grpcCfg := &grpcConfig{
		host: getEnv("GRPC_HOST", "localhost"),
		port: getEnv("GRPC_PORT", "50051"),
	}

	pgCfg := &pgConfig{
		host:     getEnv("DB_HOST", "localhost"),
		port:     getEnv("DB_PORT", "5432"),
		user:     getEnv("DB_USER", ""),
		password: getEnv("DB_PASSWORD", ""),
		dbname:   getEnv("DB_NAME", ""),
	}

	return grpcCfg, pgCfg, nil
}

func getEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
