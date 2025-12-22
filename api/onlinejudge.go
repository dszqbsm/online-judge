package api

import (
	"fmt"

	"gitlab.homerunsmartapi.com/go-zero-libs/templatex/api/internal/config"
	"gitlab.homerunsmartapi.com/go-zero-libs/templatex/api/internal/handler"
	"gitlab.homerunsmartapi.com/go-zero-libs/templatex/api/internal/svc"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

func NewCommand(configFile *string) *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "Start api service",
		Run: func(cmd *cobra.Command, args []string) {
			var c config.Config
			conf.MustLoad(*configFile, &c)
			ctx := svc.NewServiceContext(c)

			s := rest.MustNewServer(c.RestConf)
			defer s.Stop()
			handler.RegisterHandlers(s, ctx)

			fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
			s.Start()
		},
	}
}
