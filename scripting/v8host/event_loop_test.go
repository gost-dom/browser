package v8host_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("EventLoop", func() {
	It("Defers execution", func() {
		ctx := NewTestContext(IgnoreUnhandledErrors)
		Expect(
			ctx.Eval(`
				let val; 
				setTimeout(() => { val = 42 }, 1);
				val`,
			),
		).To(BeNil())
		ctx.Clock().Advance(time.Millisecond)
		Expect(ctx.Eval(`val`)).To(BeEquivalentTo(42))
	})

	It("Dispatches an 'error' event on unhandled error", func() {
		ctx := NewTestContext(IgnoreUnhandledErrors)
		Expect(
			ctx.Eval(`
				let val;
				window.addEventListener('error', () => {
					val = 42;
				});
				setTimeout(() => {
					throw new Error()
				}, 0); 
				val`,
			),
		).To(BeNil())
		Expect(ctx.Eval(`val`)).To(BeEquivalentTo(42))
	})
})
