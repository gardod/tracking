package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tratnik.net/client/internal/config"
	"tratnik.net/client/internal/repository"
	"tratnik.net/client/internal/service"
	"tratnik.net/client/pkg/nats"
)

var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "Tracking CLI Client",
	Run: func(cmd *cobra.Command, args []string) {
		c := config.GetConfigFromFile(viper.GetString("config"))
		msgBroker := nats.New(c.MessageBroker)
		messageRepo := repository.NewMessage(msgBroker, "demo")
		messageSrvc := service.NewMessage(messageRepo)
		messageSrvc.Listen(c.Filter)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatal("Unable to execute command")
	}
}

func init() {
	rootCmd.PersistentFlags().String("config", "", "config file path")
	rootCmd.Flags().Int64("account-id", 0, "filter by account_id")
	viper.BindPFlag("filter.account_id", rootCmd.Flags().Lookup("account-id"))
}
