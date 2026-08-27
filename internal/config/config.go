package config

import "os"
import "strconv"

type Config struct {
	Addr, DatabasePath, SessionSecret string
	WorkerIntervalSeconds             int
}

func Load() Config {
	c := Config{Addr: env("ADDR", ":8080"), DatabasePath: env("DATABASE_PATH", "./data.sqlite"), SessionSecret: env("SESSION_SECRET", "development-secret"), WorkerIntervalSeconds: 10}
	if v, e := strconv.Atoi(os.Getenv("WORKER_INTERVAL_SECONDS")); e == nil && v > 0 {
		c.WorkerIntervalSeconds = v
	}
	return c
}
func env(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
