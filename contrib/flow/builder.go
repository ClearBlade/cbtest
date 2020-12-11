package flow

// BuilderFunc is a function that receives a *Builder instance.
type BuilderFunc func(*Builder)

// Builder lets you create flows programatically. It works by constructing a
// tree of flows that can be executed later using Run.
//
// Builder life cycle
//
// Everything is evaluated eagerly, with the exception of the Run function, which
// is executed when the flow is run.
type Builder struct {

	// workers is the queue of workers.
	workers []Worker

	// middlware updates a Worker before adding it to the queue of workers.
	middleware []func(Worker) Worker
}

// NewBuilder returns a new *Builder instance.
func NewBuilder() *Builder {
	return &Builder{
		workers: make([]Worker, 0),
	}
}

// applyMiddleware applies all pending middleware to the given worker.
func (b *Builder) applyMiddleware(worker Worker) Worker {
	for _, m := range b.middleware {
		worker = m(worker)
	}
	b.middleware = []func(Worker) Worker{}
	return worker
}

// appendWorker adds the given worker to the queue of workers and returns it.
func (b *Builder) appendWorker(worker Worker) Worker {
	worker = b.applyMiddleware(worker)
	b.workers = append(b.workers, worker)
	return worker
}

// WithName sets the name to use for the next builder call.
func (b *Builder) WithName(name string) *Builder {
	b.middleware = append(b.middleware, func(worker Worker) Worker {
		return withName(name, worker)
	})
	return b
}

// WithContext sets the context to use for the next builder call.
func (b *Builder) WithContext(borrower Borrower) *Builder {
	b.middleware = append(b.middleware, func(worker Worker) Worker {
		return withContext(borrower, worker)
	})
	return b
}

// Sequence builds a flow of sequential workers.
func (b *Builder) Sequence(build BuilderFunc) Worker {
	builder := NewBuilder()
	build(builder)
	return b.appendWorker(sequence(builder.workers...))
}

// Parallel builds a flow of parallel workers.
func (b *Builder) Parallel(build BuilderFunc) Worker {
	builder := NewBuilder()
	build(builder)
	return b.appendWorker(parallel(builder.workers...))
}

// Run adds the given worker to the current flow.
func (b *Builder) Run(worker Worker) Worker {
	return b.appendWorker(worker)
}
