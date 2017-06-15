package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("extractRecipients", func() {
	recipients := []string{"john@example.com", "darcy@example.com", "lizzy@example.com"}

	Context("when given the 'To:' header", func() {
		It("should extract all recipients", func() {
			email := "To: john@example.com, \"Darcy\" <darcy@example.com>, <lizzy@example.com>\r\n\r\n"
			Expect(extractRecipients(email)).To(ConsistOf(recipients))
		})
	})

	Context("when given the 'Cc:' header", func() {
		It("should extract all recipients", func() {
			email := "Cc: john@example.com, \"Darcy\" <darcy@example.com>, <lizzy@example.com>\r\n\r\n"
			Expect(extractRecipients(email)).To(ConsistOf(recipients))
		})
	})

	Context("when given the 'Bcc:' header", func() {
		It("should extract all recipients", func() {
			email := "Bcc: john@example.com, \"Darcy\" <darcy@example.com>, <lizzy@example.com>\r\n\r\n"
			Expect(extractRecipients(email)).To(ConsistOf(recipients))
		})
	})

	Context("when given multiple headers", func() {
		It("should extract all recipients", func() {
			email := "To: john@example.com\r\nCc: \"Darcy\" <darcy@example.com>\r\nBcc: <lizzy@example.com>\r\n\r\n"
			Expect(extractRecipients(email)).To(ConsistOf(recipients))
		})
	})

	Context("when given repeat addresses", func() {
		It("should deduplicate all recipients", func() {
			email := "To: john@example.com, lizzy@example.com\r\nCc: \"Darcy\" <darcy@example.com>, john@example.com\r\nBcc: <lizzy@example.com>, \"John\" <john@example.com>\r\n\r\n"
			Expect(extractRecipients(email)).To(ConsistOf(recipients))
		})
	})
})
