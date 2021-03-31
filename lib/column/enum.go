package column

import (
	"fmt"
	"github.com/ClickHouse/clickhouse-go/lib/binary"
	"strings"
	"unicode"
)

type Enum struct {
	iv map[string]interface{}
	vi map[interface{}]string
	base
	baseType interface{}
}

func (enum *Enum) Read(decoder *binary.Decoder, isNull bool) (interface{}, error) {
	var (
		err   error
		ident interface{}
	)
	switch enum.baseType.(type) {
	case int16:
		if ident, err = decoder.Int16(); err != nil {
			return nil, err
		}
	default:
		if ident, err = decoder.Int8(); err != nil {
			return nil, err
		}
	}
	if ident, found := enum.vi[ident]; found || isNull {
		return ident, nil
	}
	return nil, fmt.Errorf("invalid Enum value: %v", ident)
}

func (enum *Enum) Write(encoder *binary.Encoder, v interface{}) error {
	switch v := v.(type) {
	case string:
		ident, found := enum.iv[v]
		if !found {
			return fmt.Errorf("invalid Enum ident: %s", v)
		}
		switch ident := ident.(type) {
		case int8:
			return encoder.Int8(ident)
		case int16:
			return encoder.Int16(ident)
		}
	case uint8:
		if _, ok := enum.baseType.(int8); ok {
			return encoder.Int8(int8(v))
		}
	case int8:
		if _, ok := enum.baseType.(int8); ok {
			return encoder.Int8(v)
		}
	case uint16:
		if _, ok := enum.baseType.(int16); ok {
			return encoder.Int16(int16(v))
		}
	case int16:
		if _, ok := enum.baseType.(int16); ok {
			return encoder.Int16(v)
		}
	case int64:
		switch enum.baseType.(type) {
		case int8:
			return encoder.Int8(int8(v))
		case int16:
			return encoder.Int16(int16(v))
		}
	}
	return &ErrUnexpectedType{
		T:      v,
		Column: enum,
	}
}

func (enum *Enum) defaultValue() interface{} {
	return enum.baseType
}

func parseEnum(name, chType string) (*Enum, error) {
	var (
		data     string
		isEnum16 bool
	)
	if len(chType) < 8 {
		return nil, fmt.Errorf("invalid Enum format: %s", chType)
	}
	switch {
	case strings.HasPrefix(chType, "Enum8"):
		data = chType[6:]
	case strings.HasPrefix(chType, "Enum16"):
		data = chType[7:]
		isEnum16 = true
	default:
		return nil, fmt.Errorf("'%s' is not Enum type", chType)
	}
	enum := Enum{
		base: base{
			name:    name,
			chType:  chType,
			valueOf: columnBaseTypes[string("")],
		},
		iv: make(map[string]interface{}),
		vi: make(map[interface{}]string),
	}

	var err = ParseEnum([]byte(data[:len(data)-1]), func(str []byte, num int) {
		var (
			ident = string(str)
			value = int64(num)
		)
		{
			var (
				ident             = ident[1 : len(ident)-1]
				value interface{} = int16(value)
			)
			if !isEnum16 {
				value = int8(value.(int16))
			}
			if enum.baseType == nil {
				enum.baseType = value
			}
			enum.iv[ident] = value
			enum.vi[value] = ident
		}
	})
	if err != nil {
		return nil, err
	}
	return &enum, nil
}

type parser struct {
	b []byte
}

func ParseEnum(v []byte, fn func(str []byte, num int)) error {
	var p = parser{b: v}
	for {
		p.SkipSpaces()
		str, err := p.MustReadStr()
		if err != nil {
			return err
		}
		p.SkipSpaces()
		if err := p.MustReadChar('='); err != nil {
			return err
		}
		p.SkipSpaces()
		num, err := p.MustReadInt()
		if err != nil {
			return err
		}
		fn(str, num)

		p.SkipSpaces()
		if p.Done() {
			break
		}
		if err := p.MustReadChar(','); err != nil {
			return err
		}

	}
	return nil
}

func (p *parser) SkipSpaces() {
	for len(p.b) > 0 && unicode.IsSpace(rune(p.b[0])) {
		p.b = p.b[1:]
	}
}

func (p *parser) MustReadStr() ([]byte, error) {
	if err := p.MustReadChar('\''); err != nil {
		return nil, err
	}
	var isEscape bool
	var out []byte

	for i, v := range p.b {
		if isEscape {
			if p.b[i-1] == '\'' && v != '\'' {
				p.b = p.b[i:]
				break
			}
			out = append(out, v)
			isEscape = false
		} else if v == '\\' {
			isEscape = true
		} else if v == '\'' {
			isEscape = true
		} else {
			out = append(out, v)
		}
	}
	return out, nil

}
func (p *parser) Done() bool {
	return len(p.b) == 0
}

func (p *parser) MustReadChar(b byte) error {
	if len(p.b) == 0 || p.b[0] != b {
		return fmt.Errorf("invalid char '%s' (%s)", string([]byte{b}), p.b)
	}
	p.b = p.b[1:]
	return nil
}

func (p *parser) MustReadInt() (int, error) {
	if len(p.b) == 0 {
		return 0, fmt.Errorf("expected number")
	}
	var isNegative bool
	if p.b[0] == '-' {
		isNegative = true
		p.b = p.b[1:]
	}
	var num int
	for i := 0; len(p.b) > 0; i++ {
		var c = p.b[0]

		if c >= '0' && c <= '9' {
			if i == 0 && c == 0 {
				return 0, fmt.Errorf("expected non zero digit")
			}
			num = num*10 + int(c-'0')
		} else {
			break
		}
		p.b = p.b[1:]
	}
	if isNegative {
		num = -num
	}
	return num, nil
}
