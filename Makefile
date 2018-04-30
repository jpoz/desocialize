run:
	export CGO_ENABLED=1
	export PKG_CONFIG_PATH=/usr/local/opt/opencv@2/lib/pkgconfig
	go build -o desocialize cmd/desocialize/main.go
	./desocialize test/test1.jpg desocialize_test1.jpg
	./desocialize test/test2.jpg desocialize_test2.jpg
	./desocialize test/test3.png desocialize_test3.jpg
	./desocialize test/test4.jpg desocialize_test4.jpg
	./desocialize test/test5.png desocialize_test5.jpg
