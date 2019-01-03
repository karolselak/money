# `Money`
Manage your finances from the terminal   
![alt-text](./.i.png)

## Install
you need to have go installed, please refer to [Download
Go](https://golang.org/dl)   

To install `money`:
```bash
$ go get -u github.com/mohfunk/money
```

## Usage
```bash
$ money <command> <args ...>
```

## Commands

### list
Lists all your assets
```bash
$ money list # or money l 
```

### add
Adds an asset
```bash
$ money add <Symbol> <Name> # or networth a

# Example
$ money add BTC bitcoin
```

### modify
```bash
#                           +/-
$ money modify <Symbol> <Sign> <Quantity> # or networth mod/m

# Example
$ money m BTC + 2.3
$ money m DCR - 14
```
Modify your holdings.



