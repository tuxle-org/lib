package tuxle

import (
	"github.com/tuxle-org/lib/stream"
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

func (server *Server) Read(reader *stream.Reader) error {
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

func (server *Server) Write(writer *stream.Writer) error {
	return field.Parameters{
		"Name":        server.Name,
		"Description": server.Description,
		"PictureId":   server.PictureId,
	}.Write(writer)
}
