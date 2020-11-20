package utils

type Args struct {
	A, B int
}

const MathServiceName = "MathService"

type MathServiceInterface = interface {
	Mutiply(args *Args, reply *int) error
	Divide(args *Args, reply *int) error
}
