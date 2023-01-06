package admin

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	connect "github.com/go-zoox/connect/app"
	"github.com/go-zoox/connect/app/config"
	"github.com/go-zoox/lighthouse/admin/api"
	"github.com/go-zoox/lighthouse/core"
	"github.com/go-zoox/zoox"
)

//go:embed static
var StaticFS embed.FS

func flushDNS(cacheManager *core.IPS) {
	dnsList := api.DBAll()
	for _, dns := range dnsList {
		fmt.Printf("[flush] load dns: %s => %s\n", dns.Host, dns.Value)

		ips := strings.Split(dns.Value, ",")
		cacheManager.SetCache(dns.Host, ips)
	}
}

func Start(cfg *config.Config, dnsCfg *core.Config) {
	server := connect.New()

	ipsSearch := core.NewIPS(dnsCfg)
	ipsSearch.ClearCache()

	//
	flushDNS(ipsSearch)

	server.RegisterStatic(func(a *zoox.Application) string {
		indexHTML, _ := StaticFS.ReadFile("static/index.html")

		staticfs, _ := fs.Sub(StaticFS, "static")
		a.StaticFS("/static/", http.FS(staticfs))

		return string(indexHTML)
	})

	server.RegisterApi(func(a *zoox.Application) {
		r := a.Group("/api/v1/i/dns")

		r.Get("/", api.List())
		r.Post("/", func(ctx *zoox.Context) {
			api.Create()(ctx)

			// @TODO
			flushDNS(ipsSearch)
		})

		r.Put("/:id", func(ctx *zoox.Context) {
			api.Update()(ctx)

			// @TODO
			flushDNS(ipsSearch)
		})

		r.Delete("/:id", func(ctx *zoox.Context) {
			api.Delete()(ctx)

			// @TODO
			ipsSearch.ClearCache()
			flushDNS(ipsSearch)
		})
	})

	server.Start(cfg)
}
