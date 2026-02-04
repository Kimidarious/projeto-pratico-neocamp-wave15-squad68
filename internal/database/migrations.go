package database

import (
    "fmt"
    "log"
    
    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(databaseURL string) error {
    log.Println("🔄 Running database migrations...")
    
    m, err := migrate.New(
        "file://migrations",
        databaseURL,
    )
    if err != nil {
        return fmt.Errorf("failed to create migrate instance: %w", err)
    }
    defer m.Close()
    
    version, dirty, err := m.Version()
    if err != nil && err != migrate.ErrNilVersion {
        return fmt.Errorf("failed to get migration version: %w", err)
    }
    
    if dirty {
        log.Printf("⚠️  Database is in dirty state at version %d. Attempting to fix...", version)
        if err := m.Force(int(version)); err != nil {
            return fmt.Errorf("failed to force version: %w", err)
        }
        log.Printf("✅ Fixed dirty state, now retrying migration from version %d", version)
    }
    
    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        return fmt.Errorf("failed to run migrations: %w", err)
    }
    
    version, dirty, err = m.Version()
    if err != nil && err != migrate.ErrNilVersion {
        return fmt.Errorf("failed to get final migration version: %w", err)
    }
    
    if dirty {
        log.Printf("⚠️  Database is still in dirty state at version %d", version)
    } else {
        log.Printf("✅ Migrations completed successfully (version: %d)", version)
    }
    
    return nil
}

func RollbackMigration(databaseURL string, steps int) error {
    m, err := migrate.New(
        "file://migrations",
        databaseURL,
    )
    if err != nil {
        return err
    }
    defer m.Close()
    
    return m.Steps(-steps)
}