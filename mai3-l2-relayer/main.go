package main

import (
	"context"
	"flag"
	"math/big"
	"os"
	"os/signal"
	"syscall"

	"github.com/mcarloai/mai-v3-broker/api"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/l2relayer"
	"golang.org/x/sync/errgroup"

	logger "github.com/sirupsen/logrus"
)

func main() {
	backgroundCtx, stop := context.WithCancel(context.Background())
	go WaitExitSignal(stop)

	group, ctx := errgroup.WithContext(backgroundCtx)

	var err error
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}

	relayer, err := l2relayer.NewL2Relayer(
		conf.Conf.BlockChain.ProviderURL,
		big.NewInt(conf.Conf.BlockChain.ChainID),
		conf.Conf.L2Relayer.Key,
		conf.Conf.BrokerAddress,
		conf.Conf.L2Relayer.GasPrice,
		conf.Conf.L2Relayer.CallFunctionFeePercent,
		conf.Conf.L2Relayer.TradeFee)

	if err != nil {
		logger.Errorf("create l2 relayer fail:%s", err.Error())
		os.Exit(-5)
	}

	// start api server
	server, err := api.NewL2RelayerServer(ctx, relayer)
	if err != nil {
		logger.Errorf("create l2 relayer server fail:%s", err.Error())
		os.Exit(-5)
	}

	group.Go(func() error {
		return server.Start()
	})

	if err := group.Wait(); err != nil {
		logger.Fatalf("service stopped: %s", err)
	}
}

func WaitExitSignal(ctxStop context.CancelFunc) {
	var exitSignal = make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGTERM)
	signal.Notify(exitSignal, syscall.SIGINT)

	sig := <-exitSignal
	logger.Infof("caught sig: %+v, Stopping...\n", sig)
	ctxStop()
}
