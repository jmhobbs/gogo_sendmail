package main

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("readEmail", func() {
	It("should split the email at the first blank line", func() {
		r := strings.NewReader("Subject: Test\r\nTo: john@example.com\r\n\r\nHello!")

		headers, body := readEmail(r)

		Expect(headers).To(Equal([]string{"Subject: Test", "To: john@example.com"}))
		Expect(body).To(Equal([]string{"Hello!"}))
	})
})
