// Package im code generated by lark suite oapi sdk gen
package im

import (
	"bytes"
	"context"
	"github.com/larksuite/oapi-sdk-go/v2"
	"net/http"
)

type Im struct {
	Messages *messages
	Files    *files
}

func New(app *lark.App) *Im {
	s := &Im{}
	s.Messages = &messages{app: app}
	s.Files = &files{app: app}
	return s
}

type messages struct {
	app *lark.App
}

type files struct {
	app *lark.App
}

func (s *messages) Create(ctx context.Context, req *MessageCreateReq,
	options ...lark.RequestOptionFunc) (*MessageCreateResp, error) {
	rawResp, err := s.app.SendRequestWithAccessTokenTypes(ctx, "POST",
		"/open-apis/im/v1/messages", req, []lark.AccessTokenType{lark.AccessTokenTypeTenant}, options...)
	if err != nil {
		return nil, err
	}

	messageCreateResp := &MessageCreateResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(messageCreateResp)
	if err != nil {
		return nil, err
	}
	return messageCreateResp, err
}

func (fileService *files) Create(ctx context.Context, req *FileCreateReq,
	options ...lark.RequestOptionFunc) (*FileCreateResp, error) {
	options = append(options, lark.WithFileUpload())
	rawResp, err := fileService.app.SendRequestWithAccessTokenTypes(ctx, http.MethodPost,
		"/open-apis/im/v1/files", req, []lark.AccessTokenType{lark.AccessTokenTypeTenant}, options...)
	if err != nil {
		return nil, err
	}
	fileCreateResp := &FileCreateResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(fileCreateResp)
	if err != nil {
		return nil, err
	}
	return fileCreateResp, err
}

func (fileService *files) Get(ctx context.Context, req *FileGetReq,
	options ...lark.RequestOptionFunc) (*FileGetResp, error) {
	rawResp, err := fileService.app.SendRequestWithAccessTokenTypes(ctx, "GET",
		"/open-apis/im/v1/files/:file_key", req, []lark.AccessTokenType{lark.AccessTokenTypeTenant}, options...)
	if err != nil {
		return nil, err
	}
	fileGetResp := &FileGetResp{RawResponse: rawResp}
	if rawResp.StatusCode == http.StatusOK {
		fileGetResp.File = bytes.NewBuffer(rawResp.RawBody)
		fileGetResp.FileName = lark.FileNameByHeader(rawResp.Header)
		return fileGetResp, err
	}
	err = rawResp.JSONUnmarshalBody(fileGetResp)
	if err != nil {
		return nil, err
	}
	return fileGetResp, err
}
