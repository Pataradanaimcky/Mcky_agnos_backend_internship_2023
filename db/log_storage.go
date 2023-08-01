package db

import (
	"log"
	"time"
)

func LogRequestInfo(clientIP, requestMethod, requestPath, requestProto string, statusCode int, latency time.Duration) {
	log.Printf("Logging request info: clientIP=%s, requestMethod=%s, requestPath=%s, requestProto=%s, statusCode=%d, latency=%d",
		clientIP, requestMethod, requestPath, requestProto, statusCode, latency.Milliseconds())

	_, err := db.Exec("INSERT INTO logs (client_ip, request_method, request_path, request_proto, status_code, latency) VALUES ($1, $2, $3, $4, $5, $6)",
		clientIP, requestMethod, requestPath, requestProto, statusCode, latency.Milliseconds())
	if err != nil {
		log.Println("Failed to log request info: ", err)
	}
}

func LogRequest(password, action string) {
	log.Printf("Logging request: password=%s, action=%s", password, action)

	_, err := db.Exec("INSERT INTO logs (password, action) VALUES ($1, $2)", password, action)
	if err != nil {
		log.Println("Failed to log request: ", err)
	}
}
