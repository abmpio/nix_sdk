package sdk

import (
	"fmt"
	"time"

	"google.golang.org/grpc/keepalive"
)

const (
	defaultGrpcPort = 9028
	// send pings every 10 seconds if there is no activity
	defaultKeepaliveTime = 10 * time.Second
	// wait 1 second for ping ack before considering the connection dead
	defaultKeepaliveTimeout = time.Second
	// send pings even without active streams
	defaultKeepalivePermitWithoutStream = true
)

type Options struct {
	Host           string
	Port           int32
	keepaliveParam *keepalive.ClientParameters
}

func newDefaultOptions() *Options {
	newOptions := &Options{
		Host: "localhost",
		Port: defaultGrpcPort,
		keepaliveParam: &keepalive.ClientParameters{
			Time:                defaultKeepaliveTime,
			Timeout:             defaultKeepaliveTimeout,
			PermitWithoutStream: defaultKeepalivePermitWithoutStream,
		},
	}
	return newOptions
}

// 获取目标地址,返回为host:port格式
func (o *Options) getHostTarget() string {
	return fmt.Sprintf("%s:%d", o.Host, o.Port)
}

type Option func(options *Options)

func WithHost(host string) Option {
	return func(options *Options) {
		options.Host = host
	}
}

func WithPort(port int32) Option {
	return func(options *Options) {
		options.Port = port
	}
}

// 设置keepalive参数
func WithKeepalive(keepaliveTime time.Duration, keepaliveTimeout time.Duration, permitWithoutStream bool) Option {
	return func(options *Options) {
		options.keepaliveParam = &keepalive.ClientParameters{
			Time:                keepaliveTime,
			Timeout:             keepaliveTimeout,
			PermitWithoutStream: permitWithoutStream,
		}
	}
}
