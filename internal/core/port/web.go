package port

import "io"

type (
	Renderer interface {
		Load(path string) error
		Render(page string, wr io.Writer, data ...interface{}) error
	}
)
