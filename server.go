
package main

import (
	"context"
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"os"
	"strings"
	"thriftClient/example"
	"thriftClient/gen-go/base"
	"thriftClient/thrift_unix_domain"
)


type FormatDataImpl struct {}

func (fdi *FormatDataImpl)DoFormat(ctx context.Context,data *example.Data) (r *example.Data, err error){
	var rData example.Data
	fmt.Fprintln(os.Stderr, "Server called with : ", data.Text)
	rData.Text = strings.ToUpper(data.Text)
	return &rData, nil
}



var _ = base.GoUnusedProtection__


func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  StatusCode getStatus(string arg)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func Server() {

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	serverTransport, err := thrift_unix_domain.NewTServerUnixDomain("/tmp/thrift.sock")
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}
	/*handler := &Handler{} //your thrift struct
	processor := thriftMsg.NewThriftMsgProcessor(handler) //your thrift function
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", "/tmp/thrift.sock")
	server.Serve()*/

	handler := &FormatDataImpl{}
	processor := example.NewFormatDataProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on ")
	server.Serve()


	//flag.Usage = Usage
	//var host string
	//var port int
	//var protocol stringfmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
	//var urlString string
	//var framed bool
	//var useHttp bool
	//var parsedUrl *url.URL
	//var trans thrift.TTransport
	//_ = strconv.Atoi
	//_ = math.Abs
	//flag.Usage = Usage
	//flag.StringVar(&host, "h", "localhost", "Specify host and port")
	//flag.IntVar(&port, "p", 9090, "Specify port")
	//flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	//flag.StringVar(&urlString, "u", "", "Specify the url")
	//flag.BoolVar(&framed, "framed", false, "Use framed transport")
	//flag.BoolVar(&useHttp, "http", false, "Use http")
	//flag.Parse()
	//
	//if len(urlString) > 0 {
	//	var err error
	//	parsedUrl, err = url.Parse(urlString)
	//	if err != nil {
	//		fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
	//		flag.Usage()
	//	}
	//	host = parsedUrl.Host
	//	useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	//} else if useHttp {
	//	_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
	//	if err != nil {
	//		fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
	//		flag.Usage()
	//	}
	//}
	//
	//cmd := flag.Arg(0)
	//var err error
	//if useHttp {
	//	trans, err = thrift.NewTHttpClient(parsedUrl.String())
	//} else {
	//	portStr := fmt.Sprint(port)
	//	if strings.Contains(host, ":") {
	//		host, portStr, err = net.SplitHostPort(host)
	//		if err != nil {
	//			fmt.Fprintln(os.Stderr, "error with host:", err)
	//			os.Exit(1)
	//		}
	//	}
	//	trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
	//	if err != nil {
	//		fmt.Fprintln(os.Stderr, "error resolving address:", err)
	//		os.Exit(1)
	//	}
	//	if framed {
	//		trans = thrift.NewTFramedTransport(trans)
	//	}
	//}
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, "Error creating transport", err)
	//	os.Exit(1)
	//}
	//defer trans.Close()
	//var protocolFactory thrift.TProtocolFactory
	//switch protocol {
	//case "compact":
	//	protocolFactory = thrift.NewTCompactProtocolFactory()
	//	break
	//case "simplejson":
	//	protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	//	break
	//case "json":
	//	protocolFactory = thrift.NewTJSONProtocolFactory()
	//	break
	//case "binary", "":
	//	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	//	break
	//default:
	//	fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
	//	Usage()
	//	os.Exit(1)
	//}
	//iprot := protocolFactory.GetProtocol(trans)
	//oprot := protocolFactory.GetProtocol(trans)
	//client := test.NewTestClient(thrift.NewTStandardClient(iprot, oprot))
	//if err := trans.Open(); err != nil {
	//	fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
	//	os.Exit(1)
	//}
	//
	//switch cmd {
	//case "getStatus":
	//	if flag.NArg() - 1 != 1 {
	//		fmt.Fprintln(os.Stderr, "GetStatus requires 1 args")
	//		flag.Usage()
	//	}
	//	argvalue0 := flag.Arg(1)
	//	value0 := argvalue0
	//	fmt.Print(client.GetStatus(context.Background(), value0))
	//	fmt.Print("\n")
	//	break
	//case "":
	//	Usage()
	//	break
	//default:
	//	fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	//}
}