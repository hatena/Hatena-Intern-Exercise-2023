# Setup

### server の起動

```bash
docker image build -t intern2023/grpc_adventure:latest .

docker container run -p 10000:10000 intern2023/grpc_adventure
```

### grpcurl で動作確認

```bash
# grpcurl の install
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

```bash
# サービスの一覧
grpcurl -plaintext localhost:10000 list

# メソッドや型の詳細
grpcurl -plaintext localhost:10000 describe welcome.Welcome.Greet
grpcurl -plaintext localhost:10000 describe welcome.GreetRequest

# サービスの呼び出し
grpcurl -plaintext -d '{"name": "John Appleseed"}' localhost:10000 welcome.Welcome/Greet
```
