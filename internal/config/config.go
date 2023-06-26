package config

import (
	"strings"

	"github.com/spf13/viper"
)

func GetConfiguration() *viper.Viper {
	vp := viper.New()
	vp.AutomaticEnv()
	vp.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return vp
}
