package stew

import "time"

const (
	DefaultTTL            time.Duration = 24 * time.Hour
	DefaultInterval       time.Duration = 48 * time.Hour
	DefaultBackupPath     string        = "./backup"
	DefaultBackupData     bool          = false
	DefaultBackupInterval time.Duration = 12 * time.Hour
)
