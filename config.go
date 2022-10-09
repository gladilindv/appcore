package core

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

const (
	envKey = "APP_ENV"
	envFmt = "yaml"
	envDef = ".env"

	envLogLevel = "env.log_level"
)

// LoadConfig ...
func LoadConfig() error {

	fPath := filepath.Join(".cfg", "k8s")
	fName := os.Getenv(envKey)
	if fName == "" {
		fName = envDef
	}

	viper.AddConfigPath(fPath)
	viper.SetConfigName(fName)
	viper.SetConfigType(envFmt)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	keys := viper.AllKeys()
	sort.Strings(keys)
	fmt.Printf("config %s loaded with %d keys\n", viper.ConfigFileUsed(), len(keys))
	for _, k := range keys {
		v := viper.GetString(k)
		fmt.Println("\t", k, "=", v)
	}

	return nil
}

// LogLevel ...
func LogLevel() string {
	return viper.GetString(envLogLevel)
}

// Value ...
type Value struct {
	value any
}

// String ...
func (v Value) String() string {
	return cast.ToString(v.value)
}

// Bool ...
func (v Value) Bool() bool {
	return cast.ToBool(v.value)
}

// Int ...
func (v Value) Int() int {
	return cast.ToInt(v.value)
}

// Int32 ...
func (v Value) Int32() int32 {
	return cast.ToInt32(v.value)
}

// Int64 ...
func (v Value) Int64() int64 {
	return cast.ToInt64(v.value)
}

// Float64 ...
func (v Value) Float64() float64 {
	return cast.ToFloat64(v.value)
}

// Duration ...
func (v Value) Duration() time.Duration {
	return cast.ToDuration(v.value)
}

// ConfigValue ...
func ConfigValue(key string) Value {
	return Value{viper.Get(key)}
}
