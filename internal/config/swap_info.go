package config

import "os"

type SwapInfo struct {
	Knstl *KnstlInfo `json:"knstl"`
	Bsc   *BscInfo   `json:"bsc"`
}

type KnstlInfo struct {
	KnstlNodeGrpcUrl  string `json:"knstl_node_grpc_url"`
	KnstlNodeUrl      string `json:"knstl_node_url"`
	KnstlSwapAddr     string `json:"knstl_swap_addr"`
	KnstlSwapMnemonic string `json:"knstl_swap_mnemonic"`
}

func NewKnstlInfo() *KnstlInfo {
	return &KnstlInfo{
		KnstlNodeGrpcUrl:  os.Getenv("KNSTL_GRPC"),
		KnstlNodeUrl:      os.Getenv("KNSTL_RPC"),
		KnstlSwapAddr:     os.Getenv("KNSTL_CORPORATE_ADDR"),
		KnstlSwapMnemonic: os.Getenv("KNSTL_SWAP_ADDR_MNEMONIC"),
	}
}

type BscInfo struct {
	BscTransactionScanApiUrl string `json:"bsc_transaction_scan_api_url"`
	BscNodeUrl               string `json:"bsc_node_url"`
	BEP20ContractAddr        string `json:"bep20_contract_addr"`
	BscCorporateAddr         string `json:"bsc_corporate_addr"`
	BscCorporateAddrPrivKey  string `json:"bsc_corporate_addr_priv_key"`
}

func NewBscInfo() *BscInfo {
	return &BscInfo{
		BscTransactionScanApiUrl: os.Getenv("BSC_TRANSACTION_API_URL"),
		BscNodeUrl:               os.Getenv("BSC_RPC"),
		BEP20ContractAddr:        os.Getenv("BSC_BEP20_CONTRACT_ADDR"),
		BscCorporateAddr:         os.Getenv("BSC_CORPORATE_ADDR"),
		BscCorporateAddrPrivKey:  os.Getenv("BSC_CORPORATE_ADDR_PRIV_KEY"),
	}
}

func NewSwapInfo() *SwapInfo {
	return &SwapInfo{
		Knstl: NewKnstlInfo(),
		Bsc:   NewBscInfo(),
	}
}
