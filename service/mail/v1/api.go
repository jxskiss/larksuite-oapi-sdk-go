// Code generated by lark suite oapi sdk gen
package v1

import (
	"github.com/larksuite/oapi-sdk-go"
)

type Service struct {
	conf                       lark.Config
	MailgroupMembers           *MailgroupMemberService
	MailgroupPermissionMembers *MailgroupPermissionMemberService
	PublicMailboxs             *PublicMailboxService
	PublicMailboxMembers       *PublicMailboxMemberService
	Mailgroups                 *MailgroupService
}

func NewService(conf lark.Config) *Service {
	s := &Service{
		conf: conf,
	}
	s.MailgroupMembers = newMailgroupMemberService(s)
	s.MailgroupPermissionMembers = newMailgroupPermissionMemberService(s)
	s.PublicMailboxs = newPublicMailboxService(s)
	s.PublicMailboxMembers = newPublicMailboxMemberService(s)
	s.Mailgroups = newMailgroupService(s)
	return s
}

type MailgroupMemberService struct {
	service *Service
}

func newMailgroupMemberService(service *Service) *MailgroupMemberService {
	return &MailgroupMemberService{
		service: service,
	}
}

type MailgroupPermissionMemberService struct {
	service *Service
}

func newMailgroupPermissionMemberService(service *Service) *MailgroupPermissionMemberService {
	return &MailgroupPermissionMemberService{
		service: service,
	}
}

type PublicMailboxService struct {
	service *Service
}

func newPublicMailboxService(service *Service) *PublicMailboxService {
	return &PublicMailboxService{
		service: service,
	}
}

type PublicMailboxMemberService struct {
	service *Service
}

func newPublicMailboxMemberService(service *Service) *PublicMailboxMemberService {
	return &PublicMailboxMemberService{
		service: service,
	}
}

type MailgroupService struct {
	service *Service
}

func newMailgroupService(service *Service) *MailgroupService {
	return &MailgroupService{
		service: service,
	}
}

type MailgroupPermissionMemberDeleteReqCall struct {
	ctx                        *lark.Context
	mailgroupPermissionMembers *MailgroupPermissionMemberService
	pathParams                 map[string]interface{}
	opts                       []lark.APIRequestOpt
}

func (rc *MailgroupPermissionMemberDeleteReqCall) SetMailgroupId(mailgroupId string) {
	rc.pathParams["mailgroup_id"] = mailgroupId
}
func (rc *MailgroupPermissionMemberDeleteReqCall) SetPermissionMemberId(permissionMemberId string) {
	rc.pathParams["permission_member_id"] = permissionMemberId
}

func (rc *MailgroupPermissionMemberDeleteReqCall) Do() (*lark.NoData, error) {
	rc.opts = append(rc.opts, lark.SetPathParams(rc.pathParams))
	var result = &lark.NoData{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members/:permission_member_id", "DELETE",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, nil, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.mailgroupPermissionMembers.service.conf, req)
	return result, err
}

func (mailgroupPermissionMembers *MailgroupPermissionMemberService) Delete(ctx *lark.Context, opts ...lark.APIRequestOpt) *MailgroupPermissionMemberDeleteReqCall {
	return &MailgroupPermissionMemberDeleteReqCall{
		ctx:                        ctx,
		mailgroupPermissionMembers: mailgroupPermissionMembers,
		pathParams:                 map[string]interface{}{},
		opts:                       opts,
	}
}

type PublicMailboxMemberGetReqCall struct {
	ctx                  *lark.Context
	publicMailboxMembers *PublicMailboxMemberService
	pathParams           map[string]interface{}
	queryParams          map[string]interface{}
	opts                 []lark.APIRequestOpt
}

func (rc *PublicMailboxMemberGetReqCall) SetPublicMailboxId(publicMailboxId string) {
	rc.pathParams["public_mailbox_id"] = publicMailboxId
}
func (rc *PublicMailboxMemberGetReqCall) SetMemberId(memberId string) {
	rc.pathParams["member_id"] = memberId
}
func (rc *PublicMailboxMemberGetReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *PublicMailboxMemberGetReqCall) Do() (*PublicMailboxMember, error) {
	rc.opts = append(rc.opts, lark.SetPathParams(rc.pathParams))
	rc.opts = append(rc.opts, lark.SetQueryParams(rc.queryParams))
	var result = &PublicMailboxMember{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/:member_id", "GET",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, nil, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.publicMailboxMembers.service.conf, req)
	return result, err
}

func (publicMailboxMembers *PublicMailboxMemberService) Get(ctx *lark.Context, opts ...lark.APIRequestOpt) *PublicMailboxMemberGetReqCall {
	return &PublicMailboxMemberGetReqCall{
		ctx:                  ctx,
		publicMailboxMembers: publicMailboxMembers,
		pathParams:           map[string]interface{}{},
		queryParams:          map[string]interface{}{},
		opts:                 opts,
	}
}

type PublicMailboxMemberDeleteReqCall struct {
	ctx                  *lark.Context
	publicMailboxMembers *PublicMailboxMemberService
	pathParams           map[string]interface{}
	opts                 []lark.APIRequestOpt
}

func (rc *PublicMailboxMemberDeleteReqCall) SetPublicMailboxId(publicMailboxId string) {
	rc.pathParams["public_mailbox_id"] = publicMailboxId
}
func (rc *PublicMailboxMemberDeleteReqCall) SetMemberId(memberId string) {
	rc.pathParams["member_id"] = memberId
}

func (rc *PublicMailboxMemberDeleteReqCall) Do() (*lark.NoData, error) {
	rc.opts = append(rc.opts, lark.SetPathParams(rc.pathParams))
	var result = &lark.NoData{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/:member_id", "DELETE",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, nil, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.publicMailboxMembers.service.conf, req)
	return result, err
}

func (publicMailboxMembers *PublicMailboxMemberService) Delete(ctx *lark.Context, opts ...lark.APIRequestOpt) *PublicMailboxMemberDeleteReqCall {
	return &PublicMailboxMemberDeleteReqCall{
		ctx:                  ctx,
		publicMailboxMembers: publicMailboxMembers,
		pathParams:           map[string]interface{}{},
		opts:                 opts,
	}
}

type MailgroupMemberListReqCall struct {
	ctx              *lark.Context
	mailgroupMembers *MailgroupMemberService
	pathParams       map[string]interface{}
	queryParams      map[string]interface{}
	opts             []lark.APIRequestOpt
}

func (rc *MailgroupMemberListReqCall) SetMailgroupId(mailgroupId string) {
	rc.pathParams["mailgroup_id"] = mailgroupId
}
func (rc *MailgroupMemberListReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}
func (rc *MailgroupMemberListReqCall) SetDepartmentIdType(departmentIdType string) {
	rc.queryParams["department_id_type"] = departmentIdType
}
func (rc *MailgroupMemberListReqCall) SetPageToken(pageToken string) {
	rc.queryParams["page_token"] = pageToken
}
func (rc *MailgroupMemberListReqCall) SetPageSize(pageSize int) {
	rc.queryParams["page_size"] = pageSize
}

func (rc *MailgroupMemberListReqCall) Do() (*MailgroupMemberListResult, error) {
	rc.opts = append(rc.opts, lark.SetPathParams(rc.pathParams))
	rc.opts = append(rc.opts, lark.SetQueryParams(rc.queryParams))
	var result = &MailgroupMemberListResult{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/mailgroups/:mailgroup_id/members", "GET",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, nil, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.mailgroupMembers.service.conf, req)
	return result, err
}

func (mailgroupMembers *MailgroupMemberService) List(ctx *lark.Context, opts ...lark.APIRequestOpt) *MailgroupMemberListReqCall {
	return &MailgroupMemberListReqCall{
		ctx:              ctx,
		mailgroupMembers: mailgroupMembers,
		pathParams:       map[string]interface{}{},
		queryParams:      map[string]interface{}{},
		opts:             opts,
	}
}

type MailgroupPermissionMemberGetReqCall struct {
	ctx                        *lark.Context
	mailgroupPermissionMembers *MailgroupPermissionMemberService
	pathParams                 map[string]interface{}
	queryParams                map[string]interface{}
	opts                       []lark.APIRequestOpt
}

func (rc *MailgroupPermissionMemberGetReqCall) SetMailgroupId(mailgroupId string) {
	rc.pathParams["mailgroup_id"] = mailgroupId
}
func (rc *MailgroupPermissionMemberGetReqCall) SetPermissionMemberId(permissionMemberId string) {
	rc.pathParams["permission_member_id"] = permissionMemberId
}
func (rc *MailgroupPermissionMemberGetReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}
func (rc *MailgroupPermissionMemberGetReqCall) SetDepartmentIdType(departmentIdType string) {
	rc.queryParams["department_id_type"] = departmentIdType
}

func (rc *MailgroupPermissionMemberGetReqCall) Do() (*MailgroupPermissionMember, error) {
	rc.opts = append(rc.opts, lark.SetPathParams(rc.pathParams))
	rc.opts = append(rc.opts, lark.SetQueryParams(rc.queryParams))
	var result = &MailgroupPermissionMember{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members/:permission_member_id", "GET",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, nil, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.mailgroupPermissionMembers.service.conf, req)
	return result, err
}

func (mailgroupPermissionMembers *MailgroupPermissionMemberService) Get(ctx *lark.Context, opts ...lark.APIRequestOpt) *MailgroupPermissionMemberGetReqCall {
	return &MailgroupPermissionMemberGetReqCall{
		ctx:                        ctx,
		mailgroupPermissionMembers: mailgroupPermissionMembers,
		pathParams:                 map[string]interface{}{},
		queryParams:                map[string]interface{}{},
		opts:                       opts,
	}
}

type MailgroupPermissionMemberListReqCall struct {
	ctx                        *lark.Context
	mailgroupPermissionMembers *MailgroupPermissionMemberService
	pathParams                 map[string]interface{}
	queryParams                map[string]interface{}
	opts                       []lark.APIRequestOpt
}

func (rc *MailgroupPermissionMemberListReqCall) SetMailgroupId(mailgroupId string) {
	rc.pathParams["mailgroup_id"] = mailgroupId
}
func (rc *MailgroupPermissionMemberListReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}
func (rc *MailgroupPermissionMemberListReqCall) SetDepartmentIdType(departmentIdType string) {
	rc.queryParams["department_id_type"] = departmentIdType
}
func (rc *MailgroupPermissionMemberListReqCall) SetPageToken(pageToken string) {
	rc.queryParams["page_token"] = pageToken
}
func (rc *MailgroupPermissionMemberListReqCall) SetPageSize(pageSize int) {
	rc.queryParams["page_size"] = pageSize
}

func (rc *MailgroupPermissionMemberListReqCall) Do() (*MailgroupPermissionMemberListResult, error) {
	rc.opts = append(rc.opts, lark.SetPathParams(rc.pathParams))
	rc.opts = append(rc.opts, lark.SetQueryParams(rc.queryParams))
	var result = &MailgroupPermissionMemberListResult{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members", "GET",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, nil, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.mailgroupPermissionMembers.service.conf, req)
	return result, err
}

func (mailgroupPermissionMembers *MailgroupPermissionMemberService) List(ctx *lark.Context, opts ...lark.APIRequestOpt) *MailgroupPermissionMemberListReqCall {
	return &MailgroupPermissionMemberListReqCall{
		ctx:                        ctx,
		mailgroupPermissionMembers: mailgroupPermissionMembers,
		pathParams:                 map[string]interface{}{},
		queryParams:                map[string]interface{}{},
		opts:                       opts,
	}
}

type PublicMailboxUpdateReqCall struct {
	ctx            *lark.Context
	publicMailboxs *PublicMailboxService
	body           *PublicMailbox
	pathParams     map[string]interface{}
	opts           []lark.APIRequestOpt
}

func (rc *PublicMailboxUpdateReqCall) SetPublicMailboxId(publicMailboxId string) {
	rc.pathParams["public_mailbox_id"] = publicMailboxId
}

func (rc *PublicMailboxUpdateReqCall) Do() (*PublicMailbox, error) {
	rc.opts = append(rc.opts, lark.SetPathParams(rc.pathParams))
	var result = &PublicMailbox{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/public_mailboxes/:public_mailbox_id", "PUT",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, rc.body, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.publicMailboxs.service.conf, req)
	return result, err
}

func (publicMailboxs *PublicMailboxService) Update(ctx *lark.Context, body *PublicMailbox, opts ...lark.APIRequestOpt) *PublicMailboxUpdateReqCall {
	return &PublicMailboxUpdateReqCall{
		ctx:            ctx,
		publicMailboxs: publicMailboxs,
		body:           body,
		pathParams:     map[string]interface{}{},
		opts:           opts,
	}
}

type PublicMailboxMemberClearReqCall struct {
	ctx                  *lark.Context
	publicMailboxMembers *PublicMailboxMemberService
	pathParams           map[string]interface{}
	opts                 []lark.APIRequestOpt
}

func (rc *PublicMailboxMemberClearReqCall) SetPublicMailboxId(publicMailboxId string) {
	rc.pathParams["public_mailbox_id"] = publicMailboxId
}

func (rc *PublicMailboxMemberClearReqCall) Do() (*lark.NoData, error) {
	rc.opts = append(rc.opts, lark.SetPathParams(rc.pathParams))
	var result = &lark.NoData{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/clear", "POST",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, nil, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.publicMailboxMembers.service.conf, req)
	return result, err
}

func (publicMailboxMembers *PublicMailboxMemberService) Clear(ctx *lark.Context, opts ...lark.APIRequestOpt) *PublicMailboxMemberClearReqCall {
	return &PublicMailboxMemberClearReqCall{
		ctx:                  ctx,
		publicMailboxMembers: publicMailboxMembers,
		pathParams:           map[string]interface{}{},
		opts:                 opts,
	}
}

type PublicMailboxMemberListReqCall struct {
	ctx                  *lark.Context
	publicMailboxMembers *PublicMailboxMemberService
	pathParams           map[string]interface{}
	queryParams          map[string]interface{}
	opts                 []lark.APIRequestOpt
}

func (rc *PublicMailboxMemberListReqCall) SetPublicMailboxId(publicMailboxId string) {
	rc.pathParams["public_mailbox_id"] = publicMailboxId
}
func (rc *PublicMailboxMemberListReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}
func (rc *PublicMailboxMemberListReqCall) SetPageToken(pageToken string) {
	rc.queryParams["page_token"] = pageToken
}
func (rc *PublicMailboxMemberListReqCall) SetPageSize(pageSize int) {
	rc.queryParams["page_size"] = pageSize
}

func (rc *PublicMailboxMemberListReqCall) Do() (*PublicMailboxMemberListResult, error) {
	rc.opts = append(rc.opts, lark.SetPathParams(rc.pathParams))
	rc.opts = append(rc.opts, lark.SetQueryParams(rc.queryParams))
	var result = &PublicMailboxMemberListResult{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members", "GET",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, nil, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.publicMailboxMembers.service.conf, req)
	return result, err
}

func (publicMailboxMembers *PublicMailboxMemberService) List(ctx *lark.Context, opts ...lark.APIRequestOpt) *PublicMailboxMemberListReqCall {
	return &PublicMailboxMemberListReqCall{
		ctx:                  ctx,
		publicMailboxMembers: publicMailboxMembers,
		pathParams:           map[string]interface{}{},
		queryParams:          map[string]interface{}{},
		opts:                 opts,
	}
}

type PublicMailboxMemberCreateReqCall struct {
	ctx                  *lark.Context
	publicMailboxMembers *PublicMailboxMemberService
	body                 *PublicMailboxMember
	pathParams           map[string]interface{}
	queryParams          map[string]interface{}
	opts                 []lark.APIRequestOpt
}

func (rc *PublicMailboxMemberCreateReqCall) SetPublicMailboxId(publicMailboxId string) {
	rc.pathParams["public_mailbox_id"] = publicMailboxId
}
func (rc *PublicMailboxMemberCreateReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *PublicMailboxMemberCreateReqCall) Do() (*PublicMailboxMember, error) {
	rc.opts = append(rc.opts, lark.SetPathParams(rc.pathParams))
	rc.opts = append(rc.opts, lark.SetQueryParams(rc.queryParams))
	var result = &PublicMailboxMember{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members", "POST",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, rc.body, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.publicMailboxMembers.service.conf, req)
	return result, err
}

func (publicMailboxMembers *PublicMailboxMemberService) Create(ctx *lark.Context, body *PublicMailboxMember, opts ...lark.APIRequestOpt) *PublicMailboxMemberCreateReqCall {
	return &PublicMailboxMemberCreateReqCall{
		ctx:                  ctx,
		publicMailboxMembers: publicMailboxMembers,
		body:                 body,
		pathParams:           map[string]interface{}{},
		queryParams:          map[string]interface{}{},
		opts:                 opts,
	}
}

type MailgroupGetReqCall struct {
	ctx        *lark.Context
	mailgroups *MailgroupService
	pathParams map[string]interface{}
	opts       []lark.APIRequestOpt
}

func (rc *MailgroupGetReqCall) SetMailgroupId(mailgroupId string) {
	rc.pathParams["mailgroup_id"] = mailgroupId
}

func (rc *MailgroupGetReqCall) Do() (*Mailgroup, error) {
	rc.opts = append(rc.opts, lark.SetPathParams(rc.pathParams))
	var result = &Mailgroup{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/mailgroups/:mailgroup_id", "GET",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, nil, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.mailgroups.service.conf, req)
	return result, err
}

func (mailgroups *MailgroupService) Get(ctx *lark.Context, opts ...lark.APIRequestOpt) *MailgroupGetReqCall {
	return &MailgroupGetReqCall{
		ctx:        ctx,
		mailgroups: mailgroups,
		pathParams: map[string]interface{}{},
		opts:       opts,
	}
}

type MailgroupUpdateReqCall struct {
	ctx        *lark.Context
	mailgroups *MailgroupService
	body       *Mailgroup
	pathParams map[string]interface{}
	opts       []lark.APIRequestOpt
}

func (rc *MailgroupUpdateReqCall) SetMailgroupId(mailgroupId string) {
	rc.pathParams["mailgroup_id"] = mailgroupId
}

func (rc *MailgroupUpdateReqCall) Do() (*Mailgroup, error) {
	rc.opts = append(rc.opts, lark.SetPathParams(rc.pathParams))
	var result = &Mailgroup{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/mailgroups/:mailgroup_id", "PUT",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, rc.body, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.mailgroups.service.conf, req)
	return result, err
}

func (mailgroups *MailgroupService) Update(ctx *lark.Context, body *Mailgroup, opts ...lark.APIRequestOpt) *MailgroupUpdateReqCall {
	return &MailgroupUpdateReqCall{
		ctx:        ctx,
		mailgroups: mailgroups,
		body:       body,
		pathParams: map[string]interface{}{},
		opts:       opts,
	}
}

type PublicMailboxListReqCall struct {
	ctx            *lark.Context
	publicMailboxs *PublicMailboxService
	queryParams    map[string]interface{}
	opts           []lark.APIRequestOpt
}

func (rc *PublicMailboxListReqCall) SetPageToken(pageToken string) {
	rc.queryParams["page_token"] = pageToken
}
func (rc *PublicMailboxListReqCall) SetPageSize(pageSize int) {
	rc.queryParams["page_size"] = pageSize
}

func (rc *PublicMailboxListReqCall) Do() (*PublicMailboxListResult, error) {
	rc.opts = append(rc.opts, lark.SetQueryParams(rc.queryParams))
	var result = &PublicMailboxListResult{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/public_mailboxes", "GET",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, nil, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.publicMailboxs.service.conf, req)
	return result, err
}

func (publicMailboxs *PublicMailboxService) List(ctx *lark.Context, opts ...lark.APIRequestOpt) *PublicMailboxListReqCall {
	return &PublicMailboxListReqCall{
		ctx:            ctx,
		publicMailboxs: publicMailboxs,
		queryParams:    map[string]interface{}{},
		opts:           opts,
	}
}

type MailgroupDeleteReqCall struct {
	ctx        *lark.Context
	mailgroups *MailgroupService
	pathParams map[string]interface{}
	opts       []lark.APIRequestOpt
}

func (rc *MailgroupDeleteReqCall) SetMailgroupId(mailgroupId string) {
	rc.pathParams["mailgroup_id"] = mailgroupId
}

func (rc *MailgroupDeleteReqCall) Do() (*lark.NoData, error) {
	rc.opts = append(rc.opts, lark.SetPathParams(rc.pathParams))
	var result = &lark.NoData{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/mailgroups/:mailgroup_id", "DELETE",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, nil, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.mailgroups.service.conf, req)
	return result, err
}

func (mailgroups *MailgroupService) Delete(ctx *lark.Context, opts ...lark.APIRequestOpt) *MailgroupDeleteReqCall {
	return &MailgroupDeleteReqCall{
		ctx:        ctx,
		mailgroups: mailgroups,
		pathParams: map[string]interface{}{},
		opts:       opts,
	}
}

type MailgroupPermissionMemberCreateReqCall struct {
	ctx                        *lark.Context
	mailgroupPermissionMembers *MailgroupPermissionMemberService
	body                       *MailgroupPermissionMember
	pathParams                 map[string]interface{}
	queryParams                map[string]interface{}
	opts                       []lark.APIRequestOpt
}

func (rc *MailgroupPermissionMemberCreateReqCall) SetMailgroupId(mailgroupId string) {
	rc.pathParams["mailgroup_id"] = mailgroupId
}
func (rc *MailgroupPermissionMemberCreateReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}
func (rc *MailgroupPermissionMemberCreateReqCall) SetDepartmentIdType(departmentIdType string) {
	rc.queryParams["department_id_type"] = departmentIdType
}

func (rc *MailgroupPermissionMemberCreateReqCall) Do() (*MailgroupPermissionMember, error) {
	rc.opts = append(rc.opts, lark.SetPathParams(rc.pathParams))
	rc.opts = append(rc.opts, lark.SetQueryParams(rc.queryParams))
	var result = &MailgroupPermissionMember{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members", "POST",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, rc.body, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.mailgroupPermissionMembers.service.conf, req)
	return result, err
}

func (mailgroupPermissionMembers *MailgroupPermissionMemberService) Create(ctx *lark.Context, body *MailgroupPermissionMember, opts ...lark.APIRequestOpt) *MailgroupPermissionMemberCreateReqCall {
	return &MailgroupPermissionMemberCreateReqCall{
		ctx:                        ctx,
		mailgroupPermissionMembers: mailgroupPermissionMembers,
		body:                       body,
		pathParams:                 map[string]interface{}{},
		queryParams:                map[string]interface{}{},
		opts:                       opts,
	}
}

type PublicMailboxCreateReqCall struct {
	ctx            *lark.Context
	publicMailboxs *PublicMailboxService
	body           *PublicMailbox
	opts           []lark.APIRequestOpt
}

func (rc *PublicMailboxCreateReqCall) Do() (*PublicMailbox, error) {
	var result = &PublicMailbox{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/public_mailboxes", "POST",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, rc.body, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.publicMailboxs.service.conf, req)
	return result, err
}

func (publicMailboxs *PublicMailboxService) Create(ctx *lark.Context, body *PublicMailbox, opts ...lark.APIRequestOpt) *PublicMailboxCreateReqCall {
	return &PublicMailboxCreateReqCall{
		ctx:            ctx,
		publicMailboxs: publicMailboxs,
		body:           body,
		opts:           opts,
	}
}

type MailgroupMemberGetReqCall struct {
	ctx              *lark.Context
	mailgroupMembers *MailgroupMemberService
	pathParams       map[string]interface{}
	queryParams      map[string]interface{}
	opts             []lark.APIRequestOpt
}

func (rc *MailgroupMemberGetReqCall) SetMailgroupId(mailgroupId string) {
	rc.pathParams["mailgroup_id"] = mailgroupId
}
func (rc *MailgroupMemberGetReqCall) SetMemberId(memberId string) {
	rc.pathParams["member_id"] = memberId
}
func (rc *MailgroupMemberGetReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}
func (rc *MailgroupMemberGetReqCall) SetDepartmentIdType(departmentIdType string) {
	rc.queryParams["department_id_type"] = departmentIdType
}

func (rc *MailgroupMemberGetReqCall) Do() (*MailgroupMember, error) {
	rc.opts = append(rc.opts, lark.SetPathParams(rc.pathParams))
	rc.opts = append(rc.opts, lark.SetQueryParams(rc.queryParams))
	var result = &MailgroupMember{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/mailgroups/:mailgroup_id/members/:member_id", "GET",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, nil, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.mailgroupMembers.service.conf, req)
	return result, err
}

func (mailgroupMembers *MailgroupMemberService) Get(ctx *lark.Context, opts ...lark.APIRequestOpt) *MailgroupMemberGetReqCall {
	return &MailgroupMemberGetReqCall{
		ctx:              ctx,
		mailgroupMembers: mailgroupMembers,
		pathParams:       map[string]interface{}{},
		queryParams:      map[string]interface{}{},
		opts:             opts,
	}
}

type PublicMailboxGetReqCall struct {
	ctx            *lark.Context
	publicMailboxs *PublicMailboxService
	pathParams     map[string]interface{}
	opts           []lark.APIRequestOpt
}

func (rc *PublicMailboxGetReqCall) SetPublicMailboxId(publicMailboxId string) {
	rc.pathParams["public_mailbox_id"] = publicMailboxId
}

func (rc *PublicMailboxGetReqCall) Do() (*PublicMailbox, error) {
	rc.opts = append(rc.opts, lark.SetPathParams(rc.pathParams))
	var result = &PublicMailbox{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/public_mailboxes/:public_mailbox_id", "GET",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, nil, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.publicMailboxs.service.conf, req)
	return result, err
}

func (publicMailboxs *PublicMailboxService) Get(ctx *lark.Context, opts ...lark.APIRequestOpt) *PublicMailboxGetReqCall {
	return &PublicMailboxGetReqCall{
		ctx:            ctx,
		publicMailboxs: publicMailboxs,
		pathParams:     map[string]interface{}{},
		opts:           opts,
	}
}

type PublicMailboxPatchReqCall struct {
	ctx            *lark.Context
	publicMailboxs *PublicMailboxService
	body           *PublicMailbox
	pathParams     map[string]interface{}
	opts           []lark.APIRequestOpt
}

func (rc *PublicMailboxPatchReqCall) SetPublicMailboxId(publicMailboxId string) {
	rc.pathParams["public_mailbox_id"] = publicMailboxId
}

func (rc *PublicMailboxPatchReqCall) Do() (*PublicMailbox, error) {
	rc.opts = append(rc.opts, lark.SetPathParams(rc.pathParams))
	var result = &PublicMailbox{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/public_mailboxes/:public_mailbox_id", "PATCH",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, rc.body, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.publicMailboxs.service.conf, req)
	return result, err
}

func (publicMailboxs *PublicMailboxService) Patch(ctx *lark.Context, body *PublicMailbox, opts ...lark.APIRequestOpt) *PublicMailboxPatchReqCall {
	return &PublicMailboxPatchReqCall{
		ctx:            ctx,
		publicMailboxs: publicMailboxs,
		body:           body,
		pathParams:     map[string]interface{}{},
		opts:           opts,
	}
}

type MailgroupCreateReqCall struct {
	ctx        *lark.Context
	mailgroups *MailgroupService
	body       *Mailgroup
	opts       []lark.APIRequestOpt
}

func (rc *MailgroupCreateReqCall) Do() (*Mailgroup, error) {
	var result = &Mailgroup{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/mailgroups", "POST",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, rc.body, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.mailgroups.service.conf, req)
	return result, err
}

func (mailgroups *MailgroupService) Create(ctx *lark.Context, body *Mailgroup, opts ...lark.APIRequestOpt) *MailgroupCreateReqCall {
	return &MailgroupCreateReqCall{
		ctx:        ctx,
		mailgroups: mailgroups,
		body:       body,
		opts:       opts,
	}
}

type MailgroupListReqCall struct {
	ctx         *lark.Context
	mailgroups  *MailgroupService
	queryParams map[string]interface{}
	opts        []lark.APIRequestOpt
}

func (rc *MailgroupListReqCall) SetPageToken(pageToken string) {
	rc.queryParams["page_token"] = pageToken
}
func (rc *MailgroupListReqCall) SetPageSize(pageSize int) {
	rc.queryParams["page_size"] = pageSize
}

func (rc *MailgroupListReqCall) Do() (*MailgroupListResult, error) {
	rc.opts = append(rc.opts, lark.SetQueryParams(rc.queryParams))
	var result = &MailgroupListResult{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/mailgroups", "GET",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, nil, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.mailgroups.service.conf, req)
	return result, err
}

func (mailgroups *MailgroupService) List(ctx *lark.Context, opts ...lark.APIRequestOpt) *MailgroupListReqCall {
	return &MailgroupListReqCall{
		ctx:         ctx,
		mailgroups:  mailgroups,
		queryParams: map[string]interface{}{},
		opts:        opts,
	}
}

type MailgroupPatchReqCall struct {
	ctx        *lark.Context
	mailgroups *MailgroupService
	body       *Mailgroup
	pathParams map[string]interface{}
	opts       []lark.APIRequestOpt
}

func (rc *MailgroupPatchReqCall) SetMailgroupId(mailgroupId string) {
	rc.pathParams["mailgroup_id"] = mailgroupId
}

func (rc *MailgroupPatchReqCall) Do() (*Mailgroup, error) {
	rc.opts = append(rc.opts, lark.SetPathParams(rc.pathParams))
	var result = &Mailgroup{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/mailgroups/:mailgroup_id", "PATCH",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, rc.body, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.mailgroups.service.conf, req)
	return result, err
}

func (mailgroups *MailgroupService) Patch(ctx *lark.Context, body *Mailgroup, opts ...lark.APIRequestOpt) *MailgroupPatchReqCall {
	return &MailgroupPatchReqCall{
		ctx:        ctx,
		mailgroups: mailgroups,
		body:       body,
		pathParams: map[string]interface{}{},
		opts:       opts,
	}
}

type MailgroupMemberCreateReqCall struct {
	ctx              *lark.Context
	mailgroupMembers *MailgroupMemberService
	body             *MailgroupMember
	pathParams       map[string]interface{}
	queryParams      map[string]interface{}
	opts             []lark.APIRequestOpt
}

func (rc *MailgroupMemberCreateReqCall) SetMailgroupId(mailgroupId string) {
	rc.pathParams["mailgroup_id"] = mailgroupId
}
func (rc *MailgroupMemberCreateReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}
func (rc *MailgroupMemberCreateReqCall) SetDepartmentIdType(departmentIdType string) {
	rc.queryParams["department_id_type"] = departmentIdType
}

func (rc *MailgroupMemberCreateReqCall) Do() (*MailgroupMember, error) {
	rc.opts = append(rc.opts, lark.SetPathParams(rc.pathParams))
	rc.opts = append(rc.opts, lark.SetQueryParams(rc.queryParams))
	var result = &MailgroupMember{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/mailgroups/:mailgroup_id/members", "POST",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, rc.body, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.mailgroupMembers.service.conf, req)
	return result, err
}

func (mailgroupMembers *MailgroupMemberService) Create(ctx *lark.Context, body *MailgroupMember, opts ...lark.APIRequestOpt) *MailgroupMemberCreateReqCall {
	return &MailgroupMemberCreateReqCall{
		ctx:              ctx,
		mailgroupMembers: mailgroupMembers,
		body:             body,
		pathParams:       map[string]interface{}{},
		queryParams:      map[string]interface{}{},
		opts:             opts,
	}
}

type MailgroupMemberDeleteReqCall struct {
	ctx              *lark.Context
	mailgroupMembers *MailgroupMemberService
	pathParams       map[string]interface{}
	opts             []lark.APIRequestOpt
}

func (rc *MailgroupMemberDeleteReqCall) SetMailgroupId(mailgroupId string) {
	rc.pathParams["mailgroup_id"] = mailgroupId
}
func (rc *MailgroupMemberDeleteReqCall) SetMemberId(memberId string) {
	rc.pathParams["member_id"] = memberId
}

func (rc *MailgroupMemberDeleteReqCall) Do() (*lark.NoData, error) {
	rc.opts = append(rc.opts, lark.SetPathParams(rc.pathParams))
	var result = &lark.NoData{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/mail/v1/mailgroups/:mailgroup_id/members/:member_id", "DELETE",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, nil, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.mailgroupMembers.service.conf, req)
	return result, err
}

func (mailgroupMembers *MailgroupMemberService) Delete(ctx *lark.Context, opts ...lark.APIRequestOpt) *MailgroupMemberDeleteReqCall {
	return &MailgroupMemberDeleteReqCall{
		ctx:              ctx,
		mailgroupMembers: mailgroupMembers,
		pathParams:       map[string]interface{}{},
		opts:             opts,
	}
}
