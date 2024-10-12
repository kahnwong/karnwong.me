---
title: Calling C from Go, Python and Rust benchmark
date: 2024-10-12T14:56:10+07:00
draft: false
ShowToc: true
images:
tags:
  - benchmark
  - programming
---

There's a consensus that generally, c is very fast, and python is very slow. But if we are talking about go and rust, you would find that rust is slightly faster than go. So from fastest to slowest: c, rust, go, python.

But what if you have go, python and rust calling c function? There would be more overhead, but how much?

## Calling C from Go, Python and Rust

### Stats

![01-stats-all.webp](images/01-stats-all.webp)

Turns out python has the most overhead.

---

![02-stats-exclude-python.webp](images/02-stats-exclude-python.webp)

But with go and rust, overheads are more or less negligible, with rust having the least overhead.

### Execution Time

![03-times-all.webp](images/03-times-all.webp)

You can see that python execution time plateaued, but go and rust are within the same range as c.

---

![04-times-c-go-rust.webp](images/04-times-c-go-rust.webp)

You would also notice that go and c have random execution time peaks, but rust is the most stable in terms of execution time. Additionally, calling the same c code, using rust results in the most constant memory usage footprint.

## The Same Function But Implemented Natively

![07-native-stats-all.webp](images/05-native-stats-all.webp)

It's the same pattern if native functions are used instead of c foreign function interface.

## Calling C vs Native Function

![06-all-stats-c-rust.webp](images/06-all-stats-c-rust.webp)

But if we really have to compare CFFI and native function implementation, `pure rust` is faster than `rust calling c`

### Side-by-side Comparison for Go and Rust

![07-go-comparison.webp](images/07-go-comparison.webp)

Comparing `pure go` (in blue) and `go calling c`, pure go execution time is faster.

---

![08-rust-comparison.webp](images/08-rust-comparison.webp)

With `pure rust` (in rust) and `rust calling c` it's still native implementation faster than CFFI, although with rust, the overhead is more or less negligible.

---

![09-all-times-c-rust.webp](images/09-all-times-c-rust.webp)

But if we really have to compare `c` and `pure rust` (in rust), c is still faster, but the difference here  is still negligible.

## Bonus

![10-all-stats-all.webp](images/10-all-stats-all.webp)

You can also call rust from python via pyo3. In which, for the same function implementation, it is waaaaay faster than CFFI!
