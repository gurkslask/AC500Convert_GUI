package main

import (
	"encoding/json"
	"strings"

	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
	"github.com/gurkslask/AC500Convert"
)

var choice string

// handleMessages handles messages
func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "init":
		var ch Choices
		ch.Protocol = []string{"COMLI", "Modbus"}
		payload = ch
		return
	case "set":
		var data string
		if len(m.Payload) > 0 {
			// Unmarshal payload
			if err = json.Unmarshal(m.Payload, &data); err != nil {
				payload = err.Error()
				return
			}
			choice = data
		}

	case "update":
		// Unmarshal payload
		var c Communication
		var data string
		if len(m.Payload) > 0 {
			// Unmarshal payload
			if err = json.Unmarshal(m.Payload, &data); err != nil {
				payload = err.Error()
				return
			}
		}
		switch choice {
		case "COMLI":
			res, err := AC500Convert.GenerateAccessComli(strings.Split(data, "\n"))
			if err != nil {
				c.Access = err.Error()
			} else {
				var s strings.Builder
				for _, j := range res {
					s.WriteString(j + "<br>")
				}
				c.Access = s.String()
			}
			resPanel, err := AC500Convert.ExtractDataComli(res)
			if err != nil {
				c.Panel = err.Error()
			} else {
				var s strings.Builder
				s.WriteString("//Name,DataType,GlobalDataType,Address_1,Description //<br>")
				for _, j := range AC500Convert.OutputToText(resPanel) {
					s.WriteString(j + "<br>")
				}
				c.Panel = s.String()
			}
		case "Modbus":
			res, err := AC500Convert.GenerateAccessModbus(strings.Split(data, "\n"))
			if err != nil {
				c.Access = err.Error()
			} else {
				var s strings.Builder
				for _, j := range res {
					s.WriteString(j + "<br>")
				}
				c.Access = s.String()
			}
			resPanel, err := AC500Convert.ExtractDataModbus(res)
			if err != nil {
				c.Panel = err.Error()
			} else {
				var s strings.Builder
				s.WriteString("//Name,DataType,GlobalDataType,Address_1,Description //<br>")
				for _, j := range AC500Convert.OutputToText(resPanel) {
					s.WriteString(j + "<br>")
				}
				c.Panel = s.String()
			}
		}
		payload = c

	}
	return
}

// Communication represents the returned strings from conversion
type Communication struct {
	Access string `json:"access"`
	Panel  string `json:"panel"`
}

// Choices defines what protocol choices exists
type Choices struct {
	Protocol []string `json:"protocol"`
}
