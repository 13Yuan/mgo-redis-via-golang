package db

import (
	"path"
	"os"
	"log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
    Io *IO
)

type (
	IO struct {
		path string
		Org int
		Inst int
		Sale int
	}
	BreakPoint struct {
		Collection int // 0 org_reference, 1 inst_reference, 2 sale_reference
		Number int
	}
)

func init() {
	dir, _ := os.Getwd()
	if Io == nil {
        Io = &IO{path: path.Join(dir, "../src/config/settings.yml")}
    }
}

func (io *IO) Read() (IO, error) {
	raw, err := ioutil.ReadFile(io.path)
	if err != nil {
		return IO{}, err
	}
	var _io IO
	if err := yaml.Unmarshal(raw, &_io); err != nil {
		return IO{}, err
	}

	return _io, nil
}
func (io *IO) Write(collection string, num int) {
	io.Org = -1
	io.Inst = -1
	io.Sale = -1
	switch collection {
		case "org_reference":
			io.Org = num
			io.Inst = 0
			io.Sale = 0
		break
		case "inst_reference":
			io.Inst = num
			io.Sale = 0
		break
		case "sale_reference":
			io.Sale = num
		break
	}
	data, err := yaml.Marshal(io)
	if err != nil {
		log.Println("write file num error.")
	} else {
		ioutil.WriteFile(io.path, data, 0644)
	}
}