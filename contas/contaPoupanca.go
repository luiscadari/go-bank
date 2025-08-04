package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, Operacao, NumeroConta int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valor float64) string {
	if valor >= c.saldo {
		return "saldo insuficiente para saque."
	}
	c.saldo = c.saldo - valor
	return "Saque realizado com sucesso!"
}

func (c *ContaPoupanca) Depositar(valor float64) (string, float64) {
	if valor <= 0 {
		return "O valor a ser depositado deve ser maior que 0", valor
	}
	c.saldo = c.saldo + valor
	return "Saldo foi depositado!", valor
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
