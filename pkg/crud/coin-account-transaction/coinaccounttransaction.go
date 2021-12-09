package coinaccounttransaction

import (
	"context"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-billing/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/const"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/coinaccounttransaction" //nolint

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateCoinAccountTransaction(info *npool.CoinAccountTransaction) error {
	if _, err := uuid.Parse(info.CoinTypeID); err != nil {
		return xerrors.Errorf("invalid coin type id: %v", err)
	}
	if _, err := uuid.Parse(info.AppID); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.UserID); err != nil {
		return xerrors.Errorf("invalid user id: %v", err)
	}
	if _, err := uuid.Parse(info.FromAddressID); err != nil {
		return xerrors.Errorf("invalid from address id: %v", err)
	}
	if _, err := uuid.Parse(info.ToAddressID); err != nil {
		return xerrors.Errorf("invalid to address id: %v", err)
	}
	if _, err := uuid.Parse(info.PlatformTransactionID); err != nil {
		return xerrors.Errorf("invalid platform transaction id: %v", err)
	}
	return nil
}

func dbRowToCoinAccountTransaction(row *ent.CoinAccountTransaction) *npool.CoinAccountTransaction {
	return &npool.CoinAccountTransaction{
		ID:                    row.ID.String(),
		UserID:                row.UserID.String(),
		AppID:                 row.AppID.String(),
		FromAddressID:         row.FromAddressID.String(),
		ToAddressID:           row.ToAddressID.String(),
		CoinTypeID:            row.CoinTypeID.String(),
		Amount:                price.DBPriceToVisualPrice(row.Amount),
		Message:               row.Message,
		ChainTransactionID:    row.ChainTransactionID,
		PlatformTransactionID: row.PlatformTransactionID.String(),
		CreateAt:              row.CreateAt,
		State:                 string(row.State),
	}
}

func Create(ctx context.Context, in *npool.CreateCoinAccountTransactionRequest) (*npool.CreateCoinAccountTransactionResponse, error) {
	if err := validateCoinAccountTransaction(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		CoinAccountTransaction.
		Create().
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetFromAddressID(uuid.MustParse(in.GetInfo().GetFromAddressID())).
		SetToAddressID(uuid.MustParse(in.GetInfo().GetToAddressID())).
		SetCoinTypeID(uuid.MustParse(in.GetInfo().GetCoinTypeID())).
		SetAmount(price.VisualPriceToDBPrice(in.GetInfo().GetAmount())).
		SetMessage(in.GetInfo().GetMessage()).
		SetChainTransactionID(in.GetInfo().GetChainTransactionID()).
		SetPlatformTransactionID(uuid.MustParse(in.GetInfo().GetPlatformTransactionID())).
		SetState(constant.CoinTransactionStateWait).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create coin account: %v", err)
	}

	return &npool.CreateCoinAccountTransactionResponse{
		Info: dbRowToCoinAccountTransaction(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetCoinAccountTransactionRequest) (*npool.GetCoinAccountTransactionResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid coin account transaction id: %v", err)
	}

	infos, err := db.Client().
		CoinAccountTransaction.
		Query().
		Where(
			coinaccounttransaction.And(
				coinaccounttransaction.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query coin account transaction: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty coin account transaction")
	}

	return &npool.GetCoinAccountTransactionResponse{
		Info: dbRowToCoinAccountTransaction(infos[0]),
	}, nil
}

func GetCoinAccountTransactionsByCoinAccount(ctx context.Context, in *npool.GetCoinAccountTransactionsByCoinAccountRequest) (*npool.GetCoinAccountTransactionsByCoinAccountResponse, error) {
	if _, err := uuid.Parse(in.GetCoinTypeID()); err != nil {
		return nil, xerrors.Errorf("invalid coin type id: %v", err)
	}
	if _, err := uuid.Parse(in.GetAddressID()); err != nil {
		return nil, xerrors.Errorf("invalid address id: %v", err)
	}

	infos, err := db.Client().
		CoinAccountTransaction.
		Query().
		Where(
			coinaccounttransaction.And(
				coinaccounttransaction.CoinTypeID(uuid.MustParse(in.GetCoinTypeID())),
				coinaccounttransaction.Or(
					coinaccounttransaction.FromAddressID(uuid.MustParse(in.GetAddressID())),
					coinaccounttransaction.ToAddressID(uuid.MustParse(in.GetAddressID())),
				),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query coin account transaction: %v", err)
	}

	transactions := []*npool.CoinAccountTransaction{}
	for _, info := range infos {
		transactions = append(transactions, dbRowToCoinAccountTransaction(info))
	}

	return &npool.GetCoinAccountTransactionsByCoinAccountResponse{
		Infos: transactions,
	}, nil
}

func GetCoinAccountTransactionsByState(ctx context.Context, in *npool.GetCoinAccountTransactionsByStateRequest) (*npool.GetCoinAccountTransactionsByStateResponse, error) {
	infos, err := db.Client().
		CoinAccountTransaction.
		Query().
		Where(
			coinaccounttransaction.And(
				coinaccounttransaction.StateEQ(coinaccounttransaction.State(in.GetState())),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query coin account transaction: %v", err)
	}

	transactions := []*npool.CoinAccountTransaction{}
	for _, info := range infos {
		transactions = append(transactions, dbRowToCoinAccountTransaction(info))
	}

	return &npool.GetCoinAccountTransactionsByStateResponse{
		Infos: transactions,
	}, nil
}

func GetCoinAccountTransactionsByCoin(ctx context.Context, in *npool.GetCoinAccountTransactionsByCoinRequest) (*npool.GetCoinAccountTransactionsByCoinResponse, error) {
	if _, err := uuid.Parse(in.GetCoinTypeID()); err != nil {
		return nil, xerrors.Errorf("invalid coin type id: %v", err)
	}

	infos, err := db.Client().
		CoinAccountTransaction.
		Query().
		Where(
			coinaccounttransaction.And(
				coinaccounttransaction.CoinTypeID(uuid.MustParse(in.GetCoinTypeID())),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query coin transaction: %v", err)
	}

	transactions := []*npool.CoinAccountTransaction{}
	for _, info := range infos {
		transactions = append(transactions, dbRowToCoinAccountTransaction(info))
	}

	return &npool.GetCoinAccountTransactionsByCoinResponse{
		Infos: transactions,
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateCoinAccountTransactionRequest) (*npool.UpdateCoinAccountTransactionResponse, error) {
	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	if err := validateCoinAccountTransaction(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		CoinAccountTransaction.
		UpdateOneID(uuid.MustParse(in.GetInfo().GetID())).
		SetState(coinaccounttransaction.State(in.GetInfo().GetState())).
		SetChainTransactionID(in.GetInfo().GetChainTransactionID()).
		SetPlatformTransactionID(uuid.MustParse(in.GetInfo().GetPlatformTransactionID())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update coin account: %v", err)
	}

	return &npool.UpdateCoinAccountTransactionResponse{
		Info: dbRowToCoinAccountTransaction(info),
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteCoinAccountTransactionRequest) (*npool.DeleteCoinAccountTransactionResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	info, err := db.Client().
		CoinAccountTransaction.
		UpdateOneID(uuid.MustParse(in.GetID())).
		SetDeleteAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail delete coin account: %v", err)
	}

	return &npool.DeleteCoinAccountTransactionResponse{
		Info: dbRowToCoinAccountTransaction(info),
	}, nil
}
