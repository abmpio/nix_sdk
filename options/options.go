package options

import (
	"fmt"
	"sync"

	"github.com/abmpio/abmp/pkg/log"
	"github.com/abmpio/configurationx"
	"github.com/go-viper/mapstructure/v2"
	"go.uber.org/zap"
)

const (
	ConfigurationKey string = "nix_sdk"
)

var (
	_options NixOptions
	_once    sync.Once
)

type NixOptions struct {
	Host     string `json:"host"`
	Port     int32  `json:"port"`
	Disabled bool   `json:"disabled"`

	// 多少秒没数据就发一个 ping(秒)，默认值是30s
	KeepaliveTimeSec *int32 `json:"keepaliveTimeSec"`
	// ping 发出去后，等对方回 ACK 的最大时间(秒),默认值是5s
	KeepaliveTimeoutSec *int32 `json:"keepaliveTimeoutSec"`
	// 允许在没有流的情况下发送 keepalive,空闲也能保活，防止 NAT 回收
	KeepalivePermitWithoutStream *bool `json:"keepalivePermitWithoutStream"`
}

func (o *NixOptions) normalize() {
	if len(o.Host) <= 0 {
		o.Host = "127.0.0.1"
	}
	if o.Port <= 0 {
		o.Port = 9028
	}
	if o.KeepalivePermitWithoutStream == nil {
		v := true
		o.KeepalivePermitWithoutStream = &v
	}
	if o.KeepaliveTimeSec == nil {
		v := int32(30)
		o.KeepaliveTimeSec = &v
	}
	if o.KeepaliveTimeoutSec == nil {
		v := int32(5)
		o.KeepaliveTimeoutSec = &v
	}
}

func (o *NixOptions) String() string {
	KeepaliveTimeSec, KeepaliveTimeoutSec := 0, 0
	if o.KeepaliveTimeSec != nil {
		KeepaliveTimeSec = int(*o.KeepaliveTimeSec)
	}
	if o.KeepaliveTimeoutSec != nil {
		KeepaliveTimeoutSec = int(*o.KeepaliveTimeoutSec)
	}
	KeepalivePermitWithoutStream := false
	if o.KeepalivePermitWithoutStream != nil {
		KeepalivePermitWithoutStream = *o.KeepalivePermitWithoutStream
	}
	return fmt.Sprintf("%s:%d,keepaliveTimeSec:%d,keepaliveTimeoutSec:%d,keepalivePermitWithoutStream:%t",
		o.Host,
		o.Port,
		KeepaliveTimeSec,
		KeepaliveTimeoutSec,
		KeepalivePermitWithoutStream)
}

func GetOptions() *NixOptions {
	_once.Do(func() {
		if err := configurationx.GetInstance().UnmarshFromKey(ConfigurationKey, &_options, func(dc *mapstructure.DecoderConfig) {
			dc.TagName = "json"
		}); err != nil {
			err = fmt.Errorf("无效的配置文件,%s", err)
			log.Logger.Error(err.Error())
			panic(err)
		}
		_options.normalize()
		log.Logger.Info("nix sdk options", zap.Any("options", &_options))
	})
	return &_options
}
