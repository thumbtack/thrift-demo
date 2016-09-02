// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package demo

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type Operation int64

const (
	Operation_ADD Operation = 0
	Operation_SUB Operation = 1
	Operation_MUL Operation = 2
	Operation_DIV Operation = 3
)

func (p Operation) String() string {
	switch p {
	case Operation_ADD:
		return "ADD"
	case Operation_SUB:
		return "SUB"
	case Operation_MUL:
		return "MUL"
	case Operation_DIV:
		return "DIV"
	}
	return "<UNSET>"
}

func OperationFromString(s string) (Operation, error) {
	switch s {
	case "ADD":
		return Operation_ADD, nil
	case "SUB":
		return Operation_SUB, nil
	case "MUL":
		return Operation_MUL, nil
	case "DIV":
		return Operation_DIV, nil
	}
	return Operation(0), fmt.Errorf("not a valid Operation string")
}

func OperationPtr(v Operation) *Operation { return &v }

func (p Operation) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *Operation) UnmarshalText(text []byte) error {
	q, err := OperationFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

// Attributes:
//  - Number1
//  - Number2
//  - Op
type CalculateInput struct {
	Number1 float64   `thrift:"number1,1,required" json:"number1"`
	Number2 float64   `thrift:"number2,2,required" json:"number2"`
	Op      Operation `thrift:"op,3,required" json:"op"`
}

func NewCalculateInput() *CalculateInput {
	return &CalculateInput{}
}

func (p *CalculateInput) GetNumber1() float64 {
	return p.Number1
}

func (p *CalculateInput) GetNumber2() float64 {
	return p.Number2
}

func (p *CalculateInput) GetOp() Operation {
	return p.Op
}
func (p *CalculateInput) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetNumber1 bool = false
	var issetNumber2 bool = false
	var issetOp bool = false

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
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetNumber1 = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetNumber2 = true
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
			issetOp = true
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
	if !issetNumber1 {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Number1 is not set"))
	}
	if !issetNumber2 {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Number2 is not set"))
	}
	if !issetOp {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Op is not set"))
	}
	return nil
}

func (p *CalculateInput) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Number1 = v
	}
	return nil
}

func (p *CalculateInput) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Number2 = v
	}
	return nil
}

func (p *CalculateInput) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		temp := Operation(v)
		p.Op = temp
	}
	return nil
}

func (p *CalculateInput) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("CalculateInput"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *CalculateInput) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("number1", thrift.DOUBLE, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:number1: ", p), err)
	}
	if err := oprot.WriteDouble(float64(p.Number1)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.number1 (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:number1: ", p), err)
	}
	return err
}

func (p *CalculateInput) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("number2", thrift.DOUBLE, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:number2: ", p), err)
	}
	if err := oprot.WriteDouble(float64(p.Number2)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.number2 (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:number2: ", p), err)
	}
	return err
}

func (p *CalculateInput) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("op", thrift.I32, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:op: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Op)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.op (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:op: ", p), err)
	}
	return err
}

func (p *CalculateInput) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CalculateInput(%+v)", *p)
}

// Attributes:
//  - Message
type InputException struct {
	Message string `thrift:"message,1,required" json:"message"`
}

func NewInputException() *InputException {
	return &InputException{}
}

func (p *InputException) GetMessage() string {
	return p.Message
}
func (p *InputException) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetMessage bool = false

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
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetMessage = true
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
	if !issetMessage {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Message is not set"))
	}
	return nil
}

func (p *InputException) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Message = v
	}
	return nil
}

func (p *InputException) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("InputException"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *InputException) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("message", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:message: ", p), err)
	}
	if err := oprot.WriteString(string(p.Message)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.message (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:message: ", p), err)
	}
	return err
}

func (p *InputException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("InputException(%+v)", *p)
}

func (p *InputException) Error() string {
	return p.String()
}