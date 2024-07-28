//go:build darwin

package main

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/google/uuid"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"log"
	"sync"

	"github.com/tetratelabs/wazero"
)

//go:embed process.wasm
var processWasm []byte

func main() {
	ctx := context.Background()

	r := wazero.NewRuntime(ctx)
	// This closes everything this Runtime created.
	defer func() {
		_ = r.Close(ctx)
	}()

	// instantiate go wasi module
	if err := initialise(ctx, r); err != nil {
		log.Fatalf("error initialising runtime: %v", err)
	}

	compiledWasm, err := r.CompileModule(ctx, processWasm)
	if err != nil {
		log.Panicf("failed to compile Wasm binary: %v", err)
	}

	var wg sync.WaitGroup
	const goroutines = 50
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(i int) {
			defer wg.Done()

			// initialise a new context so that they are not shared by different go routines
			ctx = context.Background()

			instance, err := r.InstantiateModule(ctx, compiledWasm, wazero.NewModuleConfig().WithName(""))
			if err != nil {
				log.Panicf("[%d] failed to instantiate %v", i, err)
			}

			tmp := &temp{a: uuid.New().String()}
			ctx = context.WithValue(ctx, "key", tmp)

			_, err = instance.ExportedFunction("Process").Call(ctx)
			if err != nil {
				log.Panicf("[%d] failed to invoke \"Process\": %v", i, err)
			}
			fmt.Printf("Processed Value: %s\n", tmp.Get())

		}(i)
	}

	wg.Wait()
}

func initialise(ctx context.Context, r wazero.Runtime) error {
	_, err := wasi_snapshot_preview1.Instantiate(ctx, r)
	if err != nil {
		return err
	}

	_, err = r.InstantiateModule(ctx, getFunctionModule(ctx, r), wazero.NewModuleConfig().WithName("env"))
	if err != nil {
		fmt.Println("func instantiation is the error:", err)
		return err
	}

	return nil
}
