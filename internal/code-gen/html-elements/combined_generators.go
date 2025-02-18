package htmlelements

// PackageGeneratorSpecs contains specifactions for files to generate in a
// package. The key is the filename.
type PackageGeneratorSpecs map[string]GeneratorConfig

var PackageConfigs = PackageGeneratorSpecs{
	"dom":           DOMPackageConfig,
	"html":          HTMLPackageConfig,
	"urlinterfaces": URLPackageConfig,
}
