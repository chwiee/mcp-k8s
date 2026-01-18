package output

import (
	"fmt"

	"github.com/chwiee/mcp-k8s/internal/diag"
)

func Print(findings []diag.Finding, namespace string) {
	if len(findings) == 0 {
		fmt.Printf("✅ Namespace %s sem problemas detectados.\n", namespace)
		return
	}

	fmt.Printf("❌ Problemas detectados no namespace %s\n\n", namespace)

	for _, f := range findings {
		fmt.Printf("🔴 Pod: %s\n", f.PodName)
		fmt.Printf("   Tipo: %s\n", f.Reason)
		fmt.Printf("   Explicação: %s\n", f.Description)
		fmt.Println("   Recomendação:")
		fmt.Println("   - Aumentar o memory limit do container")
		fmt.Println("   - Avaliar consumo de memória da aplicação")
		fmt.Println("   - Verificar se há vazamento de memória")
		fmt.Println()
	}
}
