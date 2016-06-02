package application_test

import (
	"errors"

	"github.com/cloudfoundry/cli/cf/api/applications/applicationsfakes"
	"github.com/cloudfoundry/cli/cf/commandregistry"
	"github.com/cloudfoundry/cli/cf/configuration/coreconfig"
	"github.com/cloudfoundry/cli/cf/models"
	testcmd "github.com/cloudfoundry/cli/testhelpers/commands"
	testconfig "github.com/cloudfoundry/cli/testhelpers/configuration"
	testreq "github.com/cloudfoundry/cli/testhelpers/requirements"
	testterm "github.com/cloudfoundry/cli/testhelpers/terminal"

	. "github.com/cloudfoundry/cli/testhelpers/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("enable-ssh command", func() {
	var (
		ui                  *testterm.FakeUI
		requirementsFactory *testreq.FakeReqFactory
		appRepo             *applicationsfakes.FakeRepository
		configRepo          coreconfig.Repository
		deps                commandregistry.Dependency
	)

	BeforeEach(func() {
		ui = &testterm.FakeUI{}
		configRepo = testconfig.NewRepositoryWithDefaults()
		requirementsFactory = &testreq.FakeReqFactory{}
		appRepo = new(applicationsfakes.FakeRepository)
	})

	updateCommandDependency := func(pluginCall bool) {
		deps.UI = ui
		deps.Config = configRepo
		deps.RepoLocator = deps.RepoLocator.SetApplicationRepository(appRepo)
		commandregistry.Commands.SetCommand(commandregistry.Commands.FindCommand("enable-ssh").SetDependency(deps, pluginCall))
	}

	runCommand := func(args ...string) bool {
		return testcmd.RunCLICommand("enable-ssh", args, requirementsFactory, updateCommandDependency, false, ui)
	}

	Describe("requirements", func() {
		It("fails with usage when called without enough arguments", func() {
			requirementsFactory.LoginSuccess = true

			runCommand()
			Expect(ui.Outputs).To(ContainSubstrings(
				[]string{"Incorrect Usage", "Requires", "argument"},
			))

		})

		It("fails requirements when not logged in", func() {
			Expect(runCommand("my-app", "none")).To(BeFalse())
		})

		It("fails if a space is not targeted", func() {
			requirementsFactory.LoginSuccess = true
			requirementsFactory.TargetedSpaceSuccess = false
			Expect(runCommand("my-app", "none")).To(BeFalse())
		})
	})

	Describe("enable-ssh", func() {
		var (
			app models.Application
		)

		BeforeEach(func() {
			requirementsFactory.LoginSuccess = true
			requirementsFactory.TargetedSpaceSuccess = true

			app = models.Application{}
			app.Name = "my-app"
			app.GUID = "my-app-guid"
			app.EnableSSH = false

			requirementsFactory.Application = app
		})

		Context("when enable_ssh is already set to the true", func() {
			BeforeEach(func() {
				app.EnableSSH = true
				requirementsFactory.Application = app
			})

			It("notifies the user", func() {
				runCommand("my-app")

				Expect(ui.Outputs).To(ContainSubstrings([]string{"ssh support is already enabled for 'my-app'"}))
			})
		})

		Context("Updating enable_ssh when not already set to true", func() {
			Context("Update successfully", func() {
				BeforeEach(func() {
					app = models.Application{}
					app.Name = "my-app"
					app.GUID = "my-app-guid"
					app.EnableSSH = true

					appRepo.UpdateReturns(app, nil)
				})

				It("updates the app's enable_ssh", func() {
					runCommand("my-app")

					Expect(appRepo.UpdateCallCount()).To(Equal(1))
					appGUID, params := appRepo.UpdateArgsForCall(0)
					Expect(appGUID).To(Equal("my-app-guid"))
					Expect(*params.EnableSSH).To(Equal(true))
					Expect(ui.Outputs).To(ContainSubstrings([]string{"Enabling ssh support for 'my-app'"}))
					Expect(ui.Outputs).To(ContainSubstrings([]string{"OK"}))
				})
			})

			Context("Update fails", func() {
				It("notifies user of any api error", func() {
					appRepo.UpdateReturns(models.Application{}, errors.New("Error updating app."))
					runCommand("my-app")

					Expect(appRepo.UpdateCallCount()).To(Equal(1))
					Expect(ui.Outputs).To(ContainSubstrings(
						[]string{"FAILED"},
						[]string{"Error enabling ssh support"},
					))

				})

				It("notifies user when updated result is not in the desired state", func() {
					app = models.Application{}
					app.Name = "my-app"
					app.GUID = "my-app-guid"
					app.EnableSSH = false
					appRepo.UpdateReturns(app, nil)

					runCommand("my-app")

					Expect(appRepo.UpdateCallCount()).To(Equal(1))
					Expect(ui.Outputs).To(ContainSubstrings(
						[]string{"FAILED"},
						[]string{"ssh support is not enabled for my-app"},
					))

				})
			})

		})
	})

})
