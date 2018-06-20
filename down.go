package main

import (
	"github.com/urfave/cli"
	"log"
	"os/exec"
	"bufio"
	"os"
	"fmt"
	"strings"
)

// destroyNode will tear down the deployed darknode, but keep the config file.
func destroyNode(ctx *cli.Context) error {
	// FIXME : currently it only supports tear down AWS deployment.
	// Needs to figure out way which suits for all kinds of cloud service.
	skip := ctx.Bool("skip-prompt")
	if !skip {
		for {
			fmt.Println("Have you deregistered your node and withdrawn all fees? (Yes/No)")
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			if strings.ToLower(strings.TrimSpace(text)) == "yes"{
				break
			}
			if strings.ToLower(strings.TrimSpace(text)) == "no"{
				return nil
			}
		}
	}

	return destroyAwsNode()
}

// destroyAwsNode tear down the AWS instance.
func destroyAwsNode() error {
	log.Println("Destroying your darknode ...")
	destroy := exec.Command("./terraform", "destroy", "--force")
	pipeToStd(destroy)
	if err := destroy.Start(); err != nil {
		return err
	}

	return destroy.Wait()
}
