/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/multierr"
)

type paths struct {
	ConfigPath       string
	CommonConfigPath string
	SecretsPath      string
}

var (
	p       paths
	err     error
	rootCmd = &cobra.Command{
		Use:   "server",
		Short: "Using go gen tools",
		Long:  `For using go gen tools`,
	}
)

func Execute() {
	ctx := context.Background()
	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		log.Panicf("can't load configs: %v", err)
	}
}
func init() {
	cobra.OnInitialize(initConfig)

	val := reflect.ValueOf(&p).Elem()
	for i := 0; i < val.NumField(); i++ {
		name := val.Type().Field(i).Name
		ptr := val.Field(i).Addr().Interface().(*string)
		text := strings.Join(strings.Split(strcase.ToSnake(name), "_"), " ")
		rootCmd.PersistentFlags().StringVar(
			ptr,
			name,
			"",
			fmt.Sprintf("%s, usually used for configuration", text),
		)
		err = multierr.Append(err, rootCmd.MarkPersistentFlagRequired("commonConfigPath"))
	}

	if err != nil {
		log.Fatalf("failed to set up root cobra.Command: %s", err)
	}

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	val := reflect.ValueOf(&p).Elem()
	for i := 0; i < val.NumField(); i++ {
		name := val.Type().Field(i).Name
		viper.SetConfigFile(name)
		viper.AutomaticEnv()
		if err := viper.MergeInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}

}
