package intgr

import (
	"sync"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

func Test_Runner(t *testing.T) {

	t.Parallel()

	wg := &sync.WaitGroup{}
	suits := []runner.TestSuite{
		// &AuthSuite{},
		// &HouseSuite{},
		&NoticeSuite{},
		// &ApartmentSuite{},
	}
	wg.Add(len(suits))

	for _, s := range suits {
		go func() {
			suite.RunSuite(t, s)
			wg.Done()
		}()
	}

	wg.Wait()
}