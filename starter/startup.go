package starter

import (
	"context"
	"fmt"
	"time"

	"github.com/abmpio/abmp/pkg/log"
	"github.com/abmpio/app"
	"github.com/abmpio/app/host"
	"github.com/abmpio/configurationx/options/nix"
	"github.com/abmpio/nix_sdk"
	pb "github.com/abmpio/nix_sdk/proto"
)

func init() {
	fmt.Println("abmpio.nix_sdk.starter init")

	app.RegisterOneStartupAction(clientIniStartupAction)
}

func clientIniStartupAction() app.IStartupAction {
	return app.NewStartupAction(func() {
		if app.HostApplication.SystemConfig().App.IsRunInCli {
			return
		}
		opt := nix.GetNixOptions()
		normalizeOptions(opt)
		_client := nix_sdk.NewClient(nix_sdk.WithHost(opt.Host), nix_sdk.WithPort(opt.Port))
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

func normalizeOptions(o *nix.NixOptions) {
	if o != nil && len(o.DefaultApp) <= 0 {
		appName := host.GetHostEnvironment().GetEnvString(host.ENV_AppName)
		o.DefaultApp = appName
	}
}
