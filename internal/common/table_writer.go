package common

import (
	"io"

	"github.com/olekukonko/tablewriter"
)

func NewTableWriter(writer io.Writer) *tablewriter.Table {
	table := tablewriter.NewWriter(writer)

	table.SetBorder(false)
	table.SetRowLine(false)
	table.SetColumnSeparator("\t")
	table.SetColWidth(80)

	return table
}
