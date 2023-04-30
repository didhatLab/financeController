package entrypoints

import (
	"main/finances/entrypoints/middleware"
	"main/finances/entrypoints/webmodels"
	"main/finances/services/group"
	"net/http"
)

type GroupEntryPoint struct {
	AddMemberService    group.AddGroupMemberService
	DeleteMemberService group.DeleteGroupMemberService
}

func (gr GroupEntryPoint) GroupEntryPoint() *http.ServeMux {
	groupAction := gr.groupEntrypoint()
	return groupAction

}

func (gr GroupEntryPoint) groupEntrypoint() *http.ServeMux {
	groupMux := http.NewServeMux()

	groupMux.Handle("/member/add", middleware.AuthMiddleware(http.HandlerFunc(gr.AddNewMember)))

	return groupMux
}

func (gr GroupEntryPoint) AddNewMember(w http.ResponseWriter, req *http.Request) {
	var body webmodels.NewMemberRequest

	ctx := req.Context()

	realUser, ok := middleware.UserFromContext(ctx)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err := webmodels.DecodeJSONBody(w, req, &body)

	if err != nil {
		webmodels.EncodeJSONResponseBody(w, http.StatusBadRequest, struct{ Err string }{Err: err.Error()})
		return
	}

	err = gr.AddMemberService.AddMemberToGroup(ctx, realUser.UserId, body.NewMemberId, body.GroupId)

	if err != nil {
		webmodels.EncodeJSONResponseBody(w, http.StatusBadRequest, struct{ Err string }{Err: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func (gr GroupEntryPoint) DeleteMember(w http.ResponseWriter, req *http.Request) {
	var body webmodels.DeleteMemberRequest

	ctx := req.Context()

	realUser, ok := middleware.UserFromContext(ctx)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err := webmodels.DecodeJSONBody(w, req, &body)

	if err != nil {
		webmodels.EncodeJSONResponseBody(w, http.StatusBadRequest, struct{ Err string }{Err: err.Error()})
	}

	err = gr.DeleteMemberService.DeleteGroupMember(ctx, body.GroupId, realUser.UserId, body.UsernameForDelete)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	return

}
