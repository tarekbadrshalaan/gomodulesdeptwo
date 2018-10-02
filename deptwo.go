package gomodulesdeptwo

import (
	"github.com/tarekbadrshalaan/gomodulesdepone/v2/subpkg"
)

// GetDatadepone : get data from depone
func GetDatadepone() string {

	return subpkg.GetExtraData()
}

