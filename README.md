# boba-party-kit

![boba-party-kit](https://target.scene7.com/is/image/Target/GUEST_14845f22-401e-492f-b498-f9a5979e855b?wid=400&hei=400&qlt=80)

Config file only [bubbletea](https://github.com/charmbracelet/bubbletea) TUI generator

# Concepts

## Named searches

It's possible to have multiple searches in a single `config.yaml`.

This makes sense when you're searching for one resource, find it and then want to search through its child resources.

In this case, we want to define multiple _named searches_ in our `config.yaml` file under the `search:` key so we can
reference one search from another in a given search's `select:` directive.

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
