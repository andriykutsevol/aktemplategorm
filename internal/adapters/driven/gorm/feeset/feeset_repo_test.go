package feeset

import (
	"context"
	"fmt"
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driven/gorm"
	"github.com/andriykusevol/aktemplategorm/internal/domain/entity/psp"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// https://tanutaran.medium.com/golang-unit-testing-with-gorm-and-sqlmock-postgresql-simplest-setup-67ccc7c056ef
// https://github.com/DATA-DOG/go-sqlmock
// https://medium.com/@adamszpilewicz/golang-mocking-sql-connection-in-unit-tests-with-sqlmock-a-practical-guide-33891efec439
// https://www.codingexplorations.com/blog/mastering-sqlmock-for-effective-database-testing
// https://www.codingexplorations.com/blog/testing-gorm-with-sqlmock#:~:text=Conclusion,your%20application%20logic%20is%20correct.

func Test_AddPSP_Create_Sucess(t *testing.T) {
	db, mock := orm.NewMockDB()
	_ = db
	_ = mock

	repo := NewRepository(db)
	_ = repo
	dto := PspDto{
		PspCode: "TEST1234",
	}
	_ = dto
	mock.ExpectBegin() // Gorm usually starts a transaction.

	// mock.ExpectExec("INSERT INTO psp_dtos \\(psp_code\\) VALUES \\('TEST1234'\\)").
	// 	WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO `Psp` \\(`psp_code`,`psp_short_name`,`psp_country_code`,`created_at`,`updated_at`,`deleted_at`,`created_by`,`updated_by`,`deleted_by`\\) VALUES \\(\\?,\\?,\\?,\\?,\\?,\\?,\\?,\\?,\\?\\)").
		WithArgs(
			"sucess",         // PspCode
			sqlmock.AnyArg(), // PspShortName
			sqlmock.AnyArg(), // PspCountryCode
			sqlmock.AnyArg(), // CreatedAt
			sqlmock.AnyArg(), // UpdatedAt
			sqlmock.AnyArg(), // DeletedAt
			sqlmock.AnyArg(), // CreatedBy
			sqlmock.AnyArg(), // UpdatedBy
			sqlmock.AnyArg(), // DeletedBy
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	//result := repo.Create(&dto)
	ctx := context.Background()

	p := psp.PSP{
		ID:      func(s uint) *uint { return &s }(uint(10)),
		PspCode: "sucess",
	}

	result, err := repo.AddPSP(ctx, p)
	_ = result

	assert.NoError(t, err)                    // Ensure no error occurred
	assert.NotNil(t, result)                  // Ensure result is not nil
	assert.Equal(t, "sucess", result.PspCode) // Ensure expected PspCode

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())

}

func Test_AddPSP_Create_Error(t *testing.T) {
	db, mock := orm.NewMockDB()
	_ = db
	_ = mock

	repo := NewRepository(db)
	_ = repo
	dto := PspDto{
		PspCode: "TEST1234",
	}
	_ = dto
	mock.ExpectBegin() // Gorm usually starts a transaction.

	// mock.ExpectExec("INSERT INTO psp_dtos \\(psp_code\\) VALUES \\('TEST1234'\\)").
	// 	WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO `Psp` \\(`psp_code`,`psp_short_name`,`psp_country_code`,`created_at`,`updated_at`,`deleted_at`,`created_by`,`updated_by`,`deleted_by`\\) VALUES \\(\\?,\\?,\\?,\\?,\\?,\\?,\\?,\\?,\\?\\)").
		WithArgs(
			"failure",        // PspCode (could be anything to differentiate error case)
			sqlmock.AnyArg(), // PspShortName
			sqlmock.AnyArg(), // PspCountryCode
			sqlmock.AnyArg(), // CreatedAt
			sqlmock.AnyArg(), // UpdatedAt
			sqlmock.AnyArg(), // DeletedAt
			sqlmock.AnyArg(), // CreatedBy
			sqlmock.AnyArg(), // UpdatedBy
			sqlmock.AnyArg(), // DeletedBy
		).
		WillReturnError(fmt.Errorf("mocked database error")) // Simulate a database error

	mock.ExpectRollback() // Expect rollback since transaction should fail

	//result := repo.Create(&dto)
	ctx := context.Background()

	p := psp.PSP{
		ID:      func(s uint) *uint { return &s }(uint(10)),
		PspCode: "sucess",
	}

	result, err := repo.AddPSP(ctx, p)

	assert.Nil(t, result)
	assert.Error(t, err)

}

func Test_repository_AddPSP(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		psp psp.PSP
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *psp.PSP
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: tt.fields.db,
			}
			got, err := r.AddPSP(tt.args.ctx, tt.args.psp)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.AddPSP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.AddPSP() = %v, want %v", got, tt.want)
			}
		})
	}
}
