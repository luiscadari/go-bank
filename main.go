package main

import "fmt"
import (
	"banco/clientes"
	"banco/contas"
)

// PagarBoleto realiza o pagamento de um boleto a partir de uma conta que implementa a interface verificarConta
// Parâmetros:
//   - conta: objeto que implementa a interface verificarConta (ContaCorrente ou ContaPoupanca)
//   - valorBoleto: valor a ser sacado da conta para pagamento
func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

// verificarConta é uma interface que define o método Sacar que deve ser implementado por tipos de conta
type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	// Exemplo 1: Criando uma conta corrente para Bruno

	// Cria um titular para a conta
	bruno := clientes.Titular{
		Nome:      "Bruno",
		CPF:       "123.456.789-09",
		Profissao: "Desamassador de vidro",
	}

	// Cria uma conta corrente associada ao titular Bruno
	conta := contas.ContaCorrente{
		Titular:       bruno,
		NumeroAgencia: 999,
		NumeroConta:   999,
		Saldo:         100, // Saldo inicial de R$ 100
	}

	// Exibe informações do titular da conta
	fmt.Println(conta.Titular)

	// Realiza um depósito de R$ 60 na conta
	conta.Depositar(60)

	// Exemplo 2: Criando uma conta poupança para Denis

	// Cria um titular para a conta poupança
	denis := clientes.Titular{
		Nome:      "Denis",
		CPF:       "123.456.789-10",
		Profissao: "Pedreiro",
	}

	// Cria uma conta poupança associada ao titular Denis
	contaDenis := contas.ContaPoupanca{
		Titular:       denis,
		NumeroAgencia: 123,
		Operacao:      123,
		NumeroConta:   123,
	}

	// Exibe informações do titular da conta poupança
	fmt.Println(contaDenis.Titular)

	// Realiza um depósito de R$ 100 na conta poupança
	contaDenis.Depositar(100)

	// Realiza o pagamento de um boleto no valor de R$ 60 usando a conta poupança
	// Note que passamos o ponteiro (&contaDenis) para que as alterações sejam refletidas
	PagarBoleto(&contaDenis, 60)

	// Exibe o saldo atual da conta poupança
	fmt.Println("Saldo da conta do Denis:", contaDenis.ObterSaldo())

	// Realiza o pagamento de um boleto no valor de R$ 50 usando a conta corrente
	// Também passamos o ponteiro (&conta) aqui
	PagarBoleto(&conta, 50)

	// Exibe o saldo atual da conta corrente
	fmt.Println("Saldo da conta do Bruno:", conta.Saldo)
}
