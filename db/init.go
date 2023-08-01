package db

import "log"

func CreateLogsTable() error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS logs (
        id SERIAL PRIMARY KEY,
        client_ip VARCHAR(50),
        request_method VARCHAR(10),
        request_path VARCHAR(200),
        request_proto VARCHAR(10),
        status_code INT,
        latency INT,
        password VARCHAR(40),
        action VARCHAR(50)
    )`)
	if err != nil {
		log.Println("Failed to create logs table: ", err)
	}
	return err
}
