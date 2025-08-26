package stew

import "time"

type Option func(*Config)

func NewConfig(option ...func(*Config)) *Config {
	// creating an instance of the struct
	config := &Config{
		TTL:            DefaultTTL,
		Interval:       DefaultInterval,
		BackupData:     DefaultBackupData,
		BackupInterval: DefaultBackupInterval,
		BackupPath:     DefaultBackupPath,
	}
	// iterating through the provided config options
	for _, o := range option {
		// updating the value of the given struct with the initialized config
		// this calls the function and updates the value
		o(config)
	}
	return config
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
