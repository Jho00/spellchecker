package pkg

type Checker interface {
	Check(text string, config Config)
}
