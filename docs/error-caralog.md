## E001 – OOMKilled

### Sintomas
- Pod reiniciando
- Reason: OOMKilled
- ExitCode: 137

### Causa
O container excedeu o limite de memória configurado.

### Onde aparece
- Pod status
- Events
- Container state

### Correção
- Aumentar memory limit
- Revisar uso de memória da aplicação
