package gdig

import "go.uber.org/dig"

var globalContainer = dig.New()

func Invoke(function interface{}, opts ...dig.InvokeOption) error {
	return globalContainer.Invoke(function, opts...)
}

func Provide(constructor interface{}, opts ...dig.ProvideOption) error {
	return globalContainer.Provide(constructor, opts...)
}

func String() string {
	return globalContainer.String()
}
