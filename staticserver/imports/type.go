package imports

import (
	"os/exec"
	"io"
	"bytes"
)

type DirOpt struct {
	name string
	mode bool
}

type ExecOpt struct {
	name string
	mode bool
}

type SourceOpt struct {
	name string
	mode bool
}

type PackageOpt struct {
	name string
	mode bool
}

type EraseOpt struct {
	file bool
}

type ImportsOpt struct {
	Dir DirOpt
	Exec ExecOpt
	Source SourceOpt
	Package *PackageOpt
	Erase *EraseOpt
}

type Plugins struct {
	Cmd *exec.Cmd
	Stderr bytes.Buffer
}

type Imports struct {
	Opt ImportsOpt
	Cmd *exec.Cmd
	Stdin io.Reader
	Stdout bytes.Buffer
	Stderr bytes.Buffer
}
