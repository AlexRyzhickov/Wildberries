.PHONY: run-redis
run-redis:
	@docker run -p 6379:6379 --name some-redis -d redis

.PHONY: run-telnet
run-telnet:
	 @go run telnet.go alex-G5-5587 6379
