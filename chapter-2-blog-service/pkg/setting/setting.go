package setting

import "github.com/spf13/viper"

type Setting struct {
	vip *viper.Viper
}

func NewSetting() (*Setting, error) {
	vip := viper.New()
	vip.AddConfigPath("configs/")
	vip.SetConfigName("config")
	vip.SetConfigType("yaml")
	err := vip.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vip}, nil
}

