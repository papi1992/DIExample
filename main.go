package diExample

import (
	"diExample/dep"
)

func main() {
	//Solution No. 1 - Global Variable

	dep.DoM = dep.DependencyOfMethodImp{}
	dep.Method_S1()

	//...

	//Solution No. 2 - Method Paramter

	domImp := dep.DependencyOfMethodImp{}
	dep.Method_S2(domImp)

	//...

	//Solution No. 3 - Struct Field

	domImp2 := dep.DependencyOfMethodImp{}
	se := dep.StructExampleD{DoM: domImp2}
	se.Method_S3()

	//...
}
