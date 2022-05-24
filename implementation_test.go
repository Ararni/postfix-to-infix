package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestPostfixToInfixResults1(c *C) {
	res1, err1 := PostfixToInfix("2 - 3 * 5 +")
	c.Check(res1, Equals, "(((4-2)*3)+5)")
	c.Check(err1, IsNil)
}

func (s *MySuite) TestPostfixToInfixResults2(c *C) {
	res2, err2 := PostfixToInfix("8 3 + 4 * 6 -")
	c.Check(res2, Equals, "(((8+3)*4)-6)")
	c.Check(err2, IsNil)
}

func (s *MySuite) TestPostfixToInfixResults3(c *C) {
	res3, err3 := PostfixToInfix("7 3 -")
	c.Check(res3, Equals, "(7-3)")
	c.Check(err3, IsNil)
}

func (s *MySuite) TestPostfixToInfixResults4(c *C) {
	res4, err4 := PostfixToInfix("55 13 25 + -")
	c.Check(res4, Equals, "(55-(13+25))")
	c.Check(err4, IsNil)
}

func (s *MySuite) TestPostfixToInfixResults5(c *C) {
	res5, err5 := PostfixToInfix("1 74 + 10 * 36 +")
	c.Check(res5, Equals, "(((1+74)*10)+36)")
	c.Check(err5, IsNil)
}

func (s *MySuite) TestPostfixToInfixErrors1(c *C) {
	_, err1 := PostfixToInfix("$ 2 2")
	c.Check(err1, NotNil)
	c.Check(fmt.Sprintf("%s", err1), Equals, "Invalid postfix")
}

func (s *MySuite) TestPostfixToInfixErrors2(c *C) {
	_, err2 := PostfixToInfix("+ 2 2")
	c.Check(err2, NotNil)
	c.Check(fmt.Sprintf("%s", err2), Equals, "Invalid postfix")
}

func (s *MySuite) TestPostfixToInfixErrors3(c *C) {
	_, err3 := PostfixToInfix("")
	c.Check(err3, NotNil)
	c.Check(fmt.Sprintf("%s", err3), Equals, "Empty postfix")
}

func (s *MySuite) TestPostfixToInfixErrors4(c *C) {
	_, err4 := PostfixToInfix("+ 2 / 6 2 4")
	c.Check(err4, NotNil)
	c.Check(fmt.Sprintf("%s", err4), Equals, "Invalid postfix")
}
