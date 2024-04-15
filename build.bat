set GOOS=js
set GOARCH=wasm
go build -o ./output/fingerprint.wasm main.go

cd ./output
anywhere -s

pause