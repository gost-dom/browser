package customrules

// GoType represents a go type referenced by attributes or operations - either
// as argument or return types. Name is the type name. If Package is the empty
// string, the type is assumed to belong to the package where the code is
// generated, or a native Go type, such as int or bool. Pointer is for struct
// types only, whether the type is a pointer type.
type GoType struct {
	Name    string
	Package string
	Pointer bool
}
