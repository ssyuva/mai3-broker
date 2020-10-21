package main

import (
	"context"
	"github.com/mcarloai/mai-v3-broker/common/redis"
	"github.com/mcarloai/mai-v3-broker/dao"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/mcarloai/mai-v3-broker/api"
	"github.com/mcarloai/mai-v3-broker/websocket"
	"github.com/micro/go-micro/v2/logger"
)

func main() {
	ctx, stop := context.WithCancel(context.Background())
	go WaitExitSignal(stop)

	// init redis
	err := redis.Init(os.Getenv("HSK_REDIS_URL"))
	if err != nil {
		logger.Errorf("create redis client fail:%s", err.Error())
		os.Exit(-1)
	}

	// init database
	if err = dao.ConnectPostgres(os.Getenv("HSK_DATABASE_URL")); err != nil {
		logger.Errorf("create database fail:%s", err.Error())
		os.Exit(-2)
	}

	// msg chan for websocket message
	wsChan := make(chan interface{}, 100)
	// orderbook for every perpetuals
	var orderbookMap sync.Map
	// stopbook for every perpetuals
	var stopBookMap sync.Map
	// cached perpetuals
	var perpetualsMap sync.Map

	wg := &sync.WaitGroup{}

	// start api server
	apiServer, err := api.New(ctx, wsChan, &orderbookMap)
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

	// start stop for stop limit orders every perpetual

	// start match for limit orders every perpetual

	// start watcher

	// start monitor

	select {
	case <-ctx.Done():
		wg.Wait()
		os.Exit(0)
	case err := <-wsErrChan:
		logger.Errorf("websocket server stopped error:%s", err.Error())
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
