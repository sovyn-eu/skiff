// Command skiff creates throwaway Linux VMs on Apple Silicon Macs.
package main

import (
	"os"

	"github.com/sovyn-eu/skiff/internal/cli"
)

func main() {
	os.Exit(cli.Main())
}
