package main

import (
	"fmt"
	"strings"
)

const replacementSeperator = "<=>"

type replacements struct {
	rs []replacement
}

func (rl *replacements) String() string {
	return ""
}

func (rl *replacements) Set(val string) error {
	splitted := strings.Split(val, replacementSeperator)
	if len(splitted) != 2 {
		return fmt.Errorf("%q should contain exactely one %q", val, replacementSeperator)
	}
	k, v := splitted[0], splitted[1]

	rl.rs = append(rl.rs, replacement{key: k, val: v})

	return nil
}
