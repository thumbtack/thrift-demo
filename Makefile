.PHONY: all
all:
	thrift --gen go service.thrift
	# only required if php will implement server, otherwise you can run:
	# thrift --gen php service.thrift
	thrift --gen php:server service.thrift
