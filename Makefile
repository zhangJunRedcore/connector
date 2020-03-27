build-linux-server:
	mkdir -p connector/shell
	mkdir -p connector/bin
	mkdir -p connector/etc
	chmod +x ./connector/shell/restart_connector.sh
	chmod +x ./connector/shell/run_connector.sh
	chmod +x ./connector/shell/stop_connector.sh
	cp ./conf/conf.yaml ./connector/etc/
	go build main.go
	mv ./main ./connector/bin/connectord && chmod +x ./connector/bin/connectord

start:
	build-linux-server
