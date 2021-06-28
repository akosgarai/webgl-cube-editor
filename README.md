# Cube editor with webgl

This application is a POC application for myself. It is possible to write a webgl application without writing a single line of javascript code.

## Generate js code

This tool uses the gopherjs application for compiling the go code to javascript. Use the go get command to get and install gopherjs.

```bash
go get -u github.com/gopherjs/gopherjs
```

Next [set the environment variable](https://github.com/gopherjs/gopherjs#environment-variables) that is required by the gopherjs. (Put it into your zshrc or bashrc file.)

```bash
export GOPHERJS_GOROOT="$(go env GOROOT)"
```

Compile the code with the following command:

```bash
gopherjs build -o docs/webgl-cube-editor.js
```

## Demo page

https://akosgarai.github.io/webgl-cube-editor/
