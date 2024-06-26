// Code generated by goa v3.14.0, DO NOT EDIT.
//
// users HTTP server types
//
// Command:
// $ goa gen users/design

package server

import (
	users "users/gen/users"

	goa "goa.design/goa/v3/pkg"
)

// CreateUserRequestBody is the type of the "users" service "CreateUser"
// endpoint HTTP request body.
type CreateUserRequestBody struct {
	Username    *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
	DisplayName *string `form:"display_name,omitempty" json:"display_name,omitempty" xml:"display_name,omitempty"`
}

// DeleteUserRequestBody is the type of the "users" service "DeleteUser"
// endpoint HTTP request body.
type DeleteUserRequestBody struct {
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// UpdateUsernameRequestBody is the type of the "users" service
// "UpdateUsername" endpoint HTTP request body.
type UpdateUsernameRequestBody struct {
	Username *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
}

// UpdateBioRequestBody is the type of the "users" service "UpdateBio" endpoint
// HTTP request body.
type UpdateBioRequestBody struct {
	Bio *string `form:"bio,omitempty" json:"bio,omitempty" xml:"bio,omitempty"`
}

// FollowRequestBody is the type of the "users" service "Follow" endpoint HTTP
// request body.
type FollowRequestBody struct {
	FollowerID *int `form:"follower_id,omitempty" json:"follower_id,omitempty" xml:"follower_id,omitempty"`
	FolloweeID *int `form:"followee_id,omitempty" json:"followee_id,omitempty" xml:"followee_id,omitempty"`
}

// UnfollowRequestBody is the type of the "users" service "Unfollow" endpoint
// HTTP request body.
type UnfollowRequestBody struct {
	FollowerID *int `form:"follower_id,omitempty" json:"follower_id,omitempty" xml:"follower_id,omitempty"`
	FolloweeID *int `form:"followee_id,omitempty" json:"followee_id,omitempty" xml:"followee_id,omitempty"`
}

// CreateUserResponseBody is the type of the "users" service "CreateUser"
// endpoint HTTP response body.
type CreateUserResponseBody struct {
	ID          int    `form:"id" json:"id" xml:"id"`
	Username    string `form:"username" json:"username" xml:"username"`
	DisplayName string `form:"display_name" json:"display_name" xml:"display_name"`
	Bio         string `form:"bio" json:"bio" xml:"bio"`
	CreatedAt   string `form:"created_at" json:"created_at" xml:"created_at"`
	UpdatedAt   string `form:"updated_at" json:"updated_at" xml:"updated_at"`
}

// FindUserByIDResponseBody is the type of the "users" service "FindUserByID"
// endpoint HTTP response body.
type FindUserByIDResponseBody struct {
	ID          int    `form:"id" json:"id" xml:"id"`
	Username    string `form:"username" json:"username" xml:"username"`
	DisplayName string `form:"display_name" json:"display_name" xml:"display_name"`
	Bio         string `form:"bio" json:"bio" xml:"bio"`
	CreatedAt   string `form:"created_at" json:"created_at" xml:"created_at"`
	UpdatedAt   string `form:"updated_at" json:"updated_at" xml:"updated_at"`
}

// GetFollowersResponseBody is the type of the "users" service "GetFollowers"
// endpoint HTTP response body.
type GetFollowersResponseBody []*UserResponse

// GetFollowingsResponseBody is the type of the "users" service "GetFollowings"
// endpoint HTTP response body.
type GetFollowingsResponseBody []*UserResponse

// CreateUserBadRequestResponseBody is the type of the "users" service
// "CreateUser" endpoint HTTP response body for the "BadRequest" error.
type CreateUserBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeleteUserNotFoundResponseBody is the type of the "users" service
// "DeleteUser" endpoint HTTP response body for the "NotFound" error.
type DeleteUserNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeleteUserBadRequestResponseBody is the type of the "users" service
// "DeleteUser" endpoint HTTP response body for the "BadRequest" error.
type DeleteUserBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// FindUserByIDNotFoundResponseBody is the type of the "users" service
// "FindUserByID" endpoint HTTP response body for the "NotFound" error.
type FindUserByIDNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateUsernameNotFoundResponseBody is the type of the "users" service
// "UpdateUsername" endpoint HTTP response body for the "NotFound" error.
type UpdateUsernameNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateUsernameBadRequestResponseBody is the type of the "users" service
// "UpdateUsername" endpoint HTTP response body for the "BadRequest" error.
type UpdateUsernameBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateBioNotFoundResponseBody is the type of the "users" service "UpdateBio"
// endpoint HTTP response body for the "NotFound" error.
type UpdateBioNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateBioBadRequestResponseBody is the type of the "users" service
// "UpdateBio" endpoint HTTP response body for the "BadRequest" error.
type UpdateBioBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// FollowBadRequestResponseBody is the type of the "users" service "Follow"
// endpoint HTTP response body for the "BadRequest" error.
type FollowBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UnfollowBadRequestResponseBody is the type of the "users" service "Unfollow"
// endpoint HTTP response body for the "BadRequest" error.
type UnfollowBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetFollowersBadRequestResponseBody is the type of the "users" service
// "GetFollowers" endpoint HTTP response body for the "BadRequest" error.
type GetFollowersBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetFollowingsBadRequestResponseBody is the type of the "users" service
// "GetFollowings" endpoint HTTP response body for the "BadRequest" error.
type GetFollowingsBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UserResponse is used to define fields on response body types.
type UserResponse struct {
	ID          int    `form:"id" json:"id" xml:"id"`
	Username    string `form:"username" json:"username" xml:"username"`
	DisplayName string `form:"display_name" json:"display_name" xml:"display_name"`
	Bio         string `form:"bio" json:"bio" xml:"bio"`
	CreatedAt   string `form:"created_at" json:"created_at" xml:"created_at"`
	UpdatedAt   string `form:"updated_at" json:"updated_at" xml:"updated_at"`
}

// NewCreateUserResponseBody builds the HTTP response body from the result of
// the "CreateUser" endpoint of the "users" service.
func NewCreateUserResponseBody(res *users.User) *CreateUserResponseBody {
	body := &CreateUserResponseBody{
		ID:          res.ID,
		Username:    res.Username,
		DisplayName: res.DisplayName,
		Bio:         res.Bio,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}
	return body
}

// NewFindUserByIDResponseBody builds the HTTP response body from the result of
// the "FindUserByID" endpoint of the "users" service.
func NewFindUserByIDResponseBody(res *users.User) *FindUserByIDResponseBody {
	body := &FindUserByIDResponseBody{
		ID:          res.ID,
		Username:    res.Username,
		DisplayName: res.DisplayName,
		Bio:         res.Bio,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}
	return body
}

// NewGetFollowersResponseBody builds the HTTP response body from the result of
// the "GetFollowers" endpoint of the "users" service.
func NewGetFollowersResponseBody(res []*users.User) GetFollowersResponseBody {
	body := make([]*UserResponse, len(res))
	for i, val := range res {
		body[i] = marshalUsersUserToUserResponse(val)
	}
	return body
}

// NewGetFollowingsResponseBody builds the HTTP response body from the result
// of the "GetFollowings" endpoint of the "users" service.
func NewGetFollowingsResponseBody(res []*users.User) GetFollowingsResponseBody {
	body := make([]*UserResponse, len(res))
	for i, val := range res {
		body[i] = marshalUsersUserToUserResponse(val)
	}
	return body
}

// NewCreateUserBadRequestResponseBody builds the HTTP response body from the
// result of the "CreateUser" endpoint of the "users" service.
func NewCreateUserBadRequestResponseBody(res *goa.ServiceError) *CreateUserBadRequestResponseBody {
	body := &CreateUserBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteUserNotFoundResponseBody builds the HTTP response body from the
// result of the "DeleteUser" endpoint of the "users" service.
func NewDeleteUserNotFoundResponseBody(res *goa.ServiceError) *DeleteUserNotFoundResponseBody {
	body := &DeleteUserNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteUserBadRequestResponseBody builds the HTTP response body from the
// result of the "DeleteUser" endpoint of the "users" service.
func NewDeleteUserBadRequestResponseBody(res *goa.ServiceError) *DeleteUserBadRequestResponseBody {
	body := &DeleteUserBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewFindUserByIDNotFoundResponseBody builds the HTTP response body from the
// result of the "FindUserByID" endpoint of the "users" service.
func NewFindUserByIDNotFoundResponseBody(res *goa.ServiceError) *FindUserByIDNotFoundResponseBody {
	body := &FindUserByIDNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateUsernameNotFoundResponseBody builds the HTTP response body from the
// result of the "UpdateUsername" endpoint of the "users" service.
func NewUpdateUsernameNotFoundResponseBody(res *goa.ServiceError) *UpdateUsernameNotFoundResponseBody {
	body := &UpdateUsernameNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateUsernameBadRequestResponseBody builds the HTTP response body from
// the result of the "UpdateUsername" endpoint of the "users" service.
func NewUpdateUsernameBadRequestResponseBody(res *goa.ServiceError) *UpdateUsernameBadRequestResponseBody {
	body := &UpdateUsernameBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateBioNotFoundResponseBody builds the HTTP response body from the
// result of the "UpdateBio" endpoint of the "users" service.
func NewUpdateBioNotFoundResponseBody(res *goa.ServiceError) *UpdateBioNotFoundResponseBody {
	body := &UpdateBioNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateBioBadRequestResponseBody builds the HTTP response body from the
// result of the "UpdateBio" endpoint of the "users" service.
func NewUpdateBioBadRequestResponseBody(res *goa.ServiceError) *UpdateBioBadRequestResponseBody {
	body := &UpdateBioBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewFollowBadRequestResponseBody builds the HTTP response body from the
// result of the "Follow" endpoint of the "users" service.
func NewFollowBadRequestResponseBody(res *goa.ServiceError) *FollowBadRequestResponseBody {
	body := &FollowBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUnfollowBadRequestResponseBody builds the HTTP response body from the
// result of the "Unfollow" endpoint of the "users" service.
func NewUnfollowBadRequestResponseBody(res *goa.ServiceError) *UnfollowBadRequestResponseBody {
	body := &UnfollowBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetFollowersBadRequestResponseBody builds the HTTP response body from the
// result of the "GetFollowers" endpoint of the "users" service.
func NewGetFollowersBadRequestResponseBody(res *goa.ServiceError) *GetFollowersBadRequestResponseBody {
	body := &GetFollowersBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetFollowingsBadRequestResponseBody builds the HTTP response body from
// the result of the "GetFollowings" endpoint of the "users" service.
func NewGetFollowingsBadRequestResponseBody(res *goa.ServiceError) *GetFollowingsBadRequestResponseBody {
	body := &GetFollowingsBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateUserPayload builds a users service CreateUser endpoint payload.
func NewCreateUserPayload(body *CreateUserRequestBody) *users.CreateUserPayload {
	v := &users.CreateUserPayload{
		Username:    *body.Username,
		DisplayName: *body.DisplayName,
	}

	return v
}

// NewDeleteUserPayload builds a users service DeleteUser endpoint payload.
func NewDeleteUserPayload(body *DeleteUserRequestBody) *users.DeleteUserPayload {
	v := &users.DeleteUserPayload{
		ID: *body.ID,
	}

	return v
}

// NewFindUserByIDPayload builds a users service FindUserByID endpoint payload.
func NewFindUserByIDPayload(id int) *users.FindUserByIDPayload {
	v := &users.FindUserByIDPayload{}
	v.ID = id

	return v
}

// NewUpdateUsernamePayload builds a users service UpdateUsername endpoint
// payload.
func NewUpdateUsernamePayload(body *UpdateUsernameRequestBody, id int) *users.UpdateUsernamePayload {
	v := &users.UpdateUsernamePayload{
		Username: *body.Username,
	}
	v.ID = id

	return v
}

// NewUpdateBioPayload builds a users service UpdateBio endpoint payload.
func NewUpdateBioPayload(body *UpdateBioRequestBody, id int) *users.UpdateBioPayload {
	v := &users.UpdateBioPayload{
		Bio: *body.Bio,
	}
	v.ID = id

	return v
}

// NewFollowPayload builds a users service Follow endpoint payload.
func NewFollowPayload(body *FollowRequestBody) *users.FollowPayload {
	v := &users.FollowPayload{
		FollowerID: *body.FollowerID,
		FolloweeID: *body.FolloweeID,
	}

	return v
}

// NewUnfollowPayload builds a users service Unfollow endpoint payload.
func NewUnfollowPayload(body *UnfollowRequestBody) *users.UnfollowPayload {
	v := &users.UnfollowPayload{
		FollowerID: *body.FollowerID,
		FolloweeID: *body.FolloweeID,
	}

	return v
}

// NewGetFollowersPayload builds a users service GetFollowers endpoint payload.
func NewGetFollowersPayload(id int) *users.GetFollowersPayload {
	v := &users.GetFollowersPayload{}
	v.ID = id

	return v
}

// NewGetFollowingsPayload builds a users service GetFollowings endpoint
// payload.
func NewGetFollowingsPayload(id int) *users.GetFollowingsPayload {
	v := &users.GetFollowingsPayload{}
	v.ID = id

	return v
}

// ValidateCreateUserRequestBody runs the validations defined on
// CreateUserRequestBody
func ValidateCreateUserRequestBody(body *CreateUserRequestBody) (err error) {
	if body.Username == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("username", "body"))
	}
	if body.DisplayName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("display_name", "body"))
	}
	return
}

// ValidateDeleteUserRequestBody runs the validations defined on
// DeleteUserRequestBody
func ValidateDeleteUserRequestBody(body *DeleteUserRequestBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	return
}

// ValidateUpdateUsernameRequestBody runs the validations defined on
// UpdateUsernameRequestBody
func ValidateUpdateUsernameRequestBody(body *UpdateUsernameRequestBody) (err error) {
	if body.Username == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("username", "body"))
	}
	return
}

// ValidateUpdateBioRequestBody runs the validations defined on
// UpdateBioRequestBody
func ValidateUpdateBioRequestBody(body *UpdateBioRequestBody) (err error) {
	if body.Bio == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("bio", "body"))
	}
	return
}

// ValidateFollowRequestBody runs the validations defined on FollowRequestBody
func ValidateFollowRequestBody(body *FollowRequestBody) (err error) {
	if body.FollowerID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("follower_id", "body"))
	}
	if body.FolloweeID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("followee_id", "body"))
	}
	return
}

// ValidateUnfollowRequestBody runs the validations defined on
// UnfollowRequestBody
func ValidateUnfollowRequestBody(body *UnfollowRequestBody) (err error) {
	if body.FollowerID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("follower_id", "body"))
	}
	if body.FolloweeID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("followee_id", "body"))
	}
	return
}
