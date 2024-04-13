// A module running Extism

package main

import (
	"github.com/extism/go-sdk"
	"context"
	"fmt"
	"os"
)

type Extism struct{}

// Does the Extism golang quickstart
func (m *Extism) Demo(
	ctx context.Context,
	// +optional
	// +default="https://github.com/extism/plugins/releases/latest/download/count_vowels.wasm" 
	wasmUrl string,
	// +optional
	// +default="count_vowels"
	funcToCall string,
	// +optional
	// +default="Hello, World!"
	argData string,
	) (string, error) {
	manifest := extism.Manifest{
        Wasm: []extism.Wasm{
            extism.WasmUrl{
                Url: wasmUrl,
            },
        },
    }

    config := extism.PluginConfig{}
    plugin, err := extism.NewPlugin(ctx, manifest, config, []extism.HostFunction{})

    if err != nil {
        fmt.Printf("Failed to initialize plugin: %v\n", err)
        os.Exit(1)
    }

	data := []byte(argData)
    exit, out, err := plugin.Call(funcToCall, data)
    if err != nil {
        fmt.Println(err)
        os.Exit(int(exit))
    }

    response := string(out)
	return response, nil
}
