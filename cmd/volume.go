package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var volumeCmd = &cobra.Command{
	Use:     "volume",
	Aliases: []string{"volumes"},
	Short:   "Details of Civo volumes",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := cmd.Help()
		if err != nil {
			return err
		}
		return errors.New("a valid subcommand is required")
	},
}

func init() {
	rootCmd.AddCommand(volumeCmd)
	volumeCmd.AddCommand(volumeCreateCmd)
	volumeCmd.AddCommand(volumeListCmd)
	volumeCmd.AddCommand(volumeRemoveCmd)
	// volumeCmd.AddCommand(volumeResizeCmd)
	volumeCmd.AddCommand(volumeAttachCmd)
	volumeCmd.AddCommand(volumeDetachCmd)

	volumeCreateCmd.Flags().IntVarP(&createSizeGB, "size-gb", "s", 0, "The new size in GB (required)")
	volumeCreateCmd.Flags().StringVarP(&networkVolumeID, "network", "t", "default", "The network name/ID where the volume will be created")
	volumeCreateCmd.MarkFlagRequired("size-gb")

	volumeResizeCmd.Flags().IntVarP(&newSizeGB, "size-gb", "s", 0, "The new size in GB (required)")
	volumeResizeCmd.MarkFlagRequired("size-gb")

	volumeAttachCmd.Flags().BoolVarP(&waitVolumeAttach, "wait", "w", false, "wait until the volume is attached")

	volumeDetachCmd.Flags().BoolVarP(&waitVolumeDetach, "wait", "w", false, "wait until the volume is detached")
}
