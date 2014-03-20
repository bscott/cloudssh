package main

import (
	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"testing"
)

func TestCloudSSH(t *testing.T) {
	g := Goblin(t)

	//special hook for gomega
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })
	g.Describe("Numbers Test", func() {
		g.It("Should add two numbers ", func() {
			g.Assert(1 + 1).Equal(2)
		})
	})
}
