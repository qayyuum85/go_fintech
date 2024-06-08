package helpers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"qayyuum/go_fintech/helpers"
	"qayyuum/go_fintech/interfaces"
)

var _ = Describe("Helpers", func() {
	Context("test successful validation", func() {
		It("should return true for valid user login registration info", func() {
			values := []interfaces.Validation{
				{
					Valid: "username",
					Value: "dqwuhdqwhqwub",
				},
				{
					Valid: "email",
					Value: "someone@someemail.com",
				},
				{
					Valid: "password",
					Value: "dasdasd",
				},
			}
			Expect(helpers.Validation(values)).To(BeTrue())

		})
	})

	Context("test failed validation", func() {
		It("shauld return false for username that contains symbol", func() {
			values := []interfaces.Validation{
				{
					Valid: "username",
					Value: "someuser_hey_!@#",
				},
			}
			Expect(helpers.Validation(values)).To(BeFalse())
		})
		It("shauld return false for invalid email", func() {
			values := []interfaces.Validation{
				{
					Valid: "email",
					Value: "hdasdxxx.com",
				},
			}
			Expect(helpers.Validation(values)).To(BeFalse())

		})
		It("shauld return false for short password", func() {
			values := []interfaces.Validation{
				{
					Valid: "password",
					Value: "!@#",
				},
			}
			Expect(helpers.Validation(values)).To(BeFalse())

		})
	})
})
