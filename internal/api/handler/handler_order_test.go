package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"task-level-0/internal/service"
	mock_service "task-level-0/internal/service/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestHandler_GetOrder(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockOrder, id string)

	testTable := []struct {
		name                string
		inputId             string
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:    "OK",
			inputId: "3b2a7ece-b75a-44a7-a9c5-c74331fc4e36",
			mockBehaviour: func(s *mock_service.MockOrder, id string) {
				s.EXPECT().GetOrderById(id).Return([]byte(a), nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: a,
		},
		{
			name:    "Server error",
			inputId: "4b2a7ece-b75a-44a7-a9c5-c74331fc4e36",
			mockBehaviour: func(s *mock_service.MockOrder, id string) {
				s.EXPECT().GetOrderById(id).Return(nil, errors.New("no rows in result set"))
			},
			expectedStatusCode:  http.StatusInternalServerError,
			expectedRequestBody: `{"error":"no rows in result set"}`,
		},
		{
			name:    "OK",
			inputId: "f64c245b-9535-4d2e-8ab1-be3ad2bd8099",
			mockBehaviour: func(s *mock_service.MockOrder, id string) {
				s.EXPECT().GetOrderById(id).Return([]byte(b), nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: b,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			order := mock_service.NewMockOrder(c)
			testCase.mockBehaviour(order, testCase.inputId)

			services := &service.Service{Order: order}
			handler := NewHandler(services)

			r := gin.New()
			r.GET("/api/order/:id", handler.GetOrderById)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/order/"+testCase.inputId, nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}

var a = `{"order_uid":"3b2a7ece-b75a-44a7-a9c5-c74331fc4e36","track_number":"WBILMTESTTRACK","entry":"WBIL","delivery":{"name":"Brenda Koelpin","phone":"4775902605","zip":"3568988","city":"Scottsdale","address":"88857 New Coveshaven","region":"Lakes","email":"evalyndietrich@hickle.io"},"payment":{"transaction":"3b2a7ece-b75a-44a7-a9c5-c74331fc4e36","request_id":"","currency":"EUR","provider":"wbpay","amount":7264,"payment_dt":1702419270,"bank":"tinkoff","delivery_cost":61,"goods_total":7203,"custom_fee":0},"items":[{"chrt_id":5771437,"track_number":"WBILMTESTTRACK","price":22603,"rid":"4b40d753-1332-48cf-a0e2-2d8d71ffd28f","name":"stream carbon lamp","sale":28,"size":"0","total_price":0,"nm_id":87867998,"brand":"AntiqueWhitefoot","status":202}],"locale":"ru","internal_signature":"","customer_id":"Brekke9483","delivery_service":"aluminum","shardkey":"25","sm_id":143,"date_created":"2023-12-12T22:14:30Z","oof_shard":"1"}`
var b = `{"order_uid":"f64c245b-9525-4d2e-8ab1-be3ad2bd8099","track_number":"WBILMTESTTRACK","entry":"WBIL","delivery":{"name":"Stacey Weissnat","phone":"4894788955","zip":"8214620","city":"Omaha","address":"1675 New Forksberg","region":"Meadows","email":"jaydenschaden@windler.org"},"payment":{"transaction":"f64c245b-9525-4d2e-8ab1-be3ad2bd8099","request_id":"","currency":"USD","provider":"wbpay","amount":7462,"payment_dt":1702419311,"bank":"sber","delivery_cost":127,"goods_total":7335,"custom_fee":0},"items":[{"chrt_id":1770630,"track_number":"WBILMTESTTRACK","price":26723,"rid":"8184c841-c457-4df1-9bc4-887045edb7be","name":"spark fresh smartwatch","sale":77,"size":"0","total_price":0,"nm_id":876328482,"brand":"Bowbow","status":202}],"locale":"en","internal_signature":"","customer_id":"Lehner7920","delivery_service":"suede","shardkey":"157","sm_id":131,"date_created":"2023-12-12T22:15:11Z","oof_shard":"1"}`
