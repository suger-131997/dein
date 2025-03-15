package generator

type ArgumentGenerator interface {
	GenerateArgument() string
}

type BodyGenerator interface {
	GenerateBody() string
}

const errorHandlingSegment = `
if err != nil{
	return nil, err
}`
