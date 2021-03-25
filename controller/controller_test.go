package controller

import (
	"testing"

	"github.com/lbswl/academy-go-q12021/controller/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

//mockgen -source=controller/controller.go -destination=controller/mocks/controller.go -package=mocks
func Test_Controller(t *testing.T) {

	tests := []struct {
		name                    string
		expectedParams          int
		expectedUsecaseResponse []byte
		expectUsecaseCall       bool
		expectedError           error
		wantError               bool
	}{
		{
			name:                    "OK, FindUserbyId",
			expectedParams:          1,
			expectedUsecaseResponse: []byte(`[{"id": 2, "gender": "female", "title": "Mrs", "first": "Rebecca", "last": "Stevens", "email": "rebecca.stevens@example.com", "cell": "0420-442-682", "nat": "AU"}]`),
			expectUsecaseCall:       true,
			wantError:               false,
			expectedError:           nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			u := mocks.NewMockUseCase(mockCtrl)

			if tt.expectUsecaseCall {
				u.EXPECT().FindUserById(tt.expectedParams).Return(tt.expectedUsecaseResponse, tt.expectedError)
			}

			c := New(u)

			response, err := c.useCase.FindUserById(tt.expectedParams)
			assert.Equal(t, response, tt.expectedUsecaseResponse)

			if tt.wantError {
				assert.NotNil(t, err)
				assert.Equal(t, err, tt.expectedError)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
