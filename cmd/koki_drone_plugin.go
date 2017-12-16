package cmd

import (
	"fmt"

	"github.com/kubeciio/koki/executor"
	"gopkg.in/urfave/cli.v1"
)

var (
	KokiCmd = &cli.App{
		Name:        "koki-short-drone-plugin",
		Description: "Convert Koki Kubernetes manifests to validated Kubernetes manifests",
		Usage: `
	_  _____  _  _____ 
 | |/ / _ \| |/ /_ _|
 | ' < (_) | ' < | | 
 |_|\_\___/|_|\_\___|

koki-short converts koki manifests into Kubernetes syntax.

Full documentation available at https://docs.koki.io/short`,
		Action: func(c *cli.Context) error {
			if len(files) == 0 {
				return fmt.Errorf("No files specified for translation")
			}
			return executor.Execute([]string(files), outputPrefix, inPlace, shortPath, overwrite)
		},
		Flags: []cli.Flag{
			cli.StringSliceFlag{
				Name:   "files, f",
				Usage:  "path to koki short files that need to be transformed",
				EnvVar: "PLUGIN_FILES",
				Value:  &files,
			},
			cli.BoolFlag{
				Name:        "in-place, i",
				Usage:       "translate the files in place. Should always be used with --overwrite",
				EnvVar:      "PLUGIN_IN_PLACE",
				Destination: &inPlace,
			},
			cli.StringFlag{
				Name:        "prefix, p",
				Usage:       "prefix for the translated files",
				EnvVar:      "PLUGIN_PREFIX",
				Value:       "kube_",
				Destination: &outputPrefix,
			},
			cli.StringFlag{
				Name:        "short-path, s",
				Usage:       "Absolute path to short binary. Leave empty if short is in $PATH",
				EnvVar:      "PLUGIN_SHORT_PATH",
				Destination: &shortPath,
			},
			cli.BoolFlag{
				Name:        "overwrite, w",
				Usage:       "overwrite output files if they already exist. (False by default)",
				EnvVar:      "PLUGIN_OVERWRITE",
				Destination: &overwrite,
			},
		},
		UsageText: `
# Provide input files using flag
koki -f test1.yaml -f test2.yaml

# Provide input files using environment variable
PLUGIN_FILES=test1.yaml,test2.yaml koki 

# Provide output file prefix using flag
koki -f test1.yaml -p k8s_

# Provide output file prefix using environment variable
PLUGIN_PREFIX=k8s_ PLUGIN_FILES=test1.yaml,test2.yaml koki

# Translate file in-place using flag
koki -f test1.yaml -i

# Translate file in-place using environment variable
PLUGIN_PREFIX=k8s_ PLUGIN_FILES=test1.yaml,test2.yaml PLUGIN_IN_PLACE=true koki

# Preserves the directory in which the original file was found
PLUGIN_PREFIX=k8s_ PLUGIN_FILES=/path/to/dir/test1.yaml,test2.yaml PLUGIN_IN_PLACE=true koki
 >  output file will be created in /path/to/dir/ and in current directory

# Set overwrite=true if output file already exists (using env vars)
PLUGIN_PREFIX=k8s_ PLUGIN_FILES=test1.yaml,test2.yaml PLUGIN_IN_PLACE=true PLUGIN_OVERWRITE=true koki

# Set overwrite=true if output file already exists (using flag)
PLUGIN_PREFIX=k8s_ PLUGIN_FILES=test1.yaml,test2.yaml PLUGIN_IN_PLACE=true koki -w`,
	}
)

//flags for the plugin
var (
	inPlace      bool
	outputPrefix string
	files        cli.StringSlice
	shortPath    string
	overwrite    bool
)
