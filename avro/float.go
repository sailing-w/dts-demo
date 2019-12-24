// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     record.avsc
 */
package avro

import (
	"io"
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)

  
type Float struct {

	
	
		Value float64
	

	
	
		Precision int32
	

	
	
		Scale int32
	

}

func NewFloat() (*Float) {
	return &Float{}
}

func DeserializeFloat(r io.Reader) (*Float, error) {
	t := NewFloat()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err	
	}
	return t, err
}

func DeserializeFloatFromSchema(r io.Reader, schema string) (*Float, error) {
	t := NewFloat()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err	
	}
	return t, err
}

func writeFloat(r *Float, w io.Writer) error {
	var err error
	
	err = vm.WriteDouble( r.Value, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteInt( r.Precision, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteInt( r.Scale, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *Float) Serialize(w io.Writer) error {
	return writeFloat(r, w)
}

func (r *Float) Schema() string {
	return "{\"fields\":[{\"name\":\"value\",\"type\":\"double\"},{\"name\":\"precision\",\"type\":\"int\"},{\"name\":\"scale\",\"type\":\"int\"}],\"name\":\"Float\",\"namespace\":\"com.alibaba.dts.formats.avro\",\"type\":\"record\"}"
}

func (r *Float) SchemaName() string {
	return "com.alibaba.dts.formats.avro.Float"
}

func (_ *Float) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *Float) SetInt(v int32) { panic("Unsupported operation") }
func (_ *Float) SetLong(v int64) { panic("Unsupported operation") }
func (_ *Float) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *Float) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *Float) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *Float) SetString(v string) { panic("Unsupported operation") }
func (_ *Float) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Float) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*types.Double)(&r.Value)
		
	
	case 1:
		
		
			return (*types.Int)(&r.Precision)
		
	
	case 2:
		
		
			return (*types.Int)(&r.Scale)
		
	
	}
	panic("Unknown field index")
}

func (r *Float) SetDefault(i int) {
	switch (i) {
	
        
	
        
	
        
	
	}
	panic("Unknown field index")
}

func (_ *Float) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Float) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *Float) Finalize() { }
