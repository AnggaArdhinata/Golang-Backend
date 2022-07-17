package serve

import (
	"backend-golang/src/routers"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "server",
	Short: "Start API Server",
	RunE:  serve,
}

func serve(cmd *cobra.Command, args []string) error {
	if mainRoute, err := routers.New(); err == nil {
		var addrs string = "127.0.0.1:8080"

		if pr := os.Getenv("APP_PORT"); pr != "" {
			addrs = ":" + pr
		}

		log.Println("App running on " + addrs)

		// t := cors.New(cors.Options{
		// 	AllowedOrigins: []string{"*"},                              // All origins
		// 	AllowedMethods: []string{"GET", "POST", "HEAD", "OPTIONS"}, // Allowing only get, just an example
		// })

		t := cors.AllowAll()

		if err := http.ListenAndServe(addrs, t.Handler(mainRoute)); err != nil {
			return err
		}
		return nil
	} else {
		return err
	}

}
