package app

import (
	"github.com/tendermint/tendermint/libs/log"
	"github.com/cosmos/cosmos-sdk/x/auth"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	dbm "github.com/tendermint/tendermint/libs/db"
)

const (
	appName = "nameservice"
)

type nameServiceApp struct {
	*bam.BaseApp
}

func NewNameServiceApp(logger log.Logger, db dbm.DB) *nameServiceApp {
	// First define the top level codec that will be shared by the different modules. Note: Codec will be explained later
	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

	var app = &nameServiceApp{
		BaseApp: bApp,
		cdc:     cdc
	}

	return app
}
