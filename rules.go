package crond

import (
	"strings"

	"github.com/gorhill/cronexpr"
)

const (
	NATIVE = iota
	SHELL
)

type Rule struct {
	expr *cronexpr.Expression
	cmd  string
	kind int
}

type Rules []Rule

func (r *Rule) Cmd() string { return r.cmd }

func parseCronLine(line string) (Rule, error) {
	rule := Rule{}

	elems := strings.Fields(line)

	exprString := strings.Join(elems[:5], " ")
	expr, err := cronexpr.Parse(exprString)
	if err != nil {
		return rule, err
	}

	rule.expr = expr
	rule.kind = SHELL
	rule.cmd = strings.Join(elems[5:], " ")

	return rule, nil
}
