package bahttext

import (
	"errors"
	"strings"

	"github.com/shopspring/decimal"
)

var ErrTooManyDecimals = errors.New("amount has more than 2 decimal places")

func getValueText(value rune, powerOfTenMod int, lastPowerOfTenText string) string {
	var valueText string
	switch value {
	case '0':
		// If no prefix
		if lastPowerOfTenText == "" {
			valueText = "ศูนย์"
		} else {
			valueText = ""
		}
	case '1':
		// หนึ่งสิบ
		if powerOfTenMod == 1 {
			valueText = ""
		} else if lastPowerOfTenText == "สิบ" || lastPowerOfTenText == "ร้อย" {
			valueText = "เอ็ด"
		} else {
			valueText = "หนึ่ง"
		}
	case '2':
		if powerOfTenMod == 1 {
			valueText = "ยี่"
		} else {
			valueText = "สอง"
		}
	case '3':
		valueText = "สาม"
	case '4':
		valueText = "สี่"
	case '5':
		valueText = "ห้า"
	case '6':
		valueText = "หก"
	case '7':
		valueText = "เจ็ด"
	case '8':
		valueText = "แปด"
	case '9':
		valueText = "เก้า"
	}
	return valueText
}

func getPowerOfTenText(powerOfTen int, powerOfTenMod int) string {
	// Prevent % 6 == 0 and get ล้าน
	if powerOfTen == 0 {
		return ""
	}
	var powerOfTenText string
	switch powerOfTenMod {
	case 5:
		powerOfTenText = "แสน"
	case 4:
		powerOfTenText = "หมื่น"
	case 3:
		powerOfTenText = "พัน"
	case 2:
		powerOfTenText = "ร้อย"
	case 1:
		powerOfTenText = "สิบ"
	case 0:
		powerOfTenText = "ล้าน"
	}
	return powerOfTenText
}

func spellAmountString(amountText string) string {
	var textList []string
	for index, value := range amountText {
		powerOfTen := len(amountText) - index - 1
		powerOfTenMod := powerOfTen % 6
		powerOfTenText := getPowerOfTenText(powerOfTen, powerOfTenMod)

		var last string
		if len(textList) > 0 {
			last = textList[len(textList)-1]
		}

		valueText := getValueText(value, powerOfTenMod, last)

		if valueText != "" {
			textList = append(textList, valueText)
		}
		// 0 dont need unit, except ล้าน
		if value != '0' || powerOfTen%6 == 0 {
			textList = append(textList, powerOfTenText)
		}
	}

	return strings.Join(textList, "")
}

func addSuffix(satang decimal.Decimal, text string) string {
	if satang.IsZero() {
		return text + "บาทถ้วน"
	}

	return text + "บาท" + spellAmountString(satang.String()) + "สตางค์"
}

func ToThaiBahtText(amount decimal.Decimal) (string, error) {
	// Satang shouldn't have more than 2 decimal places
	if amount.Exponent() < -2 {
		return "", ErrTooManyDecimals
	}
	// Round anyway to prevent carry problems 1.999 -> 2.00
	amount = amount.Round(2)

	// Check if negative
	prefix := ""
	if amount.IsNegative() {
		prefix = "ลบ"
		amount = amount.Abs()
	}

	// Extract Baht and Satang from input
	baht := amount.Truncate(0)
	satang := amount.Sub(baht).Mul(decimal.NewFromInt(100)).Round(0)

	// Declare text output as empty string
	text := prefix + spellAmountString(baht.String())

	// Add suffix dealing with ถ้วน and สตางค์
	text = addSuffix(satang, text)

	return text, nil
}
