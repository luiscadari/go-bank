package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	if valor >= c.Saldo {
		return "Saldo insuficiente para saque."
	}
	c.Saldo = c.Saldo - valor
	return "Saque realizado com sucesso!"
}

func (c *ContaCorrente) Transferir(valor float64, conta *ContaCorrente) bool {
	if valor <= 0 || valor > c.Saldo {
		return false
	}
	c.Saldo = c.Saldo - valor
	conta.Saldo = conta.Saldo + valor
	return true
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor <= 0 {
		return "O valor a ser depositado deve ser maior que 0", valor
	}
	c.Saldo = c.Saldo + valor
	return "Saldo foi depositado!", valor
}
