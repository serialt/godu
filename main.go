package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/serialt/crab"
	"github.com/serialt/sugar/v3"
)

func init() {
	flag.BoolVar(&appVersion, "v", false, "Display build and version messages")
	flag.StringVar(&ConfigFile, "c", "config.yaml", "Config file")
	flag.StringVar(&AesData, "d", "", "Plaintext for encryption")
	flag.Parse()

	err := sugar.LoadConfig(ConfigFile, &config)
	if err != nil {
		config = new(Config)
	}
	slog.SetDefault(sugar.New())
	config.DecryptConfig()

}
func main() {
	if appVersion {
		fmt.Printf("AppVersion: %v\nGo Version: %v\nBuild Time: %v\nGit Commit: %v\n\n",
			APPVersion,
			GoVersion,
			BuildTime,
			GitCommit,
		)
		return
	}
	if len(AesData) > 0 {
		eData, _ := crab.AESEncryptCBCBase64(AesData, AesKey)
		fmt.Printf("Encrypted string: %v\n", eData)
		dData, _ := crab.AESDecryptCBCBase64(eData, AesKey)
		fmt.Printf("Plaintext : %v\n", dData)
		return
	}
	service()

	// 进程持续运行
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	s := <-c
	slog.Info("Aborting...", "signal", s)
	os.Exit(2)
}
