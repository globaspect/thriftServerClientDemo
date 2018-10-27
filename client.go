package main

import (
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"log"
	"os"
	"thriftClient/example"
	"thriftClient/thrift_unix_domain"
)


// https://www.jianshu.com/p/a58665a38022

func Client() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err := thrift_unix_domain. NewTUnixDomain("/tmp/thrift.sock")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}
	useTransport, err := transportFactory.GetTransport(transport)
	client := example.NewFormatDataClientFactory(useTransport, protocolFactory)
	if err := transport.Open();
		err != nil { log.Fatalln("Error opening:") }
	defer transport.Close()

	data := example.Data{Text:"hello,world!"}
	var ctx = context.Background()
	d, err := client.DoFormat(ctx,&data)
	fmt.Println(d.Text)

}

