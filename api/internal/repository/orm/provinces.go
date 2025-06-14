// Code generated by SQLBoiler 4.14.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package orm

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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Province is an object representing the database table.
type Province struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Slug      string    `boil:"slug" json:"slug" toml:"slug" yaml:"slug"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *provinceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L provinceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProvinceColumns = struct {
	ID        string
	Name      string
	Slug      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Name:      "name",
	Slug:      "slug",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var ProvinceTableColumns = struct {
	ID        string
	Name      string
	Slug      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "provinces.id",
	Name:      "provinces.name",
	Slug:      "provinces.slug",
	CreatedAt: "provinces.created_at",
	UpdatedAt: "provinces.updated_at",
}

// Generated where

var ProvinceWhere = struct {
	ID        whereHelperint64
	Name      whereHelperstring
	Slug      whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "\"provinces\".\"id\""},
	Name:      whereHelperstring{field: "\"provinces\".\"name\""},
	Slug:      whereHelperstring{field: "\"provinces\".\"slug\""},
	CreatedAt: whereHelpertime_Time{field: "\"provinces\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"provinces\".\"updated_at\""},
}

// ProvinceRels is where relationship names are stored.
var ProvinceRels = struct {
	Places string
}{
	Places: "Places",
}

// provinceR is where relationships are stored.
type provinceR struct {
	Places PlaceSlice `boil:"Places" json:"Places" toml:"Places" yaml:"Places"`
}

// NewStruct creates a new relationship struct
func (*provinceR) NewStruct() *provinceR {
	return &provinceR{}
}

func (r *provinceR) GetPlaces() PlaceSlice {
	if r == nil {
		return nil
	}
	return r.Places
}

// provinceL is where Load methods for each relationship are stored.
type provinceL struct{}

var (
	provinceAllColumns            = []string{"id", "name", "slug", "created_at", "updated_at"}
	provinceColumnsWithoutDefault = []string{"id", "name", "slug"}
	provinceColumnsWithDefault    = []string{"created_at", "updated_at"}
	provincePrimaryKeyColumns     = []string{"id"}
	provinceGeneratedColumns      = []string{}
)

type (
	// ProvinceSlice is an alias for a slice of pointers to Province.
	// This should almost always be used instead of []Province.
	ProvinceSlice []*Province

	provinceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	provinceType                 = reflect.TypeOf(&Province{})
	provinceMapping              = queries.MakeStructMapping(provinceType)
	provincePrimaryKeyMapping, _ = queries.BindMapping(provinceType, provinceMapping, provincePrimaryKeyColumns)
	provinceInsertCacheMut       sync.RWMutex
	provinceInsertCache          = make(map[string]insertCache)
	provinceUpdateCacheMut       sync.RWMutex
	provinceUpdateCache          = make(map[string]updateCache)
	provinceUpsertCacheMut       sync.RWMutex
	provinceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single province record from the query.
func (q provinceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Province, error) {
	o := &Province{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: failed to execute a one query for provinces")
	}

	return o, nil
}

// All returns all Province records from the query.
func (q provinceQuery) All(ctx context.Context, exec boil.ContextExecutor) (ProvinceSlice, error) {
	var o []*Province

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "orm: failed to assign all query results to Province slice")
	}

	return o, nil
}

// Count returns the count of all Province records in the query.
func (q provinceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to count provinces rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q provinceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "orm: failed to check if provinces exists")
	}

	return count > 0, nil
}

// Places retrieves all the place's Places with an executor.
func (o *Province) Places(mods ...qm.QueryMod) placeQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"places\".\"province_id\"=?", o.ID),
	)

	return Places(queryMods...)
}

// LoadPlaces allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (provinceL) LoadPlaces(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProvince interface{}, mods queries.Applicator) error {
	var slice []*Province
	var object *Province

	if singular {
		var ok bool
		object, ok = maybeProvince.(*Province)
		if !ok {
			object = new(Province)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeProvince)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeProvince))
			}
		}
	} else {
		s, ok := maybeProvince.(*[]*Province)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeProvince)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeProvince))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &provinceR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &provinceR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`places`),
		qm.WhereIn(`places.province_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load places")
	}

	var resultSlice []*Place
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice places")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on places")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for places")
	}

	if singular {
		object.R.Places = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &placeR{}
			}
			foreign.R.Province = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ProvinceID {
				local.R.Places = append(local.R.Places, foreign)
				if foreign.R == nil {
					foreign.R = &placeR{}
				}
				foreign.R.Province = local
				break
			}
		}
	}

	return nil
}

// AddPlaces adds the given related objects to the existing relationships
// of the province, optionally inserting them as new records.
// Appends related to o.R.Places.
// Sets related.R.Province appropriately.
func (o *Province) AddPlaces(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Place) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ProvinceID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"places\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"province_id"}),
				strmangle.WhereClause("\"", "\"", 2, placePrimaryKeyColumns),
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

			rel.ProvinceID = o.ID
		}
	}

	if o.R == nil {
		o.R = &provinceR{
			Places: related,
		}
	} else {
		o.R.Places = append(o.R.Places, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &placeR{
				Province: o,
			}
		} else {
			rel.R.Province = o
		}
	}
	return nil
}

// Provinces retrieves all the records using an executor.
func Provinces(mods ...qm.QueryMod) provinceQuery {
	mods = append(mods, qm.From("\"provinces\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"provinces\".*"})
	}

	return provinceQuery{q}
}

// FindProvince retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProvince(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Province, error) {
	provinceObj := &Province{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"provinces\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, provinceObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: unable to select from provinces")
	}

	return provinceObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Province) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no provinces provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(provinceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	provinceInsertCacheMut.RLock()
	cache, cached := provinceInsertCache[key]
	provinceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			provinceAllColumns,
			provinceColumnsWithDefault,
			provinceColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(provinceType, provinceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(provinceType, provinceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"provinces\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"provinces\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "orm: unable to insert into provinces")
	}

	if !cached {
		provinceInsertCacheMut.Lock()
		provinceInsertCache[key] = cache
		provinceInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Province.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Province) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	provinceUpdateCacheMut.RLock()
	cache, cached := provinceUpdateCache[key]
	provinceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			provinceAllColumns,
			provincePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("orm: unable to update provinces, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"provinces\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, provincePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(provinceType, provinceMapping, append(wl, provincePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "orm: unable to update provinces row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by update for provinces")
	}

	if !cached {
		provinceUpdateCacheMut.Lock()
		provinceUpdateCache[key] = cache
		provinceUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q provinceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all for provinces")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected for provinces")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProvinceSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("orm: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), provincePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"provinces\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, provincePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all in province slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected all in update all province")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Province) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no provinces provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(provinceColumnsWithDefault, o)

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

	provinceUpsertCacheMut.RLock()
	cache, cached := provinceUpsertCache[key]
	provinceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			provinceAllColumns,
			provinceColumnsWithDefault,
			provinceColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			provinceAllColumns,
			provincePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("orm: unable to upsert provinces, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(provincePrimaryKeyColumns))
			copy(conflict, provincePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"provinces\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(provinceType, provinceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(provinceType, provinceMapping, ret)
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
		return errors.Wrap(err, "orm: unable to upsert provinces")
	}

	if !cached {
		provinceUpsertCacheMut.Lock()
		provinceUpsertCache[key] = cache
		provinceUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Province record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Province) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("orm: no Province provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), provincePrimaryKeyMapping)
	sql := "DELETE FROM \"provinces\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete from provinces")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by delete for provinces")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q provinceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("orm: no provinceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from provinces")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for provinces")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProvinceSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), provincePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"provinces\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, provincePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from province slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for provinces")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Province) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProvince(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProvinceSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProvinceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), provincePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"provinces\".* FROM \"provinces\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, provincePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "orm: unable to reload all in ProvinceSlice")
	}

	*o = slice

	return nil
}

// ProvinceExists checks if the Province row exists.
func ProvinceExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"provinces\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "orm: unable to check if provinces exists")
	}

	return exists, nil
}

// Exists checks if the Province row exists.
func (o *Province) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ProvinceExists(ctx, exec, o.ID)
}
