package configuration

func configureStreamsSpecs(specs *WebAPIConfig) {
	rstream := specs.Type("ReadableStream")
	rstream.MarkMembersAsNotImplemented(
		"cancel", "getReader", "pipeThrough", "pipeTo", "tee", "locked",
	)
	defaultReader := specs.Type("ReadableStreamDefaultReader")
	defaultReader.MarkMembersAsNotImplemented("read", "releaseLock")
	byobReader := specs.Type("ReadableStreamBYOBReader")
	byobReader.MarkMembersAsNotImplemented("read", "releaseLock")
}
