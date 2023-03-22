package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	MONGO_URL        string
	JWT_SECRET       string
	IDENTIFY_KEY     string
	JWT_REFRESH_TIME int64
	JWT_TIMEOUT_TIME int64
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	MONGO_URL = os.Getenv("MONGO_URL")
	JWT_SECRET = os.Getenv("JWT_SECRET")
	IDENTIFY_KEY = os.Getenv("IDENTIFY_KEY")
	JWT_REFRESH_TIME, err = strconv.ParseInt(os.Getenv("JWT_REFRESH_TIME"), 10, 64)
	JWT_TIMEOUT_TIME, err = strconv.ParseInt(os.Getenv("JWT_TIMEOUT_TIME"), 10, 64)
}
