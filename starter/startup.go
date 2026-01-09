package starter

import (
	"context"
	"fmt"
	"time"

	"github.com/abmpio/abmp/pkg/log"
	"github.com/abmpio/app"
	"github.com/abmpio/nix_sdk/options"
	pb "github.com/abmpio/nix_sdk/proto"
	"github.com/abmpio/nix_sdk/sdk"
	nix_sdk "github.com/abmpio/nix_sdk/sdk"
)

func init() {
	if app.IsServerMode() {
		fmt.Println("abmpio.nix_sdk.starter init")
	}

	app.RegisterOneStartupAction(clientIniStartupAction)
}

func clientIniStartupAction() app.IStartupAction {
	return app.NewStartupAction(func() {
		if app.HostApplication.SystemConfig().App.IsRunInCli {
			return
		}
		var _client nix_sdk.IClient

		clientOptions := options.GetOptions()
		if !clientOptions.Disabled {
			opts := []sdk.Option{
				sdk.WithHost(clientOptions.Host),
				sdk.WithPort(clientOptions.Port),
			}
			if clientOptions.KeepaliveTimeSec != nil && clientOptions.KeepaliveTimeoutSec != nil && clientOptions.KeepalivePermitWithoutStream != nil {
				opts = append(opts, sdk.WithKeepalive(
					time.Duration(*clientOptions.KeepaliveTimeSec)*time.Second,
					time.Duration(*clientOptions.KeepaliveTimeoutSec)*time.Second,
					*clientOptions.KeepalivePermitWithoutStream,
				))
			}

			grpcClient := nix_sdk.NewClient(opts...)
			//测试ping
			for {
				err := grpcClient.InitConnnection()
				if err != nil {
					log.Logger.Warn(fmt.Sprintf("初始化nix grpc连接时出现异常,option:%s, err:%s",
						clientOptions.String(),
						err.Error()))
				} else {
					res, err := grpcClient.HealthCheck(context.TODO(), &pb.NixHealthCheckRequest{})
					if err != nil {
						log.Logger.Warn(fmt.Sprintf("检测nix grpc 服务健康是否正常时出现异常,option:%s, err:%s",
							clientOptions.String(),
							err.Error()))
					} else {
						if res != nil {
							log.Logger.Info(fmt.Sprintf("nix grpc status:%s", res.Status.String()))
						}
						// set value
						_client = grpcClient
						break
					}
				}
				log.Logger.Warn("2s后重新测试...")
				time.Sleep(2 * time.Second)
			}
		} else {
			log.Logger.Warn("nix sdk client disabled")
			// nullable client
			_client = &nix_sdk.NullableClient{}
		}
		nix_sdk.SetGlobalClient(_client)
		app.Context.RegistInstanceAs(_client, new(pb.NixClient))
		app.Context.RegistInstanceAs(_client, new(nix_sdk.Client))
	})
}
