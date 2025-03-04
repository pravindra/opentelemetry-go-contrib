// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//nolint:all // Linter is executed at the same time as tests which leads to race conditions and failures.
package main

import (
	"fmt"
	__atel_context "context"
	__atel_otel "go.opentelemetry.io/otel"
)

func foo(__atel_tracing_ctx __atel_context.Context,) {
	__atel_child_tracing_ctx, __atel_span := __atel_otel.Tracer("foo").Start(__atel_tracing_ctx, "foo")
	_ = __atel_child_tracing_ctx
	defer __atel_span.End()
	fmt.Println("foo")
}

func FibonacciHelper(__atel_tracing_ctx __atel_context.Context, n uint) (uint64, error) {
	__atel_child_tracing_ctx, __atel_span := __atel_otel.Tracer("FibonacciHelper").Start(__atel_tracing_ctx, "FibonacciHelper")
	_ = __atel_child_tracing_ctx
	defer __atel_span.End()
	func() {
		__atel_child_tracing_ctx, __atel_span := __atel_otel.Tracer("anonymous").Start(__atel_child_tracing_ctx, "anonymous")
		_ = __atel_child_tracing_ctx
		defer __atel_span.End()
		foo(__atel_child_tracing_ctx)
	}()
	return Fibonacci(__atel_child_tracing_ctx, n)
}

func Fibonacci(__atel_tracing_ctx __atel_context.Context, n uint) (uint64, error) {
	__atel_child_tracing_ctx, __atel_span := __atel_otel.Tracer("Fibonacci").Start(__atel_tracing_ctx, "Fibonacci")
	_ = __atel_child_tracing_ctx
	defer __atel_span.End()
	if n <= 1 {
		return uint64(n), nil
	}

	if n > 93 {
		return 0, fmt.Errorf("unsupported fibonacci number %d: too large", n)
	}

	var n2, n1 uint64 = 0, 1
	for i := uint(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1, nil
}
