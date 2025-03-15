package generator

type BodyGenerator interface {
	Generate() string
}

const errorHandlingSegment = `
if err != nil{
	return nil, err
}`
