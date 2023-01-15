package credit

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	type args struct {
		sum       string
		term      string
		rate      string
		isAnnuity bool
	}
	tests := []struct {
		name        string
		args        args
		wantMonthly string
		wantOverpay string
		wantTotal   string
		wantErr     bool
	}{
		{name: "empty", args: args{sum: "", term: "", rate: "", isAnnuity: false}, wantMonthly: "", wantOverpay: "", wantTotal: "", wantErr: true},
		{name: "non-valid", args: args{sum: "-10", term: "15", rate: "2", isAnnuity: false}, wantMonthly: "", wantOverpay: "", wantTotal: "", wantErr: true},
		{name: "annuity", args: args{sum: "50000", term: "12", rate: "5", isAnnuity: true}, wantMonthly: "4280.37", wantOverpay: "1364.44", wantTotal: "51364.44", wantErr: false},
		{name: "differentiated", args: args{sum: "50000", term: "12", rate: "5", isAnnuity: false}, wantMonthly: "4375..4184.03", wantOverpay: "1354.17", wantTotal: "51354.17", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMonthly, gotOverpay, gotTotal, err := Calculate(tt.args.sum, tt.args.term, tt.args.rate, tt.args.isAnnuity)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMonthly != tt.wantMonthly {
				t.Errorf("Calculate() gotMonthly = %v, want %v", gotMonthly, tt.wantMonthly)
			}
			if gotOverpay != tt.wantOverpay {
				t.Errorf("Calculate() gotOverpay = %v, want %v", gotOverpay, tt.wantOverpay)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("Calculate() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}
