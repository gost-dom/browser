package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type FormDataSuite struct {
	ScriptHostSuite
}

func NewFormDataSuite(h html.ScriptEngine) *FormDataSuite {
	return &FormDataSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *FormDataSuite) TestInheritance() {
	s.Expect(
		s.Eval("Object.getPrototypeOf(FormData.prototype) === Object.prototype"),
	).To(BeTrue())
}

func (s *FormDataSuite) TestAddAndGet() {
	s.Expect(s.Eval(`
		data = new FormData();
		data.append("key", "value");
		data.get("key");
	`)).To(Equal("value"))
}

func (s *FormDataSuite) TestGetKeys() {
	s.Expect(s.Eval(`
		data = new FormData();
		data.append("key1", "value");
		data.append("key2", "value");
		Array.from(data.keys()).join(",")
	`)).To(Equal("key1,key2"))
}

func (s *FormDataSuite) TestGetEntries() {
	s.Expect(s.Eval(`
		data = new FormData();
		data.append("key1", "value1");
		data.append("key2", "value2");
		Array.from(data.entries()).map(x => x.join(";")).join(",")
	`)).To(Equal("key1;value1,key2;value2"))
}

func (s *FormDataSuite) TestForEach() {
	s.Expect(s.Eval(`
		const result = [];
		data = new FormData();
		data.append("key1", "value1");
		data.append("key2", "value2");
		data.forEach((v,k) => { result.push(k + ": " + v) })
		result.join(", ")
	`)).To(Equal("key1: value1, key2: value2"))
}

func (s *FormDataSuite) TestIterable() {
	s.Expect(s.Eval(`
		data = new FormData();
		typeof data[Symbol.iterator]`)).To(Equal("function"))

	s.Expect(s.Eval(`
		data = new FormData();
		data.append("key1", "value1");
		data.append("key2", "value2");
		Array.from(data).map(x => x.join(";")).join(",")
		`)).To(Equal("key1;value1,key2;value2"))
}
