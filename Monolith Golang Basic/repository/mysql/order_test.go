package mysql

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"golang-basic/model"
	"net/http/httptest"
	"reflect"
	"regexp"
	"testing"
)

func Test_GetOrdersByIDs(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	mockDBRepo := NewMock(mockDB)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)

	successGet := []model.Orders{
		{
			Id:              int64(1),
			GoodsName:       "Meja",
			ReceiverName:    "Audi",
			ReceiverAddress: "BSD",
			ShipperID:       1,
		},
	}

	type args struct {
		IDs []int64
	}

	tests := []struct {
		name    string
		args   args
		mock    func()
		want    []model.Orders
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				[]int64{1},
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "goods_name", "receiver_name", "receiver_address", "shipper_id"}).
					AddRow(int64(1), "Meja", "Audi", "BSD", 1)
				mock.ExpectQuery("SELECT (.+) FROM orders WHERE id IN (.+)").WithArgs(1).WillReturnRows(rows)
			},
			want: successGet,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := mockDBRepo.GetOrdersByIDs(context, tt.args.IDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrdersByIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrdersByIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Insert(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	mockDBRepo := NewMock(mockDB)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)

	insertOrder := model.Orders{
			Id:              int64(1),
			GoodsName:       "Meja",
			ReceiverName:    "Audi",
			ReceiverAddress: "BSD",
			ShipperID:       1,
	}

	type args struct {
		IDs []int64
	}

	tests := []struct {
		name    string
		args   args
		mock    func()
		want int64
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				[]int64{1},
			},
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO orders(goods_name,receiver_name,receiver_address,shipper_id) VALUES(?,?,?,?)`)).WithArgs(
					insertOrder.GoodsName,
					insertOrder.ReceiverName,
					insertOrder.ReceiverAddress,
					insertOrder.ShipperID,
				).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			id, err := mockDBRepo.Insert(context, insertOrder)
			if (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(id, tt.want) {
				t.Errorf("Insert() = %v, want %v", id, tt.want)
			}
		})
	}
}