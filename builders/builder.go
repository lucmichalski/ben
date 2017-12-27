package builders

// RuntimeBuilder is the interface that defines how to build runtime environments
type RuntimeBuilder interface {
	Init() error
	PrepareImage() error
	SetupContainer() error
	Cleanup() error
	Benchmark() error
	Display() error
}
