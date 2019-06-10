package models

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

// Todo --> struct that contains todo's fields
type Todo struct {
	Title       string
	Description string
	Deadline    time.Time
	Completed   bool
}

// EncodeQueryURL encodes the strings from Todo struct making it URL friendly
func (t Todo) EncodeQueryURL() string {

	duedate := t.Deadline.Format("01/02/2006")

	var sb strings.Builder
	sb.WriteString("?title=")
	sb.WriteString(url.QueryEscape(t.Title))
	sb.WriteString("&description=")
	sb.WriteString(url.QueryEscape(t.Description))
	sb.WriteString("&limitdate=")
	sb.WriteString(url.QueryEscape(duedate))
	return sb.String()

}

// MonthFormatted converts type Month to type Int
func (t Todo) MonthFormatted() string {
	monthInt := int(t.Deadline.Month())
	return fmt.Sprintf("%02d", monthInt)
}

// DayFormatted converts type Month to type Int
func (t Todo) DayFormatted() string {
	return fmt.Sprintf("%02d", t.Deadline.Day())
}
