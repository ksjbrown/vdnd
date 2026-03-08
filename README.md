# vdnd

A virtual Dungeons and Dragons simulator / tracker / calculator.

## Server

### Build

```sh
go -C server build -o bin/ ./cmd/vdnv
```

### Run

```sh
./server/bin/vdnd
```

### Test

```sh
go -C server test ./...
```

## Client

### Build

```sh
npx --prefix client ng build
```

### Run

```sh
npx --prefix client ng serve
```

### Test

```sh
npx --prefix client ng test
```

