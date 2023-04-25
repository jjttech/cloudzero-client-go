//go:build integration || unit

package cloudzero

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCloudZero(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CloudZero Suite")
}
