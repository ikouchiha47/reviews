package config

import "reviews/pkg/db"

func NewDBConfig() db.Config {
	return db.Config{
		AppName:         "", //gett from viper
		Driver:          "",
		User:            "",
		Password:        "",
		Host:            "",
		Port:            0,
		DBName:          "",
		MaxIdleConns:    0,
		MaxOpenConns:    0,
		ConnMaxLifeTime: 0,
	}
}
