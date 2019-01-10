package setting

import (
	"fmt"
	"log"
	"path/filepath"
	"sync"
	"time"

	"github.com/BurntSushi/toml"
)

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string

	Argon2Memory      uint32
	Argon2Iterations  uint32
	Argon2Parallelism uint8
	Argon2SaltLength  uint32
	Argon2KeyLength   uint32
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type Social struct {
	FacebookClientKey         string
	FacebookClientSecret      string
	FacebookClientCallbackURL string
	GoogleClientKey           string
	GoogleClientSecret        string
	GoogleClientCallbackURL   string
}

type config struct {
	App      App
	Server   Server
	Database Database
	Redis    Redis
	Social   Social
}

var (
	AppConfig config
	once      sync.Once
)

func Setup() {
	var cpath string = "conf/env.toml"
	if _, err := toml.DecodeFile(cpath, &AppConfig); err != nil {
		log.Fatal(err)
	}

	AppConfig.App.ImageMaxSize = AppConfig.App.ImageMaxSize * 1024 * 1024
	AppConfig.Server.ReadTimeout = AppConfig.Server.ReadTimeout * time.Second
	AppConfig.Server.WriteTimeout = AppConfig.Server.ReadTimeout * time.Second
	AppConfig.Redis.IdleTimeout = AppConfig.Redis.IdleTimeout * time.Second
}

func Config() {
	once.Do(func() {
		filePath, err := filepath.Abs("./conf/env.toml")
		if err != nil {
			panic(err)
		}
		fmt.Printf("parse toml file once. filePath: %s\n", filePath)

		if _, err := toml.DecodeFile(filePath, &AppConfig); err != nil {
			log.Fatal(err)
		}

		AppConfig.App.ImageMaxSize = AppConfig.App.ImageMaxSize * 1024 * 1024
		AppConfig.Server.ReadTimeout = AppConfig.Server.ReadTimeout * time.Second
		AppConfig.Server.WriteTimeout = AppConfig.Server.ReadTimeout * time.Second
		AppConfig.Redis.IdleTimeout = AppConfig.Redis.IdleTimeout * time.Second
	})
	// return AppConfig
}
