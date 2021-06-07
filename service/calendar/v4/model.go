// Code generated by lark suite oapi sdk gen
package v4

import (
	"github.com/larksuite/oapi-sdk-go/api/core/tools"
)

type UserId struct {
	UserId          string   `json:"user_id,omitempty"`
	OpenId          string   `json:"open_id,omitempty"`
	UnionId         string   `json:"union_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *UserId) MarshalJSON() ([]byte, error) {
	type cp UserId
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type AclScope struct {
	Type            string   `json:"type,omitempty"`
	UserId          string   `json:"user_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *AclScope) MarshalJSON() ([]byte, error) {
	type cp AclScope
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type AttendeeChatMember struct {
	RsvpStatus      string   `json:"rsvp_status,omitempty"`
	IsOptional      bool     `json:"is_optional,omitempty"`
	DisplayName     string   `json:"display_name,omitempty"`
	IsOrganizer     bool     `json:"is_organizer,omitempty"`
	IsExternal      bool     `json:"is_external,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *AttendeeChatMember) MarshalJSON() ([]byte, error) {
	type cp AttendeeChatMember
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type Calendar struct {
	CalendarId      string   `json:"calendar_id,omitempty"`
	Summary         string   `json:"summary,omitempty"`
	Description     string   `json:"description,omitempty"`
	Permissions     string   `json:"permissions,omitempty"`
	Color           int      `json:"color,omitempty"`
	Type            string   `json:"type,omitempty"`
	SummaryAlias    string   `json:"summary_alias,omitempty"`
	IsDeleted       bool     `json:"is_deleted,omitempty"`
	IsThirdParty    bool     `json:"is_third_party,omitempty"`
	Role            string   `json:"role,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *Calendar) MarshalJSON() ([]byte, error) {
	type cp Calendar
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type CalendarAcl struct {
	AclId           string    `json:"acl_id,omitempty"`
	Role            string    `json:"role,omitempty"`
	Scope           *AclScope `json:"scope,omitempty"`
	ForceSendFields []string  `json:"-"`
}

func (s *CalendarAcl) MarshalJSON() ([]byte, error) {
	type cp CalendarAcl
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type CalendarEventAttendee struct {
	Type            string                `json:"type,omitempty"`
	AttendeeId      string                `json:"attendee_id,omitempty"`
	RsvpStatus      string                `json:"rsvp_status,omitempty"`
	IsOptional      bool                  `json:"is_optional,omitempty"`
	IsOrganizer     bool                  `json:"is_organizer,omitempty"`
	IsExternal      bool                  `json:"is_external,omitempty"`
	DisplayName     string                `json:"display_name,omitempty"`
	ChatMembers     []*AttendeeChatMember `json:"chat_members,omitempty"`
	UserId          string                `json:"user_id,omitempty"`
	ChatId          string                `json:"chat_id,omitempty"`
	RoomId          string                `json:"room_id,omitempty"`
	ThirdPartyEmail string                `json:"third_party_email,omitempty"`
	ForceSendFields []string              `json:"-"`
}

func (s *CalendarEventAttendee) MarshalJSON() ([]byte, error) {
	type cp CalendarEventAttendee
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type CalendarEventAttendeeChatMember struct {
	RsvpStatus      string   `json:"rsvp_status,omitempty"`
	IsOptional      bool     `json:"is_optional,omitempty"`
	DisplayName     string   `json:"display_name,omitempty"`
	IsOrganizer     bool     `json:"is_organizer,omitempty"`
	IsExternal      bool     `json:"is_external,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *CalendarEventAttendeeChatMember) MarshalJSON() ([]byte, error) {
	type cp CalendarEventAttendeeChatMember
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type EventLocation struct {
	Name            string   `json:"name,omitempty"`
	Address         string   `json:"address,omitempty"`
	Latitude        float64  `json:"latitude,omitempty"`
	Longitude       float64  `json:"longitude,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *EventLocation) MarshalJSON() ([]byte, error) {
	type cp EventLocation
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type EventSearchFilter struct {
	StartTime       *TimeInfo `json:"start_time,omitempty"`
	EndTime         *TimeInfo `json:"end_time,omitempty"`
	UserIds         []string  `json:"user_ids,omitempty"`
	RoomIds         []string  `json:"room_ids,omitempty"`
	ChatIds         []string  `json:"chat_ids,omitempty"`
	ForceSendFields []string  `json:"-"`
}

func (s *EventSearchFilter) MarshalJSON() ([]byte, error) {
	type cp EventSearchFilter
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type Freebusy struct {
	StartTime       string   `json:"start_time,omitempty"`
	EndTime         string   `json:"end_time,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *Freebusy) MarshalJSON() ([]byte, error) {
	type cp Freebusy
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type Reminder struct {
	Minutes         int      `json:"minutes,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *Reminder) MarshalJSON() ([]byte, error) {
	type cp Reminder
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type Schema struct {
	UiName          string   `json:"ui_name,omitempty"`
	UiStatus        string   `json:"ui_status,omitempty"`
	AppLink         string   `json:"app_link,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *Schema) MarshalJSON() ([]byte, error) {
	type cp Schema
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type Setting struct {
}

type TimeInfo struct {
	Date            string   `json:"date,omitempty"`
	Timestamp       string   `json:"timestamp,omitempty"`
	Timezone        string   `json:"timezone,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *TimeInfo) MarshalJSON() ([]byte, error) {
	type cp TimeInfo
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type TimeoffEvent struct {
	TimeoffEventId  string   `json:"timeoff_event_id,omitempty"`
	UserId          string   `json:"user_id,omitempty"`
	Timezone        string   `json:"timezone,omitempty"`
	StartTime       string   `json:"start_time,omitempty"`
	EndTime         string   `json:"end_time,omitempty"`
	Title           string   `json:"title,omitempty"`
	Description     string   `json:"description,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *TimeoffEvent) MarshalJSON() ([]byte, error) {
	type cp TimeoffEvent
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type Vchat struct {
	MeetingUrl      string   `json:"meeting_url,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *Vchat) MarshalJSON() ([]byte, error) {
	type cp Vchat
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type AclScopeEvent struct {
	Type            string   `json:"type,omitempty"`
	UserId          *UserId  `json:"user_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *AclScopeEvent) MarshalJSON() ([]byte, error) {
	type cp AclScopeEvent
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type CalendarEvent struct {
	EventId          string         `json:"event_id,omitempty"`
	Summary          string         `json:"summary,omitempty"`
	Description      string         `json:"description,omitempty"`
	NeedNotification bool           `json:"need_notification,omitempty"`
	StartTime        *TimeInfo      `json:"start_time,omitempty"`
	EndTime          *TimeInfo      `json:"end_time,omitempty"`
	Vchat            *Vchat         `json:"vchat,omitempty"`
	Visibility       string         `json:"visibility,omitempty"`
	AttendeeAbility  string         `json:"attendee_ability,omitempty"`
	FreeBusyStatus   string         `json:"free_busy_status,omitempty"`
	Location         *EventLocation `json:"location,omitempty"`
	Color            int            `json:"color,omitempty"`
	Reminders        []*Reminder    `json:"reminders,omitempty"`
	Recurrence       string         `json:"recurrence,omitempty"`
	Status           string         `json:"status,omitempty"`
	IsException      bool           `json:"is_exception,omitempty"`
	RecurringEventId string         `json:"recurring_event_id,omitempty"`
	Schemas          []*Schema      `json:"schemas,omitempty"`
	ForceSendFields  []string       `json:"-"`
}

func (s *CalendarEvent) MarshalJSON() ([]byte, error) {
	type cp CalendarEvent
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type ExchangeBinding struct {
	AdminAccount      string   `json:"admin_account,omitempty"`
	ExchangeAccount   string   `json:"exchange_account,omitempty"`
	UserId            string   `json:"user_id,omitempty"`
	Status            string   `json:"status,omitempty"`
	ExchangeBindingId string   `json:"exchange_binding_id,omitempty"`
	ForceSendFields   []string `json:"-"`
}

func (s *ExchangeBinding) MarshalJSON() ([]byte, error) {
	type cp ExchangeBinding
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type CalendarCreateResult struct {
	Calendar *Calendar `json:"calendar,omitempty"`
}

type CalendarEventGetResult struct {
	Event *CalendarEvent `json:"event,omitempty"`
}

type CalendarPatchResult struct {
	Calendar *Calendar `json:"calendar,omitempty"`
}

type CalendarAclListResult struct {
	Acls      []*CalendarAcl `json:"acls,omitempty"`
	HasMore   bool           `json:"has_more,omitempty"`
	PageToken string         `json:"page_token,omitempty"`
}

type CalendarEventCreateResult struct {
	Event *CalendarEvent `json:"event,omitempty"`
}

type CalendarEventAttendeeListResult struct {
	Items     []*CalendarEventAttendee `json:"items,omitempty"`
	HasMore   bool                     `json:"has_more,omitempty"`
	PageToken string                   `json:"page_token,omitempty"`
}

type CalendarListResult struct {
	HasMore      bool        `json:"has_more,omitempty"`
	PageToken    string      `json:"page_token,omitempty"`
	SyncToken    string      `json:"sync_token,omitempty"`
	CalendarList []*Calendar `json:"calendar_list,omitempty"`
}

type CalendarEventAttendeeBatchDeleteReqBody struct {
	AttendeeIds      []string `json:"attendee_ids,omitempty"`
	NeedNotification bool     `json:"need_notification,omitempty"`
	ForceSendFields  []string `json:"-"`
}

func (s *CalendarEventAttendeeBatchDeleteReqBody) MarshalJSON() ([]byte, error) {
	type cp CalendarEventAttendeeBatchDeleteReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type CalendarEventAttendeeCreateReqBody struct {
	Attendees        []*CalendarEventAttendee `json:"attendees,omitempty"`
	NeedNotification bool                     `json:"need_notification,omitempty"`
	ForceSendFields  []string                 `json:"-"`
}

func (s *CalendarEventAttendeeCreateReqBody) MarshalJSON() ([]byte, error) {
	type cp CalendarEventAttendeeCreateReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type CalendarEventAttendeeCreateResult struct {
	Attendees []*CalendarEventAttendee `json:"attendees,omitempty"`
}

type CalendarEventListResult struct {
	HasMore   bool             `json:"has_more,omitempty"`
	PageToken string           `json:"page_token,omitempty"`
	SyncToken string           `json:"sync_token,omitempty"`
	Items     []*CalendarEvent `json:"items,omitempty"`
}

type CalendarSearchReqBody struct {
	Query           string   `json:"query,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *CalendarSearchReqBody) MarshalJSON() ([]byte, error) {
	type cp CalendarSearchReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type CalendarSearchResult struct {
	Items     []*Calendar `json:"items,omitempty"`
	PageToken string      `json:"page_token,omitempty"`
}

type FreebusyListReqBody struct {
	TimeMin         string   `json:"time_min,omitempty"`
	TimeMax         string   `json:"time_max,omitempty"`
	UserId          string   `json:"user_id,omitempty"`
	RoomId          string   `json:"room_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *FreebusyListReqBody) MarshalJSON() ([]byte, error) {
	type cp FreebusyListReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type FreebusyListResult struct {
	FreebusyList []*Freebusy `json:"freebusy_list,omitempty"`
}

type CalendarEventPatchResult struct {
	Event *CalendarEvent `json:"event,omitempty"`
}

type CalendarEventSearchReqBody struct {
	Query           string             `json:"query,omitempty"`
	Filter          *EventSearchFilter `json:"filter,omitempty"`
	ForceSendFields []string           `json:"-"`
}

func (s *CalendarEventSearchReqBody) MarshalJSON() ([]byte, error) {
	type cp CalendarEventSearchReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type CalendarEventSearchResult struct {
	Items     []*CalendarEvent `json:"items,omitempty"`
	PageToken string           `json:"page_token,omitempty"`
}

type CalendarSubscribeResult struct {
	Calendar *Calendar `json:"calendar,omitempty"`
}

type SettingGenerateCaldavConfReqBody struct {
	DeviceName      string   `json:"device_name,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *SettingGenerateCaldavConfReqBody) MarshalJSON() ([]byte, error) {
	type cp SettingGenerateCaldavConfReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type SettingGenerateCaldavConfResult struct {
	Password      string `json:"password,omitempty"`
	UserName      string `json:"user_name,omitempty"`
	ServerAddress string `json:"server_address,omitempty"`
	DeviceName    string `json:"device_name,omitempty"`
}

type CalendarEventAttendeeChatMemberListResult struct {
	Items     []*CalendarEventAttendeeChatMember `json:"items,omitempty"`
	HasMore   bool                               `json:"has_more,omitempty"`
	PageToken string                             `json:"page_token,omitempty"`
}
