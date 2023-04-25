//go:build integration || unit

package anycost

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAnyCost(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AnyCost Suite")
}
