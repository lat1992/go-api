package config

import (
	"github.com/spf13/viper"
	"strings"
)

func GetConfiguration() *viper.Viper {
	vp := viper.New()
	vp.AutomaticEnv()
	vp.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return vp
}
