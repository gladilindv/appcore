package core

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/viper"
)

const (
	envKey = "APP_ENV"
	envFmt = "yaml"

	envLogLevel = "env.log_level"
)

// LoadConfig ...
func LoadConfig() error {

	fpath := filepath.Join(".cfg", "k8s")

	viper.AddConfigPath(fpath)
	viper.SetConfigName(os.Getenv(envKey))
	viper.SetConfigType(envFmt)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	keys := viper.AllKeys()
	sort.Strings(keys)
	for _, k := range keys {
		v := viper.GetString(k)
		fmt.Println(k, "=", v)
	}

	return nil
}

// LogLevel ...
func LogLevel() string {
	return viper.GetString(envLogLevel)
}
