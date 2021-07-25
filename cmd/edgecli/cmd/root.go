package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	edge "gitlab.com/omniedge/omniedge-linux-saas-cli"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:           "omniedge",
	Short:         "",
	Long:          ``,
	SilenceErrors: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.SetEnvPrefix("omniedge")
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Fail to execute the command", err)
	}
}

func bindFlags(cmd *cobra.Command) {
	if err := viper.BindPFlags(cmd.LocalFlags()); err != nil {
		log.Fatal(CouldNotBindFlags)
	}
}

func persistAuthFile() {
	var authFile = viper.GetString(cliAuthConfigFile)
	if authFile == "" {
		authFile = authFileDefault
	}
	handledAuthFile, err := edge.HandleAuthFile(authFile)
	if err != nil {
		log.Fatalf("Fail to parse the path of the auth file")
	}
	if err = edge.HandleAuthFileStatus(handledAuthFile); err != nil {
		log.Fatalf("Fail to create omniedge file, err is %s", err.Error())
	}
	if err := viper.WriteConfigAs(handledAuthFile); err != nil {
		log.Fatalf("Fail to write config into file, err is %s", err.Error())
	}
}
