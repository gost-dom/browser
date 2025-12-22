package htmlelements

type PackageInterfaces struct {
	webApi     string
	interfaces []string
}

var PackageInterfacesConfiguration = map[string]PackageInterfaces{
	"htmlinterfaces": {
		webApi: "html",
		interfaces: []string{
			// "WindowOrWorkerGlobalScope",
			"History",
		},
	},
}
