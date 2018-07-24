package main_test

import (
	"fmt"
	"testing"
	"time"

	"bitbucket.apps.seibert-media.net/smedia/isac-test/command"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

func TestCopy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "isac-cli-prototype main suite")
}

var err error
var session *gexec.Session

var _ = BeforeSuite(func() {
	Expect(err).To(BeNil())
})

var _ = AfterEach(func() {
	session.Interrupt()
	Eventually(session).Should(gexec.Exit())
})

var cmd command.Command

var _ = Describe("ISAC CLI when called", func() {

	BeforeEach(func() {
		cmd = command.Command{
			Name: "isac-cli-prototype",
		}
	})

	Describe("with no arguments", func() {
		It("exits with non-zero status code and usage message", func() {
			By("running command", func() {
				session, err = gexec.Start(cmd.Create(), GinkgoWriter, GinkgoWriter)
				if err != nil {
					fmt.Println(err.Error())
				}
				Expect(err).To(BeNil())
			})
			By("waiting for cmd to complete", func() {
				session.Wait(10 * time.Second)
			})
			By("checking for exit code", func() {
				Expect(session.ExitCode()).To(Equal(0))
			})
			By("expecting a message containing the program name", func() {
				Expect(session.Out).To(gbytes.Say("^isac-cli"))
			})
			By("expecting a usage message", func() {
				Expect(session.Out).To(gbytes.Say("\nUsage:"))
			})
		})
	})

})
