package main

import (
	"fmt"
	"os"
	"os/user"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("getFrom", func() {
	Context("when provided an email address", func() {
		It("should return that email address", func() {
			Expect(getFrom(options{from: "john@example.com"})).To(Equal("john@example.com"))
		})
	})

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	Context("when provided just a name email", func() {
		It("should return that email address, plus the hostname", func() {
			Expect(getFrom(options{from: "john"})).To(Equal(fmt.Sprintf("john@%s", hostname)))
		})
	})

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	Context("when provided nothing", func() {
		It("should return the username, plus the host name", func() {
			Expect(getFrom(options{})).To(Equal(fmt.Sprintf("%s@%s", user.Username, hostname)))
		})
	})
})
