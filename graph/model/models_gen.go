// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type MeetUpFilter struct {
	Name *string `json:"name"`
}

type NewMeetup struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateMeetup struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}
