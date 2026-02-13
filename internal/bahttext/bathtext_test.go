package bahttext

import (
	"testing"

	"github.com/shopspring/decimal"
)

func mustDec(t *testing.T, s string) decimal.Decimal {
	t.Helper()
	d, err := decimal.NewFromString(s)
	if err != nil {
		t.Fatalf("invalid decimal %q: %v", s, err)
	}
	return d
}

func TestToThaiBahtText_Table(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		// ---- Basic / zero ----
		{name: "zero", in: "0", want: "ศูนย์บาทถ้วน"},
		{name: "zero satang explicit", in: "0.00", want: "ศูนย์บาทถ้วน"},
		{name: "satang only 1", in: "0.01", want: "ศูนย์บาทหนึ่งสตางค์"},
		{name: "satang only 10", in: "0.10", want: "ศูนย์บาทสิบสตางค์"},
		{name: "satang only 11", in: "0.11", want: "ศูนย์บาทสิบเอ็ดสตางค์"},

		// ---- 1-99 grammar ----
		{name: "one", in: "1", want: "หนึ่งบาทถ้วน"},
		{name: "ten", in: "10", want: "สิบบาทถ้วน"},
		{name: "eleven", in: "11", want: "สิบเอ็ดบาทถ้วน"},
		{name: "twenty one", in: "21", want: "ยี่สิบเอ็ดบาทถ้วน"},
		{name: "ninety nine", in: "99", want: "เก้าสิบเก้าบาทถ้วน"},

		// ---- Satang with baht ----
		{name: "one baht one satang", in: "1.01", want: "หนึ่งบาทหนึ่งสตางค์"},
		{name: "one baht ten satang", in: "1.10", want: "หนึ่งบาทสิบสตางค์"},
		{name: "two baht twenty five satang", in: "2.25", want: "สองบาทยี่สิบห้าสตางค์"},
		{name: "example 33333.75", in: "33333.75", want: "สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์"},

		// ---- Hundreds / เอ็ด rules ----
		{name: "one hundred", in: "100", want: "หนึ่งร้อยบาทถ้วน"},
		{name: "one hundred one", in: "101", want: "หนึ่งร้อยเอ็ดบาทถ้วน"},
		{name: "one hundred ten", in: "110", want: "หนึ่งร้อยสิบบาทถ้วน"},
		{name: "one hundred eleven", in: "111", want: "หนึ่งร้อยสิบเอ็ดบาทถ้วน"},
		{name: "two hundred fifty", in: "250", want: "สองร้อยห้าสิบบาทถ้วน"},
		{name: "example 1234", in: "1234", want: "หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน"},

		// ---- Thousands ----
		{name: "one thousand", in: "1000", want: "หนึ่งพันบาทถ้วน"},
		{name: "one thousand one", in: "1001", want: "หนึ่งพันหนึ่งบาทถ้วน"},
		{name: "one thousand ten", in: "1010", want: "หนึ่งพันสิบบาทถ้วน"},
		{name: "one thousand one hundred", in: "1100", want: "หนึ่งพันหนึ่งร้อยบาทถ้วน"},
		{name: "one thousand one hundred eleven", in: "1111", want: "หนึ่งพันหนึ่งร้อยสิบเอ็ดบาทถ้วน"},

		// ---- Ten-thousands / hundred-thousands ----
		{name: "ten thousand", in: "10000", want: "หนึ่งหมื่นบาทถ้วน"},
		{name: "twenty thousand", in: "20000", want: "สองหมื่นบาทถ้วน"},
		{name: "one hundred thousand", in: "100000", want: "หนึ่งแสนบาทถ้วน"},
		{name: "three hundred thirty three thousand", in: "333000", want: "สามแสนสามหมื่นสามพันบาทถ้วน"},

		// ---- ล้าน grouping ----
		{name: "one million", in: "1000000", want: "หนึ่งล้านบาทถ้วน"},
		{name: "one million one", in: "1000001", want: "หนึ่งล้านหนึ่งบาทถ้วน"},
		{name: "two million", in: "2000000", want: "สองล้านบาทถ้วน"},
		{name: "one million with satang", in: "1000000.05", want: "หนึ่งล้านบาทห้าสตางค์"},
		{name: "complex million", in: "1234567.89", want: "หนึ่งล้านสองแสนสามหมื่นสี่พันห้าร้อยหกสิบเจ็ดบาทแปดสิบเก้าสตางค์"},

		// ---- Negative ----
		{name: "negative baht", in: "-12", want: "ลบสิบสองบาทถ้วน"},
		{name: "negative with satang", in: "-1.25", want: "ลบหนึ่งบาทยี่สิบห้าสตางค์"},

		// ---- Validation: too many decimals ----
		{name: "too many decimals", in: "1.234", wantErr: true},

		// ---- Leading Zeros
		{name: "leading zeros", in: "000123.40", want: "หนึ่งร้อยยี่สิบสามบาทสี่สิบสตางค์"},

		// One million group
		{name: "one million ten", in: "1000010", want: "หนึ่งล้านสิบบาทถ้วน"},
		{name: "one million one hundred", in: "1000100", want: "หนึ่งล้านหนึ่งร้อยบาทถ้วน"},
		{name: "one million eleven", in: "1000011", want: "หนึ่งล้านสิบเอ็ดบาทถ้วน"},
		// One Billion
		{name: "one billion", in: "1000000000", want: "หนึ่งพันล้านบาทถ้วน"},

		// mixed group + satang
		{name: "2,000,000.01", in: "2000000.01", want: "สองล้านบาทหนึ่งสตางค์"},
		{name: "12,000,034.50", in: "12000034.50", want: "สิบสองล้านสามสิบสี่บาทห้าสิบสตางค์"},

		// ---- >= พันล้าน ----
		{name: "1,000,000,000", in: "1000000000", want: "หนึ่งพันล้านบาทถ้วน"},
		{name: "2,000,000,000", in: "2000000000", want: "สองพันล้านบาทถ้วน"},
		{name: "10,000,000,000", in: "10000000000", want: "หนึ่งหมื่นล้านบาทถ้วน"},
		{name: "12,345,000,000", in: "12345000000", want: "หนึ่งหมื่นสองพันสามร้อยสี่สิบห้าล้านบาทถ้วน"},

		// ---- ล้านล้าน (>= 10^12) ----
		{name: "1,000,000,000,000", in: "1000000000000", want: "หนึ่งล้านล้านบาทถ้วน"},
		{name: "1,000,000,000,001", in: "1000000000001", want: "หนึ่งล้านล้านหนึ่งบาทถ้วน"},
		{name: "1,234,567,890,123.45", in: "1234567890123.45", want: "หนึ่งล้านสองแสนสามหมื่นสี่พันห้าร้อยหกสิบเจ็ดล้านแปดแสนเก้าหมื่นหนึ่งร้อยยี่สิบสามบาทสี่สิบห้าสตางค์"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			in := mustDec(t, tc.in)
			got, err := ToThaiBahtText(in)

			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil (result=%q)", got)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}
