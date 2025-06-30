package configuration

func configureStreamsSpecs(specs *WebAPIConfig) {
	rstream := specs.Type("ReadableStream")
	rstream.MarkMembersAsNotImplemented(
		"cancel", "pipeThrough", "pipeTo", "tee", "locked",
	)
	defaultReader := specs.Type("ReadableStreamDefaultReader")
	defaultReader.MarkMembersAsNotImplemented("releaseLock")
	byobReader := specs.Type("ReadableStreamBYOBReader")
	byobReader.MarkMembersAsNotImplemented("read", "releaseLock")
}
