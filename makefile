Generate:
	protoc -I ./protomsg/ ./protomsg/*.proto -I=.  --go_out=plugins=grpc:protomsg
	protoc -I ./protofunc/ ./protofunc/*.proto -I=../ --go_out=plugins=grpc:protofunc


