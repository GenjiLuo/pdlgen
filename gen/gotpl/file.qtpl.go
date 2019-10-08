// Code generated by qtc from "file.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line gen/gotpl/file.qtpl:1
package gotpl

//line gen/gotpl/file.qtpl:1
import (
	"sort"

	"github.com/chromedp/cdproto-gen/gen/genutil"
	"github.com/chromedp/cdproto-gen/pdl"
)

// FileHeader is the file header template.

//line gen/gotpl/file.qtpl:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line gen/gotpl/file.qtpl:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line gen/gotpl/file.qtpl:9
func StreamFileHeader(qw422016 *qt422016.Writer, pkgName string, d *pdl.Domain) {
//line gen/gotpl/file.qtpl:9
	qw422016.N().S(`
`)
//line gen/gotpl/file.qtpl:10
	if d != nil {
//line gen/gotpl/file.qtpl:10
		qw422016.N().S(`// Package `)
//line gen/gotpl/file.qtpl:10
		qw422016.N().S(genutil.PackageName(d))
//line gen/gotpl/file.qtpl:10
		qw422016.N().S(` provides the Chrome DevTools Protocol
// commands, types, and events for the `)
//line gen/gotpl/file.qtpl:11
		qw422016.N().S(d.Domain.String())
//line gen/gotpl/file.qtpl:11
		qw422016.N().S(` domain.
// `)
//line gen/gotpl/file.qtpl:12
		if desc := d.Description; desc != "" {
//line gen/gotpl/file.qtpl:12
			qw422016.N().S(`
`)
//line gen/gotpl/file.qtpl:13
			qw422016.N().S(genutil.FormatComment(desc, "", ""))
//line gen/gotpl/file.qtpl:13
			qw422016.N().S(`
//`)
//line gen/gotpl/file.qtpl:14
		}
//line gen/gotpl/file.qtpl:14
		qw422016.N().S(`
// Generated by the cdproto-gen command.`)
//line gen/gotpl/file.qtpl:15
	}
//line gen/gotpl/file.qtpl:15
	qw422016.N().S(`
package `)
//line gen/gotpl/file.qtpl:16
	qw422016.N().S(pkgName)
//line gen/gotpl/file.qtpl:16
	qw422016.N().S(`

`)
//line gen/gotpl/file.qtpl:18
	qw422016.N().S("// Code generated by cdproto-gen. DO NOT EDIT.")
//line gen/gotpl/file.qtpl:18
	qw422016.N().S(`
`)
//line gen/gotpl/file.qtpl:19
}

//line gen/gotpl/file.qtpl:19
func WriteFileHeader(qq422016 qtio422016.Writer, pkgName string, d *pdl.Domain) {
//line gen/gotpl/file.qtpl:19
	qw422016 := qt422016.AcquireWriter(qq422016)
//line gen/gotpl/file.qtpl:19
	StreamFileHeader(qw422016, pkgName, d)
//line gen/gotpl/file.qtpl:19
	qt422016.ReleaseWriter(qw422016)
//line gen/gotpl/file.qtpl:19
}

//line gen/gotpl/file.qtpl:19
func FileHeader(pkgName string, d *pdl.Domain) string {
//line gen/gotpl/file.qtpl:19
	qb422016 := qt422016.AcquireByteBuffer()
//line gen/gotpl/file.qtpl:19
	WriteFileHeader(qb422016, pkgName, d)
//line gen/gotpl/file.qtpl:19
	qs422016 := string(qb422016.B)
//line gen/gotpl/file.qtpl:19
	qt422016.ReleaseByteBuffer(qb422016)
//line gen/gotpl/file.qtpl:19
	return qs422016
//line gen/gotpl/file.qtpl:19
}

// FileImportTemplate is a general import template.

//line gen/gotpl/file.qtpl:22
func StreamFileImportTemplate(qw422016 *qt422016.Writer, importMap map[string]string) {
//line gen/gotpl/file.qtpl:23
	var keys []string
	for k := range importMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

//line gen/gotpl/file.qtpl:28
	qw422016.N().S(`
import (`)
//line gen/gotpl/file.qtpl:29
	for _, k := range keys {
//line gen/gotpl/file.qtpl:30
		v := importMap[k]

//line gen/gotpl/file.qtpl:31
		qw422016.N().S(`
	`)
//line gen/gotpl/file.qtpl:32
		if k != v {
//line gen/gotpl/file.qtpl:32
			qw422016.N().S(v)
//line gen/gotpl/file.qtpl:32
			qw422016.N().S(` `)
//line gen/gotpl/file.qtpl:32
		}
//line gen/gotpl/file.qtpl:32
		qw422016.N().Q(k)
//line gen/gotpl/file.qtpl:32
	}
//line gen/gotpl/file.qtpl:32
	qw422016.N().S(`
)
`)
//line gen/gotpl/file.qtpl:34
}

//line gen/gotpl/file.qtpl:34
func WriteFileImportTemplate(qq422016 qtio422016.Writer, importMap map[string]string) {
//line gen/gotpl/file.qtpl:34
	qw422016 := qt422016.AcquireWriter(qq422016)
//line gen/gotpl/file.qtpl:34
	StreamFileImportTemplate(qw422016, importMap)
//line gen/gotpl/file.qtpl:34
	qt422016.ReleaseWriter(qw422016)
//line gen/gotpl/file.qtpl:34
}

//line gen/gotpl/file.qtpl:34
func FileImportTemplate(importMap map[string]string) string {
//line gen/gotpl/file.qtpl:34
	qb422016 := qt422016.AcquireByteBuffer()
//line gen/gotpl/file.qtpl:34
	WriteFileImportTemplate(qb422016, importMap)
//line gen/gotpl/file.qtpl:34
	qs422016 := string(qb422016.B)
//line gen/gotpl/file.qtpl:34
	qt422016.ReleaseByteBuffer(qb422016)
//line gen/gotpl/file.qtpl:34
	return qs422016
//line gen/gotpl/file.qtpl:34
}
