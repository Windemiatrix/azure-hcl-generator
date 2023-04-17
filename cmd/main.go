package main

import (
	tpl_1_1 "azureImportKV/internal/tpl"
	"flag"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	importTplFname := flag.String("import-tpl-file", "./config/import.sh", "path to import file template")
	resourceTpfFname := flag.String("resource-tpl-file", "./config/resource.tf", "path to resource file template")
	resourceType := flag.String("resource-type", "", "resource type: keyvault, sp")
	outputDir := flag.String("output-dir", "output", "directory for output files")
	flag.Parse()

	err := os.MkdirAll(*outputDir, os.ModePerm)
	if err != nil {
		log.Fatal().Stack().Msgf("initialize application: %s", errors.Unwrap(err))
	}

	Resource, err := tpl_1_1.Init(*importTplFname, *resourceTpfFname, *resourceType)
	if err != nil {
		log.Fatal().Stack().Msgf("initialize resource: %s", errors.Unwrap(err))
	}

	contentCh := make(chan tpl_1_1.ResourceResult)
	errCh := make(chan error)
	Resource.Content(contentCh, errCh)

	go func(errCh <-chan error) {
		for err := range errCh {
			log.Fatal().Stack().Msgf("initialize resource: %s", errors.Unwrap(err))
		}
	}(errCh)

	for item := range contentCh {
		err = os.MkdirAll(*outputDir+"/"+item.RG, os.ModePerm)
		if err != nil {
			log.Fatal().Stack().Msgf("main() create directory "+*outputDir+"/"+item.RG+": %s", errors.Unwrap(err))
		}
		err = os.WriteFile(
			*outputDir+"/"+item.RG+"/"+item.Name+".sh",
			[]byte(item.Content.Import),
			0744,
		)
		if err != nil {
			log.Fatal().Stack().Msgf("main() write file "+*outputDir+"/"+item.RG+"/"+item.Name+".sh"+": %s", errors.Unwrap(err))
		}
		err = os.WriteFile(
			*outputDir+"/"+item.RG+"/"+item.Name+".tf",
			[]byte(item.Content.Resource),
			0644,
		)
		if err != nil {
			log.Fatal().Stack().Msgf("main() write file "+*outputDir+"/"+item.RG+"/"+item.Name+".sh"+": %s", errors.Unwrap(err))
		}
	}
}
