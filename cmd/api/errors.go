package main

import (
	"fmt"
)

var ErrVersionEnvNotSet = fmt.Errorf("%s environment variable not set", VERSION_ENV)
