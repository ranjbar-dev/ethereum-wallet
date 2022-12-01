package test

import (
	"fmt"
	EthereumWallet "github.com/ranjbar-dev/ethereum-wallet"
	"testing"
)

// GenerateEthereumWallet test
func TestGenerateWallet(t *testing.T) {
	w := EthereumWallet.GenerateEthereumWallet(node)
	if w == nil {
		t.Errorf("GenerateEthereumWallet res was incorect, got: %q, want: %q.", w, "*EthereumWallet")
	}
	if len(w.PrivateKey) == 0 {
		t.Errorf("GenerateEthereumWallet PrivateKey was incorect, got: %q, want: %q.", w.PrivateKey, "valid PrivateKey")
	}
	if len(w.PublicKey) == 0 {
		t.Errorf("GenerateEthereumWallet PublicKey was incorect, got: %q, want: %q.", w.PublicKey, "valid PublicKey")
	}
	if len(w.Address) == 0 {
		t.Errorf("GenerateEthereumWallet Address was incorect, got: %q, want: %q.", w.Address, "valid Address")
	}
	if len(w.Address) == 0 {
		t.Errorf("GenerateEthereumWallet AddressBase58 was incorect, got: %q, want: %q.", w.Address, "valid Address")
	}
}

// CreateEthereumWallet test
func TestCreateWallet(t *testing.T) {
	_, err := EthereumWallet.CreateEthereumWallet(node, invalidPrivateKey)
	if err == nil {
		t.Errorf("CreateEthereumWallet error was incorect, got: %q, want: %q.", err, "not nil")
	}

	w, err := EthereumWallet.CreateEthereumWallet(node, validPrivateKey)
	if err != nil {
		t.Errorf("CreateEthereumWallet error was incorect, got: %q, want: %q.", err, "nil")
	}
	if len(w.PrivateKey) == 0 {
		t.Errorf("CreateEthereumWallet PrivateKey was incorect, got: %q, want: %q.", w.PrivateKey, "valid PrivateKey")
	}
	if len(w.PublicKey) == 0 {
		t.Errorf("CreateEthereumWallet PublicKey was incorect, got: %q, want: %q.", w.PublicKey, "valid PublicKey")
	}
	if len(w.Address) == 0 {
		t.Errorf("CreateEthereumWallet Address was incorect, got: %q, want: %q.", w.Address, "valid Address")
	}
	if len(w.Address) == 0 {
		t.Errorf("CreateEthereumWallet AddressBase58 was incorect, got: %q, want: %q.", w.Address, "valid Address")
	}
}

// PrivateKeyRCDSA test
func TestPrivateKeyRCDSA(t *testing.T) {
	w := wallet()

	_, err := w.PrivateKeyRCDSA()
	if err != nil {
		t.Errorf("PrivateKeyRCDSA error was incorect, got: %q, want: %q.", err, "nil")
	}
}

// PrivateKeyBytes test
func TestPrivateKeyBytes(t *testing.T) {
	w := wallet()

	bytes, err := w.PrivateKeyBytes()
	if err != nil {
		t.Errorf("PrivateKeyBytes error was incorect, got: %q, want: %q.", err, "nil")
	}
	if len(bytes) == 0 {
		t.Errorf("PrivateKeyBytes bytes len was incorect, got: %q, want: %q.", len(bytes), "more than 0")
	}
}

// Balance test
func TestBalance(t *testing.T) {
	w := wallet()

	_, err := w.Balance()
	if err != nil {
		t.Errorf("Balance error was incorect, got: %q, want: %q.", err, "nil")
	}
}

// BalanceTRC20 test
func TestBalanceTRC20(t *testing.T) {
	w := wallet()

	_, err := w.BalanceERC20(token())
	if err != nil {
		t.Errorf("BalanceTRC20 error was incorect, got: %q, want: %q.", err, "nil")
	}
}

// Transfer test
func TestTransfer(t *testing.T) {
	w := wallet()

	/// TODO : uncomment this after checking to Address added
	//_, err := w.Transfer(invalidToAddress, ethAmount)
	//if err == nil {
	//	t.Errorf("Transfer error was incorect, got: %q, want: %q.", err, "not nil becuase to address is invalid")
	//}

	txId, err := w.Transfer(validToAddress, ethAmount)
	fmt.Println(txId)
	if err != nil {
		t.Errorf("Transfer error was incorect, got: %q, want: %q.", err, "nil")
	}
	if len(txId) == 0 {
		t.Errorf("Transfer txId was incorect, got: %q, want: %q.", txId, "not nil")
	}
}

// TransferTRC20 test
func TestTransferTRC20(t *testing.T) {
	w := wallet()
	_t := token()

	/// TODO : uncomment this after checking to Address added
	//_, err := w.TransferERC20(_t, invalidToAddress, erc20Amount)
	//if err == nil {
	//	t.Errorf("TestTransferTRC20 error was incorect, got: %q, want: %q.", err, "not nil becuase to address is invalid")
	//}

	txId, err := w.TransferERC20(_t, validToAddress, erc20Amount)
	fmt.Println(txId)
	if err != nil {
		t.Errorf("Transfer error was incorect, got: %q, want: %q.", err, "nil")
	}
	if len(txId) == 0 {
		t.Errorf("Transfer txId was incorect, got: %q, want: %q.", txId, "not nil")
	}
}
