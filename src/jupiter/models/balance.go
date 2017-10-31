package models

type UserBalance struct {
	Total   float32
	Prepaid float32
	Credit  float32
}

func (u *UserBalance) HasEnoughBalance(i *Invoice) bool {
	if i.OptionCashType == CASHTYPE_ALL {
		// all
		if i.TotalPrice > u.Total {
			return false
		}
	} else if i.OptionCashType == CASHTYPE_PREPAID_ONLY {
		// prepaid only
		if i.TotalPrice > u.Prepaid {
			return false
		}
	} else if i.OptionCashType == CASHTYPE_CREDIT_ONLY {
		if i.TotalPrice > u.Credit {
			return false
		}
	}
	return true
}
