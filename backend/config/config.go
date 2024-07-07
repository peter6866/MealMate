package config

import (
	"log"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	MONGO_URI             string
	JWT_SECRET            string
	GOOGLE_RANDOM_STATE   string
	GoogleLoginConfig     oauth2.Config
	AWS_ACCESS_KEY        string
	AWS_SECRET_ACCESS_KEY string
	AWS_REGION            string
	AWS_S3_BUCKET         string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	AppConfig = Config{
		MONGO_URI:           viper.GetString("MONGO_URI"),
		JWT_SECRET:          viper.GetString("JWT_SECRET"),
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
	}
}
