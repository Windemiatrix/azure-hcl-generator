package az

import (
	"bytes"
	"github.com/pkg/errors"
	"os/exec"
)

type App struct{}

func (this *App) GetList() (string, error) {
	app := "az"

	args := []string{
		"ad",
		"app",
		"list",
		"--all",
	}

	cmd := exec.Command(app, args...)
	result, err := cmd.Output()
	if err != nil {
		return "", errors.Wrap(err, "Execute `az ad app list --all`")
	}

	return bytes.NewBuffer(result).String(), err
}

func (this *App) GetItem(appId string) (string, error) {
	app := "az"

	args := []string{
		"ad",
		"app",
		"show",
		"--id",
		appId,
	}

	cmd := exec.Command(app, args...)
	result, err := cmd.Output()
	if err != nil {
		return "", errors.Wrap(err, "Execute `az ad app show --id "+appId+"`")
	}

	return bytes.NewBuffer(result).String(), err
}
