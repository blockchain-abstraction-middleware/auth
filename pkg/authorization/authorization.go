package authorization

import (
	"github.com/blockchain-abstraction-middleware/auth/pkg/contracts/Authorization"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Manager contains the login for the Manager Contract
type Manager struct {
	ec *ethclient.Client

	auth *bind.TransactOpts

	Manager *Authorization.AuthorizationFilterer

	address string
}

// NewAuthorization Creates a authorization contract with a given address and private key
func NewAuthorization(ec *ethclient.Client, aAddress string, auth *bind.TransactOpts) (*Manager, error) {
	authorization, err := Authorization.NewAuthorizationFilterer(common.HexToAddress(aAddress), ec)

	return &Manager{
		auth:    auth,
		ec:      ec,
		Manager: authorization,
		address: aAddress,
	}, err
}

// GetAllGameJamAddresses gets all deployed game jam addresses
// func (m *Manager) GetAllGameJamAddresses() ([]common.Address, error) {
// 	return m.manager.GetAllGameJamAddresses(&bind.CallOpts{Pending: true})
// }
