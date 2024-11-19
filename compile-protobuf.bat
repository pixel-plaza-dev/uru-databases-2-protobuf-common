@echo off
echo Compiling protoc files...

for %%f in (protobuf/*.proto) do (
    echo Compiling %%f...
    protoc --go_out=./protobuf --go_opt=paths=import --go-grpc_out=./protobuf --go-grpc_opt=paths=import protobuf/%%f
)

echo Compilation complete.

