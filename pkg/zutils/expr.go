package zutils

import (
	"context"
	"github.com/Bishoptylaor/paypay/pkg/zlog"
	"github.com/expr-lang/expr"
)

// Expr 处理空规则直接通过
func Expr(ctx context.Context, rule string, input interface{}) (bool, error) {
	if rule == "" {
		return true, nil
	}
	output, err := expr.Eval(rule, input)
	if err != nil {
		return false, err
	}
	if output == nil {
		zlog.Warnf(ctx, "Expr -> output nil err; rule:%s, input:%+v", rule, input)
		return false, err
	}
	return output.(bool), nil
}
