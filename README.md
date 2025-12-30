# boba-party-kit

![boba-party-kit](https://target.scene7.com/is/image/Target/GUEST_14845f22-401e-492f-b498-f9a5979e855b?wid=400&hei=400&qlt=80)

Config file only [bubbletea](https://github.com/charmbracelet/bubbletea) TUI framework/generator

# Concepts

## Named searches

It's possible to have multiple searches in a single `config.yaml`.

This makes sense when you're searching for one resource, find it and then want to search through its child resources.

In this case, we want to define multiple _named searches_ in our `config.yaml` file under the `search:` key so we can
reference one search from another in a given search's `select:` directive.

# Concepts

## `package datadriver`

Various drivers for fetching data. All structs in this package implement the following interface:

```go
type (
  datadriver interface {
    Fetch() ([]byte, error)
  }
)
```

- `datadriver.HTTP`: makes an HTTP request and returns its response body as `[]byte`
- `datadriver.File`: reads a file on disk and returns its contents as `[]byte`

## `package dbdriver`

Drivers for fetching data from databases:

- `dbdriver.SQLite`: provides the ability to read/write data in an SQLite database.

## `package formatdriver`

Various drivers for parsing and formatting structured data. 

- `datadriver.TableJSON`: formats `[]byte` data into rows and columns, returns `([][]string, []string, error)`
- `datadriver.TableCSV`: formats `[]byte` data in rows and columns, returns `([][]string, []string, error)`

Implements the following interface:

```go
type ( 
  datadriver interface {
    Format() ([][]string, []string, error)
  }
)
```

# Components

## Search

A fuzzy-search component.

```yaml
search:
  init: # data to load on start up
  select: # action to take when a search result is selected
  results: # how to format search results
    list:
    table:

```

## Tree

A multi-node, nested tree.

## Modal

An overlay/modal component.
