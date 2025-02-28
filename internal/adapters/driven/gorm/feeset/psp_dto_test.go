package feeset

import (
	"fmt"
	domain_psp "github.com/andriykusevol/aktemplategorm/internal/domain/entity/psp"
	"reflect"
	"testing"
)

func TestDtoToDomain_PSP_Table(t *testing.T) {
	type args struct {
		dto PspDto
	}

	// Helper function to return a pointer to a uint
	uintPtr := func(v uint) *uint {
		return &v
	}

	tests := []struct {
		name    string
		args    args
		want    *domain_psp.PSP
		wantErr bool
	}{
		{
			name: "TestDtoToDomain_PSP_Sucess",
			args: args{
				dto: PspDto{
					ID: uintPtr(10),
				},
			},
			wantErr: false,
			want: &domain_psp.PSP{
				ID: uintPtr(10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DtoToDomain_PSP(tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("DtoToDomain_PSP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got.ID)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DtoToDomain_PSP() = %v, want %v", got, tt.want)
			}
		})
	}

}
