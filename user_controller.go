package main

import "GO_WEB/framework"

func UserLoginController(c *framework.Context) error {
	foo, _ := c.QueryString("foo", "def")
	c.SetOkStatus().Json("ok, UserLoginController: " + foo)
	return nil
}
