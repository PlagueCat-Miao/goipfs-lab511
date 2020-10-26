
build :
	go build main.go -o output/main

runGateway :
	go run main.go -s 2

runCloud :
	go run main.go -s 3

runEdge :
	go run cmd/edge.go

testRtmp :
	go test  -cover -v dal/rtmpffmpeg/r* -test.run TestRmptFfmpegPull -count=1

clear :
	rm -rf output/main
