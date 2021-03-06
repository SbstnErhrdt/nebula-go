// Autogenerated by Thrift Compiler (facebook)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// @generated

package storage

import (
	"bytes"
	"sync"
	"fmt"
	thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift"
	nebula0 "github.com/vesoft-inc/nebula-go/nebula"
	meta1 "github.com/vesoft-inc/nebula-go/nebula/meta"

)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = sync.Mutex{}
var _ = bytes.Equal

var _ = nebula0.GoUnusedProtection__
var _ = meta1.GoUnusedProtection__
type GeneralStorageService interface {
  // Parameters:
  //  - Req
  Get(req *KVGetRequest) (r *KVGetResponse, err error)
  // Parameters:
  //  - Req
  Put(req *KVPutRequest) (r *ExecResponse, err error)
  // Parameters:
  //  - Req
  Remove(req *KVRemoveRequest) (r *ExecResponse, err error)
}

type GeneralStorageServiceClient struct {
  Transport thrift.Transport
  ProtocolFactory thrift.ProtocolFactory
  InputProtocol thrift.Protocol
  OutputProtocol thrift.Protocol
  SeqId int32
}

func (client *GeneralStorageServiceClient) Close() error {
  return client.Transport.Close()
}

func NewGeneralStorageServiceClientFactory(t thrift.Transport, f thrift.ProtocolFactory) *GeneralStorageServiceClient {
  return &GeneralStorageServiceClient{Transport: t,
    ProtocolFactory: f,
    InputProtocol: f.GetProtocol(t),
    OutputProtocol: f.GetProtocol(t),
    SeqId: 0,
  }
}

func NewGeneralStorageServiceClient(t thrift.Transport, iprot thrift.Protocol, oprot thrift.Protocol) *GeneralStorageServiceClient {
  return &GeneralStorageServiceClient{Transport: t,
    ProtocolFactory: nil,
    InputProtocol: iprot,
    OutputProtocol: oprot,
    SeqId: 0,
  }
}

// Parameters:
//  - Req
func (p *GeneralStorageServiceClient) Get(req *KVGetRequest) (r *KVGetResponse, err error) {
  if err = p.sendGet(req); err != nil { return }
  return p.recvGet()
}

func (p *GeneralStorageServiceClient) sendGet(req *KVGetRequest)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("get", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := GeneralStorageServiceGetArgs{
  Req : req,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *GeneralStorageServiceClient) recvGet() (value *KVGetResponse, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "get" {
    err = thrift.NewApplicationException(thrift.WRONG_METHOD_NAME, "get failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewApplicationException(thrift.BAD_SEQUENCE_ID, "get failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error365 := thrift.NewApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error366 error
    error366, err = error365.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error366
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "get failed: invalid message type")
    return
  }
  result := GeneralStorageServiceGetResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}

// Parameters:
//  - Req
func (p *GeneralStorageServiceClient) Put(req *KVPutRequest) (r *ExecResponse, err error) {
  if err = p.sendPut(req); err != nil { return }
  return p.recvPut()
}

func (p *GeneralStorageServiceClient) sendPut(req *KVPutRequest)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("put", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := GeneralStorageServicePutArgs{
  Req : req,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *GeneralStorageServiceClient) recvPut() (value *ExecResponse, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "put" {
    err = thrift.NewApplicationException(thrift.WRONG_METHOD_NAME, "put failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewApplicationException(thrift.BAD_SEQUENCE_ID, "put failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error367 := thrift.NewApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error368 error
    error368, err = error367.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error368
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "put failed: invalid message type")
    return
  }
  result := GeneralStorageServicePutResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}

// Parameters:
//  - Req
func (p *GeneralStorageServiceClient) Remove(req *KVRemoveRequest) (r *ExecResponse, err error) {
  if err = p.sendRemove(req); err != nil { return }
  return p.recvRemove()
}

func (p *GeneralStorageServiceClient) sendRemove(req *KVRemoveRequest)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("remove", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := GeneralStorageServiceRemoveArgs{
  Req : req,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *GeneralStorageServiceClient) recvRemove() (value *ExecResponse, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "remove" {
    err = thrift.NewApplicationException(thrift.WRONG_METHOD_NAME, "remove failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewApplicationException(thrift.BAD_SEQUENCE_ID, "remove failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error369 := thrift.NewApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error370 error
    error370, err = error369.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error370
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "remove failed: invalid message type")
    return
  }
  result := GeneralStorageServiceRemoveResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}


type GeneralStorageServiceThreadsafeClient struct {
  Transport thrift.Transport
  ProtocolFactory thrift.ProtocolFactory
  InputProtocol thrift.Protocol
  OutputProtocol thrift.Protocol
  SeqId int32
  Mu sync.Mutex
}

func NewGeneralStorageServiceThreadsafeClientFactory(t thrift.Transport, f thrift.ProtocolFactory) *GeneralStorageServiceThreadsafeClient {
  return &GeneralStorageServiceThreadsafeClient{Transport: t,
    ProtocolFactory: f,
    InputProtocol: f.GetProtocol(t),
    OutputProtocol: f.GetProtocol(t),
    SeqId: 0,
  }
}

func NewGeneralStorageServiceThreadsafeClient(t thrift.Transport, iprot thrift.Protocol, oprot thrift.Protocol) *GeneralStorageServiceThreadsafeClient {
  return &GeneralStorageServiceThreadsafeClient{Transport: t,
    ProtocolFactory: nil,
    InputProtocol: iprot,
    OutputProtocol: oprot,
    SeqId: 0,
  }
}

func (p *GeneralStorageServiceThreadsafeClient) Threadsafe() {}

// Parameters:
//  - Req
func (p *GeneralStorageServiceThreadsafeClient) Get(req *KVGetRequest) (r *KVGetResponse, err error) {
  p.Mu.Lock()
  defer p.Mu.Unlock()
  if err = p.sendGet(req); err != nil { return }
  return p.recvGet()
}

func (p *GeneralStorageServiceThreadsafeClient) sendGet(req *KVGetRequest)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("get", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := GeneralStorageServiceGetArgs{
  Req : req,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *GeneralStorageServiceThreadsafeClient) recvGet() (value *KVGetResponse, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "get" {
    err = thrift.NewApplicationException(thrift.WRONG_METHOD_NAME, "get failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewApplicationException(thrift.BAD_SEQUENCE_ID, "get failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error371 := thrift.NewApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error372 error
    error372, err = error371.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error372
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "get failed: invalid message type")
    return
  }
  result := GeneralStorageServiceGetResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}

// Parameters:
//  - Req
func (p *GeneralStorageServiceThreadsafeClient) Put(req *KVPutRequest) (r *ExecResponse, err error) {
  p.Mu.Lock()
  defer p.Mu.Unlock()
  if err = p.sendPut(req); err != nil { return }
  return p.recvPut()
}

func (p *GeneralStorageServiceThreadsafeClient) sendPut(req *KVPutRequest)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("put", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := GeneralStorageServicePutArgs{
  Req : req,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *GeneralStorageServiceThreadsafeClient) recvPut() (value *ExecResponse, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "put" {
    err = thrift.NewApplicationException(thrift.WRONG_METHOD_NAME, "put failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewApplicationException(thrift.BAD_SEQUENCE_ID, "put failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error373 := thrift.NewApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error374 error
    error374, err = error373.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error374
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "put failed: invalid message type")
    return
  }
  result := GeneralStorageServicePutResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}

// Parameters:
//  - Req
func (p *GeneralStorageServiceThreadsafeClient) Remove(req *KVRemoveRequest) (r *ExecResponse, err error) {
  p.Mu.Lock()
  defer p.Mu.Unlock()
  if err = p.sendRemove(req); err != nil { return }
  return p.recvRemove()
}

func (p *GeneralStorageServiceThreadsafeClient) sendRemove(req *KVRemoveRequest)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("remove", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := GeneralStorageServiceRemoveArgs{
  Req : req,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *GeneralStorageServiceThreadsafeClient) recvRemove() (value *ExecResponse, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "remove" {
    err = thrift.NewApplicationException(thrift.WRONG_METHOD_NAME, "remove failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewApplicationException(thrift.BAD_SEQUENCE_ID, "remove failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error375 := thrift.NewApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error376 error
    error376, err = error375.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error376
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "remove failed: invalid message type")
    return
  }
  result := GeneralStorageServiceRemoveResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}


type GeneralStorageServiceProcessor struct {
  processorMap map[string]thrift.ProcessorFunction
  handler GeneralStorageService
}

func (p *GeneralStorageServiceProcessor) AddToProcessorMap(key string, processor thrift.ProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *GeneralStorageServiceProcessor) GetProcessorFunction(key string) (processor thrift.ProcessorFunction, err error) {
  if processor, ok := p.processorMap[key]; ok {
    return processor, nil
  }
  return nil, nil // generic error message will be sent
}

func (p *GeneralStorageServiceProcessor) ProcessorMap() map[string]thrift.ProcessorFunction {
  return p.processorMap
}

func NewGeneralStorageServiceProcessor(handler GeneralStorageService) *GeneralStorageServiceProcessor {
  self377 := &GeneralStorageServiceProcessor{handler:handler, processorMap:make(map[string]thrift.ProcessorFunction)}
  self377.processorMap["get"] = &generalStorageServiceProcessorGet{handler:handler}
  self377.processorMap["put"] = &generalStorageServiceProcessorPut{handler:handler}
  self377.processorMap["remove"] = &generalStorageServiceProcessorRemove{handler:handler}
  return self377
}

type generalStorageServiceProcessorGet struct {
  handler GeneralStorageService
}

func (p *generalStorageServiceProcessorGet) Read(iprot thrift.Protocol) (thrift.Struct, thrift.Exception) {
  args := GeneralStorageServiceGetArgs{}
  if err := args.Read(iprot); err != nil {
    return nil, err
  }
  iprot.ReadMessageEnd()
  return &args, nil
}

func (p *generalStorageServiceProcessorGet) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Protocol) (err thrift.Exception) {
  var err2 error
  messageType := thrift.REPLY
  switch result.(type) {
  case thrift.ApplicationException:
    messageType = thrift.EXCEPTION
  }
  if err2 = oprot.WriteMessageBegin("get", messageType, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  return err
}

func (p *generalStorageServiceProcessorGet) Run(argStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
  args := argStruct.(*GeneralStorageServiceGetArgs)
  var result GeneralStorageServiceGetResult
  if retval, err := p.handler.Get(args.Req); err != nil {
    switch err.(type) {
    default:
      x := thrift.NewApplicationException(thrift.INTERNAL_ERROR, "Internal error processing get: " + err.Error())
      return x, x
    }
  } else {
    result.Success = retval
  }
  return &result, nil
}

type generalStorageServiceProcessorPut struct {
  handler GeneralStorageService
}

func (p *generalStorageServiceProcessorPut) Read(iprot thrift.Protocol) (thrift.Struct, thrift.Exception) {
  args := GeneralStorageServicePutArgs{}
  if err := args.Read(iprot); err != nil {
    return nil, err
  }
  iprot.ReadMessageEnd()
  return &args, nil
}

func (p *generalStorageServiceProcessorPut) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Protocol) (err thrift.Exception) {
  var err2 error
  messageType := thrift.REPLY
  switch result.(type) {
  case thrift.ApplicationException:
    messageType = thrift.EXCEPTION
  }
  if err2 = oprot.WriteMessageBegin("put", messageType, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  return err
}

func (p *generalStorageServiceProcessorPut) Run(argStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
  args := argStruct.(*GeneralStorageServicePutArgs)
  var result GeneralStorageServicePutResult
  if retval, err := p.handler.Put(args.Req); err != nil {
    switch err.(type) {
    default:
      x := thrift.NewApplicationException(thrift.INTERNAL_ERROR, "Internal error processing put: " + err.Error())
      return x, x
    }
  } else {
    result.Success = retval
  }
  return &result, nil
}

type generalStorageServiceProcessorRemove struct {
  handler GeneralStorageService
}

func (p *generalStorageServiceProcessorRemove) Read(iprot thrift.Protocol) (thrift.Struct, thrift.Exception) {
  args := GeneralStorageServiceRemoveArgs{}
  if err := args.Read(iprot); err != nil {
    return nil, err
  }
  iprot.ReadMessageEnd()
  return &args, nil
}

func (p *generalStorageServiceProcessorRemove) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Protocol) (err thrift.Exception) {
  var err2 error
  messageType := thrift.REPLY
  switch result.(type) {
  case thrift.ApplicationException:
    messageType = thrift.EXCEPTION
  }
  if err2 = oprot.WriteMessageBegin("remove", messageType, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  return err
}

func (p *generalStorageServiceProcessorRemove) Run(argStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
  args := argStruct.(*GeneralStorageServiceRemoveArgs)
  var result GeneralStorageServiceRemoveResult
  if retval, err := p.handler.Remove(args.Req); err != nil {
    switch err.(type) {
    default:
      x := thrift.NewApplicationException(thrift.INTERNAL_ERROR, "Internal error processing remove: " + err.Error())
      return x, x
    }
  } else {
    result.Success = retval
  }
  return &result, nil
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Req
type GeneralStorageServiceGetArgs struct {
  Req *KVGetRequest `thrift:"req,1" db:"req" json:"req"`
}

func NewGeneralStorageServiceGetArgs() *GeneralStorageServiceGetArgs {
  return &GeneralStorageServiceGetArgs{}
}

var GeneralStorageServiceGetArgs_Req_DEFAULT *KVGetRequest
func (p *GeneralStorageServiceGetArgs) GetReq() *KVGetRequest {
  if !p.IsSetReq() {
    return GeneralStorageServiceGetArgs_Req_DEFAULT
  }
return p.Req
}
func (p *GeneralStorageServiceGetArgs) IsSetReq() bool {
  return p.Req != nil
}

func (p *GeneralStorageServiceGetArgs) Read(iprot thrift.Protocol) error {
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
      if err := p.ReadField1(iprot); err != nil {
        return err
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

func (p *GeneralStorageServiceGetArgs)  ReadField1(iprot thrift.Protocol) error {
  p.Req = NewKVGetRequest()
  if err := p.Req.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Req), err)
  }
  return nil
}

func (p *GeneralStorageServiceGetArgs) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("get_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *GeneralStorageServiceGetArgs) writeField1(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:req: ", p), err) }
  if err := p.Req.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Req), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:req: ", p), err) }
  return err
}

func (p *GeneralStorageServiceGetArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GeneralStorageServiceGetArgs(%+v)", *p)
}

// Attributes:
//  - Success
type GeneralStorageServiceGetResult struct {
  Success *KVGetResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewGeneralStorageServiceGetResult() *GeneralStorageServiceGetResult {
  return &GeneralStorageServiceGetResult{}
}

var GeneralStorageServiceGetResult_Success_DEFAULT *KVGetResponse
func (p *GeneralStorageServiceGetResult) GetSuccess() *KVGetResponse {
  if !p.IsSetSuccess() {
    return GeneralStorageServiceGetResult_Success_DEFAULT
  }
return p.Success
}
func (p *GeneralStorageServiceGetResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *GeneralStorageServiceGetResult) Read(iprot thrift.Protocol) error {
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
      if err := p.ReadField0(iprot); err != nil {
        return err
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

func (p *GeneralStorageServiceGetResult)  ReadField0(iprot thrift.Protocol) error {
  p.Success = NewKVGetResponse()
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *GeneralStorageServiceGetResult) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("get_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField0(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *GeneralStorageServiceGetResult) writeField0(oprot thrift.Protocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *GeneralStorageServiceGetResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GeneralStorageServiceGetResult(%+v)", *p)
}

// Attributes:
//  - Req
type GeneralStorageServicePutArgs struct {
  Req *KVPutRequest `thrift:"req,1" db:"req" json:"req"`
}

func NewGeneralStorageServicePutArgs() *GeneralStorageServicePutArgs {
  return &GeneralStorageServicePutArgs{}
}

var GeneralStorageServicePutArgs_Req_DEFAULT *KVPutRequest
func (p *GeneralStorageServicePutArgs) GetReq() *KVPutRequest {
  if !p.IsSetReq() {
    return GeneralStorageServicePutArgs_Req_DEFAULT
  }
return p.Req
}
func (p *GeneralStorageServicePutArgs) IsSetReq() bool {
  return p.Req != nil
}

func (p *GeneralStorageServicePutArgs) Read(iprot thrift.Protocol) error {
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
      if err := p.ReadField1(iprot); err != nil {
        return err
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

func (p *GeneralStorageServicePutArgs)  ReadField1(iprot thrift.Protocol) error {
  p.Req = NewKVPutRequest()
  if err := p.Req.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Req), err)
  }
  return nil
}

func (p *GeneralStorageServicePutArgs) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("put_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *GeneralStorageServicePutArgs) writeField1(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:req: ", p), err) }
  if err := p.Req.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Req), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:req: ", p), err) }
  return err
}

func (p *GeneralStorageServicePutArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GeneralStorageServicePutArgs(%+v)", *p)
}

// Attributes:
//  - Success
type GeneralStorageServicePutResult struct {
  Success *ExecResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewGeneralStorageServicePutResult() *GeneralStorageServicePutResult {
  return &GeneralStorageServicePutResult{}
}

var GeneralStorageServicePutResult_Success_DEFAULT *ExecResponse
func (p *GeneralStorageServicePutResult) GetSuccess() *ExecResponse {
  if !p.IsSetSuccess() {
    return GeneralStorageServicePutResult_Success_DEFAULT
  }
return p.Success
}
func (p *GeneralStorageServicePutResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *GeneralStorageServicePutResult) Read(iprot thrift.Protocol) error {
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
      if err := p.ReadField0(iprot); err != nil {
        return err
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

func (p *GeneralStorageServicePutResult)  ReadField0(iprot thrift.Protocol) error {
  p.Success = NewExecResponse()
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *GeneralStorageServicePutResult) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("put_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField0(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *GeneralStorageServicePutResult) writeField0(oprot thrift.Protocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *GeneralStorageServicePutResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GeneralStorageServicePutResult(%+v)", *p)
}

// Attributes:
//  - Req
type GeneralStorageServiceRemoveArgs struct {
  Req *KVRemoveRequest `thrift:"req,1" db:"req" json:"req"`
}

func NewGeneralStorageServiceRemoveArgs() *GeneralStorageServiceRemoveArgs {
  return &GeneralStorageServiceRemoveArgs{}
}

var GeneralStorageServiceRemoveArgs_Req_DEFAULT *KVRemoveRequest
func (p *GeneralStorageServiceRemoveArgs) GetReq() *KVRemoveRequest {
  if !p.IsSetReq() {
    return GeneralStorageServiceRemoveArgs_Req_DEFAULT
  }
return p.Req
}
func (p *GeneralStorageServiceRemoveArgs) IsSetReq() bool {
  return p.Req != nil
}

func (p *GeneralStorageServiceRemoveArgs) Read(iprot thrift.Protocol) error {
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
      if err := p.ReadField1(iprot); err != nil {
        return err
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

func (p *GeneralStorageServiceRemoveArgs)  ReadField1(iprot thrift.Protocol) error {
  p.Req = NewKVRemoveRequest()
  if err := p.Req.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Req), err)
  }
  return nil
}

func (p *GeneralStorageServiceRemoveArgs) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("remove_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *GeneralStorageServiceRemoveArgs) writeField1(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:req: ", p), err) }
  if err := p.Req.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Req), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:req: ", p), err) }
  return err
}

func (p *GeneralStorageServiceRemoveArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GeneralStorageServiceRemoveArgs(%+v)", *p)
}

// Attributes:
//  - Success
type GeneralStorageServiceRemoveResult struct {
  Success *ExecResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewGeneralStorageServiceRemoveResult() *GeneralStorageServiceRemoveResult {
  return &GeneralStorageServiceRemoveResult{}
}

var GeneralStorageServiceRemoveResult_Success_DEFAULT *ExecResponse
func (p *GeneralStorageServiceRemoveResult) GetSuccess() *ExecResponse {
  if !p.IsSetSuccess() {
    return GeneralStorageServiceRemoveResult_Success_DEFAULT
  }
return p.Success
}
func (p *GeneralStorageServiceRemoveResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *GeneralStorageServiceRemoveResult) Read(iprot thrift.Protocol) error {
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
      if err := p.ReadField0(iprot); err != nil {
        return err
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

func (p *GeneralStorageServiceRemoveResult)  ReadField0(iprot thrift.Protocol) error {
  p.Success = NewExecResponse()
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *GeneralStorageServiceRemoveResult) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("remove_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField0(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *GeneralStorageServiceRemoveResult) writeField0(oprot thrift.Protocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *GeneralStorageServiceRemoveResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GeneralStorageServiceRemoveResult(%+v)", *p)
}


