package template

import (
	"os"
	"path/filepath"
	"text/template"
)

type FileInfo struct {
	Template []byte
	FileName string
}

type BuildInfo struct {
	FilePath string
	Data     any
}

func (f *FileInfo) Build(info *BuildInfo) error {
	var (
		file *os.File
		tmpl *template.Template
		err  error
	)

	if file, err = os.Create(filepath.Join(info.FilePath, f.FileName)); err != nil {
		return err
	}

	if tmpl, err = template.New(f.FileName).Parse(string(f.Template)); err != nil {
		return err
	}

	return tmpl.Execute(file, info.Data)
}
