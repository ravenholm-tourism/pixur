package handlers

import (
	"context"

	"pixur.org/pixur/api"
	"pixur.org/pixur/be/schema"
	"pixur.org/pixur/be/status"
	"pixur.org/pixur/be/tasks"
)

// TODO: add tests

func (s *serv) handleLookupUser(ctx context.Context, req *api.LookupUserRequest) (
	*api.LookupUserResponse, status.S) {
	var objectUserID schema.Varint
	if req.UserId != "" {
		if err := objectUserID.DecodeAll(req.UserId); err != nil {
			return nil, status.InvalidArgument(err, "bad user id")
		}
	}

	var task = &tasks.LookupUserTask{
		DB:           s.db,
		ObjectUserID: int64(objectUserID),
		Ctx:          ctx,
	}

	if sts := s.runner.Run(task); sts != nil {
		return nil, sts
	}

	return &api.LookupUserResponse{
		User: apiUser(task.User),
	}, nil
}
