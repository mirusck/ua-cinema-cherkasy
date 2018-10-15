package scrapper_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestScrapper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scrapper Suite")
}
