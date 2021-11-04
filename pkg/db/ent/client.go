// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/coinaccountinfo"

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
		ctx:             ctx,
		config:          cfg,
		CoinAccountInfo: NewCoinAccountInfoClient(cfg),
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
		config:          cfg,
		CoinAccountInfo: NewCoinAccountInfoClient(cfg),
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