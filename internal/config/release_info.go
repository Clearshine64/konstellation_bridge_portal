package config

import "os"

// ReleaseInfo struct should contain
// information about current release
type ReleaseInfo struct {
	Service string `json:"service"`
	Env     string `json:"env"`
	Version string `json:"version"`
	Hash    string `json:"hash"`
}

func NewReleaseInfo() *ReleaseInfo {
	return &ReleaseInfo{
		Service: os.Getenv("APP_SERVICE"),
		Env:     os.Getenv("APP_ENV"),
		Version: os.Getenv("APP_VERSION"),
		Hash:    os.Getenv("APP_HASH"),
	}
}
