// cmd/grimoire/main.go

package main

import (
	"context"

	"github.com/gphorvath/grimoire/src/cmd"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd.Execute(ctx)
}
