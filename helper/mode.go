package helper

import "os"

func Mode() string {
	mode, ok := os.LookupEnv("MODE")
	if !ok {
		mode = "dev"
	}
	return mode
}
