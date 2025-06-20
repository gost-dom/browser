# gost-dom code generator

This directory is a separate Go module, used only internally in Gost-DOM to
autogenerate code from web IDL specifications.

This has its own module file as it has no code shared at runtime with the main
code. It also has different dependencies, e.g. this codebase uses tools for
codegeneration, which is irrelevant in the main library; and has no use for V8.

The types of generated code include:

- JavaScript bindings, converting JavaScript calls and arguments to a native Go
  function call. This is by far the largest part of this generator
- Mapping HTML element tag names to DOM interfaces, e.g. `<a>` maps the
  `HTMLAnchorElement`.
- Go interfaces representing IDL interfaces

This part of the code is given somewhat less attention that the rest of the
codebase. However, it is still a goal that the _generated_ code is nice and
clean. 

There are far fewer tests - the primary feedback loop consists of inspeciting
generated code. For refactoring, that generated code is unchanged.
