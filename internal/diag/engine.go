package diag

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Analyze(ctx context.Context, client *kubernetes.Clientset, namespace string) (*Result, error) {
	pods, err := client.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, pod := range pods.Items {
		for _, cs := range pod.Status.ContainerStatuses {
			if cs.LastTerminationState.Terminated != nil &&
				cs.LastTerminationState.Terminated.Reason == "OOMKilled" {

				return &Result{
					Severity:  "high",
					ErrorType: "OOMKilled",
					Summary:   "O pod está sendo finalizado por falta de memória.",
					Details: []string{
						"Motivo: OOMKilled",
						"ExitCode: 137",
						"Pod: " + pod.Name,
						"Namespace: " + namespace,
					},
					Recommendations: []string{
						"Aumentar o memory limit do container",
						"Verificar consumo de memória da aplicação",
						"Investigar possível vazamento de memória",
					},
				}, nil
			}
		}
	}

	return &Result{
		Severity:  "low",
		ErrorType: "none",
		Summary:   "Nenhum problema crítico detectado no namespace.",
	}, nil
}
