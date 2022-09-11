package fuzztraversal

import (
	"github.com/Cracked5pider/Havoc/teamserver/pkg/profile/yaotl"
	"github.com/Cracked5pider/Havoc/teamserver/pkg/profile/yaotl/hclsyntax"
)

func Fuzz(data []byte) int {
	_, diags := hclsyntax.ParseTraversalAbs(data, "<fuzz-trav>", hcl.Pos{Line: 1, Column: 1})

	if diags.HasErrors() {
		return 0
	}

	return 1
}