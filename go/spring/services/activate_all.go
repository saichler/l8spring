package services

import (
	"sync"

	"github.com/saichler/l8spring/go/spring/bids"
	"github.com/saichler/l8spring/go/spring/categories"
	"github.com/saichler/l8spring/go/spring/deals"
	"github.com/saichler/l8spring/go/spring/listings"
	"github.com/saichler/l8spring/go/spring/reviews"
	"github.com/saichler/l8types/go/ifs"
)

func ActivateAllServices(creds, dbname string, nic ifs.IVNic) {
	all := []func(){
		func() { categories.Activate(creds, dbname, nic) },
		func() { listings.Activate(creds, dbname, nic) },
		func() { bids.Activate(creds, dbname, nic) },
		func() { deals.Activate(creds, dbname, nic) },
		func() { reviews.Activate(creds, dbname, nic) },
	}

	sem := make(chan struct{}, 5)
	var wg sync.WaitGroup

	for _, fn := range all {
		wg.Add(1)
		sem <- struct{}{}
		go func(f func()) {
			defer wg.Done()
			defer func() { <-sem }()
			f()
		}(fn)
	}
	wg.Wait()
}
