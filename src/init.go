package src

func NewKtparser() *GoKtParser {
	Ktparser := &GoKtParser{
		SourcePath: "src/source.json",
	}

	_, err := Ktparser.ReadResource()
	if err != nil {
		panic(err)
	}

	Ktparser.ParseResource()
	return Ktparser
}
