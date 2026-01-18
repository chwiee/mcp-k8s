package output

import (
	"fmt"

	"github.com/chwiee/mcp-k8s/internal/diag"
)

func Print(result *diag.Result, namespace string) {
	fmt.Printf("📊 Diagnóstico do namespace %s\n\n", namespace)

	fmt.Printf("Severidade: %s\n", result.Severity)
	fmt.Printf("Tipo: %s\n", result.ErrorType)
	fmt.Printf("Resumo: %s\n\n", result.Summary)

	if len(result.Details) > 0 {
		fmt.Println("Detalhes:")
		for _, d := range result.Details {
			fmt.Println("-", d)
		}
		fmt.Println()
	}

	if len(result.Recommendations) > 0 {
		fmt.Println("Recomendações:")
		for _, r := range result.Recommendations {
			fmt.Println("-", r)
		}
		fmt.Println()
	}
}
