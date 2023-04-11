package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	MONGO_URL        string
	REDIS_URL 		 string
	REDIS_SECRET	 string
	JWT_SECRET       string
	RF_JWT_SECRET    string
	IDENTIFY_KEY     string
	JWT_REFRESH_TIME int64
	JWT_TIMEOUT_TIME int64
	MAX_REQUEST_PER_HOUR int64
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	MONGO_URL = os.Getenv("MONGO_URL")
	REDIS_URL = os.Getenv("REDIS_URL")
	REDIS_SECRET = os.Getenv("REDIS_SECRET")
	JWT_SECRET = os.Getenv("JWT_SECRET")
	IDENTIFY_KEY = os.Getenv("IDENTIFY_KEY")
	RF_JWT_SECRET = os.Getenv("RF_JWT_SECRET")
	MAX_REQUEST_PER_HOUR, err = strconv.ParseInt(os.Getenv("MAX_REQUEST_PER_HOUR"), 10, 64)
	JWT_REFRESH_TIME, err = strconv.ParseInt(os.Getenv("JWT_REFRESH_TIME"), 10, 64)
	JWT_TIMEOUT_TIME, err = strconv.ParseInt(os.Getenv("JWT_TIMEOUT_TIME"), 10, 64)
}
