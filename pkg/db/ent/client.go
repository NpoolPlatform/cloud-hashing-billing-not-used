// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/coinaccountinfo"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/coinaccounttransaction"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/platformbenefit"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/platformsetting"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userbenefit"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// CoinAccountInfo is the client for interacting with the CoinAccountInfo builders.
	CoinAccountInfo *CoinAccountInfoClient
	// CoinAccountTransaction is the client for interacting with the CoinAccountTransaction builders.
	CoinAccountTransaction *CoinAccountTransactionClient
	// PlatformBenefit is the client for interacting with the PlatformBenefit builders.
	PlatformBenefit *PlatformBenefitClient
	// PlatformSetting is the client for interacting with the PlatformSetting builders.
	PlatformSetting *PlatformSettingClient
	// UserBenefit is the client for interacting with the UserBenefit builders.
	UserBenefit *UserBenefitClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.CoinAccountInfo = NewCoinAccountInfoClient(c.config)
	c.CoinAccountTransaction = NewCoinAccountTransactionClient(c.config)
	c.PlatformBenefit = NewPlatformBenefitClient(c.config)
	c.PlatformSetting = NewPlatformSettingClient(c.config)
	c.UserBenefit = NewUserBenefitClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                    ctx,
		config:                 cfg,
		CoinAccountInfo:        NewCoinAccountInfoClient(cfg),
		CoinAccountTransaction: NewCoinAccountTransactionClient(cfg),
		PlatformBenefit:        NewPlatformBenefitClient(cfg),
		PlatformSetting:        NewPlatformSettingClient(cfg),
		UserBenefit:            NewUserBenefitClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:                    ctx,
		config:                 cfg,
		CoinAccountInfo:        NewCoinAccountInfoClient(cfg),
		CoinAccountTransaction: NewCoinAccountTransactionClient(cfg),
		PlatformBenefit:        NewPlatformBenefitClient(cfg),
		PlatformSetting:        NewPlatformSettingClient(cfg),
		UserBenefit:            NewUserBenefitClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		CoinAccountInfo.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.CoinAccountInfo.Use(hooks...)
	c.CoinAccountTransaction.Use(hooks...)
	c.PlatformBenefit.Use(hooks...)
	c.PlatformSetting.Use(hooks...)
	c.UserBenefit.Use(hooks...)
}

// CoinAccountInfoClient is a client for the CoinAccountInfo schema.
type CoinAccountInfoClient struct {
	config
}

// NewCoinAccountInfoClient returns a client for the CoinAccountInfo from the given config.
func NewCoinAccountInfoClient(c config) *CoinAccountInfoClient {
	return &CoinAccountInfoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `coinaccountinfo.Hooks(f(g(h())))`.
func (c *CoinAccountInfoClient) Use(hooks ...Hook) {
	c.hooks.CoinAccountInfo = append(c.hooks.CoinAccountInfo, hooks...)
}

// Create returns a create builder for CoinAccountInfo.
func (c *CoinAccountInfoClient) Create() *CoinAccountInfoCreate {
	mutation := newCoinAccountInfoMutation(c.config, OpCreate)
	return &CoinAccountInfoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CoinAccountInfo entities.
func (c *CoinAccountInfoClient) CreateBulk(builders ...*CoinAccountInfoCreate) *CoinAccountInfoCreateBulk {
	return &CoinAccountInfoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CoinAccountInfo.
func (c *CoinAccountInfoClient) Update() *CoinAccountInfoUpdate {
	mutation := newCoinAccountInfoMutation(c.config, OpUpdate)
	return &CoinAccountInfoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CoinAccountInfoClient) UpdateOne(cai *CoinAccountInfo) *CoinAccountInfoUpdateOne {
	mutation := newCoinAccountInfoMutation(c.config, OpUpdateOne, withCoinAccountInfo(cai))
	return &CoinAccountInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CoinAccountInfoClient) UpdateOneID(id uuid.UUID) *CoinAccountInfoUpdateOne {
	mutation := newCoinAccountInfoMutation(c.config, OpUpdateOne, withCoinAccountInfoID(id))
	return &CoinAccountInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CoinAccountInfo.
func (c *CoinAccountInfoClient) Delete() *CoinAccountInfoDelete {
	mutation := newCoinAccountInfoMutation(c.config, OpDelete)
	return &CoinAccountInfoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CoinAccountInfoClient) DeleteOne(cai *CoinAccountInfo) *CoinAccountInfoDeleteOne {
	return c.DeleteOneID(cai.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CoinAccountInfoClient) DeleteOneID(id uuid.UUID) *CoinAccountInfoDeleteOne {
	builder := c.Delete().Where(coinaccountinfo.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CoinAccountInfoDeleteOne{builder}
}

// Query returns a query builder for CoinAccountInfo.
func (c *CoinAccountInfoClient) Query() *CoinAccountInfoQuery {
	return &CoinAccountInfoQuery{
		config: c.config,
	}
}

// Get returns a CoinAccountInfo entity by its id.
func (c *CoinAccountInfoClient) Get(ctx context.Context, id uuid.UUID) (*CoinAccountInfo, error) {
	return c.Query().Where(coinaccountinfo.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CoinAccountInfoClient) GetX(ctx context.Context, id uuid.UUID) *CoinAccountInfo {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CoinAccountInfoClient) Hooks() []Hook {
	return c.hooks.CoinAccountInfo
}

// CoinAccountTransactionClient is a client for the CoinAccountTransaction schema.
type CoinAccountTransactionClient struct {
	config
}

// NewCoinAccountTransactionClient returns a client for the CoinAccountTransaction from the given config.
func NewCoinAccountTransactionClient(c config) *CoinAccountTransactionClient {
	return &CoinAccountTransactionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `coinaccounttransaction.Hooks(f(g(h())))`.
func (c *CoinAccountTransactionClient) Use(hooks ...Hook) {
	c.hooks.CoinAccountTransaction = append(c.hooks.CoinAccountTransaction, hooks...)
}

// Create returns a create builder for CoinAccountTransaction.
func (c *CoinAccountTransactionClient) Create() *CoinAccountTransactionCreate {
	mutation := newCoinAccountTransactionMutation(c.config, OpCreate)
	return &CoinAccountTransactionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CoinAccountTransaction entities.
func (c *CoinAccountTransactionClient) CreateBulk(builders ...*CoinAccountTransactionCreate) *CoinAccountTransactionCreateBulk {
	return &CoinAccountTransactionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CoinAccountTransaction.
func (c *CoinAccountTransactionClient) Update() *CoinAccountTransactionUpdate {
	mutation := newCoinAccountTransactionMutation(c.config, OpUpdate)
	return &CoinAccountTransactionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CoinAccountTransactionClient) UpdateOne(cat *CoinAccountTransaction) *CoinAccountTransactionUpdateOne {
	mutation := newCoinAccountTransactionMutation(c.config, OpUpdateOne, withCoinAccountTransaction(cat))
	return &CoinAccountTransactionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CoinAccountTransactionClient) UpdateOneID(id uuid.UUID) *CoinAccountTransactionUpdateOne {
	mutation := newCoinAccountTransactionMutation(c.config, OpUpdateOne, withCoinAccountTransactionID(id))
	return &CoinAccountTransactionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CoinAccountTransaction.
func (c *CoinAccountTransactionClient) Delete() *CoinAccountTransactionDelete {
	mutation := newCoinAccountTransactionMutation(c.config, OpDelete)
	return &CoinAccountTransactionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CoinAccountTransactionClient) DeleteOne(cat *CoinAccountTransaction) *CoinAccountTransactionDeleteOne {
	return c.DeleteOneID(cat.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CoinAccountTransactionClient) DeleteOneID(id uuid.UUID) *CoinAccountTransactionDeleteOne {
	builder := c.Delete().Where(coinaccounttransaction.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CoinAccountTransactionDeleteOne{builder}
}

// Query returns a query builder for CoinAccountTransaction.
func (c *CoinAccountTransactionClient) Query() *CoinAccountTransactionQuery {
	return &CoinAccountTransactionQuery{
		config: c.config,
	}
}

// Get returns a CoinAccountTransaction entity by its id.
func (c *CoinAccountTransactionClient) Get(ctx context.Context, id uuid.UUID) (*CoinAccountTransaction, error) {
	return c.Query().Where(coinaccounttransaction.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CoinAccountTransactionClient) GetX(ctx context.Context, id uuid.UUID) *CoinAccountTransaction {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CoinAccountTransactionClient) Hooks() []Hook {
	return c.hooks.CoinAccountTransaction
}

// PlatformBenefitClient is a client for the PlatformBenefit schema.
type PlatformBenefitClient struct {
	config
}

// NewPlatformBenefitClient returns a client for the PlatformBenefit from the given config.
func NewPlatformBenefitClient(c config) *PlatformBenefitClient {
	return &PlatformBenefitClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `platformbenefit.Hooks(f(g(h())))`.
func (c *PlatformBenefitClient) Use(hooks ...Hook) {
	c.hooks.PlatformBenefit = append(c.hooks.PlatformBenefit, hooks...)
}

// Create returns a create builder for PlatformBenefit.
func (c *PlatformBenefitClient) Create() *PlatformBenefitCreate {
	mutation := newPlatformBenefitMutation(c.config, OpCreate)
	return &PlatformBenefitCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PlatformBenefit entities.
func (c *PlatformBenefitClient) CreateBulk(builders ...*PlatformBenefitCreate) *PlatformBenefitCreateBulk {
	return &PlatformBenefitCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PlatformBenefit.
func (c *PlatformBenefitClient) Update() *PlatformBenefitUpdate {
	mutation := newPlatformBenefitMutation(c.config, OpUpdate)
	return &PlatformBenefitUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PlatformBenefitClient) UpdateOne(pb *PlatformBenefit) *PlatformBenefitUpdateOne {
	mutation := newPlatformBenefitMutation(c.config, OpUpdateOne, withPlatformBenefit(pb))
	return &PlatformBenefitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PlatformBenefitClient) UpdateOneID(id uuid.UUID) *PlatformBenefitUpdateOne {
	mutation := newPlatformBenefitMutation(c.config, OpUpdateOne, withPlatformBenefitID(id))
	return &PlatformBenefitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PlatformBenefit.
func (c *PlatformBenefitClient) Delete() *PlatformBenefitDelete {
	mutation := newPlatformBenefitMutation(c.config, OpDelete)
	return &PlatformBenefitDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PlatformBenefitClient) DeleteOne(pb *PlatformBenefit) *PlatformBenefitDeleteOne {
	return c.DeleteOneID(pb.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PlatformBenefitClient) DeleteOneID(id uuid.UUID) *PlatformBenefitDeleteOne {
	builder := c.Delete().Where(platformbenefit.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PlatformBenefitDeleteOne{builder}
}

// Query returns a query builder for PlatformBenefit.
func (c *PlatformBenefitClient) Query() *PlatformBenefitQuery {
	return &PlatformBenefitQuery{
		config: c.config,
	}
}

// Get returns a PlatformBenefit entity by its id.
func (c *PlatformBenefitClient) Get(ctx context.Context, id uuid.UUID) (*PlatformBenefit, error) {
	return c.Query().Where(platformbenefit.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PlatformBenefitClient) GetX(ctx context.Context, id uuid.UUID) *PlatformBenefit {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PlatformBenefitClient) Hooks() []Hook {
	return c.hooks.PlatformBenefit
}

// PlatformSettingClient is a client for the PlatformSetting schema.
type PlatformSettingClient struct {
	config
}

// NewPlatformSettingClient returns a client for the PlatformSetting from the given config.
func NewPlatformSettingClient(c config) *PlatformSettingClient {
	return &PlatformSettingClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `platformsetting.Hooks(f(g(h())))`.
func (c *PlatformSettingClient) Use(hooks ...Hook) {
	c.hooks.PlatformSetting = append(c.hooks.PlatformSetting, hooks...)
}

// Create returns a create builder for PlatformSetting.
func (c *PlatformSettingClient) Create() *PlatformSettingCreate {
	mutation := newPlatformSettingMutation(c.config, OpCreate)
	return &PlatformSettingCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PlatformSetting entities.
func (c *PlatformSettingClient) CreateBulk(builders ...*PlatformSettingCreate) *PlatformSettingCreateBulk {
	return &PlatformSettingCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PlatformSetting.
func (c *PlatformSettingClient) Update() *PlatformSettingUpdate {
	mutation := newPlatformSettingMutation(c.config, OpUpdate)
	return &PlatformSettingUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PlatformSettingClient) UpdateOne(ps *PlatformSetting) *PlatformSettingUpdateOne {
	mutation := newPlatformSettingMutation(c.config, OpUpdateOne, withPlatformSetting(ps))
	return &PlatformSettingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PlatformSettingClient) UpdateOneID(id uuid.UUID) *PlatformSettingUpdateOne {
	mutation := newPlatformSettingMutation(c.config, OpUpdateOne, withPlatformSettingID(id))
	return &PlatformSettingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PlatformSetting.
func (c *PlatformSettingClient) Delete() *PlatformSettingDelete {
	mutation := newPlatformSettingMutation(c.config, OpDelete)
	return &PlatformSettingDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PlatformSettingClient) DeleteOne(ps *PlatformSetting) *PlatformSettingDeleteOne {
	return c.DeleteOneID(ps.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PlatformSettingClient) DeleteOneID(id uuid.UUID) *PlatformSettingDeleteOne {
	builder := c.Delete().Where(platformsetting.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PlatformSettingDeleteOne{builder}
}

// Query returns a query builder for PlatformSetting.
func (c *PlatformSettingClient) Query() *PlatformSettingQuery {
	return &PlatformSettingQuery{
		config: c.config,
	}
}

// Get returns a PlatformSetting entity by its id.
func (c *PlatformSettingClient) Get(ctx context.Context, id uuid.UUID) (*PlatformSetting, error) {
	return c.Query().Where(platformsetting.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PlatformSettingClient) GetX(ctx context.Context, id uuid.UUID) *PlatformSetting {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PlatformSettingClient) Hooks() []Hook {
	return c.hooks.PlatformSetting
}

// UserBenefitClient is a client for the UserBenefit schema.
type UserBenefitClient struct {
	config
}

// NewUserBenefitClient returns a client for the UserBenefit from the given config.
func NewUserBenefitClient(c config) *UserBenefitClient {
	return &UserBenefitClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `userbenefit.Hooks(f(g(h())))`.
func (c *UserBenefitClient) Use(hooks ...Hook) {
	c.hooks.UserBenefit = append(c.hooks.UserBenefit, hooks...)
}

// Create returns a create builder for UserBenefit.
func (c *UserBenefitClient) Create() *UserBenefitCreate {
	mutation := newUserBenefitMutation(c.config, OpCreate)
	return &UserBenefitCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserBenefit entities.
func (c *UserBenefitClient) CreateBulk(builders ...*UserBenefitCreate) *UserBenefitCreateBulk {
	return &UserBenefitCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserBenefit.
func (c *UserBenefitClient) Update() *UserBenefitUpdate {
	mutation := newUserBenefitMutation(c.config, OpUpdate)
	return &UserBenefitUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserBenefitClient) UpdateOne(ub *UserBenefit) *UserBenefitUpdateOne {
	mutation := newUserBenefitMutation(c.config, OpUpdateOne, withUserBenefit(ub))
	return &UserBenefitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserBenefitClient) UpdateOneID(id uuid.UUID) *UserBenefitUpdateOne {
	mutation := newUserBenefitMutation(c.config, OpUpdateOne, withUserBenefitID(id))
	return &UserBenefitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserBenefit.
func (c *UserBenefitClient) Delete() *UserBenefitDelete {
	mutation := newUserBenefitMutation(c.config, OpDelete)
	return &UserBenefitDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserBenefitClient) DeleteOne(ub *UserBenefit) *UserBenefitDeleteOne {
	return c.DeleteOneID(ub.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserBenefitClient) DeleteOneID(id uuid.UUID) *UserBenefitDeleteOne {
	builder := c.Delete().Where(userbenefit.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserBenefitDeleteOne{builder}
}

// Query returns a query builder for UserBenefit.
func (c *UserBenefitClient) Query() *UserBenefitQuery {
	return &UserBenefitQuery{
		config: c.config,
	}
}

// Get returns a UserBenefit entity by its id.
func (c *UserBenefitClient) Get(ctx context.Context, id uuid.UUID) (*UserBenefit, error) {
	return c.Query().Where(userbenefit.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserBenefitClient) GetX(ctx context.Context, id uuid.UUID) *UserBenefit {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserBenefitClient) Hooks() []Hook {
	return c.hooks.UserBenefit
}
