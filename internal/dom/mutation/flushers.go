package mutation

type FlusherSet struct {
	flushers map[Flusher]struct{}
}

func (r *FlusherSet) ensureFlushers() {
	if r.flushers == nil {
		r.flushers = make(map[Flusher]struct{})
	}
}

func (r *FlusherSet) AddFlusher(f Flusher) {
	if f == nil {
		panic("FlusherSet.AddFlusher: f is nil")
	}
	r.ensureFlushers()
	r.flushers[f] = struct{}{}
}

func (r *FlusherSet) RemoveFlusher(f Flusher) {
	if _, found := r.flushers[f]; !found {
		panic("FlusherSet.RemoveFlusher: flusher is not added")
	}
	delete(r.flushers, f)
}

func (r *FlusherSet) Flush() {
	for f := range r.flushers {
		f.Flush()
	}
}
