// SPDX-License-Identifier: MIT
//
// Copyright (c) 2017-2021 Mark Cornick
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

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/markcornick/linenoise"
)

var (
	version = "dev"
)

func main() {
	var (
		length      = 16
		upper       = true
		lower       = true
		digit       = true
		showVersion = false
	)

	// Process command line flags.
	flag.IntVar(&length, "length", 16, "Length to generate")
	flag.BoolVar(&upper, "upper", true, "Include uppercase letters")
	flag.BoolVar(&lower, "lower", true, "Include lowercase letters")
	flag.BoolVar(&digit, "digit", true, "Include digits")
	flag.BoolVar(&showVersion, "version", false, "Show version and exit")
	flag.Parse()

	// If -version was passed, print version and exit.
	if showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	// Assemble params.
	p := linenoise.Parameters{Length: length, Upper: upper, Lower: lower, Digit: digit}

	// Otherwise, generate a password, then print it.
	result, err := linenoise.Noise(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
