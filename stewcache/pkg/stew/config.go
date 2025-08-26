package stew

import "time"

type Option func(*Config)



func defaultConfig() *Config{
	return &Config{
		TTL:            DefaultTTL,
		Interval:       DefaultInterval,
		BackupData:     DefaultBackupData,
		BackupInterval: DefaultBackupInterval,
		BackupPath:     DefaultBackupPath,
	}
}

func WithTTL(ttl time.Duration) Option {
	return func(c *Config) {
		c.TTL = ttl
	}
}

func WithInterval(interval time.Duration) Option {
	return func(c *Config) {
		c.Interval = interval
	}
}

func WithBackupInterval(bckInterval time.Duration) Option {
	return func(c *Config) {
		c.BackupInterval = bckInterval
	}
}

func WithBackupData(shouldbackup bool) Option {
	return func(c *Config) {
		c.BackupData = shouldbackup
	}
}

func WithBackupPath(bckpath string) Option {
	return func(c *Config) {
		c.BackupPath = bckpath
	}
}
