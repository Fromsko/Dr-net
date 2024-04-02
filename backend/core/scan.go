package core

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

//go:embed lib/scan.exe
var scanProgramData []byte

// embeddedProgram 封装了嵌入的二进制程序
type embeddedProgram struct {
	Data []byte
	Path string
	Arg  string
}

// NewProgram 创建一个封装了嵌入的二进制程序的实例
func NewProgram(arg string) *embeddedProgram {
	return &embeddedProgram{
		Data: scanProgramData,
		Arg:  arg,
	}
}

func (p *embeddedProgram) createTempFile() error {
	tmpFile, err := os.CreateTemp("", "embeddedProgram-*.exe")
	if err != nil {
		return err
	}
	p.Path = tmpFile.Name()
	_, err = tmpFile.Write(p.Data)
	if err != nil {
		tmpFile.Close()
		return err
	}
	return tmpFile.Close()
}

func (p *embeddedProgram) Run() (string, error) {
	if err := p.createTempFile(); err != nil {
		return "", err
	}

	cmd := exec.Command(p.Path, "-t", p.Arg)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running embedded program: %w", err)
	}
	return string(output), nil
}
