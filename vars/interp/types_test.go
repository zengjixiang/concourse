package interp_test

import (
	"encoding/json"
	"reflect"

	"github.com/concourse/concourse/vars"
	"github.com/concourse/concourse/vars/interp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Interpolation", func() {
	Describe("UnmarshalJSON", func() {
		var str interp.String
		var ref interp.Var

		for _, tt := range []struct {
			desc   string
			body   string
			dst    interface{}
			result interface{}
			err    string
		}{
			{
				desc: "String",
				body: `"something"`,
				dst:  &str,

				result: interp.String("something"),
			},
			{
				desc: "String with single var",
				body: `"((hello))"`,
				dst:  &str,

				result: interp.String("((hello))"),
			},
			{
				desc: "String with vars interspersed",
				body: `"something-((hello))-else-((world))"`,
				dst:  &str,

				result: interp.String("something-((hello))-else-((world))"),
			},
			{
				desc: "simple Var",
				body: `"((hello))"`,
				dst:  &ref,

				result: interp.Var(vars.VariableReference{Name: "hello", Path: "hello"}),
			},
			{
				desc: "complex Var",
				body: `"((source:\"hello.world\".path1.path2))"`,
				dst:  &ref,

				result: interp.Var(vars.VariableReference{
					Name:   "source:\"hello.world\".path1.path2",
					Source: "source",
					Path:   "hello.world",
					Fields: []string{"path1", "path2"},
				}),
			},
			{
				desc: "invalid Var",
				body: `"no var here"`,
				dst:  &ref,

				err: "assigned value is not a var reference",
			},
			{
				desc: "non-anchored Var",
				body: `"foo((bar))baz"`,
				dst:  &ref,

				err: "assigned value is not a var reference",
			},
		} {
			tt := tt

			It(tt.desc, func() {
				err := json.Unmarshal([]byte(tt.body), tt.dst)
				if tt.err != "" {
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(MatchRegexp(tt.err))
				} else {
					Expect(err).ToNot(HaveOccurred())
					Expect(reflect.ValueOf(tt.dst).Elem().Interface()).To(Equal(tt.result))
				}
			})
		}
	})

	Describe("Interpolate", func() {
		Describe("String", func() {
			for _, tt := range []struct {
				desc   string
				str    interp.String
				vars   vars.Variables
				result string
				err    string
			}{
				{
					desc: "no var",
					str:  "hello",

					result: "hello",
				},
				{
					desc: "anchored var",
					str:  "((hello))",
					vars: vars.StaticVariables{"hello": "world"},

					result: "world",
				},
				{
					desc: "anchored var with non-string",
					str:  "((hello))",
					vars: vars.StaticVariables{"hello": 123},

					err: "cannot unmarshal number into.*string",
				},
				{
					desc: "interspersed vars",
					str:  "abc-((hello))-ghi-((world))-((blah))",
					vars: vars.StaticVariables{"hello": "def", "world": 123, "blah": true},

					result: "abc-def-ghi-123-true",
				},
				{
					desc: "non-string/non-number var",
					str:  interp.String("((hello))-abc"),
					vars: vars.StaticVariables{"hello": []string{"a", "b", "c"}},

					err: "cannot interpolate slice into a string",
				},
				{
					desc: "missing vars",
					str:  "something-((hello))-else-((world))",
					vars: vars.StaticVariables{},

					err: `"hello" was not found.*\n.*` +
						`"world" was not found`,
				},
			} {
				tt := tt

				It(tt.desc, func() {
					result, err := tt.str.Interpolate(interp.VarsResolver{Variables: tt.vars})
					if tt.err != "" {
						Expect(err).To(HaveOccurred())
						Expect(err.Error()).To(MatchRegexp(tt.err))
					} else {
						Expect(err).ToNot(HaveOccurred())
						Expect(result).To(Equal(tt.result))
					}
				})
			}
		})

		Describe("Var", func() {
			for _, tt := range []struct {
				desc   string
				ref    interp.Var
				vars   vars.Variables
				result interface{}
				err    string
			}{
				{
					desc: "string var",
					ref:  interp.Var{Name: "hello", Path: "hello"},
					vars: vars.StaticVariables{"hello": "world"},

					result: "world",
				},
				{
					desc: "list var",
					ref:  interp.Var{Name: "hello", Path: "hello"},
					vars: vars.StaticVariables{"hello": []string{"abc", "def"}},

					result: []interface{}{"abc", "def"},
				},
				{
					desc: "missing var",
					ref:  interp.Var{Name: "hello", Path: "hello"},
					vars: vars.StaticVariables{},

					err: `"hello" was not found`,
				},
			} {
				tt := tt

				It(tt.desc, func() {
					var dst interface{}
					err := tt.ref.InterpolateInto(interp.VarsResolver{Variables: tt.vars}, &dst)
					if tt.err != "" {
						Expect(err).To(HaveOccurred())
						Expect(err.Error()).To(MatchRegexp(tt.err))
					} else {
						Expect(err).ToNot(HaveOccurred())
						Expect(dst).To(Equal(tt.result))
					}
				})
			}
		})
	})
})
