package main

import (
	"context"
	"fmt"
	"github.com/tetratelabs/wazero"

	"github.com/tetratelabs/wazero/api"
)

const (
	contextKey      = "key"
	setFunctionName = "Set"
	getFunctionName = "Get"
)

var setFunction = api.GoModuleFunc(func(ctx context.Context, module api.Module, stack []uint64) {
	keyPos := uint32(stack[0])
	keyLen := uint32(stack[1])

	key, ok := module.Memory().Read(keyPos, keyLen)
	if !ok {
		//stack[0] = 0
		return
	}

	obj := ctx.Value(contextKey)
	if obj != nil {
		obj := obj.(*temp)
		obj.Set(string(key))
	}

})

var getFunction = api.GoModuleFunc(func(ctx context.Context, module api.Module, stack []uint64) {
	alloc := module.ExportedFunction("Allocate")
	if alloc == nil {
		fmt.Println("Could not allocate WASM allocate function")
		return
	}

	obj := ctx.Value(contextKey)
	if obj != nil {
		obj := obj.(*temp)
		ret := []byte(obj.Get())
		retLen := uint64(len(ret))

		out, err := alloc.Call(ctx, retLen)
		if err != nil {
			return
		}
		retPtr := out[0]
		if !module.Memory().Write(uint32(retPtr), ret) {
			fmt.Printf("Memory.Write(%d, %d) out of range of memory size %d\n", retPtr, retLen, module.Memory().Size())
			return
		}

		stack[0] = (retPtr << uint64(32)) | retLen

	}
})

type temp struct {
	a string
}

func (t *temp) Set(val string) {
	t.a = val
}

func (t *temp) Get() string {
	return t.a
}

func getFunctionModule(ctx context.Context, runtime wazero.Runtime) wazero.CompiledModule {
	ret, _ := runtime.NewHostModuleBuilder("env").
		NewFunctionBuilder().
		WithGoModuleFunction(setFunction,
			[]api.ValueType{
				api.ValueTypeI32, // pointer position
				api.ValueTypeI32, // pointer length
			},
			[]api.ValueType{}).
		Export(setFunctionName).
		NewFunctionBuilder().
		WithGoModuleFunction(getFunction,
			[]api.ValueType{},
			[]api.ValueType{
				api.ValueTypeI64, // pointer position + length
			}).
		Export(getFunctionName).
		Compile(ctx)

	return ret
}
