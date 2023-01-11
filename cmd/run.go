package cmd

import (
	"github.com/patrickbrown-dev/pbdb/db"
	"github.com/patrickbrown-dev/pbdb/http"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the pbdb server",
	Run:   run,
}

func init() {
	viper.SetDefault("port", "1728")
	viper.SetDefault("data", "/etc/pbdb/data")

	runCmd.Flags().StringP("port", "p", viper.GetString("port"), "Port pbdb will bind on")
	runCmd.Flags().StringP("data", "d", viper.GetString("data"), "Data file path for pbdb to use")

	viper.BindPFlag("port", runCmd.Flags().Lookup("port"))
	viper.BindPFlag("data", runCmd.Flags().Lookup("data"))
}

func run(cmd *cobra.Command, args []string) {
	db.Initialize()
	http.Serve()
}
