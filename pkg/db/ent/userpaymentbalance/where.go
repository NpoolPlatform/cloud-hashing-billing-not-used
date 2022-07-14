// Code generated by ent, DO NOT EDIT.

package userpaymentbalance

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// PaymentID applies equality check predicate on the "payment_id" field. It's identical to PaymentIDEQ.
func PaymentID(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaymentID), v))
	})
}

// UsedByPaymentID applies equality check predicate on the "used_by_payment_id" field. It's identical to UsedByPaymentIDEQ.
func UsedByPaymentID(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsedByPaymentID), v))
	})
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v uint64) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// CoinUsdCurrency applies equality check predicate on the "coin_usd_currency" field. It's identical to CoinUsdCurrencyEQ.
func CoinUsdCurrency(v uint64) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinUsdCurrency), v))
	})
}

// CoinTypeID applies equality check predicate on the "coin_type_id" field. It's identical to CoinTypeIDEQ.
func CoinTypeID(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v uint32) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v uint32) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v uint32) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.UserPaymentBalance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.UserPaymentBalance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.UserPaymentBalance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.UserPaymentBalance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// PaymentIDEQ applies the EQ predicate on the "payment_id" field.
func PaymentIDEQ(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaymentID), v))
	})
}

// PaymentIDNEQ applies the NEQ predicate on the "payment_id" field.
func PaymentIDNEQ(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPaymentID), v))
	})
}

// PaymentIDIn applies the In predicate on the "payment_id" field.
func PaymentIDIn(vs ...uuid.UUID) predicate.UserPaymentBalance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPaymentID), v...))
	})
}

// PaymentIDNotIn applies the NotIn predicate on the "payment_id" field.
func PaymentIDNotIn(vs ...uuid.UUID) predicate.UserPaymentBalance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPaymentID), v...))
	})
}

// PaymentIDGT applies the GT predicate on the "payment_id" field.
func PaymentIDGT(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPaymentID), v))
	})
}

// PaymentIDGTE applies the GTE predicate on the "payment_id" field.
func PaymentIDGTE(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPaymentID), v))
	})
}

// PaymentIDLT applies the LT predicate on the "payment_id" field.
func PaymentIDLT(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPaymentID), v))
	})
}

// PaymentIDLTE applies the LTE predicate on the "payment_id" field.
func PaymentIDLTE(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPaymentID), v))
	})
}

// UsedByPaymentIDEQ applies the EQ predicate on the "used_by_payment_id" field.
func UsedByPaymentIDEQ(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsedByPaymentID), v))
	})
}

// UsedByPaymentIDNEQ applies the NEQ predicate on the "used_by_payment_id" field.
func UsedByPaymentIDNEQ(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsedByPaymentID), v))
	})
}

// UsedByPaymentIDIn applies the In predicate on the "used_by_payment_id" field.
func UsedByPaymentIDIn(vs ...uuid.UUID) predicate.UserPaymentBalance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUsedByPaymentID), v...))
	})
}

// UsedByPaymentIDNotIn applies the NotIn predicate on the "used_by_payment_id" field.
func UsedByPaymentIDNotIn(vs ...uuid.UUID) predicate.UserPaymentBalance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUsedByPaymentID), v...))
	})
}

// UsedByPaymentIDGT applies the GT predicate on the "used_by_payment_id" field.
func UsedByPaymentIDGT(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUsedByPaymentID), v))
	})
}

// UsedByPaymentIDGTE applies the GTE predicate on the "used_by_payment_id" field.
func UsedByPaymentIDGTE(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUsedByPaymentID), v))
	})
}

// UsedByPaymentIDLT applies the LT predicate on the "used_by_payment_id" field.
func UsedByPaymentIDLT(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUsedByPaymentID), v))
	})
}

// UsedByPaymentIDLTE applies the LTE predicate on the "used_by_payment_id" field.
func UsedByPaymentIDLTE(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUsedByPaymentID), v))
	})
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v uint64) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v uint64) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...uint64) predicate.UserPaymentBalance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...uint64) predicate.UserPaymentBalance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v uint64) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v uint64) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v uint64) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v uint64) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// CoinUsdCurrencyEQ applies the EQ predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyEQ(v uint64) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinUsdCurrency), v))
	})
}

// CoinUsdCurrencyNEQ applies the NEQ predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyNEQ(v uint64) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinUsdCurrency), v))
	})
}

// CoinUsdCurrencyIn applies the In predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyIn(vs ...uint64) predicate.UserPaymentBalance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCoinUsdCurrency), v...))
	})
}

// CoinUsdCurrencyNotIn applies the NotIn predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyNotIn(vs ...uint64) predicate.UserPaymentBalance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCoinUsdCurrency), v...))
	})
}

// CoinUsdCurrencyGT applies the GT predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyGT(v uint64) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinUsdCurrency), v))
	})
}

// CoinUsdCurrencyGTE applies the GTE predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyGTE(v uint64) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinUsdCurrency), v))
	})
}

// CoinUsdCurrencyLT applies the LT predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyLT(v uint64) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinUsdCurrency), v))
	})
}

// CoinUsdCurrencyLTE applies the LTE predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyLTE(v uint64) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinUsdCurrency), v))
	})
}

// CoinTypeIDEQ applies the EQ predicate on the "coin_type_id" field.
func CoinTypeIDEQ(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDNEQ applies the NEQ predicate on the "coin_type_id" field.
func CoinTypeIDNEQ(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIn applies the In predicate on the "coin_type_id" field.
func CoinTypeIDIn(vs ...uuid.UUID) predicate.UserPaymentBalance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDNotIn applies the NotIn predicate on the "coin_type_id" field.
func CoinTypeIDNotIn(vs ...uuid.UUID) predicate.UserPaymentBalance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDGT applies the GT predicate on the "coin_type_id" field.
func CoinTypeIDGT(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDGTE applies the GTE predicate on the "coin_type_id" field.
func CoinTypeIDGTE(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLT applies the LT predicate on the "coin_type_id" field.
func CoinTypeIDLT(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLTE applies the LTE predicate on the "coin_type_id" field.
func CoinTypeIDLTE(v uuid.UUID) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinTypeID), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v uint32) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v uint32) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...uint32) predicate.UserPaymentBalance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateAt), v...))
	})
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...uint32) predicate.UserPaymentBalance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateAt), v...))
	})
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v uint32) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v uint32) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v uint32) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v uint32) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v uint32) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v uint32) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...uint32) predicate.UserPaymentBalance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...uint32) predicate.UserPaymentBalance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v uint32) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v uint32) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v uint32) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v uint32) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateAt), v))
	})
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v uint32) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v uint32) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...uint32) predicate.UserPaymentBalance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtNotIn applies the NotIn predicate on the "delete_at" field.
func DeleteAtNotIn(vs ...uint32) predicate.UserPaymentBalance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtGT applies the GT predicate on the "delete_at" field.
func DeleteAtGT(v uint32) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v uint32) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v uint32) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v uint32) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserPaymentBalance) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserPaymentBalance) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserPaymentBalance) predicate.UserPaymentBalance {
	return predicate.UserPaymentBalance(func(s *sql.Selector) {
		p(s.Not())
	})
}
