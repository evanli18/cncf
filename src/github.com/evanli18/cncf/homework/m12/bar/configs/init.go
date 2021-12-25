package configs

import "os"

func init() {
	if os.Getenv("AUTHSERVER") != "" {
		AUTHSERVER = os.Getenv("AUTHSERVER")
	}
	if os.Getenv("USERSERVER") != "" {
		USERSERVER = os.Getenv("USERSERVER")
	}
	if os.Getenv("ORDERSERVER") != "" {
		ORDERSERVER = os.Getenv("ORDERSERVER")
	}
	if os.Getenv("BFFSERVER") != "" {
		BFFSERVER = os.Getenv("BFFSERVER")
	}
}
