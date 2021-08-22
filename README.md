# Installation

First, make sure you have Go properly installed and setup.

```bash
git clone https://github.com/bopher/bopher
cd bopher
go install .
```

## Create New Project

```bash
bopher new myApp
```

And configure what you want!

## Usage

For library usage see [bopher](https://github.com/bopher) libraries docs.

### Access App Dependencies

For accessing app dependencies (config driver, cache driver, etc.) you must use `app` namespace functions.

```go
// github.com/myapp is your app namespace
import "github.com/myapp/src/appp"
app.Config()
```
