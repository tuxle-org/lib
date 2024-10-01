package entities

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/bbfh-dev/go-tools/tools/tbin"
	"github.com/bbfh-dev/go-tools/tools/terr"
)

type Server struct {
	ServerId    int64  `db:"server_id"   create:"INTEGER PRIMARY KEY"`
	Name        string `db:"name"        create:"TEXT NOT NULL DEFAULT ''"`
	Description string `db:"description" create:"TEXT NOT NULL DEFAULT ''"`
	IconURI     string `db:"icon_uri"    create:"TEXT NOT NULL DEFAULT ''"`
}

func (table *Server) SQL() string {
	return "server"
}

func (table *Server) Serialize(writer io.Writer) error {
	return terr.Join(
		tbin.WriteTuxleString(writer, "server:v1"),
		tbin.WriteUint64(writer, uint64(table.ServerId)),
		tbin.WriteTuxleString(writer, table.Name),
		tbin.WriteTuxleString(writer, table.Description),
		tbin.WriteTuxleString(writer, table.IconURI),
	)
}

func (table *Server) Deserialize(reader *bufio.Reader) error {
	_, err := table.readHeader(reader)
	if err != nil {
		return err
	}

	table.ServerId, err = tbin.Int64(reader)
	if err != nil {
		return err
	}

	table.Name, err = tbin.TuxleString(reader)
	if err != nil {
		return err
	}

	table.Description, err = tbin.TuxleString(reader)
	if err != nil {
		return err
	}

	table.IconURI, err = tbin.TuxleString(reader)
	if err != nil {
		return err
	}

	return nil
}

func (table *Server) readHeader(reader *bufio.Reader) (string, error) {
	var supportedVersions = []string{"v1"}

	checkVersion, err := tbin.TuxleString(reader)
	if err != nil {
		return "", terr.Prefix("Server.Deserialize.version", err)
	}
	if !strings.HasPrefix(checkVersion, "server:") {
		return "", fmt.Errorf("Server.Deserialize expected 'server:...' got: %s", checkVersion)
	}

	version := strings.TrimPrefix(checkVersion, "server:")
	if !slices.Contains(supportedVersions, version) {
		return "", fmt.Errorf(
			"Server.Deserialize unsupported version. Please update your server/client Got: %q",
			version,
		)
	}

	return version, nil
}
