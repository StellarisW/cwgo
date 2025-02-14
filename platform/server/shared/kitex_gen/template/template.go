// Code generated by thriftgo (0.3.4). DO NOT EDIT.

package template

import (
	"context"
	"fmt"
	"github.com/cloudwego/cwgo/platform/server/shared/kitex_gen/model"
)

type AddTemplateReq struct {
	Name string `thrift:"name,1" frugal:"1,default,string" form:"name,required" json:"name,required" vd:"len($)>0"`
	Type int32  `thrift:"type,2" frugal:"2,default,i32" form:"type,required" json:"type,required"`
}

func NewAddTemplateReq() *AddTemplateReq {
	return &AddTemplateReq{}
}

func (p *AddTemplateReq) InitDefault() {
	*p = AddTemplateReq{}
}

func (p *AddTemplateReq) GetName() (v string) {
	return p.Name
}

func (p *AddTemplateReq) GetType() (v int32) {
	return p.Type
}
func (p *AddTemplateReq) SetName(val string) {
	p.Name = val
}
func (p *AddTemplateReq) SetType(val int32) {
	p.Type = val
}

func (p *AddTemplateReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddTemplateReq(%+v)", *p)
}

type AddTemplateRes struct {
	Code int32  `thrift:"code,1" frugal:"1,default,i32" form:"code" json:"code" query:"code"`
	Msg  string `thrift:"msg,2" frugal:"2,default,string" form:"msg" json:"msg" query:"msg"`
}

func NewAddTemplateRes() *AddTemplateRes {
	return &AddTemplateRes{}
}

func (p *AddTemplateRes) InitDefault() {
	*p = AddTemplateRes{}
}

func (p *AddTemplateRes) GetCode() (v int32) {
	return p.Code
}

func (p *AddTemplateRes) GetMsg() (v string) {
	return p.Msg
}
func (p *AddTemplateRes) SetCode(val int32) {
	p.Code = val
}
func (p *AddTemplateRes) SetMsg(val string) {
	p.Msg = val
}

func (p *AddTemplateRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddTemplateRes(%+v)", *p)
}

type DeleteTemplateReq struct {
	Ids []int64 `thrift:"ids,1" frugal:"1,default,list<i64>" form:"ids,required" json:"ids,required" vd:"len($)>0"`
}

func NewDeleteTemplateReq() *DeleteTemplateReq {
	return &DeleteTemplateReq{}
}

func (p *DeleteTemplateReq) InitDefault() {
	*p = DeleteTemplateReq{}
}

func (p *DeleteTemplateReq) GetIds() (v []int64) {
	return p.Ids
}
func (p *DeleteTemplateReq) SetIds(val []int64) {
	p.Ids = val
}

func (p *DeleteTemplateReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DeleteTemplateReq(%+v)", *p)
}

type DeleteTemplateRes struct {
	Code int32  `thrift:"code,1" frugal:"1,default,i32" form:"code" json:"code" query:"code"`
	Msg  string `thrift:"msg,2" frugal:"2,default,string" form:"msg" json:"msg" query:"msg"`
}

func NewDeleteTemplateRes() *DeleteTemplateRes {
	return &DeleteTemplateRes{}
}

func (p *DeleteTemplateRes) InitDefault() {
	*p = DeleteTemplateRes{}
}

func (p *DeleteTemplateRes) GetCode() (v int32) {
	return p.Code
}

func (p *DeleteTemplateRes) GetMsg() (v string) {
	return p.Msg
}
func (p *DeleteTemplateRes) SetCode(val int32) {
	p.Code = val
}
func (p *DeleteTemplateRes) SetMsg(val string) {
	p.Msg = val
}

func (p *DeleteTemplateRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DeleteTemplateRes(%+v)", *p)
}

type UpdateTemplateReq struct {
	Id   int64  `thrift:"id,1" frugal:"1,default,i64" form:"id,required" json:"id,required"`
	Name string `thrift:"name,2" frugal:"2,default,string" form:"name,required" json:"name,required" vd:"len($)>0"`
}

func NewUpdateTemplateReq() *UpdateTemplateReq {
	return &UpdateTemplateReq{}
}

func (p *UpdateTemplateReq) InitDefault() {
	*p = UpdateTemplateReq{}
}

func (p *UpdateTemplateReq) GetId() (v int64) {
	return p.Id
}

func (p *UpdateTemplateReq) GetName() (v string) {
	return p.Name
}
func (p *UpdateTemplateReq) SetId(val int64) {
	p.Id = val
}
func (p *UpdateTemplateReq) SetName(val string) {
	p.Name = val
}

func (p *UpdateTemplateReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UpdateTemplateReq(%+v)", *p)
}

type UpdateTemplateRes struct {
	Code int32  `thrift:"code,1" frugal:"1,default,i32" form:"code" json:"code" query:"code"`
	Msg  string `thrift:"msg,2" frugal:"2,default,string" form:"msg" json:"msg" query:"msg"`
}

func NewUpdateTemplateRes() *UpdateTemplateRes {
	return &UpdateTemplateRes{}
}

func (p *UpdateTemplateRes) InitDefault() {
	*p = UpdateTemplateRes{}
}

func (p *UpdateTemplateRes) GetCode() (v int32) {
	return p.Code
}

func (p *UpdateTemplateRes) GetMsg() (v string) {
	return p.Msg
}
func (p *UpdateTemplateRes) SetCode(val int32) {
	p.Code = val
}
func (p *UpdateTemplateRes) SetMsg(val string) {
	p.Msg = val
}

func (p *UpdateTemplateRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UpdateTemplateRes(%+v)", *p)
}

type GetTemplatesReq struct {
	Page    int32  `thrift:"page,1" frugal:"1,default,i32" json:"page" query:"page" vd:"$>=0"`
	Limit   int32  `thrift:"limit,2" frugal:"2,default,i32" json:"limit" query:"limit" vd:"$>=0"`
	Order   int32  `thrift:"order,3" frugal:"3,default,i32" json:"order" query:"order" vd:"$>=0"`
	OrderBy string `thrift:"order_by,4" frugal:"4,default,string" json:"order_by" query:"order_by"`
}

func NewGetTemplatesReq() *GetTemplatesReq {
	return &GetTemplatesReq{}
}

func (p *GetTemplatesReq) InitDefault() {
	*p = GetTemplatesReq{}
}

func (p *GetTemplatesReq) GetPage() (v int32) {
	return p.Page
}

func (p *GetTemplatesReq) GetLimit() (v int32) {
	return p.Limit
}

func (p *GetTemplatesReq) GetOrder() (v int32) {
	return p.Order
}

func (p *GetTemplatesReq) GetOrderBy() (v string) {
	return p.OrderBy
}
func (p *GetTemplatesReq) SetPage(val int32) {
	p.Page = val
}
func (p *GetTemplatesReq) SetLimit(val int32) {
	p.Limit = val
}
func (p *GetTemplatesReq) SetOrder(val int32) {
	p.Order = val
}
func (p *GetTemplatesReq) SetOrderBy(val string) {
	p.OrderBy = val
}

func (p *GetTemplatesReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetTemplatesReq(%+v)", *p)
}

type GetTemplatesRes struct {
	Code int32                `thrift:"code,1" frugal:"1,default,i32" form:"code" json:"code" query:"code"`
	Msg  string               `thrift:"msg,2" frugal:"2,default,string" form:"msg" json:"msg" query:"msg"`
	Data *GetTemplatesResData `thrift:"data,3" frugal:"3,default,GetTemplatesResData" form:"data" json:"data" query:"data"`
}

func NewGetTemplatesRes() *GetTemplatesRes {
	return &GetTemplatesRes{}
}

func (p *GetTemplatesRes) InitDefault() {
	*p = GetTemplatesRes{}
}

func (p *GetTemplatesRes) GetCode() (v int32) {
	return p.Code
}

func (p *GetTemplatesRes) GetMsg() (v string) {
	return p.Msg
}

var GetTemplatesRes_Data_DEFAULT *GetTemplatesResData

func (p *GetTemplatesRes) GetData() (v *GetTemplatesResData) {
	if !p.IsSetData() {
		return GetTemplatesRes_Data_DEFAULT
	}
	return p.Data
}
func (p *GetTemplatesRes) SetCode(val int32) {
	p.Code = val
}
func (p *GetTemplatesRes) SetMsg(val string) {
	p.Msg = val
}
func (p *GetTemplatesRes) SetData(val *GetTemplatesResData) {
	p.Data = val
}

func (p *GetTemplatesRes) IsSetData() bool {
	return p.Data != nil
}

func (p *GetTemplatesRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetTemplatesRes(%+v)", *p)
}

type GetTemplatesResData struct {
	Templates []*model.Template `thrift:"templates,1" frugal:"1,default,list<model.Template>" form:"templates" json:"templates" query:"templates"`
}

func NewGetTemplatesResData() *GetTemplatesResData {
	return &GetTemplatesResData{}
}

func (p *GetTemplatesResData) InitDefault() {
	*p = GetTemplatesResData{}
}

func (p *GetTemplatesResData) GetTemplates() (v []*model.Template) {
	return p.Templates
}
func (p *GetTemplatesResData) SetTemplates(val []*model.Template) {
	p.Templates = val
}

func (p *GetTemplatesResData) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetTemplatesResData(%+v)", *p)
}

type AddTemplateItemReq struct {
	TemplateId int64  `thrift:"template_id,1" frugal:"1,default,i64" form:"template_id" json:"template_id" query:"template_id"`
	Name       string `thrift:"name,2" frugal:"2,default,string" form:"name,required" json:"name,required" vd:"len($)>0"`
	Content    string `thrift:"content,3" frugal:"3,default,string" form:"content,required" json:"content,required"`
}

func NewAddTemplateItemReq() *AddTemplateItemReq {
	return &AddTemplateItemReq{}
}

func (p *AddTemplateItemReq) InitDefault() {
	*p = AddTemplateItemReq{}
}

func (p *AddTemplateItemReq) GetTemplateId() (v int64) {
	return p.TemplateId
}

func (p *AddTemplateItemReq) GetName() (v string) {
	return p.Name
}

func (p *AddTemplateItemReq) GetContent() (v string) {
	return p.Content
}
func (p *AddTemplateItemReq) SetTemplateId(val int64) {
	p.TemplateId = val
}
func (p *AddTemplateItemReq) SetName(val string) {
	p.Name = val
}
func (p *AddTemplateItemReq) SetContent(val string) {
	p.Content = val
}

func (p *AddTemplateItemReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddTemplateItemReq(%+v)", *p)
}

type AddTemplateItemRes struct {
	Code int32  `thrift:"code,1" frugal:"1,default,i32" form:"code" json:"code" query:"code"`
	Msg  string `thrift:"msg,2" frugal:"2,default,string" form:"msg" json:"msg" query:"msg"`
}

func NewAddTemplateItemRes() *AddTemplateItemRes {
	return &AddTemplateItemRes{}
}

func (p *AddTemplateItemRes) InitDefault() {
	*p = AddTemplateItemRes{}
}

func (p *AddTemplateItemRes) GetCode() (v int32) {
	return p.Code
}

func (p *AddTemplateItemRes) GetMsg() (v string) {
	return p.Msg
}
func (p *AddTemplateItemRes) SetCode(val int32) {
	p.Code = val
}
func (p *AddTemplateItemRes) SetMsg(val string) {
	p.Msg = val
}

func (p *AddTemplateItemRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddTemplateItemRes(%+v)", *p)
}

type DeleteTemplateItemReq struct {
	Ids []int64 `thrift:"ids,1" frugal:"1,default,list<i64>" form:"ids,required" json:"ids,required" vd:"len($)>0"`
}

func NewDeleteTemplateItemReq() *DeleteTemplateItemReq {
	return &DeleteTemplateItemReq{}
}

func (p *DeleteTemplateItemReq) InitDefault() {
	*p = DeleteTemplateItemReq{}
}

func (p *DeleteTemplateItemReq) GetIds() (v []int64) {
	return p.Ids
}
func (p *DeleteTemplateItemReq) SetIds(val []int64) {
	p.Ids = val
}

func (p *DeleteTemplateItemReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DeleteTemplateItemReq(%+v)", *p)
}

type DeleteTemplateItemRes struct {
	Code int32  `thrift:"code,1" frugal:"1,default,i32" form:"code" json:"code" query:"code"`
	Msg  string `thrift:"msg,2" frugal:"2,default,string" form:"msg" json:"msg" query:"msg"`
}

func NewDeleteTemplateItemRes() *DeleteTemplateItemRes {
	return &DeleteTemplateItemRes{}
}

func (p *DeleteTemplateItemRes) InitDefault() {
	*p = DeleteTemplateItemRes{}
}

func (p *DeleteTemplateItemRes) GetCode() (v int32) {
	return p.Code
}

func (p *DeleteTemplateItemRes) GetMsg() (v string) {
	return p.Msg
}
func (p *DeleteTemplateItemRes) SetCode(val int32) {
	p.Code = val
}
func (p *DeleteTemplateItemRes) SetMsg(val string) {
	p.Msg = val
}

func (p *DeleteTemplateItemRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DeleteTemplateItemRes(%+v)", *p)
}

type UpdateTemplateItemReq struct {
	Id      int64  `thrift:"id,1" frugal:"1,default,i64" form:"id,required" json:"id,required"`
	Name    string `thrift:"name,2" frugal:"2,default,string" form:"name" json:"name" vd:"len($)>0"`
	Content string `thrift:"content,3" frugal:"3,default,string" form:"content" json:"content"`
}

func NewUpdateTemplateItemReq() *UpdateTemplateItemReq {
	return &UpdateTemplateItemReq{}
}

func (p *UpdateTemplateItemReq) InitDefault() {
	*p = UpdateTemplateItemReq{}
}

func (p *UpdateTemplateItemReq) GetId() (v int64) {
	return p.Id
}

func (p *UpdateTemplateItemReq) GetName() (v string) {
	return p.Name
}

func (p *UpdateTemplateItemReq) GetContent() (v string) {
	return p.Content
}
func (p *UpdateTemplateItemReq) SetId(val int64) {
	p.Id = val
}
func (p *UpdateTemplateItemReq) SetName(val string) {
	p.Name = val
}
func (p *UpdateTemplateItemReq) SetContent(val string) {
	p.Content = val
}

func (p *UpdateTemplateItemReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UpdateTemplateItemReq(%+v)", *p)
}

type UpdateTemplateItemRes struct {
	Code int32  `thrift:"code,1" frugal:"1,default,i32" form:"code" json:"code" query:"code"`
	Msg  string `thrift:"msg,2" frugal:"2,default,string" form:"msg" json:"msg" query:"msg"`
}

func NewUpdateTemplateItemRes() *UpdateTemplateItemRes {
	return &UpdateTemplateItemRes{}
}

func (p *UpdateTemplateItemRes) InitDefault() {
	*p = UpdateTemplateItemRes{}
}

func (p *UpdateTemplateItemRes) GetCode() (v int32) {
	return p.Code
}

func (p *UpdateTemplateItemRes) GetMsg() (v string) {
	return p.Msg
}
func (p *UpdateTemplateItemRes) SetCode(val int32) {
	p.Code = val
}
func (p *UpdateTemplateItemRes) SetMsg(val string) {
	p.Msg = val
}

func (p *UpdateTemplateItemRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UpdateTemplateItemRes(%+v)", *p)
}

type GetTemplateItemsReq struct {
	Id      int64  `thrift:"id,1" frugal:"1,default,i64" json:"id" query:"id" vd:"$>=0"`
	Page    int32  `thrift:"page,2" frugal:"2,default,i32" json:"page" query:"page" vd:"$>=0"`
	Limit   int32  `thrift:"limit,3" frugal:"3,default,i32" json:"limit" query:"limit" vd:"$>=0"`
	Order   int32  `thrift:"order,4" frugal:"4,default,i32" json:"order" query:"order" vd:"$>=0"`
	OrderBy string `thrift:"order_by,5" frugal:"5,default,string" json:"order_by" query:"order_by"`
}

func NewGetTemplateItemsReq() *GetTemplateItemsReq {
	return &GetTemplateItemsReq{}
}

func (p *GetTemplateItemsReq) InitDefault() {
	*p = GetTemplateItemsReq{}
}

func (p *GetTemplateItemsReq) GetId() (v int64) {
	return p.Id
}

func (p *GetTemplateItemsReq) GetPage() (v int32) {
	return p.Page
}

func (p *GetTemplateItemsReq) GetLimit() (v int32) {
	return p.Limit
}

func (p *GetTemplateItemsReq) GetOrder() (v int32) {
	return p.Order
}

func (p *GetTemplateItemsReq) GetOrderBy() (v string) {
	return p.OrderBy
}
func (p *GetTemplateItemsReq) SetId(val int64) {
	p.Id = val
}
func (p *GetTemplateItemsReq) SetPage(val int32) {
	p.Page = val
}
func (p *GetTemplateItemsReq) SetLimit(val int32) {
	p.Limit = val
}
func (p *GetTemplateItemsReq) SetOrder(val int32) {
	p.Order = val
}
func (p *GetTemplateItemsReq) SetOrderBy(val string) {
	p.OrderBy = val
}

func (p *GetTemplateItemsReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetTemplateItemsReq(%+v)", *p)
}

type GetTemplateItemsRes struct {
	Code int32                    `thrift:"code,1" frugal:"1,default,i32" form:"code" json:"code" query:"code"`
	Msg  string                   `thrift:"msg,2" frugal:"2,default,string" form:"msg" json:"msg" query:"msg"`
	Data *GetTemplateItemsResData `thrift:"data,3" frugal:"3,default,GetTemplateItemsResData" form:"data" json:"data" query:"data"`
}

func NewGetTemplateItemsRes() *GetTemplateItemsRes {
	return &GetTemplateItemsRes{}
}

func (p *GetTemplateItemsRes) InitDefault() {
	*p = GetTemplateItemsRes{}
}

func (p *GetTemplateItemsRes) GetCode() (v int32) {
	return p.Code
}

func (p *GetTemplateItemsRes) GetMsg() (v string) {
	return p.Msg
}

var GetTemplateItemsRes_Data_DEFAULT *GetTemplateItemsResData

func (p *GetTemplateItemsRes) GetData() (v *GetTemplateItemsResData) {
	if !p.IsSetData() {
		return GetTemplateItemsRes_Data_DEFAULT
	}
	return p.Data
}
func (p *GetTemplateItemsRes) SetCode(val int32) {
	p.Code = val
}
func (p *GetTemplateItemsRes) SetMsg(val string) {
	p.Msg = val
}
func (p *GetTemplateItemsRes) SetData(val *GetTemplateItemsResData) {
	p.Data = val
}

func (p *GetTemplateItemsRes) IsSetData() bool {
	return p.Data != nil
}

func (p *GetTemplateItemsRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetTemplateItemsRes(%+v)", *p)
}

type GetTemplateItemsResData struct {
	TemplateItems []*model.TemplateItem `thrift:"template_items,1" frugal:"1,default,list<model.TemplateItem>" form:"template_items" json:"template_items" query:"template_items"`
}

func NewGetTemplateItemsResData() *GetTemplateItemsResData {
	return &GetTemplateItemsResData{}
}

func (p *GetTemplateItemsResData) InitDefault() {
	*p = GetTemplateItemsResData{}
}

func (p *GetTemplateItemsResData) GetTemplateItems() (v []*model.TemplateItem) {
	return p.TemplateItems
}
func (p *GetTemplateItemsResData) SetTemplateItems(val []*model.TemplateItem) {
	p.TemplateItems = val
}

func (p *GetTemplateItemsResData) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetTemplateItemsResData(%+v)", *p)
}

type TemplateService interface {
	AddTemplate(ctx context.Context, req *AddTemplateReq) (r *AddTemplateRes, err error)

	DeleteTemplate(ctx context.Context, req *DeleteTemplateReq) (r *DeleteTemplateRes, err error)

	UpdateTemplate(ctx context.Context, req *UpdateTemplateReq) (r *UpdateTemplateRes, err error)

	GetTemplates(ctx context.Context, req *GetTemplatesReq) (r *GetTemplatesRes, err error)

	AddTemplateItem(ctx context.Context, req *AddTemplateItemReq) (r *AddTemplateItemRes, err error)

	DeleteTemplateItem(ctx context.Context, req *DeleteTemplateItemReq) (r *DeleteTemplateItemRes, err error)

	UpdateTemplateItem(ctx context.Context, req *UpdateTemplateItemReq) (r *UpdateTemplateItemRes, err error)

	GetTemplateItems(ctx context.Context, req *GetTemplateItemsReq) (r *GetTemplateItemsRes, err error)
}
