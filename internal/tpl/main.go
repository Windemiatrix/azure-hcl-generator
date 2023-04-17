package tpl

import (
	"azureImportKV/internal/az"
	"bytes"
	"encoding/json"
	"github.com/cheggaaa/pb/v3"
	"github.com/pkg/errors"
	"os"
	"strings"
	"text/template"
)

type IResource interface {
	Content(chan<- ResourceResult, chan<- error)
}

type IResourceItem interface {
	RG() string
	Name() string
	getValues() interface{}
}

type ResourceResult struct {
	RG      string
	Name    string
	Content struct {
		Resource string
		Import   string
	}
}

type resourceBase struct {
	resourceTpl string
	importTpl   string
	cloud       az.ICloud
}

func Init(importTpl, resourceTpl, resourceType string) (Resource IResource, err error) {
	switch resourceType {
	case "keyvault":
		//Resource = &Keyvault{
		//	resourceBase: resourceBase{
		//		resourceTpl: resourceTpl,
		//		importTpl:   importTpl,
		//		cloud:       &az.Keyvault{},
		//	},
		//}
	case "sp":
		Resource = &Sp{
			resourceBase{
				resourceTpl: resourceTpl,
				importTpl:   importTpl,
				cloud:       &az.Sp{},
			},
		}
	default:
		err = errors.New("Unknown resource type")
	}
	return
}

func listLoad(o any, p az.ICloud) (err error) {
	jsonContent, err := p.GetList()
	if err != nil {
		return errors.Wrap(err, "listLoad() load item list")
	}
	err = json.Unmarshal([]byte(jsonContent), o)
	if err != nil {
		return errors.Wrap(err, "listLoad() unmarshal json")
	}
	return
}

func itemLoad(id string, o any, p az.ICloud) (err error) {
	jsonContent, err := p.GetItem(id)
	if err != nil {
		return errors.Wrap(err, "listLoad() load item")
	}
	err = json.Unmarshal([]byte(jsonContent), o)
	if err != nil {
		return errors.Wrap(err, "listLoad() unmarshal json")
	}
	return
}

func tplFileRead(tplName string) (string, error) {
	fcontent, err := os.ReadFile(tplName)
	if err != nil {
		return "", errors.Wrap(err, "read import file template")
	}
	return bytes.NewBuffer(fcontent).String(), nil
}

func tplFileParce(data any, fname string) (string, error) {
	importTpl, err := tplFileRead(fname)
	if err != nil {
		return "", errors.Wrap(err, "tplFileParce()")
	}

	buf := new(bytes.Buffer)
	t := template.Must(template.New(fname).Parse(importTpl))
	err = t.Execute(buf, data)
	if err != nil {
		return "", errors.Wrap(err, "tplFileParce() execute import template")
	}

	return buf.String(), nil
}

func workerBar(in <-chan ResourceResult, out chan<- ResourceResult, bar **pb.ProgressBar) {
	defer close(out)
	for item := range in {
		out <- item
		(*bar).Increment()
	}
}

func resourceName(name string) (result string) {
	result = name
	result = strings.ToLower(result)
	result = strings.ReplaceAll(result, "-", "_")
	result = strings.ReplaceAll(result, "/", "_")
	result = strings.ReplaceAll(result, " ", "_")
	result = strings.ReplaceAll(result, ".", "_")
	return
}
