package diag

import (
	v1 "k8s.io/api/core/v1"
)

type Finding struct {
	PodName     string
	Namespace   string
	Reason      string
	Description string
}

func DetectOOM(pod v1.Pod) *Finding {
	for _, cs := range pod.Status.ContainerStatuses {
		if cs.LastTerminationState.Terminated != nil {
			term := cs.LastTerminationState.Terminated
			if term.Reason == "OOMKilled" {
				return &Finding{
					PodName:     pod.Name,
					Namespace:   pod.Namespace,
					Reason:      "OOMKilled",
					Description: "O container excedeu o limite de memória configurado.",
				}
			}
		}
	}
	return nil
}
