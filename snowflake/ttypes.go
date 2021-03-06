// Autogenerated by Thrift Compiler (0.9.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package snowflake

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
)

// (needed to ensure safety because of naive import list construction.)
var _ = math.MinInt32
var _ = thrift.ZERO
var _ = fmt.Printf

var GoUnusedProtection__ int

type InvalidSystemClock struct {
	Message string `thrift:"message,1"`
}

func NewInvalidSystemClock() *InvalidSystemClock {
	return &InvalidSystemClock{}
}

func (p *InvalidSystemClock) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *InvalidSystemClock) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.Message = v
	}
	return nil
}

func (p *InvalidSystemClock) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("InvalidSystemClock"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *InvalidSystemClock) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("message", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:message: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Message)); err != nil {
		return fmt.Errorf("%T.message (1) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:message: %s", p, err)
	}
	return err
}

func (p *InvalidSystemClock) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("InvalidSystemClock(%+v)", *p)
}

type InvalidUserAgentError struct {
	Message string `thrift:"message,1"`
}

func NewInvalidUserAgentError() *InvalidUserAgentError {
	return &InvalidUserAgentError{}
}

func (p *InvalidUserAgentError) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *InvalidUserAgentError) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.Message = v
	}
	return nil
}

func (p *InvalidUserAgentError) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("InvalidUserAgentError"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *InvalidUserAgentError) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("message", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:message: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Message)); err != nil {
		return fmt.Errorf("%T.message (1) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:message: %s", p, err)
	}
	return err
}

func (p *InvalidUserAgentError) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("InvalidUserAgentError(%+v)", *p)
}

type AuditLogEntry struct {
	Id        int64  `thrift:"id,1"`
	Useragent string `thrift:"useragent,2"`
	Tag       int64  `thrift:"tag,3"`
}

func NewAuditLogEntry() *AuditLogEntry {
	return &AuditLogEntry{}
}

func (p *AuditLogEntry) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *AuditLogEntry) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.Id = v
	}
	return nil
}

func (p *AuditLogEntry) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s")
	} else {
		p.Useragent = v
	}
	return nil
}

func (p *AuditLogEntry) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 3: %s")
	} else {
		p.Tag = v
	}
	return nil
}

func (p *AuditLogEntry) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("AuditLogEntry"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
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
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *AuditLogEntry) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("id", thrift.I64, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:id: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Id)); err != nil {
		return fmt.Errorf("%T.id (1) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:id: %s", p, err)
	}
	return err
}

func (p *AuditLogEntry) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("useragent", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:useragent: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Useragent)); err != nil {
		return fmt.Errorf("%T.useragent (2) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:useragent: %s", p, err)
	}
	return err
}

func (p *AuditLogEntry) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("tag", thrift.I64, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:tag: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Tag)); err != nil {
		return fmt.Errorf("%T.tag (3) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:tag: %s", p, err)
	}
	return err
}

func (p *AuditLogEntry) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AuditLogEntry(%+v)", *p)
}
