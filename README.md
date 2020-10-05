# Wire: A minimalistic library for dependency injection in Go
Wire defines a minimal collection of types to support 
[dependency injection](https://en.wikipedia.org/wiki/Dependency_injection) without magic. 

Organising code for 
dependency injection is extremely beneficial for many reasons, however automatic resolution of dependencies at 
compile time or runtime is the source of much confusion and frustration, not to mention bugs. 

Manually controlling which dependencies are injected is cumbersome but with the right patterns in place this avoids the 
confusion. 

## Installing
```bash
go get github.com/clausthrane/wire
```

## Documentation
- [User Guide][]
- [Example application][]

[User Guide]: ./docs/guide.md
[Example application]: ./example
