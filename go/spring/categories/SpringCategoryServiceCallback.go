package categories

import (
	common "github.com/saichler/l8spring/go/spring/common"
	"github.com/saichler/l8spring/go/types/spring"
	"github.com/saichler/l8types/go/ifs"
)

func newSpringCategoryServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&spring.SpringCategory{}, vnic).
		Require(func(v interface{}) string { return v.(*spring.SpringCategory).Name }, "Name").
		Build()
}
