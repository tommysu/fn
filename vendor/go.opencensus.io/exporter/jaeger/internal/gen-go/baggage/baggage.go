// Autogenerated by Thrift Compiler (0.11.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package baggage

import (
	"bytes"
	"context"
	"fmt"
	"reflect"

	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

// Attributes:
//  - BaggageKey
//  - MaxValueLength
type BaggageRestriction struct {
	BaggageKey     string `thrift:"baggageKey,1,required" db:"baggageKey" json:"baggageKey"`
	MaxValueLength int32  `thrift:"maxValueLength,2,required" db:"maxValueLength" json:"maxValueLength"`
}

func NewBaggageRestriction() *BaggageRestriction {
	return &BaggageRestriction{}
}

func (p *BaggageRestriction) GetBaggageKey() string {
	return p.BaggageKey
}

func (p *BaggageRestriction) GetMaxValueLength() int32 {
	return p.MaxValueLength
}
func (p *BaggageRestriction) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetBaggageKey bool = false
	var issetMaxValueLength bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
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
			issetBaggageKey = true
		case 2:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField2(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
			issetMaxValueLength = true
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
	if !issetBaggageKey {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field BaggageKey is not set"))
	}
	if !issetMaxValueLength {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field MaxValueLength is not set"))
	}
	return nil
}

func (p *BaggageRestriction) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.BaggageKey = v
	}
	return nil
}

func (p *BaggageRestriction) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.MaxValueLength = v
	}
	return nil
}

func (p *BaggageRestriction) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("BaggageRestriction"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *BaggageRestriction) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("baggageKey", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:baggageKey: ", p), err)
	}
	if err := oprot.WriteString(string(p.BaggageKey)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.baggageKey (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:baggageKey: ", p), err)
	}
	return err
}

func (p *BaggageRestriction) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("maxValueLength", thrift.I32, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:maxValueLength: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.MaxValueLength)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.maxValueLength (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:maxValueLength: ", p), err)
	}
	return err
}

func (p *BaggageRestriction) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BaggageRestriction(%+v)", *p)
}

type BaggageRestrictionManager interface {
	// getBaggageRestrictions retrieves the baggage restrictions for a specific service.
	// Usually, baggageRestrictions apply to all services however there may be situations
	// where a baggageKey might only be allowed to be set by a specific service.
	//
	// Parameters:
	//  - ServiceName
	GetBaggageRestrictions(ctx context.Context, serviceName string) (r []*BaggageRestriction, err error)
}

type BaggageRestrictionManagerClient struct {
	c thrift.TClient
}

// Deprecated: Use NewBaggageRestrictionManager instead
func NewBaggageRestrictionManagerClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *BaggageRestrictionManagerClient {
	return &BaggageRestrictionManagerClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

// Deprecated: Use NewBaggageRestrictionManager instead
func NewBaggageRestrictionManagerClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *BaggageRestrictionManagerClient {
	return &BaggageRestrictionManagerClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewBaggageRestrictionManagerClient(c thrift.TClient) *BaggageRestrictionManagerClient {
	return &BaggageRestrictionManagerClient{
		c: c,
	}
}

// getBaggageRestrictions retrieves the baggage restrictions for a specific service.
// Usually, baggageRestrictions apply to all services however there may be situations
// where a baggageKey might only be allowed to be set by a specific service.
//
// Parameters:
//  - ServiceName
func (p *BaggageRestrictionManagerClient) GetBaggageRestrictions(ctx context.Context, serviceName string) (r []*BaggageRestriction, err error) {
	var _args0 BaggageRestrictionManagerGetBaggageRestrictionsArgs
	_args0.ServiceName = serviceName
	var _result1 BaggageRestrictionManagerGetBaggageRestrictionsResult
	if err = p.c.Call(ctx, "getBaggageRestrictions", &_args0, &_result1); err != nil {
		return
	}
	return _result1.GetSuccess(), nil
}

type BaggageRestrictionManagerProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      BaggageRestrictionManager
}

func (p *BaggageRestrictionManagerProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *BaggageRestrictionManagerProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *BaggageRestrictionManagerProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewBaggageRestrictionManagerProcessor(handler BaggageRestrictionManager) *BaggageRestrictionManagerProcessor {

	self2 := &BaggageRestrictionManagerProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self2.processorMap["getBaggageRestrictions"] = &baggageRestrictionManagerProcessorGetBaggageRestrictions{handler: handler}
	return self2
}

func (p *BaggageRestrictionManagerProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x3.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x3

}

type baggageRestrictionManagerProcessorGetBaggageRestrictions struct {
	handler BaggageRestrictionManager
}

func (p *baggageRestrictionManagerProcessorGetBaggageRestrictions) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := BaggageRestrictionManagerGetBaggageRestrictionsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getBaggageRestrictions", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := BaggageRestrictionManagerGetBaggageRestrictionsResult{}
	var retval []*BaggageRestriction
	var err2 error
	if retval, err2 = p.handler.GetBaggageRestrictions(ctx, args.ServiceName); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getBaggageRestrictions: "+err2.Error())
		oprot.WriteMessageBegin("getBaggageRestrictions", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getBaggageRestrictions", thrift.REPLY, seqId); err2 != nil {
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
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - ServiceName
type BaggageRestrictionManagerGetBaggageRestrictionsArgs struct {
	ServiceName string `thrift:"serviceName,1" db:"serviceName" json:"serviceName"`
}

func NewBaggageRestrictionManagerGetBaggageRestrictionsArgs() *BaggageRestrictionManagerGetBaggageRestrictionsArgs {
	return &BaggageRestrictionManagerGetBaggageRestrictionsArgs{}
}

func (p *BaggageRestrictionManagerGetBaggageRestrictionsArgs) GetServiceName() string {
	return p.ServiceName
}
func (p *BaggageRestrictionManagerGetBaggageRestrictionsArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
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

func (p *BaggageRestrictionManagerGetBaggageRestrictionsArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ServiceName = v
	}
	return nil
}

func (p *BaggageRestrictionManagerGetBaggageRestrictionsArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getBaggageRestrictions_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *BaggageRestrictionManagerGetBaggageRestrictionsArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("serviceName", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:serviceName: ", p), err)
	}
	if err := oprot.WriteString(string(p.ServiceName)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.serviceName (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:serviceName: ", p), err)
	}
	return err
}

func (p *BaggageRestrictionManagerGetBaggageRestrictionsArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BaggageRestrictionManagerGetBaggageRestrictionsArgs(%+v)", *p)
}

// Attributes:
//  - Success
type BaggageRestrictionManagerGetBaggageRestrictionsResult struct {
	Success []*BaggageRestriction `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewBaggageRestrictionManagerGetBaggageRestrictionsResult() *BaggageRestrictionManagerGetBaggageRestrictionsResult {
	return &BaggageRestrictionManagerGetBaggageRestrictionsResult{}
}

var BaggageRestrictionManagerGetBaggageRestrictionsResult_Success_DEFAULT []*BaggageRestriction

func (p *BaggageRestrictionManagerGetBaggageRestrictionsResult) GetSuccess() []*BaggageRestriction {
	return p.Success
}
func (p *BaggageRestrictionManagerGetBaggageRestrictionsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *BaggageRestrictionManagerGetBaggageRestrictionsResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.LIST {
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

func (p *BaggageRestrictionManagerGetBaggageRestrictionsResult) ReadField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*BaggageRestriction, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		_elem4 := &BaggageRestriction{}
		if err := _elem4.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem4), err)
		}
		p.Success = append(p.Success, _elem4)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *BaggageRestrictionManagerGetBaggageRestrictionsResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getBaggageRestrictions_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *BaggageRestrictionManagerGetBaggageRestrictionsResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Success)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Success {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *BaggageRestrictionManagerGetBaggageRestrictionsResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BaggageRestrictionManagerGetBaggageRestrictionsResult(%+v)", *p)
}