package main

import(
.	"github.com/onsi/ginkgo"
.	"github.com/onsi/gomega"
	"testing"
)


func TestCloudSSH(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CloudSSH Suite")
}


