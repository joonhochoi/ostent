package main

import (
	"log"
	"net"
	"os"

	"github.com/ostrost/ostent/assets"
	"github.com/ostrost/ostent/ostent"
	shareassets "github.com/ostrost/ostent/share/assets"
	sharetemplates "github.com/ostrost/ostent/share/templates"
)

func init() {
	ostent.UsePercentTemplate = sharetemplates.UsePercentTemplate
	ostent.TooltipableTemplate = sharetemplates.TooltipableTemplate
}

func Serve(listener net.Listener, production bool, extramap ostent.Muxmap) error {
	server := ostent.NewServer(listener, production)
	access := server.Access
	chain := server.Chain
	mux := server.MUX
	recovery := mux.Recovery

	logger := log.New(os.Stderr, "[ostent] ", 0)
	for _, path := range shareassets.AssetNames() {
		hf := chain.Then(ostent.ServeContentFunc(
			AssetReadFunc(shareassets.Asset),
			AssetInfoFunc(shareassets.AssetInfo),
			path, logger))
		mux.Handle("GET", "/"+path, hf)
		mux.Handle("HEAD", "/"+path, hf)
	}

	// no logger-wrapping for this: it logs itself once a query received via websocket
	mux.Handle("GET", "/ws", recovery.ConstructorFunc(ostent.SlashwsFunc(access, periodFlag.Duration)))

	index := chain.ThenFunc(ostent.IndexFunc(sharetemplates.IndexTemplate, assets.JsAssetNames(shareassets.AssetNames()), periodFlag.Duration))
	mux.Handle("GET", "/", index)
	mux.Handle("HEAD", "/", index)

	/* panics := func(http.ResponseWriter, *http.Request) {
		panic(fmt.Errorf("I'm panicing"))
	}
	mux.Handle("GET", "/panic", chain.ThenFunc(panics)) // */

	ostent.Banner(listener.Addr().String(), "ostent", logger)
	return server.ServeExtra(listener, extramap)
}
