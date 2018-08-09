package controller

type Controller struct {
	message string
}

func NewController() *Controller {
	return &Controller{
		message: "Hello World",
	}
}
