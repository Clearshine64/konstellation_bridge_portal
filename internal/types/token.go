// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BEP20Token

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BEP20TokenMetaData contains all meta data concerning the BEP20Token contract.
var BEP20TokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506000620000246200024060201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3506040518060400160405280600981526020017f44524320546f6b656e0000000000000000000000000000000000000000000000815250600690805190602001906200010f92919062000248565b506040518060400160405280600381526020017f4452430000000000000000000000000000000000000000000000000000000000815250600590805190602001906200015d92919062000248565b506012600460006101000a81548160ff021916908360ff1602179055506a52b7d2dcc80cd2e4000000600381905550600354600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055503373ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6003546040518082815260200191505060405180910390a3620002f7565b600033905090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200028b57805160ff1916838001178555620002bc565b82800160010185558215620002bc579182015b82811115620002bb5782518255916020019190600101906200029e565b5b509050620002cb9190620002cf565b5090565b620002f491905b80821115620002f0576000816000905550600101620002d6565b5090565b90565b611c9380620003076000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c806379cc6790116100ad578063a9059cbb11610071578063a9059cbb146105c7578063b09f12661461062d578063d28d8852146106b0578063dd62ed3e14610733578063f2fde38b146107ab5761012c565b806379cc6790146103fc578063893d20e81461044a5780638da5cb5b1461049457806395d89b41146104de578063a457c2d7146105615761012c565b806332424aa3116100f457806332424aa3146102e2578063395093511461030657806342966c681461036c57806370a082311461039a578063715018a6146103f25761012c565b806306fdde0314610131578063095ea7b3146101b457806318160ddd1461021a57806323b872dd14610238578063313ce567146102be575b600080fd5b6101396107ef565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561017957808201518184015260208101905061015e565b50505050905090810190601f1680156101a65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610200600480360360408110156101ca57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610891565b604051808215151515815260200191505060405180910390f35b6102226108af565b6040518082815260200191505060405180910390f35b6102a46004803603606081101561024e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506108b9565b604051808215151515815260200191505060405180910390f35b6102c6610992565b604051808260ff1660ff16815260200191505060405180910390f35b6102ea6109a9565b604051808260ff1660ff16815260200191505060405180910390f35b6103526004803603604081101561031c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506109bc565b604051808215151515815260200191505060405180910390f35b6103986004803603602081101561038257600080fd5b8101908080359060200190929190505050610a6f565b005b6103dc600480360360208110156103b057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a83565b6040518082815260200191505060405180910390f35b6103fa610acc565b005b6104486004803603604081101561041257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610c54565b005b610452610d29565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61049c610d38565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6104e6610d61565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561052657808201518184015260208101905061050b565b50505050905090810190601f1680156105535780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6105ad6004803603604081101561057757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610e03565b604051808215151515815260200191505060405180910390f35b610613600480360360408110156105dd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610ed0565b604051808215151515815260200191505060405180910390f35b610635610eee565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561067557808201518184015260208101905061065a565b50505050905090810190601f1680156106a25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6106b8610f8c565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106f85780820151818401526020810190506106dd565b50505050905090810190601f1680156107255780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6107956004803603604081101561074957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061102a565b6040518082815260200191505060405180910390f35b6107ed600480360360208110156107c157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506110b1565b005b606060068054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108875780601f1061085c57610100808354040283529160200191610887565b820191906000526020600020905b81548152906001019060200180831161086a57829003601f168201915b5050505050905090565b60006108a561089e611186565b848461118e565b6001905092915050565b6000600354905090565b60006108c6848484611385565b610987846108d2611186565b61098285604051806060016040528060288152602001611b3f60289139600260008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000610938611186565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461163f9092919063ffffffff16565b61118e565b600190509392505050565b6000600460009054906101000a900460ff16905090565b600460009054906101000a900460ff1681565b6000610a656109c9611186565b84610a6085600260006109da611186565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546116ff90919063ffffffff16565b61118e565b6001905092915050565b610a80610a7a611186565b82611787565b50565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b610ad4611186565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610b95576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000610d0682604051806060016040528060248152602001611c1860249139600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000610cbc611186565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461163f9092919063ffffffff16565b9050610d1a83610d14611186565b8361118e565b610d248383611787565b505050565b6000610d33610d38565b905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606060058054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610df95780601f10610dce57610100808354040283529160200191610df9565b820191906000526020600020905b815481529060010190602001808311610ddc57829003601f168201915b5050505050905090565b6000610ec6610e10611186565b84610ec185604051806060016040528060258152602001611bb06025913960026000610e3a611186565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461163f9092919063ffffffff16565b61118e565b6001905092915050565b6000610ee4610edd611186565b8484611385565b6001905092915050565b60058054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f845780601f10610f5957610100808354040283529160200191610f84565b820191906000526020600020905b815481529060010190602001808311610f6757829003601f168201915b505050505081565b60068054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110225780601f10610ff757610100808354040283529160200191611022565b820191906000526020600020905b81548152906001019060200180831161100557829003601f168201915b505050505081565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6110b9611186565b73ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461117a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b61118381611941565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611214576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180611af56024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561129a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180611c3c6022913960400191505060405180910390fd5b80600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561140b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180611ad06025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611491576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180611b8d6023913960400191505060405180910390fd5b6114fd81604051806060016040528060268152602001611b6760269139600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461163f9092919063ffffffff16565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061159281600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546116ff90919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b60008383111582906116ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156116b1578082015181840152602081019050611696565b50505050905090810190601f1680156116de5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b60008082840190508381101561177d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561180d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180611bd56021913960400191505060405180910390fd5b61187981604051806060016040528060228152602001611bf660229139600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461163f9092919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506118d181600354611a8590919063ffffffff16565b600381905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156119c7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180611b196026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000611ac783836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061163f565b90509291505056fe42455032303a207472616e736665722066726f6d20746865207a65726f206164647265737342455032303a20617070726f76652066726f6d20746865207a65726f20616464726573734f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737342455032303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636542455032303a207472616e7366657220616d6f756e7420657863656564732062616c616e636542455032303a207472616e7366657220746f20746865207a65726f206164647265737342455032303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726f42455032303a206275726e2066726f6d20746865207a65726f206164647265737342455032303a206275726e20616d6f756e7420657863656564732062616c616e636542455032303a206275726e20616d6f756e74206578636565647320616c6c6f77616e636542455032303a20617070726f766520746f20746865207a65726f2061646472657373a2646970667358221220761fbad0ef8bd2f15bcf2b8b5ec271a2ddfe7e9702d0f95c6d436bede92450cc64736f6c63430006080033",
}

// BEP20TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use BEP20TokenMetaData.ABI instead.
var BEP20TokenABI = BEP20TokenMetaData.ABI

// BEP20TokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BEP20TokenMetaData.Bin instead.
var BEP20TokenBin = BEP20TokenMetaData.Bin

// DeployBEP20Token deploys a new Ethereum contract, binding an instance of BEP20Token to it.
func DeployBEP20Token(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BEP20Token, error) {
	parsed, err := BEP20TokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BEP20TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BEP20Token{BEP20TokenCaller: BEP20TokenCaller{contract: contract}, BEP20TokenTransactor: BEP20TokenTransactor{contract: contract}, BEP20TokenFilterer: BEP20TokenFilterer{contract: contract}}, nil
}

// BEP20Token is an auto generated Go binding around an Ethereum contract.
type BEP20Token struct {
	BEP20TokenCaller     // Read-only binding to the contract
	BEP20TokenTransactor // Write-only binding to the contract
	BEP20TokenFilterer   // Log filterer for contract events
}

// BEP20TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BEP20TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BEP20TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BEP20TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BEP20TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BEP20TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BEP20TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BEP20TokenSession struct {
	Contract     *BEP20Token       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BEP20TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BEP20TokenCallerSession struct {
	Contract *BEP20TokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BEP20TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BEP20TokenTransactorSession struct {
	Contract     *BEP20TokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BEP20TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BEP20TokenRaw struct {
	Contract *BEP20Token // Generic contract binding to access the raw methods on
}

// BEP20TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BEP20TokenCallerRaw struct {
	Contract *BEP20TokenCaller // Generic read-only contract binding to access the raw methods on
}

// BEP20TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BEP20TokenTransactorRaw struct {
	Contract *BEP20TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBEP20Token creates a new instance of BEP20Token, bound to a specific deployed contract.
func NewBEP20Token(address common.Address, backend bind.ContractBackend) (*BEP20Token, error) {
	contract, err := bindBEP20Token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BEP20Token{BEP20TokenCaller: BEP20TokenCaller{contract: contract}, BEP20TokenTransactor: BEP20TokenTransactor{contract: contract}, BEP20TokenFilterer: BEP20TokenFilterer{contract: contract}}, nil
}

// NewBEP20TokenCaller creates a new read-only instance of BEP20Token, bound to a specific deployed contract.
func NewBEP20TokenCaller(address common.Address, caller bind.ContractCaller) (*BEP20TokenCaller, error) {
	contract, err := bindBEP20Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BEP20TokenCaller{contract: contract}, nil
}

// NewBEP20TokenTransactor creates a new write-only instance of BEP20Token, bound to a specific deployed contract.
func NewBEP20TokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BEP20TokenTransactor, error) {
	contract, err := bindBEP20Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BEP20TokenTransactor{contract: contract}, nil
}

// NewBEP20TokenFilterer creates a new log filterer instance of BEP20Token, bound to a specific deployed contract.
func NewBEP20TokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BEP20TokenFilterer, error) {
	contract, err := bindBEP20Token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BEP20TokenFilterer{contract: contract}, nil
}

// bindBEP20Token binds a generic wrapper to an already deployed contract.
func bindBEP20Token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BEP20TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BEP20Token *BEP20TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BEP20Token.Contract.BEP20TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BEP20Token *BEP20TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20Token.Contract.BEP20TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BEP20Token *BEP20TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BEP20Token.Contract.BEP20TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BEP20Token *BEP20TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BEP20Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BEP20Token *BEP20TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BEP20Token *BEP20TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BEP20Token.Contract.contract.Transact(opts, method, params...)
}

// PrivateDecimals is a free data retrieval call binding the contract method 0x32424aa3.
//
// Solidity: function _decimals() view returns(uint8)
func (_BEP20Token *BEP20TokenCaller) PrivateDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BEP20Token.contract.Call(opts, &out, "_decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PrivateDecimals is a free data retrieval call binding the contract method 0x32424aa3.
//
// Solidity: function _decimals() view returns(uint8)
func (_BEP20Token *BEP20TokenSession) PrivateDecimals() (uint8, error) {
	return _BEP20Token.Contract.PrivateDecimals(&_BEP20Token.CallOpts)
}

// PrivateDecimals is a free data retrieval call binding the contract method 0x32424aa3.
//
// Solidity: function _decimals() view returns(uint8)
func (_BEP20Token *BEP20TokenCallerSession) PrivateDecimals() (uint8, error) {
	return _BEP20Token.Contract.PrivateDecimals(&_BEP20Token.CallOpts)
}

// PrivateName is a free data retrieval call binding the contract method 0xd28d8852.
//
// Solidity: function _name() view returns(string)
func (_BEP20Token *BEP20TokenCaller) PrivateName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BEP20Token.contract.Call(opts, &out, "_name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PrivateName is a free data retrieval call binding the contract method 0xd28d8852.
//
// Solidity: function _name() view returns(string)
func (_BEP20Token *BEP20TokenSession) PrivateName() (string, error) {
	return _BEP20Token.Contract.PrivateName(&_BEP20Token.CallOpts)
}

// PrivateName is a free data retrieval call binding the contract method 0xd28d8852.
//
// Solidity: function _name() view returns(string)
func (_BEP20Token *BEP20TokenCallerSession) PrivateName() (string, error) {
	return _BEP20Token.Contract.PrivateName(&_BEP20Token.CallOpts)
}

// PrivateSymbol is a free data retrieval call binding the contract method 0xb09f1266.
//
// Solidity: function _symbol() view returns(string)
func (_BEP20Token *BEP20TokenCaller) PrivateSymbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BEP20Token.contract.Call(opts, &out, "_symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PrivateSymbol is a free data retrieval call binding the contract method 0xb09f1266.
//
// Solidity: function _symbol() view returns(string)
func (_BEP20Token *BEP20TokenSession) PrivateSymbol() (string, error) {
	return _BEP20Token.Contract.PrivateSymbol(&_BEP20Token.CallOpts)
}

// PrivateSymbol is a free data retrieval call binding the contract method 0xb09f1266.
//
// Solidity: function _symbol() view returns(string)
func (_BEP20Token *BEP20TokenCallerSession) PrivateSymbol() (string, error) {
	return _BEP20Token.Contract.PrivateSymbol(&_BEP20Token.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BEP20Token *BEP20TokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BEP20Token.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BEP20Token *BEP20TokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BEP20Token.Contract.Allowance(&_BEP20Token.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BEP20Token *BEP20TokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BEP20Token.Contract.Allowance(&_BEP20Token.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BEP20Token *BEP20TokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BEP20Token.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BEP20Token *BEP20TokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BEP20Token.Contract.BalanceOf(&_BEP20Token.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BEP20Token *BEP20TokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BEP20Token.Contract.BalanceOf(&_BEP20Token.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BEP20Token *BEP20TokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BEP20Token.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BEP20Token *BEP20TokenSession) Decimals() (uint8, error) {
	return _BEP20Token.Contract.Decimals(&_BEP20Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BEP20Token *BEP20TokenCallerSession) Decimals() (uint8, error) {
	return _BEP20Token.Contract.Decimals(&_BEP20Token.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BEP20Token *BEP20TokenCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BEP20Token.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BEP20Token *BEP20TokenSession) GetOwner() (common.Address, error) {
	return _BEP20Token.Contract.GetOwner(&_BEP20Token.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BEP20Token *BEP20TokenCallerSession) GetOwner() (common.Address, error) {
	return _BEP20Token.Contract.GetOwner(&_BEP20Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BEP20Token *BEP20TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BEP20Token.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BEP20Token *BEP20TokenSession) Name() (string, error) {
	return _BEP20Token.Contract.Name(&_BEP20Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BEP20Token *BEP20TokenCallerSession) Name() (string, error) {
	return _BEP20Token.Contract.Name(&_BEP20Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BEP20Token *BEP20TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BEP20Token.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BEP20Token *BEP20TokenSession) Owner() (common.Address, error) {
	return _BEP20Token.Contract.Owner(&_BEP20Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BEP20Token *BEP20TokenCallerSession) Owner() (common.Address, error) {
	return _BEP20Token.Contract.Owner(&_BEP20Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BEP20Token *BEP20TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BEP20Token.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BEP20Token *BEP20TokenSession) Symbol() (string, error) {
	return _BEP20Token.Contract.Symbol(&_BEP20Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BEP20Token *BEP20TokenCallerSession) Symbol() (string, error) {
	return _BEP20Token.Contract.Symbol(&_BEP20Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BEP20Token *BEP20TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BEP20Token.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BEP20Token *BEP20TokenSession) TotalSupply() (*big.Int, error) {
	return _BEP20Token.Contract.TotalSupply(&_BEP20Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BEP20Token *BEP20TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BEP20Token.Contract.TotalSupply(&_BEP20Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BEP20Token *BEP20TokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20Token.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BEP20Token *BEP20TokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20Token.Contract.Approve(&_BEP20Token.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BEP20Token *BEP20TokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20Token.Contract.Approve(&_BEP20Token.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_BEP20Token *BEP20TokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BEP20Token.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_BEP20Token *BEP20TokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BEP20Token.Contract.Burn(&_BEP20Token.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_BEP20Token *BEP20TokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BEP20Token.Contract.Burn(&_BEP20Token.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_BEP20Token *BEP20TokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20Token.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_BEP20Token *BEP20TokenSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20Token.Contract.BurnFrom(&_BEP20Token.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_BEP20Token *BEP20TokenTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20Token.Contract.BurnFrom(&_BEP20Token.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BEP20Token *BEP20TokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BEP20Token.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BEP20Token *BEP20TokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BEP20Token.Contract.DecreaseAllowance(&_BEP20Token.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BEP20Token *BEP20TokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BEP20Token.Contract.DecreaseAllowance(&_BEP20Token.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BEP20Token *BEP20TokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BEP20Token.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BEP20Token *BEP20TokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BEP20Token.Contract.IncreaseAllowance(&_BEP20Token.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BEP20Token *BEP20TokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BEP20Token.Contract.IncreaseAllowance(&_BEP20Token.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BEP20Token *BEP20TokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20Token.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BEP20Token *BEP20TokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _BEP20Token.Contract.RenounceOwnership(&_BEP20Token.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BEP20Token *BEP20TokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BEP20Token.Contract.RenounceOwnership(&_BEP20Token.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BEP20Token *BEP20TokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20Token.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BEP20Token *BEP20TokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20Token.Contract.Transfer(&_BEP20Token.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BEP20Token *BEP20TokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20Token.Contract.Transfer(&_BEP20Token.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BEP20Token *BEP20TokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20Token.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BEP20Token *BEP20TokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20Token.Contract.TransferFrom(&_BEP20Token.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BEP20Token *BEP20TokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20Token.Contract.TransferFrom(&_BEP20Token.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BEP20Token *BEP20TokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BEP20Token.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BEP20Token *BEP20TokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BEP20Token.Contract.TransferOwnership(&_BEP20Token.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BEP20Token *BEP20TokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BEP20Token.Contract.TransferOwnership(&_BEP20Token.TransactOpts, newOwner)
}

// BEP20TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BEP20Token contract.
type BEP20TokenApprovalIterator struct {
	Event *BEP20TokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BEP20TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20TokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BEP20TokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BEP20TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20TokenApproval represents a Approval event raised by the BEP20Token contract.
type BEP20TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BEP20Token *BEP20TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BEP20TokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BEP20Token.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BEP20TokenApprovalIterator{contract: _BEP20Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BEP20Token *BEP20TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BEP20TokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BEP20Token.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20TokenApproval)
				if err := _BEP20Token.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BEP20Token *BEP20TokenFilterer) ParseApproval(log types.Log) (*BEP20TokenApproval, error) {
	event := new(BEP20TokenApproval)
	if err := _BEP20Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BEP20TokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BEP20Token contract.
type BEP20TokenOwnershipTransferredIterator struct {
	Event *BEP20TokenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BEP20TokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20TokenOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BEP20TokenOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BEP20TokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20TokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20TokenOwnershipTransferred represents a OwnershipTransferred event raised by the BEP20Token contract.
type BEP20TokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BEP20Token *BEP20TokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BEP20TokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BEP20Token.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BEP20TokenOwnershipTransferredIterator{contract: _BEP20Token.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BEP20Token *BEP20TokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BEP20TokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BEP20Token.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20TokenOwnershipTransferred)
				if err := _BEP20Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BEP20Token *BEP20TokenFilterer) ParseOwnershipTransferred(log types.Log) (*BEP20TokenOwnershipTransferred, error) {
	event := new(BEP20TokenOwnershipTransferred)
	if err := _BEP20Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BEP20TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BEP20Token contract.
type BEP20TokenTransferIterator struct {
	Event *BEP20TokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BEP20TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20TokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BEP20TokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BEP20TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20TokenTransfer represents a Transfer event raised by the BEP20Token contract.
type BEP20TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BEP20Token *BEP20TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BEP20TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BEP20Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BEP20TokenTransferIterator{contract: _BEP20Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BEP20Token *BEP20TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BEP20TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BEP20Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20TokenTransfer)
				if err := _BEP20Token.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BEP20Token *BEP20TokenFilterer) ParseTransfer(log types.Log) (*BEP20TokenTransfer, error) {
	event := new(BEP20TokenTransfer)
	if err := _BEP20Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
