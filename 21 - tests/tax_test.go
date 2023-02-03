package tax

import "testing"

/*
	Comandos de teste
	go test -v
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out => coverage em HTML
	go test -bench=.   => radar teste de benchmark
	go test -bench=. -run=ˆ#
	go test -bench=. -run=ˆ# -count=10 => rodar apenas 10 vezes
	go test -bench=. -run=ˆ# -count=10 -benchtime=3s radar por 3 segundos
	go test -bench=. -run=ˆ# -benchmem  => ver a alocação de memoria
	go help test
*/

// Teste simples
func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0
	result := CalculateTax(amount)
	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

// Test in Catch
func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expect float64
	}
	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		{0.0, 0.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expect {
			t.Errorf("Expected %f but got %f", item.expect, result)
		}
	}
}

// Teste de Benchmark
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(5000.0)
	}
}
