package goodpayment

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/goodpayment"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateGoodPayment(info *npool.GoodPayment) error {
	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		return xerrors.Errorf("invalid good id: %v", err)
	}
	if _, err := uuid.Parse(info.GetPaymentCoinTypeID()); err != nil {
		return xerrors.Errorf("invalid coin type id: %v", err)
	}
	if _, err := uuid.Parse(info.GetAccountID()); err != nil {
		return xerrors.Errorf("invalid account id: %v", err)
	}
	return nil
}

func dbRowToGoodPayment(row *ent.GoodPayment) *npool.GoodPayment {
	return &npool.GoodPayment{
		ID:                row.ID.String(),
		GoodID:            row.GoodID.String(),
		PaymentCoinTypeID: row.PaymentCoinTypeID.String(),
		AccountID:         row.AccountID.String(),
		Idle:              row.Idle,
	}
}

func Create(ctx context.Context, in *npool.CreateGoodPaymentRequest) (*npool.CreateGoodPaymentResponse, error) {
	if err := validateGoodPayment(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		GoodPayment.
		Create().
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetPaymentCoinTypeID(uuid.MustParse(in.GetInfo().GetPaymentCoinTypeID())).
		SetAccountID(uuid.MustParse(in.GetInfo().GetAccountID())).
		SetIdle(in.GetInfo().GetIdle()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create good payment: %v", err)
	}

	return &npool.CreateGoodPaymentResponse{
		Info: dbRowToGoodPayment(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateGoodPaymentRequest) (*npool.UpdateGoodPaymentResponse, error) {
	if err := validateGoodPayment(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		GoodPayment.
		UpdateOneID(id).
		SetIdle(in.GetInfo().GetIdle()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update good payment: %v", err)
	}

	return &npool.UpdateGoodPaymentResponse{
		Info: dbRowToGoodPayment(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetGoodPaymentRequest) (*npool.GetGoodPaymentResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		GoodPayment.
		Query().
		Where(
			goodpayment.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query good payment: %v", err)
	}

	var payment *npool.GoodPayment
	for _, info := range infos {
		payment = dbRowToGoodPayment(info)
		break
	}

	return &npool.GetGoodPaymentResponse{
		Info: payment,
	}, nil
}

func GetByAccount(ctx context.Context, in *npool.GetGoodPaymentByAccountRequest) (*npool.GetGoodPaymentByAccountResponse, error) {
	accountID, err := uuid.Parse(in.GetAccountID())
	if err != nil {
		return nil, xerrors.Errorf("invalid account id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		GoodPayment.
		Query().
		Where(
			goodpayment.AccountID(accountID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query good payment: %v", err)
	}

	var payment *npool.GoodPayment
	for _, info := range infos {
		payment = dbRowToGoodPayment(info)
		break
	}

	return &npool.GetGoodPaymentByAccountResponse{
		Info: payment,
	}, nil
}

func GetByGood(ctx context.Context, in *npool.GetGoodPaymentsByGoodRequest) (*npool.GetGoodPaymentsByGoodResponse, error) {
	goodID, err := uuid.Parse(in.GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		GoodPayment.
		Query().
		Where(
			goodpayment.GoodID(goodID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query good payment: %v", err)
	}

	payments := []*npool.GoodPayment{}
	for _, info := range infos {
		payments = append(payments, dbRowToGoodPayment(info))
	}

	return &npool.GetGoodPaymentsByGoodResponse{
		Infos: payments,
	}, nil
}

func GetIdleByGood(ctx context.Context, in *npool.GetIdleGoodPaymentsByGoodRequest) (*npool.GetIdleGoodPaymentsByGoodResponse, error) {
	goodID, err := uuid.Parse(in.GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		GoodPayment.
		Query().
		Where(
			goodpayment.And(
				goodpayment.GoodID(goodID),
				goodpayment.Idle(true),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query good payment: %v", err)
	}

	payments := []*npool.GoodPayment{}
	for _, info := range infos {
		payments = append(payments, dbRowToGoodPayment(info))
	}

	return &npool.GetIdleGoodPaymentsByGoodResponse{
		Infos: payments,
	}, nil
}

func GetIdleByGoodPaymentCoin(ctx context.Context, in *npool.GetIdleGoodPaymentsByGoodPaymentCoinRequest) (*npool.GetIdleGoodPaymentsByGoodPaymentCoinResponse, error) {
	goodID, err := uuid.Parse(in.GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	coinTypeID, err := uuid.Parse(in.GetPaymentCoinTypeID())
	if err != nil {
		return nil, xerrors.Errorf("invalid payment coin type id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		GoodPayment.
		Query().
		Where(
			goodpayment.And(
				goodpayment.GoodID(goodID),
				goodpayment.PaymentCoinTypeID(coinTypeID),
				goodpayment.Idle(true),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query good payment: %v", err)
	}

	payments := []*npool.GoodPayment{}
	for _, info := range infos {
		payments = append(payments, dbRowToGoodPayment(info))
	}

	return &npool.GetIdleGoodPaymentsByGoodPaymentCoinResponse{
		Infos: payments,
	}, nil
}
