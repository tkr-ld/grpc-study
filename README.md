# grpc-weather

## ボイラーテンプレート作成方法

```zsh
protoc --go_out=server/gen/api --go_opt=paths=source_relative \
    --go-grpc_out=server/gen/api --go-grpc_opt=paths=source_relative \
    proto/weather.proto
```

NGな例
```zsh
# --go_outが.で指定されているため、protoと同じ場所に生成される,opt=pathを指定していないため、パッケージのディレクトリ改装になってしまう
protoc --go_out=. \
    --go-grpc_out=. \
    proto/weather.proto

```