package config

import "os"

type Config struct {
	PostgresHost               string
	PostgresPort               string
	PostgresUser               string
	PostgresPass               string
	PostgresDB                 string
	PostgresCA                 string
	PostgresCert               string
	PostgresKey                string
	FirebaseCredentialFilePath string
	GoogleMapsAPIKey           string
}

var config *Config

func init() {
	env := os.Getenv("APP_ENV")
	switch env {
	case "prd":
		config = &Config{
			PostgresHost:               "giraffe-402013:asia-northeast1:giraffe-db",
			PostgresUser:               "postgres",
			PostgresPass:               "[[hQ%Kz?]DI%Tss,",
			PostgresDB:                 "production",
			FirebaseCredentialFilePath: "config/giraffe-402013-firebase-adminsdk-oxtqo-aef79cbdb3.json",
			GoogleMapsAPIKey:           "AIzaSyDKPUE8NncfZsSa-BszPRdIHfpWsXGuFm0",
		}
	case "stg":
		config = &Config{
			PostgresHost:               "giraffe-402013:asia-northeast1:giraffe-db",
			PostgresUser:               "postgres",
			PostgresPass:               "[[hQ%Kz?]DI%Tss,",
			PostgresDB:                 "staging",
			FirebaseCredentialFilePath: "config/giraffe-402013-firebase-adminsdk-oxtqo-aef79cbdb3.json",
			GoogleMapsAPIKey:           "AIzaSyDKPUE8NncfZsSa-BszPRdIHfpWsXGuFm0",
		}
	case "dev":
		config = &Config{
			PostgresHost:               "giraffe-402013:asia-northeast1:giraffe-db",
			PostgresUser:               "postgres",
			PostgresPass:               "[[hQ%Kz?]DI%Tss,",
			PostgresDB:                 "develop",
			FirebaseCredentialFilePath: "config/giraffe-402013-firebase-adminsdk-oxtqo-aef79cbdb3.json",
			GoogleMapsAPIKey:           "AIzaSyDKPUE8NncfZsSa-BszPRdIHfpWsXGuFm0",
		}
	default:
		config = &Config{
			PostgresHost:               "34.146.103.238",
			PostgresPort:               "5432",
			PostgresUser:               "postgres",
			PostgresPass:               "[[hQ%Kz?]DI%Tss,",
			PostgresDB:                 "develop",
			PostgresCA:                 "config/server-ca.pem",
			PostgresCert:               "config/client-cert.pem",
			PostgresKey:                "config/client-key.pem",
			FirebaseCredentialFilePath: "config/giraffe-402013-firebase-adminsdk-oxtqo-aef79cbdb3.json",
			GoogleMapsAPIKey:           "AIzaSyDKPUE8NncfZsSa-BszPRdIHfpWsXGuFm0",
		}
	}
}

func Get() *Config {
	return config
}
