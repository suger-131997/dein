package generator

type BodyGenerator interface {
	GenerateBody() string
}

const errorHandlingSegment = `
if err != nil{
	return nil, err
}`
