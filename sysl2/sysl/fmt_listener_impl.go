package main

import (
	"regexp"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/anz-bank/sysl/sysl2/sysl/fmt_grammar"
)

type FmtTreeShapeListener struct {
	*parser.BaseFmtParserListener
	attrs       map[string]string
	resultStack []string
	result      string
	flagStack   []bool
}

func ParseFmt(attr map[string]string, input string) string {
	var s FmtTreeShapeListener
	s.attrs = attr
	prefixRegex := regexp.MustCompile(`((?:[^%]|%[^(\n]|\n)*?)($|%\()`)

	prefix := prefixRegex.FindStringSubmatch(input)[1]
	if prefix != "" {
		input = strings.Replace(input, prefix, "", -1)
		prefix = removePercentSymbol(prefix)
	}
	is := antlr.NewInputStream(input)
	lexer := parser.NewFmtParserLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewFmtParserParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(&s, p.Expression())
	s.result = prefix + removePercentSymbol(s.result)
	s.result = strings.Replace(s.result, "\n", "\\n", -1)

	return s.result
}

func (s *FmtTreeShapeListener) VisitErrorNode(node antlr.ErrorNode) {
	panic(node.GetText())
}

func (s *FmtTreeShapeListener) EnterSimp_exp(ctx *parser.Simp_expContext) {
	param := ctx.Param().GetText()
	parent := ctx.GetParent().(antlr.RuleContext)
	parentRuleName := checkRuleName(ctx.GetParser().GetRuleNames(), parent.GetRuleIndex())
	if parentRuleName == "expression" || parentRuleName == "" {
		if s.attrs[param] != "" {
			s.result += s.attrs[param]
		}
	} else {
		s.resultStack[len(s.resultStack)-1] += s.attrs[param]
	}
}

func (s *FmtTreeShapeListener) EnterText(ctx *parser.TextContext) {
	parent := ctx.GetParent().(antlr.RuleContext)
	parentRuleName := checkRuleName(ctx.GetParser().GetRuleNames(), parent.GetRuleIndex())
	if parentRuleName == "yes_stmt" || parentRuleName == "no_stmt" {
		s.resultStack[len(s.resultStack)-1] += ctx.GetText()
	}
	if parentRuleName == "expression" {
		s.result += ctx.GetText()
	}
}

func (s *FmtTreeShapeListener) EnterBool_exp_body(ctx *parser.Bool_exp_bodyContext) {
	param := ctx.Param().GetText()

	if s.attrs[param] != "" {
		s.flagStack = append(s.flagStack, true)
	} else {
		s.flagStack = append(s.flagStack, false)
	}
}

func (s *FmtTreeShapeListener) EnterPattern_exp_body(ctx *parser.Pattern_exp_bodyContext) {
	param := ctx.Param().GetText()
	pattern := strings.Replace(strings.Replace(ctx.Value().GetText(), "~/", "", -1), "/", "", -1)

	if s.attrs[param] != "" {
		if strings.Contains(s.attrs[param], pattern) {
			s.flagStack = append(s.flagStack, true)
		} else {
			s.flagStack = append(s.flagStack, false)
		}
	} else {
		s.flagStack = append(s.flagStack, false)
	}
}

func (s *FmtTreeShapeListener) ExitPattern_exp_body(ctx *parser.Pattern_exp_bodyContext) {
	if ctx.No_stmt() == nil {
		s.resultStack = append(s.resultStack, "")
	}
	s.exitFmt()
}

func (s *FmtTreeShapeListener) ExitBool_exp_body(ctx *parser.Bool_exp_bodyContext) {
	if ctx.No_stmt() == nil {
		s.resultStack = append(s.resultStack, "")
	}
	s.exitFmt()
}

func (s *FmtTreeShapeListener) EnterCmp_exp_body(ctx *parser.Cmp_exp_bodyContext) {
	param := ctx.Param().GetText()
	eq := ctx.Eq().GetText()
	value := strings.Replace(ctx.Value().GetText(), "'", "", -1)

	if s.attrs[param] != "" {
		if eq == "==" {
			if s.attrs[param] == value {
				s.flagStack = append(s.flagStack, true)
			} else {
				s.flagStack = append(s.flagStack, false)
			}
		} else {
			if s.attrs[param] == value {
				s.flagStack = append(s.flagStack, false)
			} else {
				s.flagStack = append(s.flagStack, true)
			}
		}
	} else {
		s.flagStack = append(s.flagStack, false)
	}
}

func (s *FmtTreeShapeListener) ExitCmp_exp_body(ctx *parser.Cmp_exp_bodyContext) {
	if ctx.No_stmt() == nil {
		s.resultStack = append(s.resultStack, "")
	}
	s.exitFmt()
}

func (s *FmtTreeShapeListener) EnterYes_stmt(ctx *parser.Yes_stmtContext) {
	s.resultStack = append(s.resultStack, "")
}

func (s *FmtTreeShapeListener) EnterNo_stmt(ctx *parser.No_stmtContext) {
	s.resultStack = append(s.resultStack, "")
}

func checkRuleName(ruleNames []string, index int) string {
	return ruleNames[index]
}

func (s *FmtTreeShapeListener) popStmt() string {
	if len(s.resultStack) < 1 {
		panic("stack is empty unable to pop")
	}
	// Get the last value from the stack.
	result := s.resultStack[len(s.resultStack)-1]
	// Remove the last element from the stack.
	s.resultStack = s.resultStack[:len(s.resultStack)-1]

	return result
}

func (s *FmtTreeShapeListener) popFlag() bool {
	if len(s.flagStack) < 1 {
		panic("stack is empty unable to pop")
	}
	// Get the last value from the stack.
	result := s.flagStack[len(s.flagStack)-1]
	// Remove the last element from the stack.
	s.flagStack = s.flagStack[:len(s.flagStack)-1]

	return result
}

func (s *FmtTreeShapeListener) exitFmt() {
	noStmt := s.popStmt()
	yesStmt := s.popStmt()

	flag := s.popFlag()
	resultLen := len(s.resultStack)
	if flag {
		if resultLen == 0 {
			s.result += yesStmt
		} else {
			s.resultStack[resultLen-1] = s.resultStack[resultLen-1] + yesStmt
		}
	} else {
		if resultLen == 0 {
			s.result += noStmt
		} else {
			s.resultStack[resultLen-1] = s.resultStack[resultLen-1] + noStmt
		}
	}
}
