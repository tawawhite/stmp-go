// STMP - a lightweight real-time bidirectional framework
//
// The MIT License (MIT)
//
// Copyright (c) 2016 acrazing
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// protoc-gen-stmp is a plugin for the Google protocol buffer compiler to generate
// stmp relative code, supported languages: golang and javascript.
//
// To use protoc-gen-stmp, you need to build the project and put the result to your
// path, the most simple way is run the command with go:
//  go get -u github.com/acrazing/stmp-go/protoc-gen-stmp
//
// For golang, protoc-gen-stmp dependents on any one go code generate, such as
// protoc-gen-go, or protoc-gen-gogo, you can choose according to your preference.
// the emitted file name ends with `.stmp.go`, and the path is same to `protoc-gen-go`,
// if with the output configuration. There is no custom options for emit go code.
// for example, with protoc-gen-go:
//  protoc --go_out=$GOPATH/src --stmp_out=lang=go:$GOPATH/src ./pb/*.proto
//
// For javascript, protoc-gen-stmp supports emit .js and .d.ts for TypeScript.
// It dependents on protobufjs (https://www.npmjs.com/package/protobufjs).
// Like protobufjs, protoc-gen-stmp merges all input .proto files to one file,
// the file name is specified by custom option js.out. We recommend you to let the
// file name ends with .stmp.js. At the same time, protoc-gen-stmp will emit types
// declaration file automatically, the file name is replace the end of the js.out
// with .d.ts. For example:
//  pbjs -o ./pb/chat.pb.js ./pb/*.proto
//  pbts -o ./pb/chat.pb.d.ts ./pb/chat.pb.js
//  protoc --stmp_out=lang=js,js.pb=./pb/chat.pb.js,js.out=./pb/chat.stmp.js:. ./pb/*.proto
//
// As the example upon, protoc-gen-stmp supports some custom options for javascript,
// includes:
//  js.pb: string,               required, the pbjs emitted file path relative to current directory
//  js.out: string,              required, the output file path relative to the current directory
//  js.dts: "" | "0",            optional, if set as "0", will not emit .d.ts file
//  js.module: "esm" | "cjs",    optional, the module kind, default is "esm", you can set as "cjs" to run the generated code in Node.js directly
//  js.root: string,             optional, the root namespace name, default is "stmp", it is useful for avoid .d.ts ns conflict
package main