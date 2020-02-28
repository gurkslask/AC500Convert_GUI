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

/*
// Exploration represents the results of an exploration
type Exploration struct {
	Dirs       []Dir              `json:"dirs"`
	Files      *astichartjs.Chart `json:"files,omitempty"`
	FilesCount int                `json:"files_count"`
	FilesSize  string             `json:"files_size"`
	Path       string             `json:"path"`
}

// PayloadDir represents a dir payload
type Dir struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// explore explores a path.
// If path is empty, it explores the user's home directory
func explore(path string) (e Exploration, err error) {
	// If no path is provided, use the user's home dir
	if len(path) == 0 {
		var u *user.User
		if u, err = user.Current(); err != nil {
			return
		}
		path = u.HomeDir
	}

	// Read dir
	var files []os.FileInfo
	if files, err = ioutil.ReadDir(path); err != nil {
		return
	}

	// Init exploration
	e = Exploration{
		Dirs: []Dir{},
		Path: path,
	}

	// Add previous dir
	if filepath.Dir(path) != path {
		e.Dirs = append(e.Dirs, Dir{
			Name: "..",
			Path: filepath.Dir(path),
		})
	}

	// Loop through files
	var sizes []int
	var sizesMap = make(map[int][]string)
	var filesSize int64
	for _, f := range files {
		if f.IsDir() {
			e.Dirs = append(e.Dirs, Dir{
				Name: f.Name(),
				Path: filepath.Join(path, f.Name()),
			})
		} else {
			var s = int(f.Size())
			sizes = append(sizes, s)
			sizesMap[s] = append(sizesMap[s], f.Name())
			e.FilesCount++
			filesSize += f.Size()
		}
	}

	// Prepare files size
	if filesSize < 1e3 {
		e.FilesSize = strconv.Itoa(int(filesSize)) + "b"
	} else if filesSize < 1e6 {
		e.FilesSize = strconv.FormatFloat(float64(filesSize)/float64(1024), 'f', 0, 64) + "kb"
	} else if filesSize < 1e9 {
		e.FilesSize = strconv.FormatFloat(float64(filesSize)/float64(1024*1024), 'f', 0, 64) + "Mb"
	} else {
		e.FilesSize = strconv.FormatFloat(float64(filesSize)/float64(1024*1024*1024), 'f', 0, 64) + "Gb"
	}

	// Prepare files chart
	sort.Ints(sizes)
	if len(sizes) > 0 {
		e.Files = &astichartjs.Chart{
			Data: &astichartjs.Data{Datasets: []astichartjs.Dataset{{
				BackgroundColor: []string{
					astichartjs.ChartBackgroundColorYellow,
					astichartjs.ChartBackgroundColorGreen,
					astichartjs.ChartBackgroundColorRed,
					astichartjs.ChartBackgroundColorBlue,
					astichartjs.ChartBackgroundColorPurple,
				},
				BorderColor: []string{
					astichartjs.ChartBorderColorYellow,
					astichartjs.ChartBorderColorGreen,
					astichartjs.ChartBorderColorRed,
					astichartjs.ChartBorderColorBlue,
					astichartjs.ChartBorderColorPurple,
				},
			}}},
			Type: astichartjs.ChartTypePie,
		}
		var sizeOther int
		for i := len(sizes) - 1; i >= 0; i-- {
			for _, l := range sizesMap[sizes[i]] {
				if len(e.Files.Data.Labels) < 4 {
					e.Files.Data.Datasets[0].Data = append(e.Files.Data.Datasets[0].Data, sizes[i])
					e.Files.Data.Labels = append(e.Files.Data.Labels, l)
				} else {
					sizeOther += sizes[i]
				}
			}
		}
		if sizeOther > 0 {
			e.Files.Data.Datasets[0].Data = append(e.Files.Data.Datasets[0].Data, sizeOther)
			e.Files.Data.Labels = append(e.Files.Data.Labels, "other")
		}
	}
	return
}
*/
