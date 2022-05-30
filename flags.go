package main

import (
	"fmt"
	"strings"
)

type replacers struct {
	rs []replacer
}

func (rl *replacers) String() string {
	return ""
}

func (rl *replacers) Set(val string) error {
	splitted := strings.Split(val, ":")
	if len(splitted) != 2 {
		return fmt.Errorf("%q should contain exactely one ':'", val)
	}
	k, v := splitted[0], splitted[1]

	rl.rs = append(rl.rs, replacer{key: k, val: v})

	return nil
}
