package core

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/viper"
)

// LoadConfig ...
func LoadConfig(path string) error {

	viper.AddConfigPath(path)
	viper.SetConfigName(os.Getenv("ENV_APP"))
	viper.SetConfigType("yaml")
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
