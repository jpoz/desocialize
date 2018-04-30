run:
	export CGO_ENABLED=1
	export PKG_CONFIG_PATH=/usr/local/opt/opencv@2/lib/pkgconfig
	go run cmd/desocialize/main.go IMG_4346.jpg
