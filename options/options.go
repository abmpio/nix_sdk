package options

import (
	"fmt"
	"sync"

	"github.com/abmpio/abmp/pkg/log"
	"github.com/abmpio/app/host"
	"github.com/abmpio/configurationx"
	"github.com/mitchellh/mapstructure"
)

const (
	ConfigurationKey string = "plugins.nixsdk"
)

var (
	_options NixSdkOptions
	_once    sync.Once
)

type NixSdkOptions struct {
	// 插件所在的应用名称，当直接使用service中的FindOneByCode函数时默认将获取此app范围内的值
	// 如果不指定，则获取所有的
	DefaultApp    string `json:"defaultApp"`
	NixServerHost string `json:"nixServerHost"`
	Port          int32  `json:"port"`
}

func (o *NixSdkOptions) normalize() {
	if len(o.DefaultApp) <= 0 {
		appName := host.GetHostEnvironment().GetEnvString(host.ENV_AppName)
		o.DefaultApp = appName
	}
	if len(o.NixServerHost) <= 0 {
		o.NixServerHost = "127.0.0.1"
	}
	if o.Port <= 0 {
		o.Port = 9028
	}
}

func GetOptions() *NixSdkOptions {
	_once.Do(func() {
		if err := configurationx.GetInstance().UnmarshFromKey(ConfigurationKey, &_options, func(dc *mapstructure.DecoderConfig) {
			dc.TagName = "json"
		}); err != nil {
			err = fmt.Errorf("无效的配置文件,%s", err)
			log.Logger.Error(err.Error())
			panic(err)
		}
		_options.normalize()
	})
	return &_options
}
