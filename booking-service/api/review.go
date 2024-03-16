package apiBooking

import (
	"booking-service/internal/model"
	"booking-service/internal/util"
	protoBooking "booking-service/proto/booking"
	protoSdk "booking-service/proto/sdk"
	"context"
	"fmt"
	"reflect"
	"strings"
	"github.com/hadanhtuan/go-sdk/common"
	"github.com/hadanhtuan/go-sdk/db/orm"
)

func (bc *BookingController) CreateReview(ctx context.Context, req *protoBooking.MsgCreateReview) (*protoSdk.BaseResponse, error) {
	review := &model.Review{
		UserId:     req.UserId,
		PropertyId: req.PropertyId,
		Rating:     req.Rating,
		Comment:    req.Comment,
		ImageUrl:   req.ImageUrl,
	}

	if req.ParentId != "" {
		review.ParentId = &req.ParentId
	}

	result := model.ReviewDB.Create(review)
	if result.Status != common.APIStatus.Created {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.ServerError,
			Message: "Create Review Failed.",
		})
	}

	return util.ConvertToGRPC(&common.APIResponse{
		Status:  common.APIStatus.Ok,
		Message: "Create Review Successfully.",
	})
}

func (bc *BookingController) UpdateReview(ctx context.Context, req *protoBooking.MsgUpdateReview) (*protoSdk.BaseResponse, error) {
	review := &model.Review{
		ID: req.ReviewId,
	}

	reviewUpdated := &model.Review{
		Rating:   req.Rating,
		Comment:  req.Comment,
		ImageUrl: req.ImageUrl,
	}

	result := model.ReviewDB.Update(review, reviewUpdated)
	return util.ConvertToGRPC(result)

}

func (bc *BookingController) DeleteReview(ctx context.Context, req *protoBooking.MsgDeleteReview) (*protoSdk.BaseResponse, error) {
	review := &model.Review{
		ID: req.ReviewId,
	}

	result := model.ReviewDB.Delete(review)
	return util.ConvertToGRPC(result)
}

func (bc *BookingController) GetReview(ctx context.Context, req *protoBooking.MessageQueryReview) (*protoSdk.BaseResponse, error) {
	filter := &model.Review{}

	if req.QueryFields.PropertyId != nil {
		filter.PropertyId = req.QueryFields.PropertyId
	}

	result := model.ReviewDB.Query(filter, req.Paginate.Offset, req.Paginate.Limit, &orm.QueryOption{
		Order: []string{"created_at asc"},
	})

	data := result.Data.([]*model.Review)
	fmt.Println(data)

	return util.ConvertToGRPC(result)
}

// func newListObject(m *orm.Instance, limit int) interface{} {
// 	t := reflect.TypeOf(m.Model)
// 	slice := reflect.MakeSlice(reflect.SliceOf(t), 0, limit)
// 	slicePtr := reflect.New(slice.Type())

// 	slicePtr.Elem()
// 	return slicePtr.Interface()

// }
func newListObject(m *orm.Instance, limit int) interface{} {
	t := reflect.TypeOf(m.Model) // Get the element type of the slice
	slice := reflect.MakeSlice(reflect.SliceOf(t), 0, limit)

	slicePtr := reflect.New(slice.Type())

	return slicePtr.Interface()
}

type optionQuery struct {
	Preload []string
	Order   []string
}

func Query(m *orm.Instance, query interface{}, offset int32, limit int32, option *optionQuery) *common.APIResponse {
	// check table
	if m.DB == nil {
		return &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "DB error: Table " + m.TableName + " is not init.",
		}
	}

	entities := newListObject(m, int(limit))
	var total int64

	db := m.DB.WithContext(context.TODO()).Table(m.TableName).Model(m.Model)

	if option != nil {
		if option.Preload != nil {
			for _, preload := range option.Preload {
				db.Preload(preload)
			}
		}

		if option.Order != nil {
			orders := strings.Join(option.Order, ", ")
			db.Order(orders)
		}
	}

	err := db.Where(query).Count(&total). // count
						Offset(int((offset - 1) * limit)).Limit(int(limit)). // paginate
						Where(query).Find(entities).Error

	if err != nil {
		return &common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Cannot find item in table " + ". Error detail: " + err.Error(),
		}
	}
	var data interface{}
	v := reflect.ValueOf(entities)

	if v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem() // Dereference the pointer
		// Now 'v' holds the value that the pointer points to
		data = v.Interface()
	} else {
		data = nil
	}

	return &common.APIResponse{
		Status:  common.APIStatus.Ok,
		Data:    data,
		Total:   total,
		Message: "Query " + m.TableName + " successfully.",
	}
}
