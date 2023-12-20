package goutils

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigLoad(t *testing.T) {

	viper.AddConfigPath("./testdata")
	err := viper.ReadInConfig()
	PanicError(err)
	allConfigs := viper.AllSettings()
	assert.Equal(t, "map[server:map[port:3000]]", fmt.Sprint(allConfigs))

}
