package imports

import (
	"os/exec"
)

func NewSOpt(a string, b bool) (e SourceOpt) {
	return SourceOpt{
		name: a,
		mode: b,
	}
}

func NewEOpt(a string, b bool) (e ExecOpt) {
	return ExecOpt{
		name: a,
		mode: b,
	}
}

func NewPOpt(a string, b bool) (e *PackageOpt) {
	return &PackageOpt{
		name: a,
		mode: b,
	}
}

func NewDOpt(a string, b bool) (e DirOpt) {
	return DirOpt{
		name: a,
		mode: b,
	}
}

func NewEraseOpt(a bool) (e *EraseOpt) {
	return &EraseOpt {
		file: a,
	}
}

func NewOpt(a DirOpt, b SourceOpt, c ExecOpt, d *PackageOpt, f *EraseOpt) (e *ImportsOpt) {
	e = &ImportsOpt{
		Dir: a,
		Source: b,
		Exec: c,
		Package: d,
		Erase: f,
	}
	return
}

func NewPlugins(a, b string) (p *Plugins) {
	p = &Plugins {
		Cmd: exec.Command("go", "build", "-buildmode=plugin", "-o", a + ".so", b),
	}
	p.Cmd.Stderr = &p.Stderr
	return
}

func NewImports(name, dest, directory string) (f *Imports) {
	f = new(Imports)
	f.Cmd = exec.Command("go-bindata", "-o", name + ".go", dest)
	f.Cmd.Stderr = &f.Stderr
	return
}
