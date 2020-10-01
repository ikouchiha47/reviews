package main

import (
	"fmt"
	"math/rand"
	"os"
	"reviews/internal/app/reviewsapp"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if err := NewRootCommand().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "reviewsapp",
	}
	cmd.AddCommand(NewApp())

	return cmd
}

func NewApp() *cobra.Command {
	app := reviewsapp.NewApp()

	cmd := &cobra.Command{
		Use: "start",
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.Start()
		},
		// },
		// PostRunE: func(cmd *cobra.Command, args []string) error {
		// 	return app.Close()
		// },
	}

	return cmd
}
