package htmlelements

// PackageGeneratorSpecs contains specifactions for files to generate in a
// package. The key is the filename.
type PackageGeneratorSpecs map[string]GeneratorConfig

var PackageConfigs = PackageGeneratorSpecs{
	"html":           HTMLPackageConfig,
	"urlinterfaces":  URLPackageConfig,
	"dominterfaces":  DOMInterfacesPackageConfig,
	"htmlinterfaces": HTMLInterfacesPackageConfig,
}
