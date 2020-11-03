package main

import (
	"context"
	"flag"
	"github.com/mcarloai/mai-v3-broker/common/redis"
	"github.com/mcarloai/mai-v3-broker/dao"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/mcarloai/mai-v3-broker/api"
	"github.com/mcarloai/mai-v3-broker/common/chain"
	"github.com/mcarloai/mai-v3-broker/common/chain/ethereum"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/launcher"
	"github.com/mcarloai/mai-v3-broker/match"
	"github.com/mcarloai/mai-v3-broker/watcher"
	"github.com/mcarloai/mai-v3-broker/websocket"
	logger "github.com/sirupsen/logrus"
)

func main() {
	ctx, stop := context.WithCancel(context.Background())
	go WaitExitSignal(stop)

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
	chainCli, err = ethereum.NewClient(ctx, conf.Conf.BlockChain.ProviderURL)
        if err != nil {
                logger.Errorf("init ethereum client error:%s", err.Error())
                os.Exit(-3)
        }

	// msg chan for websocket message
	wsChan := make(chan interface{}, 100)
	// match chan for match message
	matchChan := make(chan interface{}, 100)

	wg := &sync.WaitGroup{}

	// start api server
	apiServer, err := api.New(ctx, chainCli, dao, wsChan, matchChan)
	if err != nil {
		logger.Errorf("create api server fail:%s", err.Error())
		os.Exit(-3)
	}
	apiErrChan := make(chan error, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := apiServer.Start(); err != nil {
			apiErrChan <- err
		}
	}()

	// start websocket server
	wsErrChan := make(chan error, 1)
	wg.Add(1)
	wsServer := websocket.New(ctx, wsChan)
	go func() {
		defer wg.Done()
		if err := wsServer.Start(); err != nil {
			wsErrChan <- err
		}
	}()

	// start match server
	matchErrChan := make(chan error, 1)
	wg.Add(1)
	matchServer := match.New(ctx, chainCli, dao, wsChan, matchChan)
	go func() {
		defer wg.Done()
		if err := matchServer.Start(); err != nil {
			matchErrChan <- err
		}
	}()

	// start launcher
	launcherErrChan := make(chan error, 1)
	wg.Add(1)
	launch := launcher.NewLaunch(ctx, dao, chainCli, wsChan, matchChan)
	go func() {
		defer wg.Done()
		if err := launch.Start(); err != nil {
			launcherErrChan <- err
		}
	}()

	// start watcher
	watcherErrChan := make(chan error, 1)
	wg.Add(1)
	watcherSrv := watcher.New(ctx, chainCli, dao, conf.Conf.FactoryAddress, wsChan, matchChan)
	go func() {
		defer wg.Done()
		if err := watcherSrv.Start(); err != nil {
			watcherErrChan <- err
		}
	}()

	select {
	case <-ctx.Done():
		wg.Wait()
		os.Exit(0)
	case err := <-apiErrChan:
		logger.Errorf("api server stop error:%s", err.Error())
		stop()
	case err := <-wsErrChan:
		logger.Errorf("websocket server stop error:%s", err.Error())
		stop()
	case err := <-matchErrChan:
		logger.Errorf("match server stop error:%s", err.Error())
		stop()
	case err := <-watcherErrChan:
		logger.Errorf("watcher server stop error:%s", err.Error())
		stop()
	}
	wg.Wait()
	os.Exit(1)
}

func WaitExitSignal(ctxStop context.CancelFunc) {
	var exitSignal = make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGTERM)
	signal.Notify(exitSignal, syscall.SIGINT)

	sig := <-exitSignal
	logger.Infof("caught sig: %+v, Stopping...\n", sig)
	ctxStop()
}
