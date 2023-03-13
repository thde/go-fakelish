const go = new Go() // Defined in wasm_exec.js
const WASM_URL = "web.wasm"

if ("instantiateStreaming" in WebAssembly) {
  WebAssembly.instantiateStreaming(fetch(WASM_URL), go.importObject).then(
    (obj) => {
      run(obj.instance)
    }
  )
} else {
  fetch(WASM_URL)
    .then((resp) => resp.arrayBuffer())
    .then((bytes) =>
      WebAssembly.instantiate(bytes, go.importObject).then((obj) => {
        run(obj.instance)
      })
    )
}

const run = (wasm) => {
  fetch("dictionaries/en.txt")
    .then((r) => r.text())
    .then((t) => {
      document.getElementById("words").textContent = t
      go.run(wasm)
    })
}
