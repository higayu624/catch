// Code generated by SQLBoiler 4.18.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Customer is an object representing the database table.
type Customer struct {
	ID                int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt         null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt         null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt         null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Name              null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Email             null.String `boil:"email" json:"email,omitempty" toml:"email" yaml:"email,omitempty"`
	Gender            null.String `boil:"gender" json:"gender,omitempty" toml:"gender" yaml:"gender,omitempty"`
	Age               null.Int64  `boil:"age" json:"age,omitempty" toml:"age" yaml:"age,omitempty"`
	GoogleAccessToken null.String `boil:"google_access_token" json:"google_access_token,omitempty" toml:"google_access_token" yaml:"google_access_token,omitempty"`

	R *customerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L customerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CustomerColumns = struct {
	ID                string
	CreatedAt         string
	UpdatedAt         string
	DeletedAt         string
	Name              string
	Email             string
	Gender            string
	Age               string
	GoogleAccessToken string
}{
	ID:                "id",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
	DeletedAt:         "deleted_at",
	Name:              "name",
	Email:             "email",
	Gender:            "gender",
	Age:               "age",
	GoogleAccessToken: "google_access_token",
}

var CustomerTableColumns = struct {
	ID                string
	CreatedAt         string
	UpdatedAt         string
	DeletedAt         string
	Name              string
	Email             string
	Gender            string
	Age               string
	GoogleAccessToken string
}{
	ID:                "customers.id",
	CreatedAt:         "customers.created_at",
	UpdatedAt:         "customers.updated_at",
	DeletedAt:         "customers.deleted_at",
	Name:              "customers.name",
	Email:             "customers.email",
	Gender:            "customers.gender",
	Age:               "customers.age",
	GoogleAccessToken: "customers.google_access_token",
}

// Generated where

var CustomerWhere = struct {
	ID                whereHelperint64
	CreatedAt         whereHelpernull_Time
	UpdatedAt         whereHelpernull_Time
	DeletedAt         whereHelpernull_Time
	Name              whereHelpernull_String
	Email             whereHelpernull_String
	Gender            whereHelpernull_String
	Age               whereHelpernull_Int64
	GoogleAccessToken whereHelpernull_String
}{
	ID:                whereHelperint64{field: "\"customers\".\"id\""},
	CreatedAt:         whereHelpernull_Time{field: "\"customers\".\"created_at\""},
	UpdatedAt:         whereHelpernull_Time{field: "\"customers\".\"updated_at\""},
	DeletedAt:         whereHelpernull_Time{field: "\"customers\".\"deleted_at\""},
	Name:              whereHelpernull_String{field: "\"customers\".\"name\""},
	Email:             whereHelpernull_String{field: "\"customers\".\"email\""},
	Gender:            whereHelpernull_String{field: "\"customers\".\"gender\""},
	Age:               whereHelpernull_Int64{field: "\"customers\".\"age\""},
	GoogleAccessToken: whereHelpernull_String{field: "\"customers\".\"google_access_token\""},
}

// CustomerRels is where relationship names are stored.
var CustomerRels = struct {
	Stores string
}{
	Stores: "Stores",
}

// customerR is where relationships are stored.
type customerR struct {
	Stores StoreSlice `boil:"Stores" json:"Stores" toml:"Stores" yaml:"Stores"`
}

// NewStruct creates a new relationship struct
func (*customerR) NewStruct() *customerR {
	return &customerR{}
}

func (r *customerR) GetStores() StoreSlice {
	if r == nil {
		return nil
	}
	return r.Stores
}

// customerL is where Load methods for each relationship are stored.
type customerL struct{}

var (
	customerAllColumns            = []string{"id", "created_at", "updated_at", "deleted_at", "name", "email", "gender", "age", "google_access_token"}
	customerColumnsWithoutDefault = []string{}
	customerColumnsWithDefault    = []string{"id", "created_at", "updated_at", "deleted_at", "name", "email", "gender", "age", "google_access_token"}
	customerPrimaryKeyColumns     = []string{"id"}
	customerGeneratedColumns      = []string{}
)

type (
	// CustomerSlice is an alias for a slice of pointers to Customer.
	// This should almost always be used instead of []Customer.
	CustomerSlice []*Customer
	// CustomerHook is the signature for custom Customer hook methods
	CustomerHook func(context.Context, boil.ContextExecutor, *Customer) error

	customerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	customerType                 = reflect.TypeOf(&Customer{})
	customerMapping              = queries.MakeStructMapping(customerType)
	customerPrimaryKeyMapping, _ = queries.BindMapping(customerType, customerMapping, customerPrimaryKeyColumns)
	customerInsertCacheMut       sync.RWMutex
	customerInsertCache          = make(map[string]insertCache)
	customerUpdateCacheMut       sync.RWMutex
	customerUpdateCache          = make(map[string]updateCache)
	customerUpsertCacheMut       sync.RWMutex
	customerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var customerAfterSelectMu sync.Mutex
var customerAfterSelectHooks []CustomerHook

var customerBeforeInsertMu sync.Mutex
var customerBeforeInsertHooks []CustomerHook
var customerAfterInsertMu sync.Mutex
var customerAfterInsertHooks []CustomerHook

var customerBeforeUpdateMu sync.Mutex
var customerBeforeUpdateHooks []CustomerHook
var customerAfterUpdateMu sync.Mutex
var customerAfterUpdateHooks []CustomerHook

var customerBeforeDeleteMu sync.Mutex
var customerBeforeDeleteHooks []CustomerHook
var customerAfterDeleteMu sync.Mutex
var customerAfterDeleteHooks []CustomerHook

var customerBeforeUpsertMu sync.Mutex
var customerBeforeUpsertHooks []CustomerHook
var customerAfterUpsertMu sync.Mutex
var customerAfterUpsertHooks []CustomerHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Customer) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Customer) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Customer) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Customer) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Customer) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Customer) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Customer) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Customer) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Customer) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCustomerHook registers your hook function for all future operations.
func AddCustomerHook(hookPoint boil.HookPoint, customerHook CustomerHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		customerAfterSelectMu.Lock()
		customerAfterSelectHooks = append(customerAfterSelectHooks, customerHook)
		customerAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		customerBeforeInsertMu.Lock()
		customerBeforeInsertHooks = append(customerBeforeInsertHooks, customerHook)
		customerBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		customerAfterInsertMu.Lock()
		customerAfterInsertHooks = append(customerAfterInsertHooks, customerHook)
		customerAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		customerBeforeUpdateMu.Lock()
		customerBeforeUpdateHooks = append(customerBeforeUpdateHooks, customerHook)
		customerBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		customerAfterUpdateMu.Lock()
		customerAfterUpdateHooks = append(customerAfterUpdateHooks, customerHook)
		customerAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		customerBeforeDeleteMu.Lock()
		customerBeforeDeleteHooks = append(customerBeforeDeleteHooks, customerHook)
		customerBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		customerAfterDeleteMu.Lock()
		customerAfterDeleteHooks = append(customerAfterDeleteHooks, customerHook)
		customerAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		customerBeforeUpsertMu.Lock()
		customerBeforeUpsertHooks = append(customerBeforeUpsertHooks, customerHook)
		customerBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		customerAfterUpsertMu.Lock()
		customerAfterUpsertHooks = append(customerAfterUpsertHooks, customerHook)
		customerAfterUpsertMu.Unlock()
	}
}

// One returns a single customer record from the query.
func (q customerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Customer, error) {
	o := &Customer{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for customers")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Customer records from the query.
func (q customerQuery) All(ctx context.Context, exec boil.ContextExecutor) (CustomerSlice, error) {
	var o []*Customer

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Customer slice")
	}

	if len(customerAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Customer records in the query.
func (q customerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count customers rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q customerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if customers exists")
	}

	return count > 0, nil
}

// Stores retrieves all the store's Stores with an executor.
func (o *Customer) Stores(mods ...qm.QueryMod) storeQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"stores\".\"customer_id\"=?", o.ID),
	)

	return Stores(queryMods...)
}

// LoadStores allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (customerL) LoadStores(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCustomer interface{}, mods queries.Applicator) error {
	var slice []*Customer
	var object *Customer

	if singular {
		var ok bool
		object, ok = maybeCustomer.(*Customer)
		if !ok {
			object = new(Customer)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCustomer)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCustomer))
			}
		}
	} else {
		s, ok := maybeCustomer.(*[]*Customer)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCustomer)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCustomer))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &customerR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &customerR{}
			}
			args[obj.ID] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`stores`),
		qm.WhereIn(`stores.customer_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load stores")
	}

	var resultSlice []*Store
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice stores")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on stores")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for stores")
	}

	if len(storeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Stores = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &storeR{}
			}
			foreign.R.Customer = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.CustomerID) {
				local.R.Stores = append(local.R.Stores, foreign)
				if foreign.R == nil {
					foreign.R = &storeR{}
				}
				foreign.R.Customer = local
				break
			}
		}
	}

	return nil
}

// AddStores adds the given related objects to the existing relationships
// of the customer, optionally inserting them as new records.
// Appends related to o.R.Stores.
// Sets related.R.Customer appropriately.
func (o *Customer) AddStores(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Store) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.CustomerID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"stores\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"customer_id"}),
				strmangle.WhereClause("\"", "\"", 2, storePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.CustomerID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &customerR{
			Stores: related,
		}
	} else {
		o.R.Stores = append(o.R.Stores, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &storeR{
				Customer: o,
			}
		} else {
			rel.R.Customer = o
		}
	}
	return nil
}

// SetStores removes all previously related items of the
// customer replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Customer's Stores accordingly.
// Replaces o.R.Stores with related.
// Sets related.R.Customer's Stores accordingly.
func (o *Customer) SetStores(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Store) error {
	query := "update \"stores\" set \"customer_id\" = null where \"customer_id\" = $1"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.Stores {
			queries.SetScanner(&rel.CustomerID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Customer = nil
		}
		o.R.Stores = nil
	}

	return o.AddStores(ctx, exec, insert, related...)
}

// RemoveStores relationships from objects passed in.
// Removes related items from R.Stores (uses pointer comparison, removal does not keep order)
// Sets related.R.Customer.
func (o *Customer) RemoveStores(ctx context.Context, exec boil.ContextExecutor, related ...*Store) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.CustomerID, nil)
		if rel.R != nil {
			rel.R.Customer = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("customer_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Stores {
			if rel != ri {
				continue
			}

			ln := len(o.R.Stores)
			if ln > 1 && i < ln-1 {
				o.R.Stores[i] = o.R.Stores[ln-1]
			}
			o.R.Stores = o.R.Stores[:ln-1]
			break
		}
	}

	return nil
}

// Customers retrieves all the records using an executor.
func Customers(mods ...qm.QueryMod) customerQuery {
	mods = append(mods, qm.From("\"customers\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"customers\".*"})
	}

	return customerQuery{q}
}

// FindCustomer retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCustomer(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Customer, error) {
	customerObj := &Customer{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"customers\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, customerObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from customers")
	}

	if err = customerObj.doAfterSelectHooks(ctx, exec); err != nil {
		return customerObj, err
	}

	return customerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Customer) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no customers provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(customerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	customerInsertCacheMut.RLock()
	cache, cached := customerInsertCache[key]
	customerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			customerAllColumns,
			customerColumnsWithDefault,
			customerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(customerType, customerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(customerType, customerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"customers\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"customers\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into customers")
	}

	if !cached {
		customerInsertCacheMut.Lock()
		customerInsertCache[key] = cache
		customerInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Customer.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Customer) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	customerUpdateCacheMut.RLock()
	cache, cached := customerUpdateCache[key]
	customerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			customerAllColumns,
			customerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update customers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"customers\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, customerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(customerType, customerMapping, append(wl, customerPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update customers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for customers")
	}

	if !cached {
		customerUpdateCacheMut.Lock()
		customerUpdateCache[key] = cache
		customerUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q customerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for customers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for customers")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CustomerSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), customerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"customers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, customerPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in customer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all customer")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Customer) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no customers provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(customerColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	customerUpsertCacheMut.RLock()
	cache, cached := customerUpsertCache[key]
	customerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			customerAllColumns,
			customerColumnsWithDefault,
			customerColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			customerAllColumns,
			customerPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert customers, could not build update column list")
		}

		ret := strmangle.SetComplement(customerAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(customerPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert customers, could not build conflict column list")
			}

			conflict = make([]string, len(customerPrimaryKeyColumns))
			copy(conflict, customerPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"customers\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(customerType, customerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(customerType, customerMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert customers")
	}

	if !cached {
		customerUpsertCacheMut.Lock()
		customerUpsertCache[key] = cache
		customerUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Customer record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Customer) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Customer provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), customerPrimaryKeyMapping)
	sql := "DELETE FROM \"customers\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from customers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for customers")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q customerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no customerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from customers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for customers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CustomerSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(customerBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), customerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"customers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, customerPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from customer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for customers")
	}

	if len(customerAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Customer) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCustomer(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CustomerSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CustomerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), customerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"customers\".* FROM \"customers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, customerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CustomerSlice")
	}

	*o = slice

	return nil
}

// CustomerExists checks if the Customer row exists.
func CustomerExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"customers\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if customers exists")
	}

	return exists, nil
}

// Exists checks if the Customer row exists.
func (o *Customer) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return CustomerExists(ctx, exec, o.ID)
}
