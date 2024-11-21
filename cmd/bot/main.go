package main

import (
	"context"

	"github.com/aasumitro/asttax/internal"
	"github.com/spf13/viper"

	// sqlite
	_ "github.com/glebarez/go-sqlite"
)

func main() {
	viper.SetConfigFile(".env")
	ctx := context.Background()
	internal.Run(ctx)
}
