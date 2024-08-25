package tuxle

import (
	"github.com/tuxle-org/lib/internal"
	"github.com/tuxle-org/lib/tuxle/field"
)

type Server struct {
	Name        string
	Description string
	PictureId   string
}

func NewServer() *Server {
	return &Server{
		Name:        "",
		Description: "",
		PictureId:   "",
	}
}

func (server *Server) Read(reader *internal.Reader) error {
	params := field.Parameters{}
	err := params.ReadUntilEOF(reader)
	if err != nil {
		return err
	}

	err = params.Validate(map[string]field.Validator{
		"Name":        field.NotEmpty,
		"Description": field.NotEmpty,
		"PictureId":   field.Exists,
	})
	if err != nil {
		return err
	}

	server.Name = params["Name"]
	server.Description = params["Description"]
	server.PictureId = params["PictureId"]

	return nil
}

func (server *Server) Write(writer *internal.Writer) error {
	return field.Parameters{
		"Name":        server.Name,
		"Description": server.Description,
		"PictureId":   server.PictureId,
	}.Write(writer)
}
