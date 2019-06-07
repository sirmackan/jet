package sqlbuilder

type BoolExpression interface {
	Expression

	EQ(expression BoolExpression) BoolExpression
	NOT_EQ(expression BoolExpression) BoolExpression
	IS_DISTINCT_FROM(rhs BoolExpression) BoolExpression
	IS_NOT_DISTINCT_FROM(rhs BoolExpression) BoolExpression

	IS_TRUE() BoolExpression
	IS_NOT_TRUE() BoolExpression
	IS_FALSE() BoolExpression
	IS_NOT_FALSE() BoolExpression
	IS_UNKNOWN() BoolExpression
	IS_NOT_UNKNOWN() BoolExpression

	AND(expression BoolExpression) BoolExpression
	OR(expression BoolExpression) BoolExpression
}

type boolInterfaceImpl struct {
	parent BoolExpression
}

func (b *boolInterfaceImpl) EQ(expression BoolExpression) BoolExpression {
	return EQ(b.parent, expression)
}

func (b *boolInterfaceImpl) NOT_EQ(expression BoolExpression) BoolExpression {
	return NOT_EQ(b.parent, expression)
}

func (b *boolInterfaceImpl) IS_DISTINCT_FROM(rhs BoolExpression) BoolExpression {
	return IS_DISTINCT_FROM(b.parent, rhs)
}

func (b *boolInterfaceImpl) IS_NOT_DISTINCT_FROM(rhs BoolExpression) BoolExpression {
	return IS_NOT_DISTINCT_FROM(b.parent, rhs)
}

func (b *boolInterfaceImpl) AND(expression BoolExpression) BoolExpression {
	return newBinaryBoolExpression(b.parent, expression, "AND")
}

func (b *boolInterfaceImpl) OR(expression BoolExpression) BoolExpression {
	return newBinaryBoolExpression(b.parent, expression, "OR")
}

func (b *boolInterfaceImpl) IS_TRUE() BoolExpression {
	return newPostifxBoolExpression(b.parent, "IS TRUE")
}

func (b *boolInterfaceImpl) IS_NOT_TRUE() BoolExpression {
	return newPostifxBoolExpression(b.parent, "IS NOT TRUE")
}

func (b *boolInterfaceImpl) IS_FALSE() BoolExpression {
	return newPostifxBoolExpression(b.parent, "IS FALSE")
}

func (b *boolInterfaceImpl) IS_NOT_FALSE() BoolExpression {
	return newPostifxBoolExpression(b.parent, "IS NOT FALSE")
}

func (b *boolInterfaceImpl) IS_UNKNOWN() BoolExpression {
	return newPostifxBoolExpression(b.parent, "IS UNKNOWN")
}

func (b *boolInterfaceImpl) IS_NOT_UNKNOWN() BoolExpression {
	return newPostifxBoolExpression(b.parent, "IS NOT UNKNOWN")
}

//---------------------------------------------------//
type binaryBoolExpression struct {
	expressionInterfaceImpl
	boolInterfaceImpl

	binaryOpExpression
}

func newBinaryBoolExpression(lhs, rhs Expression, operator string) BoolExpression {
	boolExpression := binaryBoolExpression{}

	boolExpression.binaryOpExpression = newBinaryExpression(lhs, rhs, operator)
	boolExpression.expressionInterfaceImpl.parent = &boolExpression
	boolExpression.boolInterfaceImpl.parent = &boolExpression

	return &boolExpression
}

//---------------------------------------------------//
type prefixBoolExpression struct {
	expressionInterfaceImpl
	boolInterfaceImpl

	prefixOpExpression
}

func newPrefixBoolExpression(expression Expression, operator string) BoolExpression {
	exp := prefixBoolExpression{}
	exp.prefixOpExpression = newPrefixExpression(expression, operator)

	exp.expressionInterfaceImpl.parent = &exp
	exp.boolInterfaceImpl.parent = &exp

	return &exp
}

//---------------------------------------------------//
type postfixBoolOpExpression struct {
	expressionInterfaceImpl
	boolInterfaceImpl

	postfixOpExpression
}

func newPostifxBoolExpression(expression Expression, operator string) BoolExpression {
	exp := postfixBoolOpExpression{}
	exp.postfixOpExpression = newPostfixOpExpression(expression, operator)

	exp.expressionInterfaceImpl.parent = &exp
	exp.boolInterfaceImpl.parent = &exp

	return &exp
}

//---------------------------------------------------//

type boolExpressionWrapper struct {
	boolInterfaceImpl
	Expression
}

func newBoolExpressionWrap(expression Expression) BoolExpression {
	boolExpressionWrap := boolExpressionWrapper{Expression: expression}
	boolExpressionWrap.boolInterfaceImpl.parent = &boolExpressionWrap
	return &boolExpressionWrap
}

func BoolExp(expression Expression) BoolExpression {
	return newBoolExpressionWrap(expression)
}
