package config

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

var Env ConfigDto

func init() {
	ConfigEnv()
}

type ConfigDto struct {
	port         string
	secret_key   string
	database_url string
}

func ConfigEnv() {
	Env = ConfigDto{
		port:         os.Getenv("PORT"),
		secret_key:   os.Getenv("SECRET_KEY"),
		database_url: os.Getenv("MONGO_DB_URL"),
	}
}

func LoadEnviromentVaiable() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading env Variable!")
	}
}

func accessField(key string) (string, error) {
	v := reflect.ValueOf(Env)
	t := v.Type()
	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("expected Struct")
	}
	_, ok := t.FieldByName(key)
	if !ok {
		return "", fmt.Errorf(" property could not be find")
	}
	f := v.FieldByName(key)
	return f.String(), nil
}
func GetEnvProperty(key string) string {
	if Env.port == "" {
		ConfigEnv()
	}
	value, err := accessField(key)
	if err != nil {
		fmt.Println(err)
		return value
	}
	return value
}
