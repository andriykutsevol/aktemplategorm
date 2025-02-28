package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/andriykusevol/aktemplategorm/internal/domain/entity/psp"
	"github.com/andriykusevol/aktemplategorm/internal/domain/sport"
	"reflect"
	"testing"

	feesetRepoMock "github.com/andriykusevol/aktemplategorm/mocks/internal_/domain/sport"

	"github.com/stretchr/testify/mock"
)

func Test_feeSetApp_PspAdd(t *testing.T) {
	type fields struct {
		feeSetRepo sport.FeeSetRepository
	}
	type args struct {
		ctx context.Context
		psp *psp.PSP
	}

	mockfeeset := feesetRepoMock.NewFeeSetRepository(t)
	_ = mockfeeset

	testCases := []struct {
		name                     string
		fields                   fields
		args                     args
		want                     *psp.PSP
		wantErr                  bool
		expectation_AddPSP_Ok    func()
		expectation_AddPSP_Error func()
	}{
		{
			name: "Sucess",
			fields: fields{
				feeSetRepo: mockfeeset,
			},
			expectation_AddPSP_Ok: func() {
				domainItem := &psp.PSP{
					ID:      func(s uint) *uint { return &s }(uint(10)),
					PspCode: "sucess",
				}
				mockfeeset.On("AddPSP", mock.Anything, mock.MatchedBy(func(pspArg psp.PSP) bool {
					return pspArg.PspCode == "sucess"
				})).Return(domainItem, nil).Once()
				//mockfeeset.On("AddPSP", mock.Anything, mock.AnythingOfType("psp.PSP")).Return(domainItem, nil).Once()
			},
			args: args{
				psp: &psp.PSP{
					PspCode: "sucess",
				},
				ctx: context.Background(),
			},
			want: &psp.PSP{
				ID:      func(s uint) *uint { return &s }(uint(10)),
				PspCode: "sucess",
			},
			wantErr: false,
		},
		{
			name: "Error",
			fields: fields{
				feeSetRepo: mockfeeset,
			},
			expectation_AddPSP_Ok: func() {
				expectedError := errors.New("Error: repositoory AddPSP")
				mockfeeset.On("AddPSP", mock.Anything, mock.MatchedBy(func(pspArg psp.PSP) bool {
					return pspArg.PspCode == "error"
				})).Return(nil, expectedError).Once()
				//mockfeeset.On("AddPSP", mock.Anything, mock.AnythingOfType("psp.PSP")).Return(domainItem, nil).Once()
			},
			args: args{
				psp: &psp.PSP{
					PspCode: "error",
				},
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			if tc.expectation_AddPSP_Ok != nil {
				tc.expectation_AddPSP_Ok()
			}

			app := &feeSetApp{
				feeSetRepo: tc.fields.feeSetRepo,
			}
			got, err := app.PspAdd(tc.args.ctx, tc.args.psp)
			if (err != nil) != tc.wantErr {
				fmt.Println("111")
				t.Errorf("feeSetApp.PspAdd() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				fmt.Println("333")
				t.Errorf("feeSetApp.PspAdd() = %v, want %v", got, tc.want)
			}
		})
	}
}
