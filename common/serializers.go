package common

import (
	"encoding/json"
	"errors"
	"gin_storage/utils"
)

/**
尝试实现类似Django的serializer
*/

// A block type aims to mapper the origin struct to a map[string]interface{}
// For Instance:
//   struct{
// 		Name string
//		Age int
//   }{
//		Name: "fwhezfwhez",
//		Age: 9,
//    }
// The Block of the struct would be a map[string]interface{} {"name":"ft","age":9}
//

type Block map[string]interface{}

func (b *Block) Update(key string, value interface{}) *Block {
	(*b)[key] = value
	return b
}

func (b *Block) Pop(key string) *Block {
	delete(*b, key)
	return b
}

type SerializerI interface {
	Serialize(dest interface{}, f func(b Block) (func(b Block) Block, []string)) ([]byte, error)
}

// a json serializer realization
type JsonSerializer struct {
}

func (s JsonSerializer) Serialize(dest interface{}, f func(b Block) (func(b Block) Block, []string)) ([]byte, error) {
	//1. transfer the struct to a map
	var m = make(Block, 0)
	buf, er := json.Marshal(dest)
	if er != nil {
		return nil, errors.New(utils.Translate(er))
	}
	er = json.Unmarshal(buf, &m)
	if er != nil {
		return nil, errors.New(utils.Translate(er))
	}

	//2. handle picked fields into the map
	blockHandler, picked := f(m)
	//2.1 get filtered fields , get all fields if picked is nil or len=0
	var filtPicked Block
	if picked == nil || len(picked) == 0 {
		filtPicked = m
	} else {
		filtPicked = make(Block, 0)
		for _, v := range picked {
			vm, ok := m[v]
			if !ok {
				continue
			}
			filtPicked[v] = vm
		}
	}

	//2.2 change fields with the updated and popped staff
	buf, er = json.Marshal(blockHandler(filtPicked))
	if er != nil {
		return nil, errors.New(utils.Translate(er))
	}

	return buf, nil
}
