package env

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Server struct {
		Port string
	}
	Database struct {
		Host     string
		User     string
		Password string
		Name     string
		Port     string
		Timezone string
	}
	Secret struct {
		JwtSecretKey string
	}
	RabbitMQ struct {
		Host     string
		User     string
		Password string
		Port     string
	}
	Redis struct {
		Host     string
		Port     string
		Password string
	}
	SMTP struct {
		Host string
		Port string
	}
	Email struct {
		From     string
		Password string
	}
	Report struct {
		Port string
	}
	Casbin struct {
		ModelPath  string
		PolicyPath string
	}
}

func LoadEnv() (*Env, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	env := &Env{}

	// Load Server Config
	env.Server.Port = os.Getenv("SERVER_PORT")

	// Load Database Config
	env.Database.Host = os.Getenv("DB_HOST")
	env.Database.User = os.Getenv("DB_USER")
	env.Database.Port = os.Getenv("DB_PORT")
	env.Database.Password = os.Getenv("DB_PASSWORD")
	env.Database.Name = os.Getenv("DB_NAME")
	env.Database.Timezone = os.Getenv("DB_TIMEZONE")

	// Secret
	env.Secret.JwtSecretKey = os.Getenv("JWT_SECRET_KEY")

	// RabbitMQ
	env.RabbitMQ.Host = os.Getenv("RABBITMQ_HOST")
	env.RabbitMQ.User = os.Getenv("RABBITMQ_USER")
	env.RabbitMQ.Password = os.Getenv("RABBITMQ_PASSWORD")
	env.RabbitMQ.Port = os.Getenv("RABBITMQ_PORT")

	// Redis
	env.Redis.Host = os.Getenv("REDIS_HOST")
	env.Redis.Port = os.Getenv("REDIS_PORT")
	env.Redis.Password = os.Getenv("REDIS_PASSWORD")

	// SMTP
	env.SMTP.Host = os.Getenv("SMTP_HOST")
	env.SMTP.Port = os.Getenv("SMTP_PORT")

	// Email
	env.Email.From = os.Getenv("EMAIL_FROM")
	env.Email.Password = os.Getenv("EMAIL_PASSWORD")

	// Report
	env.Report.Port = os.Getenv("REPORT_PORT")

	// Casbin
	env.Casbin.ModelPath = os.Getenv("CASBIN_MODEL_PATH")
	env.Casbin.PolicyPath = os.Getenv("CASBIN_POLICY_PATH")

	return env, nil
}
