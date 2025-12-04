package zapl_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// This is required to hook Ginkgo into the standard Go test runner.
func TestMyPackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MyPackage Suite")
}
