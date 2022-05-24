package lab2

import (
    "bytes"
    "strings"
		"fmt"

    . "gopkg.in/check.v1"
)
func (s *MySuite) TestCompute(c *C) { 
	handler := &ComputeHandler{}

	var exist string 
	var need string 
	var err error 
	var err2 error 

	var buff bytes.Buffer

	handler.Output = &buff
	handler.Input = strings.NewReader("4 2 - 3 *")

	err = handler.Compute()
	exist = buff.String()
	need, err2 = PostfixToInfix("4 2 - 3 *")
	c.Check(err2, IsNil)
	c.Check(err, IsNil)
	c.Check(exist, Equals, need)
 }

 func (s *MySuite) TestErrors(c *C){
	handler := &ComputeHandler{}

	var exist string 
	var err error
	
	var buff bytes.Buffer

	handler.Output = &buff
	handler.Input = strings.NewReader("- 5 6 * 8 ")

	err = handler.Compute()
	exist = buff.String()
	c.Check(err, NotNil)
	c.Check(fmt.Sprintf("%s", err), Equals, "Invalid postfix")
	c.Check(exist, Equals, "")
 }