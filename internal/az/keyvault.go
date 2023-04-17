package az

import (
	"bytes"
	"github.com/pkg/errors"
	"os/exec"
)

type Keyvault struct{}

func (this *Keyvault) GetList() (string, error) {
	app := "az"

	args := []string{
		"keyvault",
		"list",
	}

	cmd := exec.Command(app, args...)
	result, err := cmd.Output()
	if err != nil {
		return "", errors.Wrap(err, "Execute az keyvault list")
	}

	return bytes.NewBuffer(result).String(), err
}

func (this *Keyvault) GetItem(itemName string) (string, error) {
	app := "az"

	args := []string{
		"keyvault",
		"show",
		"--name",
		itemName,
	}

	cmd := exec.Command(app, args...)
	result, err := cmd.Output()
	if err != nil {
		return "", errors.Wrap(err, "Execute az keyvault list")
	}

	return bytes.NewBuffer(result).String(), err
}
