/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start API server",
	Long: `The serve command starts the API server for the AI Budget App with Echo`,
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()

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

		fmt.Println("Server started at http://localhost:8080")

		if err := e.Start(":8080"); err != nil {
			fmt.Printf("Error starting server: %v\n", err)
		} else {
			fmt.Println("Server started on :8080")
		}
	},

}

func init() {
	rootCmd.AddCommand(serveCmd)
}
