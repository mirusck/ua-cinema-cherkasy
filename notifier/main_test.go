package notifier

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	It("should return no errors", func() {
		Expect(Handler()).To(Succeed())
	})

	// TODO tests, actually
})
