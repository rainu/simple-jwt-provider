package main

import (
	"errors"
	"github.com/alexflint/go-arg"
)

type config struct {
	ServerAddress              string `arg:"--server-address,env:SERVER_ADDRESS,help:Server-address network-interface to bind on e.g.: '127.0.0.1:8080' default ':80'"`
	DatabaseHost               string `arg:"--database-host,env:DATABASE_HOST,help:Database-Host"`
	DatabasePort               int    `arg:"--database-port,env:DATABASE_PORT,help:Database-Port default: '5432'"`
	DatabaseName               string `arg:"--database-name,env:DATABASE_NAME,help:Database-Name default: 'simple-auth-provider'"`
	DatabaseUsername           string `arg:"--database-username,env:DATABASE_USERNAME,help:Database-Username"`
	DatabasePassword           string `arg:"--database-password,env:DATABASE_PASSWORD,help:Database-Password"`
	DatabaseMigrationsFilePath string `arg:"--database-migrations-file-path,env:DATABASE_MIGRATIONS_FILE_PATH,required,help:Database Migrations File Path"`
	JWTPrivateKey              string `arg:"--jwt-private-key,env:JWT_PRIVATE_KEY,help:JWT PrivateKey ECDSA512,required"`
	EnableAdminAPI             bool   `arg:"--enable-admin-api,env:ENABLE_ADMIN_API,help:Enable admin API to manage stored users (true / false) default: 'true'"`
	AdminAPIUsername           string `arg:"--admin-api-username,env:ADMIN_API_USERNAME,help:Basic Auth Username if enable-admin-api = true"`
	AdminAPIPassword           string `arg:"--admin-api-password,env:ADMIN_API_PASSWORD,help:Basic Auth Password if enable-admin-api = true"`
}

func newConfig() (config, error) {
	cfg := config{
		ServerAddress:  ":80",
		DatabasePort:   5432,
		DatabaseName:   "simple-auth-provider",
		EnableAdminAPI: false,
	}
	err := arg.Parse(&cfg)

	if cfg.EnableAdminAPI && (cfg.AdminAPIPassword == "" || cfg.AdminAPIUsername == "") {
		return cfg, errors.New("admin-api-password and admin-api-username must be set if api has been enabled")
	}

	return cfg, err
}
