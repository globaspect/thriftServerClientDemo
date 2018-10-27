// Autogenerated by Thrift Compiler (0.11.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package test

import (
	"bytes"
	"reflect"
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"thriftClient/gen-go/base"

)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

var _ = base.GoUnusedProtection__
type Test interface {
  // Parameters:
  //  - Arg
  GetStatus(ctx context.Context, arg string) (r base.StatusCode, err error)
}

type TestClient struct {
  c thrift.TClient
}

// Deprecated: Use NewTest instead
func NewTestClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *TestClient {
  return &TestClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

// Deprecated: Use NewTest instead
func NewTestClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *TestClient {
  return &TestClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewTestClient(c thrift.TClient) *TestClient {
  return &TestClient{
    c: c,
  }
}

// Parameters:
//  - Arg
func (p *TestClient) GetStatus(ctx context.Context, arg string) (r base.StatusCode, err error) {
  var _args0 TestGetStatusArgs
  _args0.Arg = arg
  var _result1 TestGetStatusResult
  if err = p.c.Call(ctx, "getStatus", &_args0, &_result1); err != nil {
    return
  }
  return _result1.GetSuccess(), nil
}

type TestProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler Test
}

func (p *TestProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *TestProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *TestProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewTestProcessor(handler Test) *TestProcessor {

  self2 := &TestProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self2.processorMap["getStatus"] = &testProcessorGetStatus{handler:handler}
return self2
}

func (p *TestProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x3.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush(ctx)
  return false, x3

}

type testProcessorGetStatus struct {
  handler Test
}

func (p *testProcessorGetStatus) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := TestGetStatusArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("getStatus", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := TestGetStatusResult{}
var retval base.StatusCode
  var err2 error
  if retval, err2 = p.handler.GetStatus(ctx, args.Arg); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getStatus: " + err2.Error())
    oprot.WriteMessageBegin("getStatus", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("getStatus", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Arg
type TestGetStatusArgs struct {
  Arg string `thrift:"arg,1" db:"arg" json:"arg"`
}

func NewTestGetStatusArgs() *TestGetStatusArgs {
  return &TestGetStatusArgs{}
}


func (p *TestGetStatusArgs) GetArg() string {
  return p.Arg
}
func (p *TestGetStatusArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TestGetStatusArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Arg = v
}
  return nil
}

func (p *TestGetStatusArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getStatus_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TestGetStatusArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("arg", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:arg: ", p), err) }
  if err := oprot.WriteString(string(p.Arg)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.arg (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:arg: ", p), err) }
  return err
}

func (p *TestGetStatusArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TestGetStatusArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TestGetStatusResult struct {
  Success *base.StatusCode `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewTestGetStatusResult() *TestGetStatusResult {
  return &TestGetStatusResult{}
}

var TestGetStatusResult_Success_DEFAULT base.StatusCode
func (p *TestGetStatusResult) GetSuccess() base.StatusCode {
  if !p.IsSetSuccess() {
    return TestGetStatusResult_Success_DEFAULT
  }
return *p.Success
}
func (p *TestGetStatusResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *TestGetStatusResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TestGetStatusResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  temp := base.StatusCode(v)
  p.Success = &temp
}
  return nil
}

func (p *TestGetStatusResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getStatus_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TestGetStatusResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.I32, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteI32(int32(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *TestGetStatusResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TestGetStatusResult(%+v)", *p)
}


