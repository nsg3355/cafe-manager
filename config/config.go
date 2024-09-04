package config

import (
	"encoding/json"
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// Config 기본 값 정의
type Config struct {
	ServicePort int    `json:"SERVICE_PORT" env:"SERVICE_PORT" envDefault:"8085"`
	PayhereDB   string `json:"PAYHERE_DB" env:"PAYHERE_DB" envDefault:"root:root123@tcp(127.0.0.1:3306)/payhere"`
	JwtKey      string `json:"JWT_KEY" env:"JWT_KEY" envDefault:"GOODPAYHERE"`
}

var instance *Config

func GetInstance() *Config {
	if instance == nil {
		instance = new(Config)
	}
	return instance
}

// .env 파일 로드
func (c *Config) LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := env.Parse(c); err != nil {
		log.Printf("%+v\n", err)
	}

	conf, _ := json.Marshal(c)
	log.Println("Config Info : ", string(conf))
}
