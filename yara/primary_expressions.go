package yara

import "fmt"

type Filesize struct{}

func (px Filesize) RuleString() (out string, err error) {
	return "filesize", nil
}

func (px Filesize) expression() Expression {
	return px
}

func (px Filesize) primaryExpression() PrimaryExpression {
	return px
}

type Entrypoint struct{}

func (px Entrypoint) RuleString() (out string, err error) {
	return "entrypoint", nil
}

func (px Entrypoint) expression() Expression {
	return px
}

func (px Entrypoint) primaryExpression() PrimaryExpression {
	return px
}

// Number -- base info is lost during parsing, so it will only be represented as
// base-10 when serializing back into YARA Rule text
type Number int64

func (px Number) RuleString() (out string, err error) {
	return fmt.Sprintf("%d", px), nil
}

func (px Number) expression() Expression {
	return px
}

func (px Number) primaryExpression() PrimaryExpression {
	return px
}

// Double -- same situation as Number but for float64
type Double float64

// RuleString outputs a float64 as a string
func (px Double) RuleString() (out string, err error) {
	return fmt.Sprintf("%f", px), nil
}

func (px Double) expression() Expression {
	return px
}

func (px Double) primaryExpression() PrimaryExpression {
	return px
}

type TextString string

func (px TextString) RuleString() (out string, err error) {
	return string(px), nil
}

func (px TextString) expression() Expression {
	return px
}

func (px TextString) primaryExpression() PrimaryExpression {
	return px
}
