package user

import (
	"github.com/attic-labs/noms/datastore"
	. "github.com/attic-labs/noms/dbg"
	"github.com/attic-labs/noms/types"
)

func GetUsers(ds datastore.DataStore) types.Set {
	// 1 query (for roots set)
	if ds.Roots().Len() == 0 {
		return types.NewSet()
	} else {
		// BUG 13: We don't ever want to branch the user database. Currently we can't avoid that, but we should change DataStore::Commit() to support that mode of operation.
		Chk.EqualValues(1, ds.Roots().Len())

		// 2 queries Set::Any() -> Map::Get()
		return ds.Roots().Any().(types.Map).Get(types.NewString("value")).(types.Set)
	}
}

func InsertUser(users types.Set, email string) types.Set {
	if GetUser(users, email) != nil {
		return users
	}
	user := types.NewMap(
		types.NewString("$type"), types.NewString("noms.User"),
		types.NewString("email"), types.NewString(email),
		// TODO: Need nil Value so that we can put nil appRoot in here now?
	)
	return users.Insert(user)
}

func CommitUsers(ds datastore.DataStore, users types.Set) datastore.DataStore {
	return ds.Commit(types.NewSet(
		types.NewMap(
			types.NewString("$type"), types.NewString("noms.Root"),
			types.NewString("parents"), ds.Roots(),
			types.NewString("value"), users)))
}

func GetUser(users types.Set, email string) (r types.Map) {
	// n queries
	// could parallelize using go routines. we actually want all the users in memory so we could just get them at GetUsers()
	users.Iter(func(v types.Value) (stop bool) {
		// 0 queries
		if v.(types.Map).Get(types.NewString("email")).Equals(types.NewString(email)) {
			r = v.(types.Map)
			stop = true
		}
		return
	})
	return
}

func GetAppRoot(users types.Set, userEmail string) types.Value {
	user := GetUser(users, userEmail)
	if user == nil {
		return nil
	} else {
		// 1 query
		return user.Get(types.NewString("appRoot"))
	}
}

func SetAppRoot(users types.Set, userEmail string, val types.Value) types.Set {
	user := GetUser(users, userEmail)
	Chk.NotNil(user, "Unknown user: %s", userEmail)
	return users.Remove(user).Insert(user.Set(types.NewString("appRoot"), val))
}
