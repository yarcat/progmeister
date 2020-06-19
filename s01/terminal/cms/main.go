// Program cms is a simple content management system, which demonstrates
// https://github.com/rivo/tview usage.
package main

import (
	"fmt"
	"log"

	"github.com/rivo/tview"
)

const (
	colID = iota
	colName
	colQuant
)

func main() {
	app := tview.NewApplication()

	depotItems := loadDepotItems()
	itemsTable := tview.NewTable().SetBorders(true)
	itemsTable.SetBorder(true).SetTitle("Depot Items")
	// Header.
	for col, name := range []string{"ID", "Name", "Quantity"} {
		itemsTable.SetCellSimple(0 /* row */, col, name)
	}
	// Items.
	for row, item := range depotItems {
		itemsTable.SetCellSimple(row+1, colID, fmt.Sprint(item.id))
		itemsTable.SetCellSimple(row+1, colName, item.name)
		itemsTable.SetCellSimple(row+1, colQuant, fmt.Sprint(item.quantity))
	}
	app.SetRoot(itemsTable, true /* fullscreen */)

	if err := app.Run(); err != nil {
		log.Fatalln("Failed to run the application:", err)
	}
}
