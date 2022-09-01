package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/second-state/WasmEdge-go/wasmedge"
	bindgen "github.com/second-state/wasmedge-bindgen/host/go"
)

const (
	WASM_FILE_PATH    = "js_func.wasm"
	DEFAULT_FUNC_NAME = "run"
	DEFAULT_PORT      = ":8080"
)

var tmp string

func main() {
	r := gin.Default()
	r.GET("/getInfo", getInfo)
	r.POST("/washer", collector)
	r.Run(DEFAULT_PORT) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func collector(c *gin.Context) {
	input, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("cannot read body from request, error: ", err)
		return
	}
	tmp = string(input)
}

func getInfo(c *gin.Context) {
	if tmp == "" {
		c.String(http.StatusOK, "no info")
		return
	}
	output, err := executeWasm(tmp)
	if err != nil {
		log.Println("error when executeWasm, error: ", err)
		return
	}
	c.String(http.StatusOK, output)
}

func executeWasm(input string) (string, error) {
	wasmedge.SetLogErrorLevel()

	/// Create configure
	var conf = wasmedge.NewConfigure(wasmedge.WASI)
	defer conf.Release()

	/// Create VM with configure
	var vm = wasmedge.NewVMWithConfig(conf)
	defer vm.Release()

	/// Init WASI
	var wasi = vm.GetImportModule(wasmedge.WASI)
	wasi.InitWasi(
		[]string{},      /// The args
		os.Environ(),    /// The envs
		[]string{".:."}, /// The mapping preopens
	)

	/// Load and validate the wasm
	vm.LoadWasmFile(WASM_FILE_PATH)
	vm.Validate()

	// Instantiate the bindgen and vm
	bg := bindgen.New(vm)
	bg.Instantiate()
	defer bg.Release()

	res, _, err := bg.Execute(DEFAULT_FUNC_NAME, input)
	return res[0].(string), err
}
