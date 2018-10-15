package scrapper

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {

	It("should return no errors", func() {
		Expect(Handler()).To(Succeed())
	})

	// TODO add tests, stub external requests
})
