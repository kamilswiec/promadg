package parse

import (
	"testing"
)

func TestRulesPageToMD(t *testing.T) {
	var rp = RulesPage{}
	RulesPageToMD(rp)
}