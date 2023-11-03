# How to Use Go API template


## Requirements

To run Go API, you need the following:

- Golang 1.16 or higher
- Running Redis environment
- Running mysql database environment

## Installation

1. Clone the source code from the repository:

```bash
git clone https://github.com/ikierans/asyncfiber.git
cd asyncfiber
```

2. Install the required dependencies
```bash
go mod download
```

3. Run the following command
- run api server
```bash
make s
```
- run worker server
```bash
make w
```
