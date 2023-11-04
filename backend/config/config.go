package config

type DatabaseConfig struct {
        Dialect  string
        Host     string
        Port     string
        Username string
        Password string
        Name     string
}

func LoadConfig() *DatabaseConfig {
        return &DatabaseConfig{
                Dialect:  "postgres",
                Host:     "localhost",
                Port:     "5432",
                Username: "cabontech",
                Password:  "12345cabon",
                Name:     "taskmanager",
        }
}

