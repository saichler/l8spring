package mocks

import (
	"fmt"

	l8m "github.com/saichler/l8common/go/mocks"
	spring "github.com/saichler/l8spring/go/types/spring"
)

func generateCategories() []*spring.SpringCategory {
	count := len(categoryNames)
	categories := make([]*spring.SpringCategory, 0, count)

	for i := 0; i < count; i++ {
		cat := &spring.SpringCategory{
			CategoryId:  l8m.GenID("CAT", i),
			Name:        categoryNames[i],
			Description: fmt.Sprintf("Browse %s items", categoryNames[i]),
			IsActive:    true,
			SortOrder:   int32(i + 1),
			AuditInfo:   l8m.CreateAuditInfo(),
		}
		// Make some categories children of the first few
		if i >= 10 {
			cat.ParentCategoryId = l8m.GenID("CAT", i%5)
		}
		categories = append(categories, cat)
	}
	return categories
}
