/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// signImagesCmd represents the subcommand for `krel sign images`
var signImagesCmd = &cobra.Command{
	Use:           "images",
	Short:         "Sign images",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runSignImages(singOpts, args)
	},
}

func init() {
	signCmd.AddCommand(signImagesCmd)
}

// TODO: implement me :)
func runSignImages(signOpts *signOptions, args []string) error { //nolint:unparam // keeping the parameters for reference
	logrus.Info("Not implemented")

	return nil
}
