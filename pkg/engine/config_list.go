// Copyright 2017, Pulumi Corporation.  All rights reserved.

package engine

import (
	"fmt"
	"sort"

	"github.com/pulumi/pulumi-fabric/pkg/tokens"
)

func (eng *Engine) ListConfig(envName string) error {
	info, err := eng.initEnvCmdName(tokens.QName(envName), "")
	if err != nil {
		return err
	}
	config := info.Target.Config

	if config != nil {
		fmt.Fprintf(eng.Stdout, "%-32s %-32s\n", "KEY", "VALUE")
		var keys []string
		for key := range info.Target.Config {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			v := info.Target.Config[key]
			// TODO[pulumi/pulumi-fabric#113]: print complex values.
			fmt.Fprintf(eng.Stdout, "%-32s %-32s\n", key, v)
		}
	}
	return nil
}
