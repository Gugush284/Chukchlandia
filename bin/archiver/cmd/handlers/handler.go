package handler

import (
	"fmt"
	"os"
)

func handlerErr(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
