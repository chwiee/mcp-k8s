package main

import (
	"context"
	"fmt"
	"os"

	"github.com/chwiee/mcp-k8s/internal/diag"
	"github.com/chwiee/mcp-k8s/internal/k8s"
	"github.com/chwiee/mcp-k8s/internal/output"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: mcp <namespace>")
		os.Exit(1)
	}

	namespace := os.Args[1]

	client, err := k8s.NewClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	result, err := diag.Analyze(ctx, client, namespace)
	if err != nil {
		panic(err)
	}

	output.Print([]interface{}{result, namespace})
}
