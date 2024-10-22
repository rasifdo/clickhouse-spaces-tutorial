build:
	docker build -t clickhouse-demo .     

run: build
	docker run -d --name clickhouse-demo -p 8123:8123 -p 9000:9000 -p 9009:9009 clickhouse-demo

stop :
	docker stop clickhouse-demo

clean: stop
	docker rm clickhouse-demo
