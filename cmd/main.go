package main

import (
	"fmt"
	"identity"
	"os"
	"path/filepath"
	"sort"

	"bitbucket.org/cpchain/chain/accounts/abi/bind"
	"bitbucket.org/cpchain/chain/api/cpclient"
	"bitbucket.org/cpchain/chain/commons/log"
	"bitbucket.org/cpchain/chain/tools/contract-admin/flags"
	"bitbucket.org/cpchain/chain/tools/contract-admin/utils"
	"github.com/urfave/cli"
)

var (
	// IdentityCommand identity contract
	IdentityCommand = cli.Command{
		Name:  "identity",
		Usage: "Manage Identity Contract",
		Description: `
		Manage Identity Contract
		`,
		Flags: flags.GeneralFlags,
		Subcommands: []cli.Command{
			{
				Name:        "deploy",
				Usage:       "identity deploy",
				Action:      deploy,
				Flags:       flags.GeneralFlags,
				ArgsUsage:   "string",
				Description: `deploy contract`,
			},
			{
				Name:        "disable",
				Usage:       "disbale",
				Action:      disable,
				Flags:       flags.GeneralFlags,
				Description: `disable contract`,
			},
			{
				Name:        "enable",
				Usage:       "enable",
				Action:      enable,
				Flags:       flags.GeneralFlags,
				Description: `enable contract`,
			},
		},
	}
)

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Authors = []cli.Author{
		{
			Name:  "The cpchain authors",
			Email: "info@cpchain.io",
		},
	}
	app.Copyright = "LGPL"
	app.Usage = "Executable for the cpchain official contract admin"

	app.Action = cli.ShowAppHelp

	app.Commands = []cli.Command{
		IdentityCommand,
	}

	// maintain order
	sort.Sort(cli.CommandsByName(app.Commands))

	return app
}

func main() {
	if err := newApp().Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func deploy(ctx *cli.Context) error {
	endpoint, err := flags.GetEndpoint(ctx)
	if err != nil {
		log.Fatal("endpoint error", "err", err)
	}
	client, err := utils.PrepareCpclient(endpoint)
	if err != nil {
		log.Fatal("prepare client error", "err", err)
	}
	_ = client
	keystoreFile, err := flags.GetKeystorePath(ctx)
	if err != nil {
		log.Fatal("get keystore path error", "err", err)
	}
	password := utils.GetPassword()
	_, key := utils.GetAddressAndKey(keystoreFile, password)
	auth := bind.NewKeyedTransactor(key.PrivateKey)
	_ = auth

	addr, tx, _, err := identity.DeployIdentity(auth, client)
	if err != nil {
		log.Fatal("deploy identity contract error", "err", err)
	}
	log.Info("identity contract address", "addr", addr.Hex())
	return utils.WaitMined(client, tx)
}

func showConfigs(ctx *cli.Context) error {
	instance, _, _, err := createContractInstanceAndTransactor(ctx, false)
	if err != nil {
		return err
	}
	size, err := instance.Count(nil)
	if err != nil {
		log.Error("Maybe your identity contract address is wrong, please check it.")
		return err
	}
	log.Info("identity size", "value", size)

	enabled, _ := instance.Enabled(nil)
	log.Info("contract enabled", "value", enabled)

	return nil
}

func disable(ctx *cli.Context) error {
	instance, opts, client, err := createContractInstanceAndTransactor(ctx, true)
	if err != nil {
		return err
	}
	tx, err := instance.DisableContract(opts)
	if err != nil {
		return err
	}
	return utils.WaitMined(client, tx)
}

func enable(ctx *cli.Context) error {
	instance, opts, client, err := createContractInstanceAndTransactor(ctx, true)
	if err != nil {
		return err
	}
	tx, err := instance.EnableContract(opts)
	if err != nil {
		return err
	}
	return utils.WaitMined(client, tx)
}

func createContractInstanceAndTransactor(ctx *cli.Context, withTransactor bool) (contract *identity.Identity, opts *bind.TransactOpts, client *cpclient.Client, err error) {
	contractAddr, client, key, err := utils.PrepareAll(ctx, withTransactor)
	if err != nil {
		return &identity.Identity{}, &bind.TransactOpts{}, &cpclient.Client{}, err
	}
	if withTransactor {
		opts = bind.NewKeyedTransactor(key.PrivateKey)
	}

	contract, err = identity.NewIdentity(contractAddr, client)
	if err != nil {
		log.Info("Failed to create new contract instance", "err", err)
		return &identity.Identity{}, &bind.TransactOpts{}, &cpclient.Client{}, err
	}

	return contract, opts, client, nil
}
