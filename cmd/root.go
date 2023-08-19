package cmd

import (
	"fmt"
	. "gitee.com/bbk47/bbk/v3/src"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use: "bbk",
	Run: func(cmd *cobra.Command, args []string) {

		mode := viper.GetString("mode")

		if mode == "server" {
			servopts := &ServerOpts{
				Mode:       "server",
				ListenAddr: viper.GetString("listenAddr"),
				ListenPort: viper.GetInt("listenPort"),
				LogLevel:   viper.GetString("logLevel"),
				Method:     viper.GetString("method"),
				Password:   viper.GetString("password"),
				WorkMode:   viper.GetString("workMode"),
				WorkPath:   viper.GetString("workPath"),
				SslKey:     viper.GetString("sslKey"),
				SslCrt:     viper.GetString("sslCrt"),
			}
			svr := NewServer(servopts)
			svr.Bootstrap()
		} else if mode == "client" {
			cliopts := &ClientOpts{
				Mode:           viper.GetString("mode"),
				ListenAddr:     viper.GetString("listenAddr"),
				ListenPort:     viper.GetInt("listenPort"),
				ListenHttpPort: viper.GetInt("listenHttpPort"),
				LogLevel:       viper.GetString("logLevel"),
				Ping:           viper.GetBool("ping"),
				TunnelOpts:     nil,
			}

			cliopts.TunnelOpts = &TunnelOpts{
				Protocol: viper.GetString("tunnelOpts.protocol"),
				Secure:   viper.GetBool("tunnelOpts.secure"),
				Host:     viper.GetString("tunnelOpts.host"),
				Port:     viper.GetString("tunnelOpts.port"),
				Path:     viper.GetString("tunnelOpts.path"),
				Method:   viper.GetString("tunnelOpts.method"),
				Password: viper.GetString("tunnelOpts.password"),
			}

			cli := NewClient(cliopts)
			cli.Bootstrap()
		} else {
			log.Fatalln("invalid mode config in ", cfgFile)
		}

	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of bbk",
	Long:  `All software has versions. This is xuxihai's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bbk release v3.0.0 -- HEAD")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "--config config.json")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile == "" {
		return
	}
	// Use config file from the flag.
	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
