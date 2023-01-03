# TinyThreePassCompiler

# The problem

This problem is derived from the [Tiny Three-Pass Compiler kata from codewars](https://www.codewars.com/kata/5265b0885fda8eac5900093b)
and  basically it is about writing a three-pass compiler for a simple programming
language into a small assembly language which syntax is like follows:

```bash
[ a b ] a*a + b*b
```
> a^2 + b^2

```bash
[ first second ] (first + second) / 2
```
> average of two numbers


# How is this implementation different from the original kata

The plan for this implementation is to make it work not with JSON structures but
with protocol buffers, the reason is basically just to practice those technologies
but it is also true that in a real compiler saving and transporting binaries and
not text as it would be the case for JSON could mean a significant improvement.

