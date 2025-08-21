package config

type MessageConfig struct {
	QueueDB     int    `mapstructure:"queue_db"`
	PoolDB      int    `mapstructure:"pool_db"`
	WorkerCount int    `mapstructure:"worker_count"`
	PoolName    string `mapstructure:"pool_name"`
}
