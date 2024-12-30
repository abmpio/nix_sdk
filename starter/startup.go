package starter

import (
	"context"
	"time"

	"github.com/abmpio/abmp/pkg/log"
	"github.com/abmpio/app"
	"github.com/abmpio/nix_sdk"
	"github.com/abmpio/nix_sdk/options"
	pb "github.com/abmpio/nix_sdk/proto"
)

func init() {
	app.RegisterOneStartupAction(clientIniStartupAction)
}

func clientIniStartupAction() app.IStartupAction {
	return app.NewStartupAction(func() {
		if app.HostApplication.SystemConfig().App.IsRunInCli {
			return
		}
		opt := options.GetOptions()
		_client := nix_sdk.NewClient(nix_sdk.WithHost(opt.NixServerHost), nix_sdk.WithPort(opt.Port))
		//测试ping
		for {
			err := _client.InitConnnection()
			if err == nil {
				var res *pb.NixHealthCheckResponse
				res, err = _client.HealthCheck(context.TODO(), &pb.NixHealthCheckRequest{})
				if err == nil {
					break
				}
				if res != nil {
					log.Logger.Info(res.Status.String())
				}
			}

			log.Logger.Warn(err.Error())
			log.Logger.Warn("2s后重新测试...")
			time.Sleep(2 * time.Second)
		}
		nix_sdk.SetGlobalClient(_client)
		app.Context.RegistInstanceAs(_client, new(pb.NixClient))
		app.Context.RegistInstanceAs(_client, new(nix_sdk.Client))
	})
}
