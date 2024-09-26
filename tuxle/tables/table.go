package tables

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/tuxle-org/lib/tuxle/internal"
)

var All = []Table{
	Server{},
}

type Table interface {
	Table() string
}

func CreateTable(table Table) string {
	return fmt.Sprintf(
		"CREATE TABLE %s (%s);",
		table.Table(),
		strings.Join(GetColumns(table), ", "),
	)
}

func GetColumnsAsMap(table Table) map[string]string {
	tableType := reflect.TypeOf(table)
	var out = map[string]string{}

	for i := range tableType.NumField() {
		field := tableType.Field(i)
		out[field.Tag.Get("db")] = field.Tag.Get("create")
	}

	return out
}

func SplitRow(row string) (string, string) {
	parts := strings.SplitN(row, " ", 2)
	internal.Assert(len(parts) == 2, "Row must have a space after name")

	return strings.Trim(parts[0], `'`), parts[1]
}

func RowsToMap(rows []string) map[string]string {
	var out = map[string]string{}
	for _, row := range rows {
		col, value := SplitRow(row)
		out[col] = value
	}
	return out
}

func GetColumns(table Table) []string {
	tableType := reflect.TypeOf(table)
	var out = make([]string, tableType.NumField())

	for i := range tableType.NumField() {
		field := tableType.Field(i)
		out[i] = fmt.Sprintf(
			"'%s' %s",
			field.Tag.Get("db"),
			field.Tag.Get("create"),
		)
	}

	return out
}

var Keywords = []string{
	"PRIMARY KEY",
	"NOT NULL",
	"DEFAULT ?",
	"UNIQUE",
	"CHECK?",
}

func SplitByKeywords(row string) []string {
	var result []string

	for _, keyword := range Keywords {
		if strings.Contains(keyword, "?") {
			// Replace the "?" with a regular expression that matches non-space characters
			re := regexp.MustCompile(strings.Replace(keyword, "?", `\S+`, -1))
			matches := re.FindAllString(row, -1)
			for _, match := range matches {
				result = append(result, match)
				row = strings.Replace(row, match, "", 1)
			}
			continue
		}

		if strings.Contains(row, keyword) {
			result = append(result, keyword)
			row = strings.Replace(row, keyword, "", 1)
		}
	}

	return result
}
