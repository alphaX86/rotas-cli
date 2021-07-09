/*
Copyright © 2021 Aadhitya A

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
	"fmt"

	"github.com/spf13/cobra"

	"github.com/manifoldco/promptui"

	"os/exec"

	"log"
)

// kubeCmd represents the kube command
var kubeCmd = &cobra.Command{
	Use:   "kube",
	Short: "Install a k8s or k3s cluster of your choice",
	Long:  `If you want to install a k8s or k3s cluster or just kubectl alone, then this subcommand helps you!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello there! So what do you want to install?")

		prompt := promptui.Select{
			Label: "Kube",
			Items: []string{"minikube", "k3s"},
		}
		_, result, err := prompt.Run()

		if err != nil {
			fmt.Println("No option given!")
			return
		}

		fmt.Printf("%q is chosen.. Please wait..", result)

		if result == "minikube" {
			cmd := exec.Command("curl", "-LO", "https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64")

			err := cmd.Run()

			if err != nil {
				log.Fatal(err)
			}

			cmd2 := exec.Command("sudo", "install", "minikube-linux-amd64 /usr/local/bin/minikube")

			err2 := cmd2.Run()

			if err2 != nil {
				log.Fatal(err)
			}
		}

		if result == "k3s" {
			cmd := exec.Command("curl", "-sfL", "https://get.k3s.io")

			err := cmd.Run()

			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(kubeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kubeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kubeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
