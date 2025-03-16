package logging

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"sync"
	"time"

	//"github.com/xuri/excelize/v2"
	"embed"
)

// LogEntry holds a single log record.
type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	CheckName string    `json:"check_name"`
	Message   string    `json:"message"`
	Data      string    `json:"data"`
}

//go:embed templates/log.html.tmpl
var htmlTemplates embed.FS

// loadHTMLTemplate loads the external HTML template.
func loadHTMLTemplate() (*template.Template, error) {
	return template.ParseFS(htmlTemplates, "templates/log.html.tmpl")
}

// writeJSON writes the log entries to a JSON file.
func writeJSON(logEntries []LogEntry, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(logEntries)
}

// writeTXT writes the log entries to a plain text file.
func writeTXT(logEntries []LogEntry, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, entry := range logEntries {
		line := fmt.Sprintf("%s [%s]: %s - %s\n",
			entry.Timestamp.Format(time.RFC3339),
			entry.CheckName,
			entry.Message,
			entry.Data)
		if _, err := file.WriteString(line); err != nil {
			return err
		}
	}
	return nil
}

// writeCSV writes the log entries to a CSV file.
func writeCSV(logEntries []LogEntry, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	if err := writer.Write([]string{"Timestamp", "CheckName", "Message", "Data"}); err != nil {
		return err
	}
	for _, entry := range logEntries {
		record := []string{
			entry.Timestamp.Format(time.RFC3339),
			entry.CheckName,
			entry.Message,
			entry.Data,
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}

// writeHTML writes the log entries to an HTML file using the external template.
func writeHTML(logEntries []LogEntry, filename string) error {
	tmpl, err := loadHTMLTemplate()
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, logEntries)
}

/* writeXLSX writes the log entries to an Excel file. --------Requires Library Code review. -------
func writeXLSX(logEntries []LogEntry, filename string) error {
	f := excelize.NewFile()
	sheet := "Sheet1"

	// Write headers
	headers := []string{"Timestamp", "CheckName", "Message", "Data"}
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, header)
	}

	// Write data rows
	for row, entry := range logEntries {
		rowIndex := row + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", rowIndex), entry.Timestamp.Format(time.RFC3339))
		f.SetCellValue(sheet, fmt.Sprintf("B%d", rowIndex), entry.CheckName)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", rowIndex), entry.Message)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", rowIndex), entry.Data)
	}

	return f.SaveAs(filename)
} */

// WriteAllFormats concurrently writes log entries to multiple file formats.
// The baseFilename parameter will be used as the prefix for all files.
func WriteAllFormats(logEntries []LogEntry, baseFilename string) {
	// Define file tasks for each format.
	fileTasks := []struct {
		ext    string
		writer func([]LogEntry, string) error
	}{
		{".json", writeJSON},
		{".txt", writeTXT},
		{".csv", writeCSV},
		{".html", writeHTML},
		//{".xlsx", writeXLSX},
	}

	var wg sync.WaitGroup
	wg.Add(len(fileTasks))

	for _, task := range fileTasks {
		// Capture the loop variable.
		task := task
		go func() {
			defer wg.Done()
			filename := baseFilename + task.ext
			if err := task.writer(logEntries, filename); err != nil {
				log.Printf("Error writing %s: %v", filename, err)
			} else {
				log.Printf("Successfully wrote %s", filename)
			}
		}()
	}

	wg.Wait()
}
