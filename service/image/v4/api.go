// Code generated by lark suite oapi sdk gen
package v4

import (
	"github.com/larksuite/oapi-sdk-go"
	"io"
)

type Service struct {
	conf   lark.Config
	Images *ImageService
}

func NewService(conf lark.Config) *Service {
	s := &Service{
		conf: conf,
	}
	s.Images = newImageService(s)
	return s
}

type ImageService struct {
	service *Service
}

func newImageService(service *Service) *ImageService {
	return &ImageService{
		service: service,
	}
}

type ImageGetReqCall struct {
	ctx         *lark.Context
	images      *ImageService
	queryParams map[string]interface{}
	opts        []lark.APIRequestOpt
	result      io.Writer
}

func (rc *ImageGetReqCall) SetImageKey(imageKey string) {
	rc.queryParams["image_key"] = imageKey
}
func (rc *ImageGetReqCall) SetResponseStream(result io.Writer) {
	rc.result = result
}

func (rc *ImageGetReqCall) Do() (io.Writer, error) {
	rc.opts = append(rc.opts, lark.SetQueryParams(rc.queryParams))
	rc.opts = append(rc.opts, lark.SetResponseStream())
	req := lark.NewAPIRequestWithMultiToken("/open-apis/image/v4/get", "GET",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, nil, rc.result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.images.service.conf, req)
	return rc.result, err
}

func (images *ImageService) Get(ctx *lark.Context, opts ...lark.APIRequestOpt) *ImageGetReqCall {
	return &ImageGetReqCall{
		ctx:         ctx,
		images:      images,
		queryParams: map[string]interface{}{},
		opts:        opts,
	}
}

type ImagePutReqCall struct {
	ctx    *lark.Context
	images *ImageService
	body   *lark.FormData
	opts   []lark.APIRequestOpt
}

func (rc *ImagePutReqCall) SetImage(image *lark.FormDataFile) {
	rc.body.AddFile("image", image)
}
func (rc *ImagePutReqCall) SetImageType(imageType string) {
	rc.body.AddParam("image_type", imageType)
}

func (rc *ImagePutReqCall) Do() (*Image, error) {
	var result = &Image{}
	req := lark.NewAPIRequestWithMultiToken("/open-apis/image/v4/put", "POST",
		[]lark.AccessTokenType{lark.AccessTokenTypeTenant}, rc.body, result, rc.opts...)
	err := lark.SendAPIRequest(rc.ctx, rc.images.service.conf, req)
	return result, err
}

func (images *ImageService) Put(ctx *lark.Context, opts ...lark.APIRequestOpt) *ImagePutReqCall {
	return &ImagePutReqCall{
		ctx:    ctx,
		images: images,
		body:   lark.NewFormData(),
		opts:   opts,
	}
}
