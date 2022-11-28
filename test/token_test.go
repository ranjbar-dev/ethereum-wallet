package test

import (
	"testing"
)

// Name test
func TestTokenName(t *testing.T) {
	_t := token()

	name, err := _t.GetName(node)
	if err != nil {
		t.Errorf("GetName error was incorect, got: %q, want: %q.", err, "nil")
	}
	if len(name) == 0 {
		t.Errorf("GetName name was incorect, got: %q, want: %q.", name, "len > 0")
	}
}

// Symbol test
func TestTokenSymbol(t *testing.T) {
	_t := token()

	symbol, err := _t.GetSymbol(node)
	if err != nil {
		t.Errorf("GetSymbol error was incorect, got: %q, want: %q.", err, "nil")
	}
	if len(symbol) == 0 {
		t.Errorf("GetSymbol symbol was incorect, got: %q, want: %q.", symbol, "len > 0")
	}
}

// Decimal test
func TestTokenDecimal(t *testing.T) {
	_t := token()

	decimals, err := _t.GetDecimals(node)
	if err != nil {
		t.Errorf("GetDecimal error was incorect, got: %q, want: %q.", err, "nil")
	}

	if decimals <= 0 {
		t.Errorf("GetDecimal Decimal was incorect, got: %q, want: %q.", decimals, "len > 0")
	}
}

// Balance test
func TestTokenBalance(t *testing.T) {
	_t := token()

	balance, err := _t.GetBalance(node, validOwnerAddress)
	if err != nil {
		t.Errorf("GetBalance error was incorect, got: %q, want: %q.", err, "nil")
	}
	if balance == nil {
		t.Errorf("GetBalance Balance was incorect, got: %q, want: %q.", balance, "not nil")
	}
	if balance.Int64() <= 0 {
		t.Errorf("GetBalance Balance was incorect, got: %q, want: %q.", balance, "len > 0")
	}
}
