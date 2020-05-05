# Go-Structform - Structured Formatters

go-structform provides capabilities for serializing, desirializing, and
transcoding structured data efficiently and generically.

The top-level package provides the common layer by which serializers and
deserializers interact with each other. Serializers convert the input into a
stream of events by calling methods on the
[Visitor](https://pkg.go.dev/github.com/elastic/go-structform?tab=doc#Visitor)
interfaces. The Deserializers implement the Visitor interface to convert the
stream of events into the target structure.

A type implementing the `Visitor` interface, but forwards calls to another
visitor is called a Transformer or Filter. Transformers/Filters can manipulate
the data stream or create/collect errors if some wanted validation fails.

## Data Model

The data model describes by which data types Serializers and Deserializers
interact. In this sense the data model provides a simplified type system, that supports a subset of
serializable go types (for example channels or function pointers can not be serialized).

Go-structform provides a simplified, common data model via the [Visitor](https://pkg.go.dev/github.com/elastic/go-structform?tab=doc#Visitor) interface.

### Types

- **primitives** (See [ValueVisitor](https://pkg.go.dev/github.com/elastic/go-structform?tab=doc#ValueVisitor) interface):  `Bool`, `Byte`, `String`, `Int`, `Int8/16/32/64`, `Uint`, `Uint8/16/32/64`, `Float32`, `Float64`, untyped `Nil`
- **compound**: objects ([ObjectVisitor](https://pkg.go.dev/github.com/elastic/go-structform?tab=doc#ObjectVisitor)), arrays ([ArrayVisitor](https://pkg.go.dev/github.com/elastic/go-structform?tab=doc#ArrayVisitor))

## Extended Data Model

The extended data model provides support for similar types to the `Visitor`
interface, but provides a number of optimizations, allowing users to pass a set
of common go values directly without having to serialize and deserialize those
values.

The extended data model is provided by the
[`ExtVisitor`](https://pkg.go.dev/github.com/elastic/go-structform?tab=doc#ExtVisitor)
interface.

All features in the Extended Data Model can be seamlessly converted to the Data
Model provided by the `Visitor interface.

Deserializers supporting the Extended Data Model must always implement the
common Data Model as is required by the `Visitor` interface.

Serializers wanting to interact with the Extended Data Model should still
accept the `Visitor` interface only and use `EnsureExtVisitor` in order to create an `ExtVisitor`.
`EnsureExtVisitor` wraps the `Visitor` if necessarry, allowing the `Visitor` to
implement only a subset of features in the Extended Data Model.

- **extended primitives**: The Visitor adds support for `[]byte` values as
  strings. The value must be consumed immediately or copied, as the buffer is
  not guaranteed to be stable.
- **slices**: The visitor adds support for slice types, for each supported
  primitive type. For example `[]int8/16/32/64` can be passed as is.
- **map**: The visitor adds support for `map[string]T` types, for each
  supported primitive type. For example `map[string]int8/16/32/64`.


## Data Formats

- JSON: the `json` package provides a JSON parser and JSON serializer. The serializer implements a subset of `ExtVisitor`.
- UBJSON: the `ubjson` packages provides a parser and serializer for Universal Binary JSON.
- CBOR: the `cborl` package supports a compatible subset of CBOR (for example object keys must be strings).
- Go Types: the `gotype` package provides a `Folder` to convert go values into
  a stream of events and an `Unfolder` to apply a stream of events to go
  values.
