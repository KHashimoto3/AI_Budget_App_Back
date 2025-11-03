package cmd

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
)

var logLevel string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start API server",
	Long:  `The serve command starts the API server for the AI Budget App with Echo`,
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()

		// Set log level
		switch logLevel {
		case "debug":
			e.Logger.SetLevel(log.DEBUG)
		case "info":
			e.Logger.SetLevel(log.INFO)
		case "warn":
			e.Logger.SetLevel(log.WARN)
		case "error":
			e.Logger.SetLevel(log.ERROR)
		default:
			e.Logger.SetLevel(log.INFO)
		}

		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		e.GET("/", func(c echo.Context) error {
			return c.JSON(200, map[string]string{
				"status": "ok",
			})
		})

		e.GET("/health", func(c echo.Context) error {
			return c.JSON(200, map[string]string{
				"status": "healthy",
			})
		})

		fmt.Println("Server started at http://localhost:8080 with log level:", logLevel)

		if err := e.Start(":8080"); err != nil {
			fmt.Printf("Error starting server: %v\n", err)
		} else {
			fmt.Println("Server started on :8080")
		}
	},
}

func init() {
	serveCmd.Flags().StringVarP(&logLevel, "log-level", "l", "info", "ログレベルを指定 (debug, info, warn, error)")
	rootCmd.AddCommand(serveCmd)
}
