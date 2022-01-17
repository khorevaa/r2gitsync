package db

type Config struct {
	TraceSQLCommands bool
	SQLSlowThreshold uint
	DB               string
}
