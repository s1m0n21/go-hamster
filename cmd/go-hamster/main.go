package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/s1m0n21/go-hamster/build"
	"github.com/s1m0n21/go-hamster/fs"

	"gopkg.in/urfave/cli.v2"
)

func getWD() string{
	wd, _ := os.Getwd()
	return wd
}

func SetupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}

func main() {
	app := &cli.App{
		Name:	 "go-hamster",
		Usage:	 "a file fs in go",
		Version: build.CurrentVersion(),
		Flags:	 []cli.Flag{
			&cli.StringFlag{
				Name:        "addr",
				Usage:       "bind address",
				Hidden:      false,
				Value:       "0.0.0.0",
			},
			&cli.StringFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				Usage:       "listen port",
				Hidden:      false,
				Value:       "4567",
			},
			&cli.StringFlag{
				Name:        "dir",
				Aliases:     []string{"d"},
				Usage:       "shared dir",
				Hidden:      false,
				Value:		 getWD(),
			},
		},
		Action:	 func(cctx *cli.Context) error {
			addr := cctx.String("addr")
			port := cctx.String("port")
			wd := cctx.String("dir")

			if err := fs.StartNewSrv(addr, port, wd); err != nil {
				return err
			}

			return nil
		},
	}

	SetupCloseHandler()

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("err: %v", err)
	}
}
