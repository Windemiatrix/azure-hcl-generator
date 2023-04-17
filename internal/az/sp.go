package az

import (
	"bytes"
	"github.com/pkg/errors"
	"os/exec"
)

type Sp struct{}

func (this *Sp) GetList() (string, error) {
	app := "az"

	args := []string{
		"ad",
		"sp",
		"list",
		"--all",
		"--query",
		"[?appOwnerOrganizationId=='2b8d2a3c-bc5d-4c9e-9de5-9219a69d6524' || appOwnerOrganizationId==null]",
		"-o",
		"json",
	}

	cmd := exec.Command(app, args...)
	result, err := cmd.Output()
	if err != nil {
		return "", errors.Wrap(err, "Execute `az ad sp list --all`")
	}

	return bytes.NewBuffer(result).String(), err
}

func (this *Sp) GetItem(spId string) (string, error) {
	app := "az"

	args := []string{
		"ad",
		"sp",
		"show",
		"--id",
		spId,
	}

	cmd := exec.Command(app, args...)
	result, err := cmd.Output()
	if err != nil {
		return "", errors.Wrap(err, "Execute `az ad sp show --id "+spId+"`")
	}

	return bytes.NewBuffer(result).String(), err
}
