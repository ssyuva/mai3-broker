package main

import (
	"context"
	"flag"
	"github.com/mcarloai/mai-v3-broker/common/redis"
	"github.com/mcarloai/mai-v3-broker/dao"
	"os"
	"os/signal"
	"syscall"

	"github.com/mcarloai/mai-v3-broker/api"
	"github.com/mcarloai/mai-v3-broker/common/chain"
	"github.com/mcarloai/mai-v3-broker/common/chain/ethereum"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/launcher"
	"github.com/mcarloai/mai-v3-broker/match"
	"github.com/mcarloai/mai-v3-broker/pricemonitor"
	"github.com/mcarloai/mai-v3-broker/watcher"
	"github.com/mcarloai/mai-v3-broker/websocket"
	"golang.org/x/sync/errgroup"

	logger "github.com/sirupsen/logrus"
)

func main() {
	backgroundCtx, stop := context.WithCancel(context.Background())
	go WaitExitSignal(stop)

	group, ctx := errgroup.WithContext(backgroundCtx)

	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}

	// init redis
	err := redis.Init(conf.Conf.RedisURL)
	if err != nil {
		logger.Errorf("create redis client fail:%s", err.Error())
		os.Exit(-1)
	}

	// init database
	if err = dao.ConnectPostgres(conf.Conf.DataBaseURL); err != nil {
		logger.Errorf("create database fail:%s", err.Error())
		os.Exit(-2)
	}

	dao := dao.New()

	var chainCli chain.ChainClient
	if conf.Conf.BlockChain.ChainType == "ethereum" {
		chainCli, err = ethereum.NewClient(ctx, conf.Conf.BlockChain.ProviderURL)
		if err != nil {
			logger.Errorf("init ethereum client error:%s", err.Error())
			os.Exit(-3)
		}
	}

	priceMonitor := pricemonitor.NewPriceMonitor(ctx)
	// msg chan for websocket message
	wsChan := make(chan interface{}, 100)

	// new match server
	matchServer, err := match.New(ctx, chainCli, dao, wsChan, priceMonitor)
	if err != nil {
		logger.Errorf("new match server error:%s", err)
		os.Exit(-4)
	}

	// start api server
	apiServer, err := api.New(ctx, chainCli, dao, matchServer)
	if err != nil {
		logger.Errorf("create api server fail:%s", err.Error())
		os.Exit(-3)
	}

	group.Go(func() error {
		return apiServer.Start()
	})

	// start websocket server
	wsServer := websocket.New(ctx, wsChan)
	group.Go(func() error {
		return wsServer.Start()
	})

	// start launcher
	launch := launcher.NewLaunch(ctx, dao, chainCli, matchServer, priceMonitor)
	group.Go(func() error {
		return launch.Start()
	})

	// start watcher
	watcherSrv := watcher.New(ctx, chainCli, dao, matchServer)
	group.Go(func() error {
		return watcherSrv.Start()
	})

	if err := group.Wait(); err != nil {
		logger.Fatalf("service stopped: %s", err)
	}
}

func WaitExitSignal(ctxStop context.CancelFunc) {
	var exitSignal = make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGTERM)
	signal.Notify(exitSignal, syscall.SIGINT)

	sig := <-exitSignal
	logger.Infof("caught sig: %+v, Stopping...\n", sig)
	ctxStop()
}
