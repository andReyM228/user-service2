package transfers

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"testing"
)

func TestRepository_Transfer(t *testing.T) {
	type fields struct {
		log    *logrus.Logger
		client *http.Client
	}
	type args struct {
		IDFrom int
		IDTo   int
		Amount int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				log:    logrus.New(),
				client: http.DefaultClient,
			},
			args: args{
				IDFrom: 1,
				IDTo:   2,
				Amount: 3,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Repository{
				log:    tt.fields.log,
				client: tt.fields.client,
			}
			if err := r.Transfer(tt.args.IDFrom, tt.args.IDTo, tt.args.Amount); (err != nil) != tt.wantErr {
				t.Errorf("Transfer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
