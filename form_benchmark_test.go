package formbenchmark

import (
	"github.com/cjtoolkit/form"
	"github.com/cjtoolkit/form/fields"
	"github.com/gorilla/schema"
	"github.com/monoculum/formam"
	"net/url"
	"testing"
)

type ForSchemaAndFormamSimple struct {
	Text string
}

type ForSchemaAndFormamComplex struct {
	Number int64
	Text   string
	Float  float64
}

type ForFormSimple struct {
	fields []form.FormFieldInterface

	TextModel string
	TextNorm  string
	TextErr   error
}

func (f *ForFormSimple) Fields() []form.FormFieldInterface {
	if nil == f.fields {
		f.fields = []form.FormFieldInterface{
			f.TextField(),
		}
	}

	return f.fields
}

func (f *ForFormSimple) TextField() fields.String {
	return fields.String{
		Name:  "Text",
		Label: "Text",
		Model: &f.TextModel,
		Norm:  &f.TextNorm,
		Err:   &f.TextErr,
	}
}

type ForFormComplex struct {
	fields []form.FormFieldInterface

	NumberModel int64
	NumberNorm  string
	NumberErr   error

	TextModel string
	TextNorm  string
	TextErr   error

	FloatModel float64
	FloatNorm  string
	FloatErr   error
}

func (f *ForFormComplex) Fields() []form.FormFieldInterface {
	if nil == f.fields {
		f.fields = []form.FormFieldInterface{
			f.NumberField(),
			f.TextField(),
			f.FloatField(),
		}
	}

	return f.fields
}

func (f *ForFormComplex) NumberField() fields.Int {
	return fields.Int{
		Name:  "Number",
		Label: "Number",
		Model: &f.NumberModel,
		Norm:  &f.NumberNorm,
		Err:   &f.NumberErr,
	}
}

func (f *ForFormComplex) TextField() fields.String {
	return fields.String{
		Name:    "Text",
		Label:   "Text",
		Model:   &f.TextModel,
		Norm:    &f.TextNorm,
		Err:     &f.TextErr,
	}
}

func (f *ForFormComplex) FloatField() fields.Float {
	return fields.Float{
		Name:  "Float",
		Label: "Float",
		Model: &f.FloatModel,
		Norm:  &f.FloatNorm,
		Err:   &f.FloatErr,
	}
}

var (
	schemaDecode  = schema.NewDecoder()
	formValidator = form.NewFormEnglishLanguage().DisablePreCheck()

	valuesSimple = url.Values{
		"Text": {"Hello World!"},
	}

	valuesComplex = url.Values{
		"Number": {"42"},
		"Text":   {"Hello World!"},
		"Float":  {"20.1"},
	}

	valuesSimpleWithUnusedData = url.Values{
		"Text":    {"Hello World!"},
		"Random1": {"Random1"},
		"Random2": {"Random2"},
		"Random3": {"Random3"},
		"Random4": {"Random4"},
		"Random5": {"Random5"},
		"Random6": {"Random6"},
	}

	valuesComplexWithUnusedData = url.Values{
		"Number":  {"42"},
		"Text":    {"Hello World!"},
		"Float":   {"20.1"},
		"Random1": {"Random1"},
		"Random2": {"Random2"},
		"Random3": {"Random3"},
		"Random4": {"Random4"},
		"Random5": {"Random5"},
		"Random6": {"Random6"},
	}
)

func BenchmarkSchemaSimple(b *testing.B) {
	data := &ForSchemaAndFormamSimple{}
	for i := 0; i < b.N; i++ {
		schemaDecode.Decode(data, valuesSimple)
	}
}

func BenchmarkFormamSimple(b *testing.B) {
	data := &ForSchemaAndFormamSimple{}
	for i := 0; i < b.N; i++ {
		formam.Decode(valuesSimple, data)
	}
}

func BenchmarkCjToolkitFormSimple(b *testing.B) {
	data := &ForFormSimple{}
	for i := 0; i < b.N; i++ {
		formValidator.SetForm(valuesSimple)
		formValidator.Validate(data)
	}
}

func BenchmarkSchemaComplex(b *testing.B) {
	data := &ForSchemaAndFormamComplex{}
	for i := 0; i < b.N; i++ {
		schemaDecode.Decode(data, valuesComplex)
	}
}

func BenchmarkFormamComplex(b *testing.B) {
	data := &ForSchemaAndFormamComplex{}
	for i := 0; i < b.N; i++ {
		formam.Decode(valuesComplex, data)
	}
}

func BenchmarkCjToolkitFormComplex(b *testing.B) {
	data := &ForFormComplex{}
	for i := 0; i < b.N; i++ {
		formValidator.SetForm(valuesComplex)
		formValidator.Validate(data)
	}
}

func BenchmarkSchemaSimpleWithUnusedData(b *testing.B) {
	data := &ForSchemaAndFormamSimple{}
	for i := 0; i < b.N; i++ {
		schemaDecode.Decode(data, valuesSimpleWithUnusedData)
	}
}

func BenchmarkFormamSimpleWithUnusedData(b *testing.B) {
	data := &ForSchemaAndFormamSimple{}
	for i := 0; i < b.N; i++ {
		formam.Decode(valuesSimpleWithUnusedData, data)
	}
}

func BenchmarkCjToolkitFormSimpleWithUnusedData(b *testing.B) {
	data := &ForFormSimple{}
	for i := 0; i < b.N; i++ {
		formValidator.SetForm(valuesSimpleWithUnusedData)
		formValidator.Validate(data)
	}
}

func BenchmarkSchemaComplexWithUnusedData(b *testing.B) {
	data := &ForSchemaAndFormamComplex{}
	for i := 0; i < b.N; i++ {
		schemaDecode.Decode(data, valuesComplexWithUnusedData)
	}
}

func BenchmarkFormamComplexWithUnusedData(b *testing.B) {
	data := &ForSchemaAndFormamComplex{}
	for i := 0; i < b.N; i++ {
		formam.Decode(valuesComplexWithUnusedData, data)
	}
}

func BenchmarkCjToolkitFormComplexWithUnusedData(b *testing.B) {
	data := &ForFormComplex{}
	for i := 0; i < b.N; i++ {
		formValidator.SetForm(valuesComplexWithUnusedData)
		formValidator.Validate(data)
	}
}
