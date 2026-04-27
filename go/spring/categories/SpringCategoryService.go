package categories

import (
	common "github.com/saichler/l8spring/go/spring/common"
	"github.com/saichler/l8spring/go/types/spring"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "Category"
	ServiceArea = byte(10)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	common.ActivateService(common.ServiceConfig{
		ServiceName: ServiceName, ServiceArea: ServiceArea,
		PrimaryKey: "CategoryId", Callback: newSpringCategoryServiceCallback(vnic),
	}, &spring.SpringCategory{}, &spring.SpringCategoryList{}, creds, dbname, vnic)
}
