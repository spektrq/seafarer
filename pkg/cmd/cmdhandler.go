package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/spektrq/kubectl-go-plugin/pkg/cli"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CmdHandler() *cobra.Command {

	var rootCmd = &cobra.Command{
		Use:     "deocde",
		Aliases: []string{"decodes"},
		Short:   "decode secrets",
		Long:    "decode kubernetes secrets",
		Version: "0.0.1",
		RunE: func(cmd *cobra.Command, args []string) error {
			return decodeSecret(args[0])
		},
	}

	return rootCmd
}

func decodeSecret(secretName string) error {
	client := cli.KubernetesClient()
	secret, err := client.CoreV1().Secrets("").Get(context.Background(), secretName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	for key, val := range secret.Data {
		fmt.Print(key, val)
	}

	return nil
}
