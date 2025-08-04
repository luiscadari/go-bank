# go-bank
## Simulador de Operações Bancárias em Go

Este projeto é uma simulação simples de operações bancárias implementada em Go, demonstrando conceitos como:
- Structs e métodos
- Interfaces
- Pacotes e organização de código
- Operações bancárias básicas (depósito, saque, pagamento de boletos)

## Estrutura do Projeto
```plaintext
banco/
├── clientes/   # Pacote com a definição de Titular
├── contas/     # Pacote com implementações de ContaCorrente e ContaPoupanca
main.go         # Exemplo de uso das contas bancárias
```
## Funcionalidades Implementadas

- Criação de contas correntes e poupanças
- Operações de depósito e saque
- Pagamento de boletos através da interface `verificarConta`
- Modelagem de clientes/titulares das contas

## Como Executar

1. Certifique-se de ter o Go instalado
2. Clone o repositório
3. Execute o programa:
```bash
go run main.go