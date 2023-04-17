package tpl

import (
	"azureImportKV/internal/az"
	"github.com/cheggaaa/pb/v3"
	"github.com/pkg/errors"
	"sync"
)

type Sp struct {
	resourceBase
}

type spListType []struct {
	ID                     string `json:"id"`
	AppOwnerOrganizationID string `json:"appOwnerOrganizationId"`
	AppID                  string `json:"appId"`
}

type spObjType struct {
	Sp  SpParcedItem
	App AppParcedItem
}

type SpParcedItem struct {
	AccountEnabled             bool   `json:"accountEnabled,omitempty"`
	AlternativeNames           []any  `json:"alternativeNames,omitempty"`
	AppDisplayName             string `json:"appDisplayName,omitempty"`
	AppID                      string `json:"appId,omitempty"`
	AppRoleAssignmentRequired  bool   `json:"appRoleAssignmentRequired,omitempty"`
	Description                any    `json:"description,omitempty"`
	DisplayName                string `json:"displayName,omitempty"`
	ID                         string `json:"id,omitempty"`
	LoginURL                   any    `json:"loginUrl,omitempty"`
	Notes                      any    `json:"notes,omitempty"`
	NotificationEmailAddresses []any  `json:"notificationEmailAddresses,omitempty"`
	PreferredSingleSignOnMode  any    `json:"preferredSingleSignOnMode,omitempty"`
}

func (this *Sp) Content(resCh chan<- ResourceResult, errCh chan<- error) {
	errWg := &sync.WaitGroup{}
	stageWg := &sync.WaitGroup{}
	barCh := make(chan ResourceResult)

	bar := pb.ProgressBarTemplate(`Load SP list: {{percent .}}`).Start(1)

	// load list of all items
	spList := spListType{}
	stageWg.Add(1)
	go func(errCh chan<- error, errWg *sync.WaitGroup, spList *spListType, stageWg *sync.WaitGroup) {
		defer stageWg.Done()
		err := listLoad(spList, &az.Sp{})
		if err != nil {
			errWg.Add(1)
			go func(errCh chan<- error, errWg *sync.WaitGroup) {
				defer errWg.Done()
				errCh <- errors.Wrap(err, "Content() load item list")
			}(errCh, errWg)
		}
	}(errCh, errWg, &spList, stageWg)

	go func(bar **pb.ProgressBar, stageWg *sync.WaitGroup, spList *spListType) {
		stageWg.Wait()
		(*bar).Increment()
		(*bar).Finish()
		(*bar) = pb.ProgressBarTemplate(`Load SP items: {{percent .}} {{bar .}} ({{counters .}} | {{speed .}})`).Start(len(*spList))
	}(&bar, stageWg, &spList)

	// load items
	spListCh := make(chan *spObjType)
	errWg.Add(1)
	go func(out chan<- *spObjType, errCh chan<- error, errWg *sync.WaitGroup, stageWg *sync.WaitGroup, obj *spListType) {
		defer close(out)
		defer errWg.Done()
		stageWg.Wait()
		for i := range *obj {
			if (*obj)[i].AppOwnerOrganizationID != "" && (*obj)[i].AppOwnerOrganizationID != "2b8d2a3c-bc5d-4c9e-9de5-9219a69d6524" {
				continue
			}
			item := spObjType{
				Sp:  SpParcedItem{},
				App: AppParcedItem{},
			}
			err := itemLoad((*obj)[i].ID, &item.Sp, &az.Sp{})
			if err != nil {
				errWg.Add(1)
				go func(errCh chan<- error, errWg *sync.WaitGroup) {
					defer errWg.Done()
					errCh <- errors.Wrap(err, "Content() load sp item")
				}(errCh, errWg)
			}
			// check if we have application
			if (&item.Sp).AppDisplayName != "" {
				err = itemLoad((*obj)[i].AppID, &item.App, &az.App{})
				if err != nil {
					errWg.Add(1)
					go func(errCh chan<- error, errWg *sync.WaitGroup) {
						defer errWg.Done()
						errCh <- errors.Wrap(err, "Content() load app item")
					}(errCh, errWg)
				}
			}
			out <- &item
		}
	}(spListCh, errCh, errWg, stageWg, &spList)

	// convert items to render files
	errWg.Add(1)
	go func(in <-chan *spObjType, out chan<- ResourceResult, errCh chan<- error, errWg *sync.WaitGroup, this *Sp) {
		defer close(out)
		defer errWg.Done()
		for i := range in {
			res, err := tplFileParce(i, this.resourceTpl)
			if err != nil {
				errWg.Add(1)
				go func(errCh chan<- error, errWg *sync.WaitGroup) {
					defer errWg.Done()
					errCh <- errors.Wrap(err, "Content() render resource template")
				}(errCh, errWg)
			}
			imp, err := tplFileParce(i, this.importTpl)
			if err != nil {
				errWg.Add(1)
				go func(errCh chan<- error, errWg *sync.WaitGroup) {
					defer errWg.Done()
					errCh <- errors.Wrap(err, "Content() render import template")
				}(errCh, errWg)
			}
			out <- ResourceResult{
				RG:   "global",
				Name: i.ResourceName(),
				Content: struct {
					Resource string
					Import   string
				}{res, imp},
			}
		}
	}(spListCh, barCh, errCh, errWg, this)

	go workerBar(barCh, resCh, &bar)

	go func(errCh chan<- error, errWg *sync.WaitGroup) {
		defer close(errCh)
		errWg.Wait()
	}(errCh, errWg)

	return
}

func (this spObjType) ResourceName() string {
	return resourceName(this.Sp.DisplayName)
}
