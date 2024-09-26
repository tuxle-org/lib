package tables

type Server struct {
	ServerId    int    `db:"server_id"   create:"INTEGER PRIMARY KEY"`
	Name        string `db:"name"        create:"TEXT NOT NULL DEFAULT ''"`
	Description string `db:"description" create:"TEXT NOT NULL DEFAULT ''"`
	IconURI     string `db:"icon_uri"    create:"TEXT NOT NULL DEFAULT ''"`
}

func (table Server) Table() string {
	return "server"
}
