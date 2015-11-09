// This file was generated by nomdl/codegen.

package main

import (
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __mainPackageInFile_types_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.Type{
		types.MakeStructTypeRef("Pitch",
			[]types.Field{
				types.Field{"X", types.MakePrimitiveTypeRef(types.Float64Kind), false},
				types.Field{"Z", types.MakePrimitiveTypeRef(types.Float64Kind), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{})
	__mainPackageInFile_types_CachedRef = types.RegisterPackage(&p)
}

// Pitch

type Pitch struct {
	_X float64
	_Z float64

	ref *ref.Ref
}

func NewPitch() Pitch {
	return Pitch{
		_X: float64(0),
		_Z: float64(0),

		ref: &ref.Ref{},
	}
}

type PitchDef struct {
	X float64
	Z float64
}

func (def PitchDef) New() Pitch {
	return Pitch{
		_X:  def.X,
		_Z:  def.Z,
		ref: &ref.Ref{},
	}
}

func (s Pitch) Def() (d PitchDef) {
	d.X = s._X
	d.Z = s._Z
	return
}

var __typeRefForPitch types.Type

func (m Pitch) Type() types.Type {
	return __typeRefForPitch
}

func init() {
	__typeRefForPitch = types.MakeTypeRef(__mainPackageInFile_types_CachedRef, 0)
	types.RegisterStruct(__typeRefForPitch, builderForPitch, readerForPitch)
}

func builderForPitch(values []types.Value) types.Value {
	i := 0
	s := Pitch{ref: &ref.Ref{}}
	s._X = float64(values[i].(types.Float64))
	i++
	s._Z = float64(values[i].(types.Float64))
	i++
	return s
}

func readerForPitch(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(Pitch)
	values = append(values, types.Float64(s._X))
	values = append(values, types.Float64(s._Z))
	return values
}

func (s Pitch) Equals(other types.Value) bool {
	return other != nil && __typeRefForPitch.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s Pitch) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s Pitch) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeRefForPitch.Chunks()...)
	return
}

func (s Pitch) ChildValues() (ret []types.Value) {
	ret = append(ret, types.Float64(s._X))
	ret = append(ret, types.Float64(s._Z))
	return
}

func (s Pitch) X() float64 {
	return s._X
}

func (s Pitch) SetX(val float64) Pitch {
	s._X = val
	s.ref = &ref.Ref{}
	return s
}

func (s Pitch) Z() float64 {
	return s._Z
}

func (s Pitch) SetZ(val float64) Pitch {
	s._Z = val
	s.ref = &ref.Ref{}
	return s
}

// MapOfStringToRefOfListOfPitch

type MapOfStringToRefOfListOfPitch struct {
	m   types.Map
	ref *ref.Ref
}

func NewMapOfStringToRefOfListOfPitch() MapOfStringToRefOfListOfPitch {
	return MapOfStringToRefOfListOfPitch{types.NewMap(), &ref.Ref{}}
}

type MapOfStringToRefOfListOfPitchDef map[string]ref.Ref

func (def MapOfStringToRefOfListOfPitchDef) New() MapOfStringToRefOfListOfPitch {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), NewRefOfListOfPitch(v))
	}
	return MapOfStringToRefOfListOfPitch{types.NewMap(kv...), &ref.Ref{}}
}

func (m MapOfStringToRefOfListOfPitch) Def() MapOfStringToRefOfListOfPitchDef {
	def := make(map[string]ref.Ref)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = v.(RefOfListOfPitch).TargetRef()
		return false
	})
	return def
}

func (m MapOfStringToRefOfListOfPitch) Equals(other types.Value) bool {
	return other != nil && __typeRefForMapOfStringToRefOfListOfPitch.Equals(other.Type()) && m.Ref() == other.Ref()
}

func (m MapOfStringToRefOfListOfPitch) Ref() ref.Ref {
	return types.EnsureRef(m.ref, m)
}

func (m MapOfStringToRefOfListOfPitch) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, m.Type().Chunks()...)
	chunks = append(chunks, m.m.Chunks()...)
	return
}

func (m MapOfStringToRefOfListOfPitch) ChildValues() []types.Value {
	return append([]types.Value{}, m.m.ChildValues()...)
}

// A Noms Value that describes MapOfStringToRefOfListOfPitch.
var __typeRefForMapOfStringToRefOfListOfPitch types.Type

func (m MapOfStringToRefOfListOfPitch) Type() types.Type {
	return __typeRefForMapOfStringToRefOfListOfPitch
}

func init() {
	__typeRefForMapOfStringToRefOfListOfPitch = types.MakeCompoundTypeRef(types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakeCompoundTypeRef(types.RefKind, types.MakeCompoundTypeRef(types.ListKind, types.MakeTypeRef(__mainPackageInFile_types_CachedRef, 0))))
	types.RegisterValue(__typeRefForMapOfStringToRefOfListOfPitch, builderForMapOfStringToRefOfListOfPitch, readerForMapOfStringToRefOfListOfPitch)
}

func builderForMapOfStringToRefOfListOfPitch(v types.Value) types.Value {
	return MapOfStringToRefOfListOfPitch{v.(types.Map), &ref.Ref{}}
}

func readerForMapOfStringToRefOfListOfPitch(v types.Value) types.Value {
	return v.(MapOfStringToRefOfListOfPitch).m
}

func (m MapOfStringToRefOfListOfPitch) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToRefOfListOfPitch) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToRefOfListOfPitch) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToRefOfListOfPitch) Get(p string) RefOfListOfPitch {
	return m.m.Get(types.NewString(p)).(RefOfListOfPitch)
}

func (m MapOfStringToRefOfListOfPitch) MaybeGet(p string) (RefOfListOfPitch, bool) {
	v, ok := m.m.MaybeGet(types.NewString(p))
	if !ok {
		return NewRefOfListOfPitch(ref.Ref{}), false
	}
	return v.(RefOfListOfPitch), ok
}

func (m MapOfStringToRefOfListOfPitch) Set(k string, v RefOfListOfPitch) MapOfStringToRefOfListOfPitch {
	return MapOfStringToRefOfListOfPitch{m.m.Set(types.NewString(k), v), &ref.Ref{}}
}

// TODO: Implement SetM?

func (m MapOfStringToRefOfListOfPitch) Remove(p string) MapOfStringToRefOfListOfPitch {
	return MapOfStringToRefOfListOfPitch{m.m.Remove(types.NewString(p)), &ref.Ref{}}
}

type MapOfStringToRefOfListOfPitchIterCallback func(k string, v RefOfListOfPitch) (stop bool)

func (m MapOfStringToRefOfListOfPitch) Iter(cb MapOfStringToRefOfListOfPitchIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(RefOfListOfPitch))
	})
}

type MapOfStringToRefOfListOfPitchIterAllCallback func(k string, v RefOfListOfPitch)

func (m MapOfStringToRefOfListOfPitch) IterAll(cb MapOfStringToRefOfListOfPitchIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), v.(RefOfListOfPitch))
	})
}

func (m MapOfStringToRefOfListOfPitch) IterAllP(concurrency int, cb MapOfStringToRefOfListOfPitchIterAllCallback) {
	m.m.IterAllP(concurrency, func(k, v types.Value) {
		cb(k.(types.String).String(), v.(RefOfListOfPitch))
	})
}

type MapOfStringToRefOfListOfPitchFilterCallback func(k string, v RefOfListOfPitch) (keep bool)

func (m MapOfStringToRefOfListOfPitch) Filter(cb MapOfStringToRefOfListOfPitchFilterCallback) MapOfStringToRefOfListOfPitch {
	out := m.m.Filter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(RefOfListOfPitch))
	})
	return MapOfStringToRefOfListOfPitch{out, &ref.Ref{}}
}

// ListOfRefOfMapOfStringToValue

type ListOfRefOfMapOfStringToValue struct {
	l   types.List
	ref *ref.Ref
}

func NewListOfRefOfMapOfStringToValue() ListOfRefOfMapOfStringToValue {
	return ListOfRefOfMapOfStringToValue{types.NewList(), &ref.Ref{}}
}

type ListOfRefOfMapOfStringToValueDef []ref.Ref

func (def ListOfRefOfMapOfStringToValueDef) New() ListOfRefOfMapOfStringToValue {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = NewRefOfMapOfStringToValue(d)
	}
	return ListOfRefOfMapOfStringToValue{types.NewList(l...), &ref.Ref{}}
}

func (l ListOfRefOfMapOfStringToValue) Def() ListOfRefOfMapOfStringToValueDef {
	d := make([]ref.Ref, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = l.l.Get(i).(RefOfMapOfStringToValue).TargetRef()
	}
	return d
}

func (l ListOfRefOfMapOfStringToValue) Equals(other types.Value) bool {
	return other != nil && __typeRefForListOfRefOfMapOfStringToValue.Equals(other.Type()) && l.Ref() == other.Ref()
}

func (l ListOfRefOfMapOfStringToValue) Ref() ref.Ref {
	return types.EnsureRef(l.ref, l)
}

func (l ListOfRefOfMapOfStringToValue) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, l.Type().Chunks()...)
	chunks = append(chunks, l.l.Chunks()...)
	return
}

func (l ListOfRefOfMapOfStringToValue) ChildValues() []types.Value {
	return append([]types.Value{}, l.l.ChildValues()...)
}

// A Noms Value that describes ListOfRefOfMapOfStringToValue.
var __typeRefForListOfRefOfMapOfStringToValue types.Type

func (m ListOfRefOfMapOfStringToValue) Type() types.Type {
	return __typeRefForListOfRefOfMapOfStringToValue
}

func init() {
	__typeRefForListOfRefOfMapOfStringToValue = types.MakeCompoundTypeRef(types.ListKind, types.MakeCompoundTypeRef(types.RefKind, types.MakeCompoundTypeRef(types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakePrimitiveTypeRef(types.ValueKind))))
	types.RegisterValue(__typeRefForListOfRefOfMapOfStringToValue, builderForListOfRefOfMapOfStringToValue, readerForListOfRefOfMapOfStringToValue)
}

func builderForListOfRefOfMapOfStringToValue(v types.Value) types.Value {
	return ListOfRefOfMapOfStringToValue{v.(types.List), &ref.Ref{}}
}

func readerForListOfRefOfMapOfStringToValue(v types.Value) types.Value {
	return v.(ListOfRefOfMapOfStringToValue).l
}

func (l ListOfRefOfMapOfStringToValue) Len() uint64 {
	return l.l.Len()
}

func (l ListOfRefOfMapOfStringToValue) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfRefOfMapOfStringToValue) Get(i uint64) RefOfMapOfStringToValue {
	return l.l.Get(i).(RefOfMapOfStringToValue)
}

func (l ListOfRefOfMapOfStringToValue) Slice(idx uint64, end uint64) ListOfRefOfMapOfStringToValue {
	return ListOfRefOfMapOfStringToValue{l.l.Slice(idx, end), &ref.Ref{}}
}

func (l ListOfRefOfMapOfStringToValue) Set(i uint64, val RefOfMapOfStringToValue) ListOfRefOfMapOfStringToValue {
	return ListOfRefOfMapOfStringToValue{l.l.Set(i, val), &ref.Ref{}}
}

func (l ListOfRefOfMapOfStringToValue) Append(v ...RefOfMapOfStringToValue) ListOfRefOfMapOfStringToValue {
	return ListOfRefOfMapOfStringToValue{l.l.Append(l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfRefOfMapOfStringToValue) Insert(idx uint64, v ...RefOfMapOfStringToValue) ListOfRefOfMapOfStringToValue {
	return ListOfRefOfMapOfStringToValue{l.l.Insert(idx, l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfRefOfMapOfStringToValue) Remove(idx uint64, end uint64) ListOfRefOfMapOfStringToValue {
	return ListOfRefOfMapOfStringToValue{l.l.Remove(idx, end), &ref.Ref{}}
}

func (l ListOfRefOfMapOfStringToValue) RemoveAt(idx uint64) ListOfRefOfMapOfStringToValue {
	return ListOfRefOfMapOfStringToValue{(l.l.RemoveAt(idx)), &ref.Ref{}}
}

func (l ListOfRefOfMapOfStringToValue) fromElemSlice(p []RefOfMapOfStringToValue) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

type ListOfRefOfMapOfStringToValueIterCallback func(v RefOfMapOfStringToValue, i uint64) (stop bool)

func (l ListOfRefOfMapOfStringToValue) Iter(cb ListOfRefOfMapOfStringToValueIterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(v.(RefOfMapOfStringToValue), i)
	})
}

type ListOfRefOfMapOfStringToValueIterAllCallback func(v RefOfMapOfStringToValue, i uint64)

func (l ListOfRefOfMapOfStringToValue) IterAll(cb ListOfRefOfMapOfStringToValueIterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(v.(RefOfMapOfStringToValue), i)
	})
}

func (l ListOfRefOfMapOfStringToValue) IterAllP(concurrency int, cb ListOfRefOfMapOfStringToValueIterAllCallback) {
	l.l.IterAllP(concurrency, func(v types.Value, i uint64) {
		cb(v.(RefOfMapOfStringToValue), i)
	})
}

type ListOfRefOfMapOfStringToValueFilterCallback func(v RefOfMapOfStringToValue, i uint64) (keep bool)

func (l ListOfRefOfMapOfStringToValue) Filter(cb ListOfRefOfMapOfStringToValueFilterCallback) ListOfRefOfMapOfStringToValue {
	out := l.l.Filter(func(v types.Value, i uint64) bool {
		return cb(v.(RefOfMapOfStringToValue), i)
	})
	return ListOfRefOfMapOfStringToValue{out, &ref.Ref{}}
}

// RefOfMapOfStringToValue

type RefOfMapOfStringToValue struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfMapOfStringToValue(target ref.Ref) RefOfMapOfStringToValue {
	return RefOfMapOfStringToValue{target, &ref.Ref{}}
}

func (r RefOfMapOfStringToValue) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfMapOfStringToValue) Ref() ref.Ref {
	return types.EnsureRef(r.ref, r)
}

func (r RefOfMapOfStringToValue) Equals(other types.Value) bool {
	return other != nil && __typeRefForRefOfMapOfStringToValue.Equals(other.Type()) && r.Ref() == other.Ref()
}

func (r RefOfMapOfStringToValue) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, r.Type().Chunks()...)
	chunks = append(chunks, r.target)
	return
}

func (r RefOfMapOfStringToValue) ChildValues() []types.Value {
	return nil
}

// A Noms Value that describes RefOfMapOfStringToValue.
var __typeRefForRefOfMapOfStringToValue types.Type

func (m RefOfMapOfStringToValue) Type() types.Type {
	return __typeRefForRefOfMapOfStringToValue
}

func init() {
	__typeRefForRefOfMapOfStringToValue = types.MakeCompoundTypeRef(types.RefKind, types.MakeCompoundTypeRef(types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakePrimitiveTypeRef(types.ValueKind)))
	types.RegisterRef(__typeRefForRefOfMapOfStringToValue, builderForRefOfMapOfStringToValue)
}

func builderForRefOfMapOfStringToValue(r ref.Ref) types.Value {
	return NewRefOfMapOfStringToValue(r)
}

func (r RefOfMapOfStringToValue) TargetValue(cs chunks.ChunkSource) MapOfStringToValue {
	return types.ReadValue(r.target, cs).(MapOfStringToValue)
}

func (r RefOfMapOfStringToValue) SetTargetValue(val MapOfStringToValue, cs chunks.ChunkSink) RefOfMapOfStringToValue {
	return NewRefOfMapOfStringToValue(types.WriteValue(val, cs))
}

// MapOfStringToValue

type MapOfStringToValue struct {
	m   types.Map
	ref *ref.Ref
}

func NewMapOfStringToValue() MapOfStringToValue {
	return MapOfStringToValue{types.NewMap(), &ref.Ref{}}
}

type MapOfStringToValueDef map[string]types.Value

func (def MapOfStringToValueDef) New() MapOfStringToValue {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), v)
	}
	return MapOfStringToValue{types.NewMap(kv...), &ref.Ref{}}
}

func (m MapOfStringToValue) Def() MapOfStringToValueDef {
	def := make(map[string]types.Value)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = v
		return false
	})
	return def
}

func (m MapOfStringToValue) Equals(other types.Value) bool {
	return other != nil && __typeRefForMapOfStringToValue.Equals(other.Type()) && m.Ref() == other.Ref()
}

func (m MapOfStringToValue) Ref() ref.Ref {
	return types.EnsureRef(m.ref, m)
}

func (m MapOfStringToValue) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, m.Type().Chunks()...)
	chunks = append(chunks, m.m.Chunks()...)
	return
}

func (m MapOfStringToValue) ChildValues() []types.Value {
	return append([]types.Value{}, m.m.ChildValues()...)
}

// A Noms Value that describes MapOfStringToValue.
var __typeRefForMapOfStringToValue types.Type

func (m MapOfStringToValue) Type() types.Type {
	return __typeRefForMapOfStringToValue
}

func init() {
	__typeRefForMapOfStringToValue = types.MakeCompoundTypeRef(types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakePrimitiveTypeRef(types.ValueKind))
	types.RegisterValue(__typeRefForMapOfStringToValue, builderForMapOfStringToValue, readerForMapOfStringToValue)
}

func builderForMapOfStringToValue(v types.Value) types.Value {
	return MapOfStringToValue{v.(types.Map), &ref.Ref{}}
}

func readerForMapOfStringToValue(v types.Value) types.Value {
	return v.(MapOfStringToValue).m
}

func (m MapOfStringToValue) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToValue) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToValue) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToValue) Get(p string) types.Value {
	return m.m.Get(types.NewString(p))
}

func (m MapOfStringToValue) MaybeGet(p string) (types.Value, bool) {
	v, ok := m.m.MaybeGet(types.NewString(p))
	if !ok {
		return types.Bool(false), false
	}
	return v, ok
}

func (m MapOfStringToValue) Set(k string, v types.Value) MapOfStringToValue {
	return MapOfStringToValue{m.m.Set(types.NewString(k), v), &ref.Ref{}}
}

// TODO: Implement SetM?

func (m MapOfStringToValue) Remove(p string) MapOfStringToValue {
	return MapOfStringToValue{m.m.Remove(types.NewString(p)), &ref.Ref{}}
}

type MapOfStringToValueIterCallback func(k string, v types.Value) (stop bool)

func (m MapOfStringToValue) Iter(cb MapOfStringToValueIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v)
	})
}

type MapOfStringToValueIterAllCallback func(k string, v types.Value)

func (m MapOfStringToValue) IterAll(cb MapOfStringToValueIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), v)
	})
}

func (m MapOfStringToValue) IterAllP(concurrency int, cb MapOfStringToValueIterAllCallback) {
	m.m.IterAllP(concurrency, func(k, v types.Value) {
		cb(k.(types.String).String(), v)
	})
}

type MapOfStringToValueFilterCallback func(k string, v types.Value) (keep bool)

func (m MapOfStringToValue) Filter(cb MapOfStringToValueFilterCallback) MapOfStringToValue {
	out := m.m.Filter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v)
	})
	return MapOfStringToValue{out, &ref.Ref{}}
}

// ListOfPitch

type ListOfPitch struct {
	l   types.List
	ref *ref.Ref
}

func NewListOfPitch() ListOfPitch {
	return ListOfPitch{types.NewList(), &ref.Ref{}}
}

type ListOfPitchDef []PitchDef

func (def ListOfPitchDef) New() ListOfPitch {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = d.New()
	}
	return ListOfPitch{types.NewList(l...), &ref.Ref{}}
}

func (l ListOfPitch) Def() ListOfPitchDef {
	d := make([]PitchDef, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = l.l.Get(i).(Pitch).Def()
	}
	return d
}

func (l ListOfPitch) Equals(other types.Value) bool {
	return other != nil && __typeRefForListOfPitch.Equals(other.Type()) && l.Ref() == other.Ref()
}

func (l ListOfPitch) Ref() ref.Ref {
	return types.EnsureRef(l.ref, l)
}

func (l ListOfPitch) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, l.Type().Chunks()...)
	chunks = append(chunks, l.l.Chunks()...)
	return
}

func (l ListOfPitch) ChildValues() []types.Value {
	return append([]types.Value{}, l.l.ChildValues()...)
}

// A Noms Value that describes ListOfPitch.
var __typeRefForListOfPitch types.Type

func (m ListOfPitch) Type() types.Type {
	return __typeRefForListOfPitch
}

func init() {
	__typeRefForListOfPitch = types.MakeCompoundTypeRef(types.ListKind, types.MakeTypeRef(__mainPackageInFile_types_CachedRef, 0))
	types.RegisterValue(__typeRefForListOfPitch, builderForListOfPitch, readerForListOfPitch)
}

func builderForListOfPitch(v types.Value) types.Value {
	return ListOfPitch{v.(types.List), &ref.Ref{}}
}

func readerForListOfPitch(v types.Value) types.Value {
	return v.(ListOfPitch).l
}

func (l ListOfPitch) Len() uint64 {
	return l.l.Len()
}

func (l ListOfPitch) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfPitch) Get(i uint64) Pitch {
	return l.l.Get(i).(Pitch)
}

func (l ListOfPitch) Slice(idx uint64, end uint64) ListOfPitch {
	return ListOfPitch{l.l.Slice(idx, end), &ref.Ref{}}
}

func (l ListOfPitch) Set(i uint64, val Pitch) ListOfPitch {
	return ListOfPitch{l.l.Set(i, val), &ref.Ref{}}
}

func (l ListOfPitch) Append(v ...Pitch) ListOfPitch {
	return ListOfPitch{l.l.Append(l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfPitch) Insert(idx uint64, v ...Pitch) ListOfPitch {
	return ListOfPitch{l.l.Insert(idx, l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfPitch) Remove(idx uint64, end uint64) ListOfPitch {
	return ListOfPitch{l.l.Remove(idx, end), &ref.Ref{}}
}

func (l ListOfPitch) RemoveAt(idx uint64) ListOfPitch {
	return ListOfPitch{(l.l.RemoveAt(idx)), &ref.Ref{}}
}

func (l ListOfPitch) fromElemSlice(p []Pitch) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

type ListOfPitchIterCallback func(v Pitch, i uint64) (stop bool)

func (l ListOfPitch) Iter(cb ListOfPitchIterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(v.(Pitch), i)
	})
}

type ListOfPitchIterAllCallback func(v Pitch, i uint64)

func (l ListOfPitch) IterAll(cb ListOfPitchIterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(v.(Pitch), i)
	})
}

func (l ListOfPitch) IterAllP(concurrency int, cb ListOfPitchIterAllCallback) {
	l.l.IterAllP(concurrency, func(v types.Value, i uint64) {
		cb(v.(Pitch), i)
	})
}

type ListOfPitchFilterCallback func(v Pitch, i uint64) (keep bool)

func (l ListOfPitch) Filter(cb ListOfPitchFilterCallback) ListOfPitch {
	out := l.l.Filter(func(v types.Value, i uint64) bool {
		return cb(v.(Pitch), i)
	})
	return ListOfPitch{out, &ref.Ref{}}
}

// RefOfListOfPitch

type RefOfListOfPitch struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfListOfPitch(target ref.Ref) RefOfListOfPitch {
	return RefOfListOfPitch{target, &ref.Ref{}}
}

func (r RefOfListOfPitch) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfListOfPitch) Ref() ref.Ref {
	return types.EnsureRef(r.ref, r)
}

func (r RefOfListOfPitch) Equals(other types.Value) bool {
	return other != nil && __typeRefForRefOfListOfPitch.Equals(other.Type()) && r.Ref() == other.Ref()
}

func (r RefOfListOfPitch) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, r.Type().Chunks()...)
	chunks = append(chunks, r.target)
	return
}

func (r RefOfListOfPitch) ChildValues() []types.Value {
	return nil
}

// A Noms Value that describes RefOfListOfPitch.
var __typeRefForRefOfListOfPitch types.Type

func (m RefOfListOfPitch) Type() types.Type {
	return __typeRefForRefOfListOfPitch
}

func init() {
	__typeRefForRefOfListOfPitch = types.MakeCompoundTypeRef(types.RefKind, types.MakeCompoundTypeRef(types.ListKind, types.MakeTypeRef(__mainPackageInFile_types_CachedRef, 0)))
	types.RegisterRef(__typeRefForRefOfListOfPitch, builderForRefOfListOfPitch)
}

func builderForRefOfListOfPitch(r ref.Ref) types.Value {
	return NewRefOfListOfPitch(r)
}

func (r RefOfListOfPitch) TargetValue(cs chunks.ChunkSource) ListOfPitch {
	return types.ReadValue(r.target, cs).(ListOfPitch)
}

func (r RefOfListOfPitch) SetTargetValue(val ListOfPitch, cs chunks.ChunkSink) RefOfListOfPitch {
	return NewRefOfListOfPitch(types.WriteValue(val, cs))
}
