// Code generated from FmtParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // FmtParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 146,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 3, 2, 7, 2, 30, 10, 2, 12, 2, 14, 2, 33, 11, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 5, 2, 39, 10, 2, 3, 2, 7, 2, 42, 10, 2, 12, 2, 14,
	2, 45, 11, 2, 7, 2, 47, 10, 2, 12, 2, 14, 2, 50, 11, 2, 3, 3, 7, 3, 53,
	10, 3, 12, 3, 14, 3, 56, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 62, 10, 3,
	3, 3, 7, 3, 65, 10, 3, 12, 3, 14, 3, 68, 11, 3, 7, 3, 70, 10, 3, 12, 3,
	14, 3, 73, 11, 3, 3, 4, 7, 4, 76, 10, 4, 12, 4, 14, 4, 79, 11, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 5, 4, 85, 10, 4, 3, 4, 7, 4, 88, 10, 4, 12, 4, 14, 4,
	91, 11, 4, 7, 4, 93, 10, 4, 12, 4, 14, 4, 96, 11, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 105, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 118, 10, 7, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 5, 8, 125, 10, 8, 3, 9, 3, 9, 3, 9, 5, 9, 130, 10, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 14, 3, 14, 3, 14, 2, 2, 15, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 2, 3, 3, 2, 10, 12, 2, 154, 2, 31, 3, 2, 2, 2, 4, 54, 3, 2,
	2, 2, 6, 77, 3, 2, 2, 2, 8, 97, 3, 2, 2, 2, 10, 106, 3, 2, 2, 2, 12, 110,
	3, 2, 2, 2, 14, 119, 3, 2, 2, 2, 16, 126, 3, 2, 2, 2, 18, 133, 3, 2, 2,
	2, 20, 137, 3, 2, 2, 2, 22, 139, 3, 2, 2, 2, 24, 141, 3, 2, 2, 2, 26, 143,
	3, 2, 2, 2, 28, 30, 5, 20, 11, 2, 29, 28, 3, 2, 2, 2, 30, 33, 3, 2, 2,
	2, 31, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 48, 3, 2, 2, 2, 33, 31,
	3, 2, 2, 2, 34, 39, 5, 18, 10, 2, 35, 39, 5, 16, 9, 2, 36, 39, 5, 10, 6,
	2, 37, 39, 5, 20, 11, 2, 38, 34, 3, 2, 2, 2, 38, 35, 3, 2, 2, 2, 38, 36,
	3, 2, 2, 2, 38, 37, 3, 2, 2, 2, 39, 43, 3, 2, 2, 2, 40, 42, 5, 20, 11,
	2, 41, 40, 3, 2, 2, 2, 42, 45, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44,
	3, 2, 2, 2, 44, 47, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 46, 38, 3, 2, 2, 2,
	47, 50, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 3, 3, 2,
	2, 2, 50, 48, 3, 2, 2, 2, 51, 53, 5, 20, 11, 2, 52, 51, 3, 2, 2, 2, 53,
	56, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 71, 3, 2, 2,
	2, 56, 54, 3, 2, 2, 2, 57, 62, 5, 18, 10, 2, 58, 62, 5, 16, 9, 2, 59, 62,
	5, 10, 6, 2, 60, 62, 5, 20, 11, 2, 61, 57, 3, 2, 2, 2, 61, 58, 3, 2, 2,
	2, 61, 59, 3, 2, 2, 2, 61, 60, 3, 2, 2, 2, 62, 66, 3, 2, 2, 2, 63, 65,
	5, 20, 11, 2, 64, 63, 3, 2, 2, 2, 65, 68, 3, 2, 2, 2, 66, 64, 3, 2, 2,
	2, 66, 67, 3, 2, 2, 2, 67, 70, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 61,
	3, 2, 2, 2, 70, 73, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2,
	72, 5, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 74, 76, 5, 20, 11, 2, 75, 74, 3,
	2, 2, 2, 76, 79, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78,
	94, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 80, 85, 5, 18, 10, 2, 81, 85, 5, 16,
	9, 2, 82, 85, 5, 10, 6, 2, 83, 85, 5, 20, 11, 2, 84, 80, 3, 2, 2, 2, 84,
	81, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 84, 83, 3, 2, 2, 2, 85, 89, 3, 2, 2,
	2, 86, 88, 5, 20, 11, 2, 87, 86, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87,
	3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 93, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2,
	92, 84, 3, 2, 2, 2, 93, 96, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 94, 95, 3,
	2, 2, 2, 95, 7, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 97, 98, 5, 22, 12, 2, 98,
	99, 5, 26, 14, 2, 99, 100, 5, 24, 13, 2, 100, 101, 7, 5, 2, 2, 101, 104,
	5, 4, 3, 2, 102, 103, 7, 8, 2, 2, 103, 105, 5, 6, 4, 2, 104, 102, 3, 2,
	2, 2, 104, 105, 3, 2, 2, 2, 105, 9, 3, 2, 2, 2, 106, 107, 7, 3, 2, 2, 107,
	108, 5, 8, 5, 2, 108, 109, 7, 4, 2, 2, 109, 11, 3, 2, 2, 2, 110, 111, 5,
	22, 12, 2, 111, 112, 7, 7, 2, 2, 112, 113, 5, 24, 13, 2, 113, 114, 7, 5,
	2, 2, 114, 117, 5, 4, 3, 2, 115, 116, 7, 8, 2, 2, 116, 118, 5, 6, 4, 2,
	117, 115, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 13, 3, 2, 2, 2, 119, 120,
	5, 22, 12, 2, 120, 121, 7, 5, 2, 2, 121, 124, 5, 4, 3, 2, 122, 123, 7,
	8, 2, 2, 123, 125, 5, 6, 4, 2, 124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2,
	2, 125, 15, 3, 2, 2, 2, 126, 129, 7, 3, 2, 2, 127, 130, 5, 14, 8, 2, 128,
	130, 5, 12, 7, 2, 129, 127, 3, 2, 2, 2, 129, 128, 3, 2, 2, 2, 130, 131,
	3, 2, 2, 2, 131, 132, 7, 4, 2, 2, 132, 17, 3, 2, 2, 2, 133, 134, 7, 3,
	2, 2, 134, 135, 5, 22, 12, 2, 135, 136, 7, 4, 2, 2, 136, 19, 3, 2, 2, 2,
	137, 138, 9, 2, 2, 2, 138, 21, 3, 2, 2, 2, 139, 140, 9, 2, 2, 2, 140, 23,
	3, 2, 2, 2, 141, 142, 7, 11, 2, 2, 142, 25, 3, 2, 2, 2, 143, 144, 7, 6,
	2, 2, 144, 27, 3, 2, 2, 2, 18, 31, 38, 43, 48, 54, 61, 66, 71, 77, 84,
	89, 94, 104, 117, 124, 129,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'%('", "')'", "'?'", "", "'~'", "'|'", "'@'",
}
var symbolicNames = []string{
	"", "EXP_OPEN", "EXP_CLOSE", "QM", "DEQ", "PATTERN", "BAR", "AT", "TSTR",
	"QSTRING", "STRING", "SEQ", "ESCAPE",
}

var ruleNames = []string{
	"expression", "yes_stmt", "no_stmt", "cmp_exp_body", "cmp_exp", "pattern_exp_body",
	"bool_exp_body", "bool_exp", "simp_exp", "text", "param", "value", "eq",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type FmtParserParser struct {
	*antlr.BaseParser
}

func NewFmtParserParser(input antlr.TokenStream) *FmtParserParser {
	this := new(FmtParserParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "FmtParser.g4"

	return this
}

// FmtParserParser tokens.
const (
	FmtParserParserEOF       = antlr.TokenEOF
	FmtParserParserEXP_OPEN  = 1
	FmtParserParserEXP_CLOSE = 2
	FmtParserParserQM        = 3
	FmtParserParserDEQ       = 4
	FmtParserParserPATTERN   = 5
	FmtParserParserBAR       = 6
	FmtParserParserAT        = 7
	FmtParserParserTSTR      = 8
	FmtParserParserQSTRING   = 9
	FmtParserParserSTRING    = 10
	FmtParserParserSEQ       = 11
	FmtParserParserESCAPE    = 12
)

// FmtParserParser rules.
const (
	FmtParserParserRULE_expression       = 0
	FmtParserParserRULE_yes_stmt         = 1
	FmtParserParserRULE_no_stmt          = 2
	FmtParserParserRULE_cmp_exp_body     = 3
	FmtParserParserRULE_cmp_exp          = 4
	FmtParserParserRULE_pattern_exp_body = 5
	FmtParserParserRULE_bool_exp_body    = 6
	FmtParserParserRULE_bool_exp         = 7
	FmtParserParserRULE_simp_exp         = 8
	FmtParserParserRULE_text             = 9
	FmtParserParserRULE_param            = 10
	FmtParserParserRULE_value            = 11
	FmtParserParserRULE_eq               = 12
)

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FmtParserParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FmtParserParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllText() []ITextContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITextContext)(nil)).Elem())
	var tst = make([]ITextContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITextContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Text(i int) ITextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *ExpressionContext) AllSimp_exp() []ISimp_expContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISimp_expContext)(nil)).Elem())
	var tst = make([]ISimp_expContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISimp_expContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Simp_exp(i int) ISimp_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimp_expContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISimp_expContext)
}

func (s *ExpressionContext) AllBool_exp() []IBool_expContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBool_expContext)(nil)).Elem())
	var tst = make([]IBool_expContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBool_expContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Bool_exp(i int) IBool_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBool_expContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBool_expContext)
}

func (s *ExpressionContext) AllCmp_exp() []ICmp_expContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICmp_expContext)(nil)).Elem())
	var tst = make([]ICmp_expContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICmp_expContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Cmp_exp(i int) ICmp_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICmp_expContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICmp_expContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *FmtParserParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FmtParserParserRULE_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(26)
				p.Text()
			}

		}
		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FmtParserParserEXP_OPEN)|(1<<FmtParserParserTSTR)|(1<<FmtParserParserQSTRING)|(1<<FmtParserParserSTRING))) != 0 {
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(32)
				p.Simp_exp()
			}

		case 2:
			{
				p.SetState(33)
				p.Bool_exp()
			}

		case 3:
			{
				p.SetState(34)
				p.Cmp_exp()
			}

		case 4:
			{
				p.SetState(35)
				p.Text()
			}

		}
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(38)
					p.Text()
				}

			}
			p.SetState(43)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
		}

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IYes_stmtContext is an interface to support dynamic dispatch.
type IYes_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsYes_stmtContext differentiates from other interfaces.
	IsYes_stmtContext()
}

type Yes_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyYes_stmtContext() *Yes_stmtContext {
	var p = new(Yes_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FmtParserParserRULE_yes_stmt
	return p
}

func (*Yes_stmtContext) IsYes_stmtContext() {}

func NewYes_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Yes_stmtContext {
	var p = new(Yes_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FmtParserParserRULE_yes_stmt

	return p
}

func (s *Yes_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Yes_stmtContext) AllText() []ITextContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITextContext)(nil)).Elem())
	var tst = make([]ITextContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITextContext)
		}
	}

	return tst
}

func (s *Yes_stmtContext) Text(i int) ITextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *Yes_stmtContext) AllSimp_exp() []ISimp_expContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISimp_expContext)(nil)).Elem())
	var tst = make([]ISimp_expContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISimp_expContext)
		}
	}

	return tst
}

func (s *Yes_stmtContext) Simp_exp(i int) ISimp_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimp_expContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISimp_expContext)
}

func (s *Yes_stmtContext) AllBool_exp() []IBool_expContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBool_expContext)(nil)).Elem())
	var tst = make([]IBool_expContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBool_expContext)
		}
	}

	return tst
}

func (s *Yes_stmtContext) Bool_exp(i int) IBool_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBool_expContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBool_expContext)
}

func (s *Yes_stmtContext) AllCmp_exp() []ICmp_expContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICmp_expContext)(nil)).Elem())
	var tst = make([]ICmp_expContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICmp_expContext)
		}
	}

	return tst
}

func (s *Yes_stmtContext) Cmp_exp(i int) ICmp_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICmp_expContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICmp_expContext)
}

func (s *Yes_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Yes_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Yes_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.EnterYes_stmt(s)
	}
}

func (s *Yes_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.ExitYes_stmt(s)
	}
}

func (p *FmtParserParser) Yes_stmt() (localctx IYes_stmtContext) {
	localctx = NewYes_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FmtParserParserRULE_yes_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(49)
				p.Text()
			}

		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FmtParserParserEXP_OPEN)|(1<<FmtParserParserTSTR)|(1<<FmtParserParserQSTRING)|(1<<FmtParserParserSTRING))) != 0 {
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(55)
				p.Simp_exp()
			}

		case 2:
			{
				p.SetState(56)
				p.Bool_exp()
			}

		case 3:
			{
				p.SetState(57)
				p.Cmp_exp()
			}

		case 4:
			{
				p.SetState(58)
				p.Text()
			}

		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(61)
					p.Text()
				}

			}
			p.SetState(66)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INo_stmtContext is an interface to support dynamic dispatch.
type INo_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNo_stmtContext differentiates from other interfaces.
	IsNo_stmtContext()
}

type No_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNo_stmtContext() *No_stmtContext {
	var p = new(No_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FmtParserParserRULE_no_stmt
	return p
}

func (*No_stmtContext) IsNo_stmtContext() {}

func NewNo_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *No_stmtContext {
	var p = new(No_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FmtParserParserRULE_no_stmt

	return p
}

func (s *No_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *No_stmtContext) AllText() []ITextContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITextContext)(nil)).Elem())
	var tst = make([]ITextContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITextContext)
		}
	}

	return tst
}

func (s *No_stmtContext) Text(i int) ITextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *No_stmtContext) AllSimp_exp() []ISimp_expContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISimp_expContext)(nil)).Elem())
	var tst = make([]ISimp_expContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISimp_expContext)
		}
	}

	return tst
}

func (s *No_stmtContext) Simp_exp(i int) ISimp_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimp_expContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISimp_expContext)
}

func (s *No_stmtContext) AllBool_exp() []IBool_expContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBool_expContext)(nil)).Elem())
	var tst = make([]IBool_expContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBool_expContext)
		}
	}

	return tst
}

func (s *No_stmtContext) Bool_exp(i int) IBool_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBool_expContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBool_expContext)
}

func (s *No_stmtContext) AllCmp_exp() []ICmp_expContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICmp_expContext)(nil)).Elem())
	var tst = make([]ICmp_expContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICmp_expContext)
		}
	}

	return tst
}

func (s *No_stmtContext) Cmp_exp(i int) ICmp_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICmp_expContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICmp_expContext)
}

func (s *No_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *No_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *No_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.EnterNo_stmt(s)
	}
}

func (s *No_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.ExitNo_stmt(s)
	}
}

func (p *FmtParserParser) No_stmt() (localctx INo_stmtContext) {
	localctx = NewNo_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FmtParserParserRULE_no_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(72)
				p.Text()
			}

		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FmtParserParserEXP_OPEN)|(1<<FmtParserParserTSTR)|(1<<FmtParserParserQSTRING)|(1<<FmtParserParserSTRING))) != 0 {
		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(78)
				p.Simp_exp()
			}

		case 2:
			{
				p.SetState(79)
				p.Bool_exp()
			}

		case 3:
			{
				p.SetState(80)
				p.Cmp_exp()
			}

		case 4:
			{
				p.SetState(81)
				p.Text()
			}

		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(84)
					p.Text()
				}

			}
			p.SetState(89)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICmp_exp_bodyContext is an interface to support dynamic dispatch.
type ICmp_exp_bodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCmp_exp_bodyContext differentiates from other interfaces.
	IsCmp_exp_bodyContext()
}

type Cmp_exp_bodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmp_exp_bodyContext() *Cmp_exp_bodyContext {
	var p = new(Cmp_exp_bodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FmtParserParserRULE_cmp_exp_body
	return p
}

func (*Cmp_exp_bodyContext) IsCmp_exp_bodyContext() {}

func NewCmp_exp_bodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cmp_exp_bodyContext {
	var p = new(Cmp_exp_bodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FmtParserParserRULE_cmp_exp_body

	return p
}

func (s *Cmp_exp_bodyContext) GetParser() antlr.Parser { return s.parser }

func (s *Cmp_exp_bodyContext) Param() IParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *Cmp_exp_bodyContext) Eq() IEqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqContext)
}

func (s *Cmp_exp_bodyContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Cmp_exp_bodyContext) QM() antlr.TerminalNode {
	return s.GetToken(FmtParserParserQM, 0)
}

func (s *Cmp_exp_bodyContext) Yes_stmt() IYes_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYes_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYes_stmtContext)
}

func (s *Cmp_exp_bodyContext) BAR() antlr.TerminalNode {
	return s.GetToken(FmtParserParserBAR, 0)
}

func (s *Cmp_exp_bodyContext) No_stmt() INo_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INo_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INo_stmtContext)
}

func (s *Cmp_exp_bodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cmp_exp_bodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cmp_exp_bodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.EnterCmp_exp_body(s)
	}
}

func (s *Cmp_exp_bodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.ExitCmp_exp_body(s)
	}
}

func (p *FmtParserParser) Cmp_exp_body() (localctx ICmp_exp_bodyContext) {
	localctx = NewCmp_exp_bodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FmtParserParserRULE_cmp_exp_body)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Param()
	}
	{
		p.SetState(96)
		p.Eq()
	}
	{
		p.SetState(97)
		p.Value()
	}
	{
		p.SetState(98)
		p.Match(FmtParserParserQM)
	}
	{
		p.SetState(99)
		p.Yes_stmt()
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FmtParserParserBAR {
		{
			p.SetState(100)
			p.Match(FmtParserParserBAR)
		}
		{
			p.SetState(101)
			p.No_stmt()
		}

	}

	return localctx
}

// ICmp_expContext is an interface to support dynamic dispatch.
type ICmp_expContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCmp_expContext differentiates from other interfaces.
	IsCmp_expContext()
}

type Cmp_expContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmp_expContext() *Cmp_expContext {
	var p = new(Cmp_expContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FmtParserParserRULE_cmp_exp
	return p
}

func (*Cmp_expContext) IsCmp_expContext() {}

func NewCmp_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cmp_expContext {
	var p = new(Cmp_expContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FmtParserParserRULE_cmp_exp

	return p
}

func (s *Cmp_expContext) GetParser() antlr.Parser { return s.parser }

func (s *Cmp_expContext) EXP_OPEN() antlr.TerminalNode {
	return s.GetToken(FmtParserParserEXP_OPEN, 0)
}

func (s *Cmp_expContext) Cmp_exp_body() ICmp_exp_bodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICmp_exp_bodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICmp_exp_bodyContext)
}

func (s *Cmp_expContext) EXP_CLOSE() antlr.TerminalNode {
	return s.GetToken(FmtParserParserEXP_CLOSE, 0)
}

func (s *Cmp_expContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cmp_expContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cmp_expContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.EnterCmp_exp(s)
	}
}

func (s *Cmp_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.ExitCmp_exp(s)
	}
}

func (p *FmtParserParser) Cmp_exp() (localctx ICmp_expContext) {
	localctx = NewCmp_expContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FmtParserParserRULE_cmp_exp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(FmtParserParserEXP_OPEN)
	}
	{
		p.SetState(105)
		p.Cmp_exp_body()
	}
	{
		p.SetState(106)
		p.Match(FmtParserParserEXP_CLOSE)
	}

	return localctx
}

// IPattern_exp_bodyContext is an interface to support dynamic dispatch.
type IPattern_exp_bodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPattern_exp_bodyContext differentiates from other interfaces.
	IsPattern_exp_bodyContext()
}

type Pattern_exp_bodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPattern_exp_bodyContext() *Pattern_exp_bodyContext {
	var p = new(Pattern_exp_bodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FmtParserParserRULE_pattern_exp_body
	return p
}

func (*Pattern_exp_bodyContext) IsPattern_exp_bodyContext() {}

func NewPattern_exp_bodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pattern_exp_bodyContext {
	var p = new(Pattern_exp_bodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FmtParserParserRULE_pattern_exp_body

	return p
}

func (s *Pattern_exp_bodyContext) GetParser() antlr.Parser { return s.parser }

func (s *Pattern_exp_bodyContext) Param() IParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *Pattern_exp_bodyContext) PATTERN() antlr.TerminalNode {
	return s.GetToken(FmtParserParserPATTERN, 0)
}

func (s *Pattern_exp_bodyContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Pattern_exp_bodyContext) QM() antlr.TerminalNode {
	return s.GetToken(FmtParserParserQM, 0)
}

func (s *Pattern_exp_bodyContext) Yes_stmt() IYes_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYes_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYes_stmtContext)
}

func (s *Pattern_exp_bodyContext) BAR() antlr.TerminalNode {
	return s.GetToken(FmtParserParserBAR, 0)
}

func (s *Pattern_exp_bodyContext) No_stmt() INo_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INo_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INo_stmtContext)
}

func (s *Pattern_exp_bodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pattern_exp_bodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pattern_exp_bodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.EnterPattern_exp_body(s)
	}
}

func (s *Pattern_exp_bodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.ExitPattern_exp_body(s)
	}
}

func (p *FmtParserParser) Pattern_exp_body() (localctx IPattern_exp_bodyContext) {
	localctx = NewPattern_exp_bodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FmtParserParserRULE_pattern_exp_body)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Param()
	}
	{
		p.SetState(109)
		p.Match(FmtParserParserPATTERN)
	}
	{
		p.SetState(110)
		p.Value()
	}
	{
		p.SetState(111)
		p.Match(FmtParserParserQM)
	}
	{
		p.SetState(112)
		p.Yes_stmt()
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FmtParserParserBAR {
		{
			p.SetState(113)
			p.Match(FmtParserParserBAR)
		}
		{
			p.SetState(114)
			p.No_stmt()
		}

	}

	return localctx
}

// IBool_exp_bodyContext is an interface to support dynamic dispatch.
type IBool_exp_bodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBool_exp_bodyContext differentiates from other interfaces.
	IsBool_exp_bodyContext()
}

type Bool_exp_bodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBool_exp_bodyContext() *Bool_exp_bodyContext {
	var p = new(Bool_exp_bodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FmtParserParserRULE_bool_exp_body
	return p
}

func (*Bool_exp_bodyContext) IsBool_exp_bodyContext() {}

func NewBool_exp_bodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bool_exp_bodyContext {
	var p = new(Bool_exp_bodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FmtParserParserRULE_bool_exp_body

	return p
}

func (s *Bool_exp_bodyContext) GetParser() antlr.Parser { return s.parser }

func (s *Bool_exp_bodyContext) Param() IParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *Bool_exp_bodyContext) QM() antlr.TerminalNode {
	return s.GetToken(FmtParserParserQM, 0)
}

func (s *Bool_exp_bodyContext) Yes_stmt() IYes_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYes_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYes_stmtContext)
}

func (s *Bool_exp_bodyContext) BAR() antlr.TerminalNode {
	return s.GetToken(FmtParserParserBAR, 0)
}

func (s *Bool_exp_bodyContext) No_stmt() INo_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INo_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INo_stmtContext)
}

func (s *Bool_exp_bodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bool_exp_bodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bool_exp_bodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.EnterBool_exp_body(s)
	}
}

func (s *Bool_exp_bodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.ExitBool_exp_body(s)
	}
}

func (p *FmtParserParser) Bool_exp_body() (localctx IBool_exp_bodyContext) {
	localctx = NewBool_exp_bodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FmtParserParserRULE_bool_exp_body)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Param()
	}
	{
		p.SetState(118)
		p.Match(FmtParserParserQM)
	}
	{
		p.SetState(119)
		p.Yes_stmt()
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FmtParserParserBAR {
		{
			p.SetState(120)
			p.Match(FmtParserParserBAR)
		}
		{
			p.SetState(121)
			p.No_stmt()
		}

	}

	return localctx
}

// IBool_expContext is an interface to support dynamic dispatch.
type IBool_expContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBool_expContext differentiates from other interfaces.
	IsBool_expContext()
}

type Bool_expContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBool_expContext() *Bool_expContext {
	var p = new(Bool_expContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FmtParserParserRULE_bool_exp
	return p
}

func (*Bool_expContext) IsBool_expContext() {}

func NewBool_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bool_expContext {
	var p = new(Bool_expContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FmtParserParserRULE_bool_exp

	return p
}

func (s *Bool_expContext) GetParser() antlr.Parser { return s.parser }

func (s *Bool_expContext) EXP_OPEN() antlr.TerminalNode {
	return s.GetToken(FmtParserParserEXP_OPEN, 0)
}

func (s *Bool_expContext) EXP_CLOSE() antlr.TerminalNode {
	return s.GetToken(FmtParserParserEXP_CLOSE, 0)
}

func (s *Bool_expContext) Bool_exp_body() IBool_exp_bodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBool_exp_bodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBool_exp_bodyContext)
}

func (s *Bool_expContext) Pattern_exp_body() IPattern_exp_bodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPattern_exp_bodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPattern_exp_bodyContext)
}

func (s *Bool_expContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bool_expContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bool_expContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.EnterBool_exp(s)
	}
}

func (s *Bool_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.ExitBool_exp(s)
	}
}

func (p *FmtParserParser) Bool_exp() (localctx IBool_expContext) {
	localctx = NewBool_expContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FmtParserParserRULE_bool_exp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(FmtParserParserEXP_OPEN)
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(125)
			p.Bool_exp_body()
		}

	case 2:
		{
			p.SetState(126)
			p.Pattern_exp_body()
		}

	}
	{
		p.SetState(129)
		p.Match(FmtParserParserEXP_CLOSE)
	}

	return localctx
}

// ISimp_expContext is an interface to support dynamic dispatch.
type ISimp_expContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimp_expContext differentiates from other interfaces.
	IsSimp_expContext()
}

type Simp_expContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimp_expContext() *Simp_expContext {
	var p = new(Simp_expContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FmtParserParserRULE_simp_exp
	return p
}

func (*Simp_expContext) IsSimp_expContext() {}

func NewSimp_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simp_expContext {
	var p = new(Simp_expContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FmtParserParserRULE_simp_exp

	return p
}

func (s *Simp_expContext) GetParser() antlr.Parser { return s.parser }

func (s *Simp_expContext) EXP_OPEN() antlr.TerminalNode {
	return s.GetToken(FmtParserParserEXP_OPEN, 0)
}

func (s *Simp_expContext) Param() IParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *Simp_expContext) EXP_CLOSE() antlr.TerminalNode {
	return s.GetToken(FmtParserParserEXP_CLOSE, 0)
}

func (s *Simp_expContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simp_expContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simp_expContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.EnterSimp_exp(s)
	}
}

func (s *Simp_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.ExitSimp_exp(s)
	}
}

func (p *FmtParserParser) Simp_exp() (localctx ISimp_expContext) {
	localctx = NewSimp_expContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FmtParserParserRULE_simp_exp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(FmtParserParserEXP_OPEN)
	}
	{
		p.SetState(132)
		p.Param()
	}
	{
		p.SetState(133)
		p.Match(FmtParserParserEXP_CLOSE)
	}

	return localctx
}

// ITextContext is an interface to support dynamic dispatch.
type ITextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTextContext differentiates from other interfaces.
	IsTextContext()
}

type TextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextContext() *TextContext {
	var p = new(TextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FmtParserParserRULE_text
	return p
}

func (*TextContext) IsTextContext() {}

func NewTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextContext {
	var p = new(TextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FmtParserParserRULE_text

	return p
}

func (s *TextContext) GetParser() antlr.Parser { return s.parser }

func (s *TextContext) TSTR() antlr.TerminalNode {
	return s.GetToken(FmtParserParserTSTR, 0)
}

func (s *TextContext) STRING() antlr.TerminalNode {
	return s.GetToken(FmtParserParserSTRING, 0)
}

func (s *TextContext) QSTRING() antlr.TerminalNode {
	return s.GetToken(FmtParserParserQSTRING, 0)
}

func (s *TextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.EnterText(s)
	}
}

func (s *TextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.ExitText(s)
	}
}

func (p *FmtParserParser) Text() (localctx ITextContext) {
	localctx = NewTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FmtParserParserRULE_text)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FmtParserParserTSTR)|(1<<FmtParserParserQSTRING)|(1<<FmtParserParserSTRING))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FmtParserParserRULE_param
	return p
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FmtParserParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) TSTR() antlr.TerminalNode {
	return s.GetToken(FmtParserParserTSTR, 0)
}

func (s *ParamContext) STRING() antlr.TerminalNode {
	return s.GetToken(FmtParserParserSTRING, 0)
}

func (s *ParamContext) QSTRING() antlr.TerminalNode {
	return s.GetToken(FmtParserParserQSTRING, 0)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *FmtParserParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FmtParserParserRULE_param)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FmtParserParserTSTR)|(1<<FmtParserParserQSTRING)|(1<<FmtParserParserSTRING))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FmtParserParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FmtParserParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) QSTRING() antlr.TerminalNode {
	return s.GetToken(FmtParserParserQSTRING, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *FmtParserParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FmtParserParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(FmtParserParserQSTRING)
	}

	return localctx
}

// IEqContext is an interface to support dynamic dispatch.
type IEqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqContext differentiates from other interfaces.
	IsEqContext()
}

type EqContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqContext() *EqContext {
	var p = new(EqContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FmtParserParserRULE_eq
	return p
}

func (*EqContext) IsEqContext() {}

func NewEqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqContext {
	var p = new(EqContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FmtParserParserRULE_eq

	return p
}

func (s *EqContext) GetParser() antlr.Parser { return s.parser }

func (s *EqContext) DEQ() antlr.TerminalNode {
	return s.GetToken(FmtParserParserDEQ, 0)
}

func (s *EqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.EnterEq(s)
	}
}

func (s *EqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FmtParserListener); ok {
		listenerT.ExitEq(s)
	}
}

func (p *FmtParserParser) Eq() (localctx IEqContext) {
	localctx = NewEqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FmtParserParserRULE_eq)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(FmtParserParserDEQ)
	}

	return localctx
}
