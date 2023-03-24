package entity

import "errors"

var (
	ErrTitleIsBlank        = errors.New("title cannot be blank")
	ErrUserIdNotValid      = errors.New("user id is not valid")
	ErrStatusInvalid       = errors.New("status is not valid")
	ErrTaskDeleted         = errors.New("task has been deleted")
	ErrTaskNotFound        = errors.New("task not found")
	ErrCannotCreateTask    = errors.New("cannot create task")
	ErrCannotUpdateTask    = errors.New("cannot update task")
	ErrCannotDeleteTask    = errors.New("cannot update task")
	ErrCannotListTask      = errors.New("cannot list tasks")
	ErrCannotGetTask       = errors.New("cannot get task details")
	ErrRequesterIsNotOwner = errors.New("no permission, only task owner can do this")
)
