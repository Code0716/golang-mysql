package controllers

type Context interface {
	JSON(int, interface{})
}
