package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	MONGO_URI             string
	JWT_KEY               string
	GOOGLE_RANDOM_STATE   string
	GoogleLoginConfig     oauth2.Config
	AWS_ACCESS_KEY        string
	AWS_SECRET_ACCESS_KEY string
	AWS_REGION            string
	AWS_S3_BUCKET         string
	ALLOWED_ORIGIN        string
	SENDGRID_API_KEY      string
	RABBITMQ_URL          string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Cannot find config file, falling back to environment variables")
	}

	viper.AutomaticEnv()

	AppConfig = Config{
		MONGO_URI:           viper.GetString("MONGO_URI"),
		JWT_KEY:             viper.GetString("JWT_KEY"),
		GOOGLE_RANDOM_STATE: viper.GetString("GOOGLE_RANDOM_STATE"),
		GoogleLoginConfig: oauth2.Config{
			ClientID:     viper.GetString("GOOGLE_CLIENT_ID"),
			ClientSecret: viper.GetString("GOOGLE_CLIENT_SECRET"),
			RedirectURL:  viper.GetString("GOOGLE_REDIRECT_URL"),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
			Endpoint:     google.Endpoint,
		},
		AWS_ACCESS_KEY:        viper.GetString("AWS_ACCESS_KEY"),
		AWS_SECRET_ACCESS_KEY: viper.GetString("AWS_SECRET_ACCESS_KEY"),
		AWS_REGION:            viper.GetString("AWS_REGION"),
		AWS_S3_BUCKET:         viper.GetString("AWS_S3_BUCKET"),
		ALLOWED_ORIGIN:        viper.GetString("ALLOWED_ORIGIN"),
		SENDGRID_API_KEY:      viper.GetString("SENDGRID_API_KEY"),
		RABBITMQ_URL:          viper.GetString("RABBITMQ_URL"),
	}

	// validate config
	if AppConfig.MONGO_URI == "" {
		log.Fatal("MONGO_URI is required")
	}
	if AppConfig.JWT_KEY == "" {
		log.Fatal("JWT_KEY is required")
	}
	if AppConfig.GOOGLE_RANDOM_STATE == "" {
		log.Fatal("GOOGLE_RANDOM_STATE is required")
	}
	if AppConfig.GoogleLoginConfig.ClientID == "" {
		log.Fatal("GOOGLE_CLIENT_ID is required")
	}
	if AppConfig.GoogleLoginConfig.ClientSecret == "" {
		log.Fatal("GOOGLE_CLIENT_SECRET is required")
	}
	if AppConfig.GoogleLoginConfig.RedirectURL == "" {
		log.Fatal("GOOGLE_REDIRECT_URL is required")
	}
	if AppConfig.AWS_ACCESS_KEY == "" {
		log.Fatal("AWS_ACCESS_KEY is required")
	}
	if AppConfig.AWS_SECRET_ACCESS_KEY == "" {
		log.Fatal("AWS_SECRET_ACCESS_KEY is required")
	}
	if AppConfig.AWS_REGION == "" {
		log.Fatal("AWS_REGION is required")
	}
	if AppConfig.AWS_S3_BUCKET == "" {
		log.Fatal("AWS_S3_BUCKET is required")
	}
	if AppConfig.ALLOWED_ORIGIN == "" {
		log.Fatal("ALLOWED_ORIGIN is required")
	}
	if AppConfig.SENDGRID_API_KEY == "" {
		log.Fatal("SENDGRID_API_KEY is required")
	}
	if AppConfig.RABBITMQ_URL == "" {
		log.Fatal("RABBITMQ_URL is required")
	}

}
