// Code generated from FmtParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // FmtParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FmtParserListener is a complete listener for a parse tree produced by FmtParserParser.
type FmtParserListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterYes_stmt is called when entering the yes_stmt production.
	EnterYes_stmt(c *Yes_stmtContext)

	// EnterNo_stmt is called when entering the no_stmt production.
	EnterNo_stmt(c *No_stmtContext)

	// EnterCmp_exp_body is called when entering the cmp_exp_body production.
	EnterCmp_exp_body(c *Cmp_exp_bodyContext)

	// EnterCmp_exp is called when entering the cmp_exp production.
	EnterCmp_exp(c *Cmp_expContext)

	// EnterPattern_exp_body is called when entering the pattern_exp_body production.
	EnterPattern_exp_body(c *Pattern_exp_bodyContext)

	// EnterBool_exp_body is called when entering the bool_exp_body production.
	EnterBool_exp_body(c *Bool_exp_bodyContext)

	// EnterBool_exp is called when entering the bool_exp production.
	EnterBool_exp(c *Bool_expContext)

	// EnterSimp_exp is called when entering the simp_exp production.
	EnterSimp_exp(c *Simp_expContext)

	// EnterText is called when entering the text production.
	EnterText(c *TextContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterEq is called when entering the eq production.
	EnterEq(c *EqContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitYes_stmt is called when exiting the yes_stmt production.
	ExitYes_stmt(c *Yes_stmtContext)

	// ExitNo_stmt is called when exiting the no_stmt production.
	ExitNo_stmt(c *No_stmtContext)

	// ExitCmp_exp_body is called when exiting the cmp_exp_body production.
	ExitCmp_exp_body(c *Cmp_exp_bodyContext)

	// ExitCmp_exp is called when exiting the cmp_exp production.
	ExitCmp_exp(c *Cmp_expContext)

	// ExitPattern_exp_body is called when exiting the pattern_exp_body production.
	ExitPattern_exp_body(c *Pattern_exp_bodyContext)

	// ExitBool_exp_body is called when exiting the bool_exp_body production.
	ExitBool_exp_body(c *Bool_exp_bodyContext)

	// ExitBool_exp is called when exiting the bool_exp production.
	ExitBool_exp(c *Bool_expContext)

	// ExitSimp_exp is called when exiting the simp_exp production.
	ExitSimp_exp(c *Simp_expContext)

	// ExitText is called when exiting the text production.
	ExitText(c *TextContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitEq is called when exiting the eq production.
	ExitEq(c *EqContext)
}
