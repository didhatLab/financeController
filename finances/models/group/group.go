package group

import "main/finances/entrypoints/webmodels"

type SpendGroup struct {
	Id          int
	Name        string
	Description string
	Members     []Member
}

func SpendGroupFromReq(request webmodels.CreateGroupRequest) SpendGroup {
	return SpendGroup{Name: request.GroupName, Description: request.GroupDescription}
}
