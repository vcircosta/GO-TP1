package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
		SSLMode  string
	}
	Storage struct {
		Type         string
		DataDir      string
		ContactsFile string
		SQLiteFile   string
	}
}

func LoadConfig() (*Config, error) {
	viper.SetDefault("storage.type", "json")
	viper.SetDefault("storage.data_dir", "data")
	viper.SetDefault("storage.contacts_file", "contacts.json")
	viper.SetDefault("storage.sqlite_file", "contacts.db") // <-- default

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("erreur unmarshall config: %w", err)
	}

	return &cfg, nil
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.Name,
		c.Database.SSLMode,
	)
}

func (c *Config) GetSQLiteFile() string {
	return c.Storage.SQLiteFile
}

func (c *Config) GetContactsFile() string {
	return c.Storage.ContactsFile
}

func (c *Config) GetDataDir() string {
	return c.Storage.DataDir
}

func (c *Config) GetStorageType() string {
	return c.Storage.Type
}

func (c *Config) IsSQLite() bool {
	return strings.ToLower(c.Storage.Type) == "sqlite"
}

func (c *Config) IsJSON() bool {
	return strings.ToLower(c.Storage.Type) == "json"
}
