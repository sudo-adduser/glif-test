package common

import "log"

func WriteLogger(n int, err error) {
	if err != nil {
		log.Printf("[response:writer] failed to write %v", err)
	}
}
