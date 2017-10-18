package integration_test

import (
	"net/http"
	"os"
	"os/exec"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"

	"testing"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var agoutiDriver *agouti.WebDriver

var _ = BeforeSuite(func() {
	buildCmd := exec.Command("go", "build", "..")
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	Expect(buildCmd.Run()).To(Succeed())

	agoutiDriver = agouti.ChromeDriver()
	Expect(agoutiDriver.Start()).To(Succeed())
})

var _ = AfterSuite(func() {
	Expect(agoutiDriver.Stop()).To(Succeed())
})

var snitchCmd *exec.Cmd

var _ = BeforeEach(func(done Done) {
	startSnitch()
	close(done)
}, 4.0)

var _ = AfterEach(func(done Done) {
	stopSnitch()
	close(done)
})

var page *agouti.Page

var _ = BeforeEach(func() {
	var err error
	page, err = agoutiDriver.NewPage()
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterEach(func() {
	Expect(page.Destroy()).To(Succeed())
})

func startSnitch() {

	_, err := os.Stat("db")
	if os.IsNotExist(err) {
		Expect(os.Mkdir("db", 0777)).To(Succeed())
	}

	snitchCmd = exec.Command("./snitch", "-p", "23453")
	snitchCmd.Stdout = os.Stdout
	snitchCmd.Stderr = os.Stderr
	Expect(snitchCmd.Start()).To(Succeed())
	for {
		_, err := http.Get("http://localhost:23453")
		if err != nil {
			time.Sleep(15 * time.Millisecond)
			continue
		}
		break
	}
}

func stopSnitch() {
	Expect(snitchCmd.Process.Kill()).To(Succeed())
	snitchCmd.Process.Wait()
}

var _ = Describe("Index page", func() {

	Context("When I navigate to the index page", func() {
		BeforeEach(func() {
			Expect(page.Navigate("http://localhost:23453")).To(Succeed())
		})
		It("should have 'Add Target' button", func() {
			Eventually(page.FindByLink("Add Target")).Should(BeFound())
		})
		Context("When I click on 'Add Target' button", func() {
			BeforeEach(func() {
				Eventually(page.FindByLink("Add Target")).Should(BeFound())
				Expect(page.FindByLink("Add Target").Click()).To(Succeed())
			})
			It("Should have 'Name' input", func() {
				Eventually(page.Find("input[placeholder=Name]")).Should(BeFound())
			})
			It("Should have 'Image' input", func() {
				Eventually(page.Find("input[placeholder=Image]")).Should(BeFound())
			})
			It("Should have 'Command' input", func() {
				Eventually(page.Find("input[placeholder=Command]")).Should(BeFound())
			})
			It("Should have 'Interval' input", func() {
				Eventually(page.Find("input[placeholder=Interval]")).Should(BeFound())
			})
			It("Should have 'Add Target' button", func() {
				Eventually(page.FindByButton("Add Target")).Should(BeFound())
			})

			Context("When I create a new target", func() {
				BeforeEach(func() {
					Eventually(page.Find("input[placeholder=Name]")).Should(BeFound())
					Expect(page.Find("input[placeholder=Name]").Fill("target1")).To(Succeed())
					Expect(page.Find("input[placeholder=Image]").Fill("alpine:3.6")).To(Succeed())
					Expect(page.Find("input[placeholder=Command]").Fill("true")).To(Succeed())
					Expect(page.Find("input[placeholder=Interval]").Fill("20")).To(Succeed())
					Eventually(page.FindByButton("Add Target")).Should(BeEnabled())
					Expect(page.FindByButton("Add Target").Click()).To(Succeed())
				})
				It("Should show the target details page", func() {
					Eventually(page.First("div.panel-heading")).Should(MatchText("^Target target1"))
				})
			})

		})
	})
})
