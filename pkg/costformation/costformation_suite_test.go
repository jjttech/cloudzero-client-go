//go:build integration || unit

package costformation

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCostformation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CostFormation Suite")
}
