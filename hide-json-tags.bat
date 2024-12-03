@echo off
echo Compiling .go files...
go build -o ./bin/
echo Compiling .go files... Done.

echo Executing /bin/uru-databases-2-protobuf-common.exe...
start "" "./bin/uru-databases-2-protobuf-common.exe"
echo Executing /bin/uru-databases-2-protobuf-common.exe... Done.