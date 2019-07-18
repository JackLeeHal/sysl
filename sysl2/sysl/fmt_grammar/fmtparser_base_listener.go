// Code generated from FmtParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // FmtParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFmtParserListener is a complete listener for a parse tree produced by FmtParserParser.
type BaseFmtParserListener struct{}

var _ FmtParserListener = &BaseFmtParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFmtParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFmtParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFmtParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFmtParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFmtParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFmtParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterYes_stmt is called when production yes_stmt is entered.
func (s *BaseFmtParserListener) EnterYes_stmt(ctx *Yes_stmtContext) {}

// ExitYes_stmt is called when production yes_stmt is exited.
func (s *BaseFmtParserListener) ExitYes_stmt(ctx *Yes_stmtContext) {}

// EnterNo_stmt is called when production no_stmt is entered.
func (s *BaseFmtParserListener) EnterNo_stmt(ctx *No_stmtContext) {}

// ExitNo_stmt is called when production no_stmt is exited.
func (s *BaseFmtParserListener) ExitNo_stmt(ctx *No_stmtContext) {}

// EnterCmp_exp_body is called when production cmp_exp_body is entered.
func (s *BaseFmtParserListener) EnterCmp_exp_body(ctx *Cmp_exp_bodyContext) {}

// ExitCmp_exp_body is called when production cmp_exp_body is exited.
func (s *BaseFmtParserListener) ExitCmp_exp_body(ctx *Cmp_exp_bodyContext) {}

// EnterCmp_exp is called when production cmp_exp is entered.
func (s *BaseFmtParserListener) EnterCmp_exp(ctx *Cmp_expContext) {}

// ExitCmp_exp is called when production cmp_exp is exited.
func (s *BaseFmtParserListener) ExitCmp_exp(ctx *Cmp_expContext) {}

// EnterPattern_exp_body is called when production pattern_exp_body is entered.
func (s *BaseFmtParserListener) EnterPattern_exp_body(ctx *Pattern_exp_bodyContext) {}

// ExitPattern_exp_body is called when production pattern_exp_body is exited.
func (s *BaseFmtParserListener) ExitPattern_exp_body(ctx *Pattern_exp_bodyContext) {}

// EnterBool_exp_body is called when production bool_exp_body is entered.
func (s *BaseFmtParserListener) EnterBool_exp_body(ctx *Bool_exp_bodyContext) {}

// ExitBool_exp_body is called when production bool_exp_body is exited.
func (s *BaseFmtParserListener) ExitBool_exp_body(ctx *Bool_exp_bodyContext) {}

// EnterBool_exp is called when production bool_exp is entered.
func (s *BaseFmtParserListener) EnterBool_exp(ctx *Bool_expContext) {}

// ExitBool_exp is called when production bool_exp is exited.
func (s *BaseFmtParserListener) ExitBool_exp(ctx *Bool_expContext) {}

// EnterSimp_exp is called when production simp_exp is entered.
func (s *BaseFmtParserListener) EnterSimp_exp(ctx *Simp_expContext) {}

// ExitSimp_exp is called when production simp_exp is exited.
func (s *BaseFmtParserListener) ExitSimp_exp(ctx *Simp_expContext) {}

// EnterText is called when production text is entered.
func (s *BaseFmtParserListener) EnterText(ctx *TextContext) {}

// ExitText is called when production text is exited.
func (s *BaseFmtParserListener) ExitText(ctx *TextContext) {}

// EnterParam is called when production param is entered.
func (s *BaseFmtParserListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseFmtParserListener) ExitParam(ctx *ParamContext) {}

// EnterValue is called when production value is entered.
func (s *BaseFmtParserListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseFmtParserListener) ExitValue(ctx *ValueContext) {}

// EnterEq is called when production eq is entered.
func (s *BaseFmtParserListener) EnterEq(ctx *EqContext) {}

// ExitEq is called when production eq is exited.
func (s *BaseFmtParserListener) ExitEq(ctx *EqContext) {}
