package generator

type Generator interface {
	GenerateBody() string
}

const errorHandlingSegment = `if err != nil{
	return nil, err
}`
