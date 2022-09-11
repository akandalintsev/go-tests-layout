GOPRIVATE := git@github.com:akandalintsev/go-tests-layout.git

mod:
	go mod tidy
  
mocks:
	@minimock -i ./internal/services/store.Store -s _mock.go -o ./internal/services/store
