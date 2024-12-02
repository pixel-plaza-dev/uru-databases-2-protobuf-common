@echo off

echo Compiling Pixel Plaza .proto files...
for %%f in (proto/pixel_plaza/*.proto ) do (
    protoc --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import proto/pixel_plaza/%%f
)
echo Compiling Pixel Plaza .proto files... Done.
