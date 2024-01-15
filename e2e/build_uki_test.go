package e2e_test

import (
	"os"
	"path"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("build-uki", func() {
	var resultDir string
	var resultFile string
	var image string
	var err error
	var enki *Enki

	BeforeEach(func() {
		enki = NewEnki("busybox")
		image = "busybox"
		resultDir, err = os.MkdirTemp("", "enki-build-uki-test-")
		Expect(err).ToNot(HaveOccurred())
		resultFile = path.Join(resultDir, "result.uki")
	})

	AfterEach(func() {
		os.RemoveAll(resultDir)
		enki.Cleanup()
	})

	When("some dependency is missing", func() {
		BeforeEach(func() {
			enki = NewEnki("busybox")
		})

		It("returns an error about missing deps", func() {
			out, err := enki.Run("build-uki", image, resultFile)
			Expect(err).To(HaveOccurred(), out)
			Expect(out).To(MatchRegexp("executable file not found in \\$PATH"))
		})
	})

	It("successfully builds an UKI from a Docker image", func() {
		out, err := enki.Run("build-uki", image, resultFile)
		Expect(err).ToNot(HaveOccurred(), out)

		_, err = os.Stat(resultFile)
		Expect(err).ToNot(HaveOccurred())
	})
})
