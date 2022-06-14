package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/konstellation/swap/internal/chain"
	"github.com/konstellation/swap/internal/config"
	"github.com/konstellation/swap/internal/logger"
	"github.com/konstellation/swap/internal/mongo"
	"github.com/konstellation/swap/internal/server"
	"github.com/spf13/cobra"
)

var (
	configPath string
)

var rootCmd = &cobra.Command{
	Use:   "portal",
	Short: "portal",
	Long:  `App for transferring tokens between BSC and Konstellation`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		err := godotenv.Load(configPath)
		if err != nil {
			panic(fmt.Errorf("error loading local env file %s", err))
		}

		c := config.New()
		log.Println(c.Name, c.Env)

		msgChan := make(chan string)
		mg, err := mongo.InitConnection(ctx, c.DB)
		if err != nil {
			log.Fatalln(err)
		}

		err = logger.SetLogFile(logger.LogFileName)
		if err != nil {
			log.Fatalln(err)
		}
		defer logger.LogFile.Close()
		log.SetOutput(logger.MultyWriter)

		var bscConn chain.BSCConnection
		var kConn chain.KnstlConnection
		if err := bscConn.InitConnection(ctx, c.SwapInfo.Bsc, mg, &kConn, msgChan); err != nil {
			log.Fatalln(err)
		}
		if err := kConn.InitConnection(ctx, c.SwapInfo.Knstl, mg, &bscConn, msgChan); err != nil {
			log.Fatalln(err)
		}

		go kConn.HandleMessage()
		go bscConn.FillInputData()
		go bscConn.StoreTransactions()
		go bscConn.HandleMessage()
		log.Println("****************** Portal server started")
		go func(msg chan string) {
			for {
				logMsg := <-msg
				log.Println(logMsg)
			}
		}(msgChan)
		server.New(c, mg, &bscConn, &kConn).Run()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "./config.yaml", "--config config.yaml")

	rootCmd.AddCommand(startCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
