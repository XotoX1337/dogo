package platform

import "runtime"

type Platform struct {
	os    string
	shell string
	exec  string
}

func New() Platform {
	p := Platform{}
	p.load()
	return p
}

func (p *Platform) load() {
	p.os = runtime.GOOS
	p.shell = "sh"
	p.exec = "-c"

	if runtime.GOOS == "windows" {
		p.shell = "cmd"
		p.exec = "/c"
	}
}

func (p *Platform) GetOs() string {
	return p.os
}

func (p *Platform) GetShell() string {
	return p.shell
}

func (p *Platform) GetExec() string {
	return p.exec
}
