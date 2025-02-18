package htmlelements

var PackageConfigs = map[string]GeneratorConfig{
	"dom":  DOMPackageConfig,
	"html": HTMLPackageConfig,
	"url":  URLPackageConfig,
}
