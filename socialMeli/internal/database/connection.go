package database

import (
    "fmt"
    "log"
    "os"
    
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

var DB *gorm.DB

type Config struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    SSLMode  string
}

func LoadConfig() *Config {
    return &Config{
        Host:     getEnv("DB_HOST", "localhost"),
        Port:     getEnv("DB_PORT", "5432"),
        User:     getEnv("DB_USER", "kim"),
        Password: getEnv("DB_PASSWORD", "kim78%"),
        DBName:   getEnv("DB_NAME", "apiSocialMeli"),
        SSLMode:  getEnv("DB_SSLMODE", "disable"),
    }
}

func Connect() (*gorm.DB, error) {
    cfg := LoadConfig()
    
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
    )
    
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }
    
    log.Println("✅ Database connected successfully")
    
    DB = db
    return db, nil
}

// AutoMigrate executa migrations automáticas do GORM
func AutoMigrate(db *gorm.DB, models ...interface{}) error {
    return db.AutoMigrate(models...)
}

func Close() error {
    if DB != nil {
        sqlDB, err := DB.DB()
        if err != nil {
            return err
        }
        return sqlDB.Close()
    }
    return nil
}

func getEnv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}