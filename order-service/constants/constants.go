package constants

import "os"

var (
	GRPCPort         = ":50052"
	PostgresUser     = getEnv("POSTGRES_USER", "orderuser")
	PostgresPassword = getEnv("POSTGRES_PASSWORD", "orderpass")
	PostgresDB       = getEnv("POSTGRES_DB", "orderdb")
	PostgresHost     = getEnv("POSTGRES_HOST", "order-postgres") // MUST BE CONTAINER NAME
	PostgresPort     = getEnv("POSTGRES_PORT", "5432")
	PostgresSSLMode  = getEnv("POSTGRES_SSLMODE", "disable")
)

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
