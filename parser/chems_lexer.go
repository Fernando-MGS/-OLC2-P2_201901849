// Code generated from ChemsLexer.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 33, 182,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 6, 9, 111, 10, 9, 13, 9, 14,
	9, 112, 3, 10, 3, 10, 7, 10, 117, 10, 10, 12, 10, 14, 10, 120, 11, 10,
	3, 10, 3, 10, 3, 11, 3, 11, 7, 11, 126, 10, 11, 12, 11, 14, 11, 129, 11,
	11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26,
	3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3,
	31, 3, 32, 6, 32, 174, 10, 32, 13, 32, 14, 32, 175, 3, 32, 3, 32, 3, 33,
	3, 33, 3, 33, 2, 2, 34, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17,
	10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35,
	19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53,
	28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 2, 3, 2, 8, 3, 2, 50, 59,
	3, 2, 36, 36, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97,
	97, 99, 124, 6, 2, 11, 12, 15, 15, 34, 34, 94, 94, 9, 2, 34, 35, 37, 37,
	45, 45, 47, 48, 60, 60, 66, 66, 93, 95, 2, 184, 2, 3, 3, 2, 2, 2, 2, 5,
	3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2,
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2,
	2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2,
	2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3,
	2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59,
	3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 3, 67, 3, 2, 2, 2, 5,
	75, 3, 2, 2, 2, 7, 79, 3, 2, 2, 2, 9, 86, 3, 2, 2, 2, 11, 93, 3, 2, 2,
	2, 13, 96, 3, 2, 2, 2, 15, 102, 3, 2, 2, 2, 17, 110, 3, 2, 2, 2, 19, 114,
	3, 2, 2, 2, 21, 123, 3, 2, 2, 2, 23, 130, 3, 2, 2, 2, 25, 132, 3, 2, 2,
	2, 27, 134, 3, 2, 2, 2, 29, 136, 3, 2, 2, 2, 31, 138, 3, 2, 2, 2, 33, 140,
	3, 2, 2, 2, 35, 142, 3, 2, 2, 2, 37, 145, 3, 2, 2, 2, 39, 148, 3, 2, 2,
	2, 41, 150, 3, 2, 2, 2, 43, 152, 3, 2, 2, 2, 45, 154, 3, 2, 2, 2, 47, 156,
	3, 2, 2, 2, 49, 158, 3, 2, 2, 2, 51, 160, 3, 2, 2, 2, 53, 162, 3, 2, 2,
	2, 55, 164, 3, 2, 2, 2, 57, 166, 3, 2, 2, 2, 59, 168, 3, 2, 2, 2, 61, 170,
	3, 2, 2, 2, 63, 173, 3, 2, 2, 2, 65, 179, 3, 2, 2, 2, 67, 68, 7, 101, 2,
	2, 68, 69, 7, 113, 2, 2, 69, 70, 7, 112, 2, 2, 70, 71, 7, 117, 2, 2, 71,
	72, 7, 113, 2, 2, 72, 73, 7, 110, 2, 2, 73, 74, 7, 103, 2, 2, 74, 4, 3,
	2, 2, 2, 75, 76, 7, 110, 2, 2, 76, 77, 7, 113, 2, 2, 77, 78, 7, 105, 2,
	2, 78, 6, 3, 2, 2, 2, 79, 80, 7, 112, 2, 2, 80, 81, 7, 119, 2, 2, 81, 82,
	7, 111, 2, 2, 82, 83, 7, 100, 2, 2, 83, 84, 7, 103, 2, 2, 84, 85, 7, 116,
	2, 2, 85, 8, 3, 2, 2, 2, 86, 87, 7, 117, 2, 2, 87, 88, 7, 118, 2, 2, 88,
	89, 7, 116, 2, 2, 89, 90, 7, 107, 2, 2, 90, 91, 7, 112, 2, 2, 91, 92, 7,
	105, 2, 2, 92, 10, 3, 2, 2, 2, 93, 94, 7, 107, 2, 2, 94, 95, 7, 104, 2,
	2, 95, 12, 3, 2, 2, 2, 96, 97, 7, 121, 2, 2, 97, 98, 7, 106, 2, 2, 98,
	99, 7, 107, 2, 2, 99, 100, 7, 110, 2, 2, 100, 101, 7, 103, 2, 2, 101, 14,
	3, 2, 2, 2, 102, 103, 7, 117, 2, 2, 103, 104, 7, 118, 2, 2, 104, 105, 7,
	116, 2, 2, 105, 106, 7, 119, 2, 2, 106, 107, 7, 101, 2, 2, 107, 108, 7,
	118, 2, 2, 108, 16, 3, 2, 2, 2, 109, 111, 9, 2, 2, 2, 110, 109, 3, 2, 2,
	2, 111, 112, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113,
	18, 3, 2, 2, 2, 114, 118, 7, 36, 2, 2, 115, 117, 10, 3, 2, 2, 116, 115,
	3, 2, 2, 2, 117, 120, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2,
	2, 2, 119, 121, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 121, 122, 7, 36, 2, 2,
	122, 20, 3, 2, 2, 2, 123, 127, 9, 4, 2, 2, 124, 126, 9, 5, 2, 2, 125, 124,
	3, 2, 2, 2, 126, 129, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 127, 128, 3, 2,
	2, 2, 128, 22, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 130, 131, 7, 48, 2, 2,
	131, 24, 3, 2, 2, 2, 132, 133, 7, 61, 2, 2, 133, 26, 3, 2, 2, 2, 134, 135,
	7, 46, 2, 2, 135, 28, 3, 2, 2, 2, 136, 137, 7, 60, 2, 2, 137, 30, 3, 2,
	2, 2, 138, 139, 7, 35, 2, 2, 139, 32, 3, 2, 2, 2, 140, 141, 7, 63, 2, 2,
	141, 34, 3, 2, 2, 2, 142, 143, 7, 64, 2, 2, 143, 144, 7, 63, 2, 2, 144,
	36, 3, 2, 2, 2, 145, 146, 7, 62, 2, 2, 146, 147, 7, 63, 2, 2, 147, 38,
	3, 2, 2, 2, 148, 149, 7, 64, 2, 2, 149, 40, 3, 2, 2, 2, 150, 151, 7, 62,
	2, 2, 151, 42, 3, 2, 2, 2, 152, 153, 7, 44, 2, 2, 153, 44, 3, 2, 2, 2,
	154, 155, 7, 49, 2, 2, 155, 46, 3, 2, 2, 2, 156, 157, 7, 45, 2, 2, 157,
	48, 3, 2, 2, 2, 158, 159, 7, 47, 2, 2, 159, 50, 3, 2, 2, 2, 160, 161, 7,
	42, 2, 2, 161, 52, 3, 2, 2, 2, 162, 163, 7, 43, 2, 2, 163, 54, 3, 2, 2,
	2, 164, 165, 7, 125, 2, 2, 165, 56, 3, 2, 2, 2, 166, 167, 7, 127, 2, 2,
	167, 58, 3, 2, 2, 2, 168, 169, 7, 93, 2, 2, 169, 60, 3, 2, 2, 2, 170, 171,
	7, 95, 2, 2, 171, 62, 3, 2, 2, 2, 172, 174, 9, 6, 2, 2, 173, 172, 3, 2,
	2, 2, 174, 175, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2,
	176, 177, 3, 2, 2, 2, 177, 178, 8, 32, 2, 2, 178, 64, 3, 2, 2, 2, 179,
	180, 7, 94, 2, 2, 180, 181, 9, 7, 2, 2, 181, 66, 3, 2, 2, 2, 7, 2, 112,
	118, 127, 175, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'console'", "'log'", "'number'", "'string'", "'if'", "'while'", "'struct'",
	"", "", "", "'.'", "';'", "','", "':'", "'!'", "'='", "'>='", "'<='", "'>'",
	"'<'", "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'{'", "'}'", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "CONSOLE", "LOG", "P_NUMBER", "P_STRING", "P_IF", "P_WHILE", "P_STRUCT",
	"NUMBER", "STRING", "ID", "PUNTO", "PTCOMA", "COMA", "DOSPT", "DIFERENTE",
	"IGUAL", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", "MUL", "DIV", "ADD",
	"SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "CORIZQ", "CORDER",
	"WHITESPACE",
}

var lexerRuleNames = []string{
	"CONSOLE", "LOG", "P_NUMBER", "P_STRING", "P_IF", "P_WHILE", "P_STRUCT",
	"NUMBER", "STRING", "ID", "PUNTO", "PTCOMA", "COMA", "DOSPT", "DIFERENTE",
	"IGUAL", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", "MUL", "DIV", "ADD",
	"SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "CORIZQ", "CORDER",
	"WHITESPACE", "ESC_SEQ",
}

type ChemsLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewChemsLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *ChemsLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewChemsLexer(input antlr.CharStream) *ChemsLexer {
	l := new(ChemsLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "ChemsLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ChemsLexer tokens.
const (
	ChemsLexerCONSOLE    = 1
	ChemsLexerLOG        = 2
	ChemsLexerP_NUMBER   = 3
	ChemsLexerP_STRING   = 4
	ChemsLexerP_IF       = 5
	ChemsLexerP_WHILE    = 6
	ChemsLexerP_STRUCT   = 7
	ChemsLexerNUMBER     = 8
	ChemsLexerSTRING     = 9
	ChemsLexerID         = 10
	ChemsLexerPUNTO      = 11
	ChemsLexerPTCOMA     = 12
	ChemsLexerCOMA       = 13
	ChemsLexerDOSPT      = 14
	ChemsLexerDIFERENTE  = 15
	ChemsLexerIGUAL      = 16
	ChemsLexerMAYORIGUAL = 17
	ChemsLexerMENORIGUAL = 18
	ChemsLexerMAYOR      = 19
	ChemsLexerMENOR      = 20
	ChemsLexerMUL        = 21
	ChemsLexerDIV        = 22
	ChemsLexerADD        = 23
	ChemsLexerSUB        = 24
	ChemsLexerPARIZQ     = 25
	ChemsLexerPARDER     = 26
	ChemsLexerLLAVEIZQ   = 27
	ChemsLexerLLAVEDER   = 28
	ChemsLexerCORIZQ     = 29
	ChemsLexerCORDER     = 30
	ChemsLexerWHITESPACE = 31
)
