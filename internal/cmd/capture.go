package cmd

import (
	"fmt"

	"github.com/serialberry/core-stream/pkg/capture"
	"github.com/spf13/cobra"
)

func CaptureCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "capture",
		Short: "capture camera feed",
		Long:  "capture camera feed then write captured frames to disk",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("write frames to disk")
			device := &capture.Device{} // todo : parsing args
			device.Read()
		},
	}
}
