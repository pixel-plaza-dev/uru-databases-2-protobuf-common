@echo off
echo Compiling protoc files...

for %%f in (protobuf/service/*.proto) do (
    echo Compiling %%f...
    protoc --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import protobuf/service/%%f
)

echo Compilation complete.
