package lexer

import (
	set "github.com/deckarep/golang-set/v3"
	uni "unicode"
)

//词法解析

type LexTokenType int
type Status int
type stratTable []Status

const (
	TpAny LexTokenType = iota
	TpNumber
	TpString
	TpKeyword
	TpIdentifier
	// ZeroSymbol 回零符号
	TpZeroSym
	TpOpSym
	TpStructSym
	TpError
)

const (
	StNormal Status = iota
	StNumber
	StComment
	StIdentifier
	StKeyword
	StError
	StZero
)

type Lex struct {
	status   Status
	pos      int
	code     []rune
	length   int
	nodetype LexTokenType
	depth    int
	peek     rune
	buf      []rune
	column   int
	line     int
	Tokens   []Token
}

type Position struct {
	PosLft int
	PosRgt int
	ColLft int
	ColRgt int
	LinLft int
	LinRgt int
}

type Token struct {
	NodeContent string
	NodeType    LexTokenType
	NodeDepth   int
	NodePos     Position
}

type Action interface {
	nextRead() rune
	peekRead() rune
	readNumber()
	readComment()
	readString()
	newToken() Token
}

const (
	FILE_EOF rune = -1
)

func Lexer(code string) /*[]Token*/ {
	codeRune := []rune(code)
	lex := &Lex{code: codeRune, length: len(codeRune)}
	for {
		ch := lex.nextRead()
		if ch == FILE_EOF {
			break
		}

		switch {
		case uni.IsLetter(ch):
			// DFA 有限状态机
			// 此处应该用strings.HasPrefix判断前缀
			// 例如>>= -> 看到>进解析符号状态，然后peek
			// peek后判断前缀，最后用contains匹配
			// TODO(Me) : 需要使用map先判断状态，在状态内部判断模式
			// FIXME!
			set.NewSet[string]()
		}
	}
}

func (lex *Lex) peekRead() rune {
	if lex.pos+1 >= lex.length {
		return FILE_EOF
	}
	return lex.code[lex.pos+1]
}

func (lex *Lex) nextRead() rune {
	if lex.pos >= lex.length {
		return FILE_EOF
	}
	val := lex.code[lex.pos]
	lex.pos++
	return val
}
