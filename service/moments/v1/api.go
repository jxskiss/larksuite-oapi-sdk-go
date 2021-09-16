// Code generated by lark suite oapi sdk gen
package v1

import (
	"github.com/larksuite/oapi-sdk-go/api"
	"github.com/larksuite/oapi-sdk-go/api/core/request"
	"github.com/larksuite/oapi-sdk-go/api/core/response"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/core/config"
)

type Service struct {
    conf           *config.Config
	Categorys   *CategoryService
	Comments   *CommentService
	Posts   *PostService
	Reactions   *ReactionService
	
}

func NewService(conf *config.Config) *Service {
    s := &Service{
        conf:     conf,
    }
    s.Categorys = newCategoryService(s)
    s.Comments = newCommentService(s)
    s.Posts = newPostService(s)
    s.Reactions = newReactionService(s)
    return s
}

type CategoryService struct {
    service     *Service
}

func newCategoryService(service *Service) *CategoryService {
    return &CategoryService{
        service: service,
    }
}
type CommentService struct {
    service     *Service
}

func newCommentService(service *Service) *CommentService {
    return &CommentService{
        service: service,
    }
}
type PostService struct {
    service     *Service
}

func newPostService(service *Service) *PostService {
    return &PostService{
        service: service,
    }
}
type ReactionService struct {
    service     *Service
}

func newReactionService(service *Service) *ReactionService {
    return &ReactionService{
        service: service,
    }
}



type CommentCreateReqCall struct {
	ctx    *core.Context
	comments *CommentService
	body *Comment
	queryParams map[string]interface{}
	optFns   []request.OptFn
}
func (rc *CommentCreateReqCall) SetUserIdType(userIdType string){
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *CommentCreateReqCall) Do() (*CommentCreateResult, error) {
    rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
    var result = &CommentCreateResult{}
	req := request.NewRequest("moments/v1/comments", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeTenant, request.AccessTokenTypeUser}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.comments.service.conf, req)
	return result, err
}

func (comments *CommentService) Create(ctx *core.Context, body  *Comment, optFns ...request.OptFn) *CommentCreateReqCall {
	return &CommentCreateReqCall{
		ctx:    ctx,
		comments: comments,
		body:        body,
        queryParams:        map[string]interface{}{},
		optFns:      optFns,
	}
}


type PostCreateReqCall struct {
	ctx    *core.Context
	posts *PostService
	body *Post
	queryParams map[string]interface{}
	optFns   []request.OptFn
}
func (rc *PostCreateReqCall) SetNeedPublish(needPublish bool){
	rc.queryParams["need_publish"] = needPublish
}
func (rc *PostCreateReqCall) SetUserIdType(userIdType string){
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *PostCreateReqCall) Do() (*PostCreateResult, error) {
    rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
    var result = &PostCreateResult{}
	req := request.NewRequest("moments/v1/posts", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeTenant, request.AccessTokenTypeUser}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.posts.service.conf, req)
	return result, err
}

func (posts *PostService) Create(ctx *core.Context, body  *Post, optFns ...request.OptFn) *PostCreateReqCall {
	return &PostCreateReqCall{
		ctx:    ctx,
		posts: posts,
		body:        body,
        queryParams:        map[string]interface{}{},
		optFns:      optFns,
	}
}


type ReactionCreateReqCall struct {
	ctx    *core.Context
	reactions *ReactionService
	body *Reaction
	queryParams map[string]interface{}
	optFns   []request.OptFn
}
func (rc *ReactionCreateReqCall) SetUserIdType(userIdType string){
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *ReactionCreateReqCall) Do() (*ReactionCreateResult, error) {
    rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
    var result = &ReactionCreateResult{}
	req := request.NewRequest("moments/v1/reactions", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeTenant, request.AccessTokenTypeUser}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.reactions.service.conf, req)
	return result, err
}

func (reactions *ReactionService) Create(ctx *core.Context, body  *Reaction, optFns ...request.OptFn) *ReactionCreateReqCall {
	return &ReactionCreateReqCall{
		ctx:    ctx,
		reactions: reactions,
		body:        body,
        queryParams:        map[string]interface{}{},
		optFns:      optFns,
	}
}


type CommentDeleteReqCall struct {
	ctx    *core.Context
	comments *CommentService
	pathParams  map[string]interface{}
	optFns   []request.OptFn
}
func (rc *CommentDeleteReqCall) SetCommentId(commentId string){
	rc.pathParams["comment_id"] = commentId
}

func (rc *CommentDeleteReqCall) Do() (*response.NoData, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
    var result = &response.NoData{}
	req := request.NewRequest("moments/v1/comments/:comment_id", "DELETE",
		[]request.AccessTokenType{request.AccessTokenTypeTenant, request.AccessTokenTypeUser}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.comments.service.conf, req)
	return result, err
}

func (comments *CommentService) Delete(ctx *core.Context,  optFns ...request.OptFn) *CommentDeleteReqCall {
	return &CommentDeleteReqCall{
		ctx:    ctx,
		comments: comments,
        pathParams:        map[string]interface{}{},
		optFns:      optFns,
	}
}


type PostDeleteReqCall struct {
	ctx    *core.Context
	posts *PostService
	pathParams  map[string]interface{}
	optFns   []request.OptFn
}
func (rc *PostDeleteReqCall) SetPostId(postId string){
	rc.pathParams["post_id"] = postId
}

func (rc *PostDeleteReqCall) Do() (*response.NoData, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
    var result = &response.NoData{}
	req := request.NewRequest("moments/v1/posts/:post_id", "DELETE",
		[]request.AccessTokenType{request.AccessTokenTypeTenant, request.AccessTokenTypeUser}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.posts.service.conf, req)
	return result, err
}

func (posts *PostService) Delete(ctx *core.Context,  optFns ...request.OptFn) *PostDeleteReqCall {
	return &PostDeleteReqCall{
		ctx:    ctx,
		posts: posts,
        pathParams:        map[string]interface{}{},
		optFns:      optFns,
	}
}


type ReactionDeleteReqCall struct {
	ctx    *core.Context
	reactions *ReactionService
	pathParams  map[string]interface{}
	optFns   []request.OptFn
}
func (rc *ReactionDeleteReqCall) SetReactionId(reactionId string){
	rc.pathParams["reaction_id"] = reactionId
}

func (rc *ReactionDeleteReqCall) Do() (*response.NoData, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
    var result = &response.NoData{}
	req := request.NewRequest("moments/v1/reactions/:reaction_id", "DELETE",
		[]request.AccessTokenType{request.AccessTokenTypeTenant, request.AccessTokenTypeUser}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.reactions.service.conf, req)
	return result, err
}

func (reactions *ReactionService) Delete(ctx *core.Context,  optFns ...request.OptFn) *ReactionDeleteReqCall {
	return &ReactionDeleteReqCall{
		ctx:    ctx,
		reactions: reactions,
        pathParams:        map[string]interface{}{},
		optFns:      optFns,
	}
}


type CommentGetReqCall struct {
	ctx    *core.Context
	comments *CommentService
	pathParams  map[string]interface{}
	queryParams map[string]interface{}
	optFns   []request.OptFn
}
func (rc *CommentGetReqCall) SetCommentId(commentId string){
	rc.pathParams["comment_id"] = commentId
}
func (rc *CommentGetReqCall) SetUserIdType(userIdType string){
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *CommentGetReqCall) Do() (*CommentGetResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
    rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
    var result = &CommentGetResult{}
	req := request.NewRequest("moments/v1/comments/:comment_id", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant, request.AccessTokenTypeUser}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.comments.service.conf, req)
	return result, err
}

func (comments *CommentService) Get(ctx *core.Context,  optFns ...request.OptFn) *CommentGetReqCall {
	return &CommentGetReqCall{
		ctx:    ctx,
		comments: comments,
        pathParams:        map[string]interface{}{},
        queryParams:        map[string]interface{}{},
		optFns:      optFns,
	}
}


type PostGetReqCall struct {
	ctx    *core.Context
	posts *PostService
	pathParams  map[string]interface{}
	queryParams map[string]interface{}
	optFns   []request.OptFn
}
func (rc *PostGetReqCall) SetPostId(postId string){
	rc.pathParams["post_id"] = postId
}
func (rc *PostGetReqCall) SetUserIdType(userIdType string){
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *PostGetReqCall) Do() (*PostGetResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
    rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
    var result = &PostGetResult{}
	req := request.NewRequest("moments/v1/posts/:post_id", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant, request.AccessTokenTypeUser}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.posts.service.conf, req)
	return result, err
}

func (posts *PostService) Get(ctx *core.Context,  optFns ...request.OptFn) *PostGetReqCall {
	return &PostGetReqCall{
		ctx:    ctx,
		posts: posts,
        pathParams:        map[string]interface{}{},
        queryParams:        map[string]interface{}{},
		optFns:      optFns,
	}
}


type CategoryListReqCall struct {
	ctx    *core.Context
	categorys *CategoryService
	queryParams map[string]interface{}
	optFns   []request.OptFn
}
func (rc *CategoryListReqCall) SetPageSize(pageSize int){
	rc.queryParams["page_size"] = pageSize
}
func (rc *CategoryListReqCall) SetPageToken(pageToken string){
	rc.queryParams["page_token"] = pageToken
}

func (rc *CategoryListReqCall) Do() (*CategoryListResult, error) {
    rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
    var result = &CategoryListResult{}
	req := request.NewRequest("moments/v1/categories", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant, request.AccessTokenTypeUser}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.categorys.service.conf, req)
	return result, err
}

func (categorys *CategoryService) List(ctx *core.Context,  optFns ...request.OptFn) *CategoryListReqCall {
	return &CategoryListReqCall{
		ctx:    ctx,
		categorys: categorys,
        queryParams:        map[string]interface{}{},
		optFns:      optFns,
	}
}


type CommentListReqCall struct {
	ctx    *core.Context
	comments *CommentService
	queryParams map[string]interface{}
	optFns   []request.OptFn
}
func (rc *CommentListReqCall) SetPageSize(pageSize int){
	rc.queryParams["page_size"] = pageSize
}
func (rc *CommentListReqCall) SetPageToken(pageToken string){
	rc.queryParams["page_token"] = pageToken
}
func (rc *CommentListReqCall) SetPostId(postId int64){
	rc.queryParams["post_id"] = postId
}
func (rc *CommentListReqCall) SetUserIdType(userIdType string){
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *CommentListReqCall) Do() (*CommentListResult, error) {
    rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
    var result = &CommentListResult{}
	req := request.NewRequest("moments/v1/comments", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant, request.AccessTokenTypeUser}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.comments.service.conf, req)
	return result, err
}

func (comments *CommentService) List(ctx *core.Context,  optFns ...request.OptFn) *CommentListReqCall {
	return &CommentListReqCall{
		ctx:    ctx,
		comments: comments,
        queryParams:        map[string]interface{}{},
		optFns:      optFns,
	}
}


type PostListReqCall struct {
	ctx    *core.Context
	posts *PostService
	queryParams map[string]interface{}
	optFns   []request.OptFn
}
func (rc *PostListReqCall) SetPageSize(pageSize int){
	rc.queryParams["page_size"] = pageSize
}
func (rc *PostListReqCall) SetPageToken(pageToken string){
	rc.queryParams["page_token"] = pageToken
}
func (rc *PostListReqCall) SetUserId(userId string){
	rc.queryParams["user_id"] = userId
}
func (rc *PostListReqCall) SetCategoryId(categoryId int64){
	rc.queryParams["category_id"] = categoryId
}
func (rc *PostListReqCall) SetUserIdType(userIdType string){
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *PostListReqCall) Do() (*PostListResult, error) {
    rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
    var result = &PostListResult{}
	req := request.NewRequest("moments/v1/posts", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant, request.AccessTokenTypeUser}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.posts.service.conf, req)
	return result, err
}

func (posts *PostService) List(ctx *core.Context,  optFns ...request.OptFn) *PostListReqCall {
	return &PostListReqCall{
		ctx:    ctx,
		posts: posts,
        queryParams:        map[string]interface{}{},
		optFns:      optFns,
	}
}


type ReactionListReqCall struct {
	ctx    *core.Context
	reactions *ReactionService
	queryParams map[string]interface{}
	optFns   []request.OptFn
}
func (rc *ReactionListReqCall) SetPageSize(pageSize int){
	rc.queryParams["page_size"] = pageSize
}
func (rc *ReactionListReqCall) SetPageToken(pageToken string){
	rc.queryParams["page_token"] = pageToken
}
func (rc *ReactionListReqCall) SetEntityId(entityId string){
	rc.queryParams["entity_id"] = entityId
}
func (rc *ReactionListReqCall) SetType(type1 string){
	rc.queryParams["type"] = type1
}
func (rc *ReactionListReqCall) SetUserIdType(userIdType string){
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *ReactionListReqCall) Do() (*ReactionListResult, error) {
    rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
    var result = &ReactionListResult{}
	req := request.NewRequest("moments/v1/reactions", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant, request.AccessTokenTypeUser}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.reactions.service.conf, req)
	return result, err
}

func (reactions *ReactionService) List(ctx *core.Context,  optFns ...request.OptFn) *ReactionListReqCall {
	return &ReactionListReqCall{
		ctx:    ctx,
		reactions: reactions,
        queryParams:        map[string]interface{}{},
		optFns:      optFns,
	}
}
