package config

import (
	"time"

	"github.com/spf13/viper"
)

var (
	Option = &ConfigOptions{}
)

type ConfigOptions struct {
	ConfigFile string

	SSHConfig             SSHConfig
	DataBaseConfig        DataBaseConfig
	MusicFolder           string
	MusicFolderStruct     []string
	MusicFolderStructType string
	CacheFolder           string
	InputDataFolder       string
	LogLevel              string
	LogFile               string
	IgnoreTagging         bool
	DryRun                bool
	DateReTagging         bool
	ExtractAlbumArt       bool
	ExtractAlbumArtFormat string
}

type SSHConfig struct {
	Enabled   bool
	Host      string
	Username  string
	Port      int
	KeyPath   string
	KnownHost string
}

type DataBaseConfig struct {
	URL             string
	Port            int
	Scheme          string
	User            string
	Pass            string
	DataBase        string
	SSLMode         bool
	ConnMaxLifetime time.Duration
}

func init() {
	viper.SetDefault("ssh-enabled", false)
	viper.SetDefault("ssh-host", "0.0.0.0")
	viper.SetDefault("ssh-port", 22)
	viper.SetDefault("ssh-user", "")
	viper.SetDefault("ssh-keypath", "")
	viper.SetDefault("ssh-passphrase", "")
	viper.SetDefault("music-folder", "")
	viper.SetDefault("music-folder-struct", []string{})
	viper.SetDefault("music-folder-struct-type", "")
	viper.SetDefault("cache-folder", "")
	viper.SetDefault("input-data-folder", "")
	viper.SetDefault("log-level", "info")
	viper.SetDefault("log-file", "")
	viper.SetDefault("ignore-tagging", false)
	viper.SetDefault("dry-run", false)
	viper.SetDefault("date-retagging", false)
	viper.SetDefault("extract-album-art", false)
	viper.SetDefault("extract-album-art-format", "jpg")
}
