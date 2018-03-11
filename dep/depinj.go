package dep

import (
	"diExample/alib"
	"diExample/blib"
	"fmt"
)

// Method has 2 dependencies.
// So It cannot be tested.
func Method() {
	pageContent, err := alib.GetPageContent()
	if err != nil {
		fmt.Println(err)
	}
	err = blib.RenderWebPage(pageContent)
	if err != nil {
		fmt.Println(err)
	}

	//...
}

type StructExample struct {
	//...
}

func (se StructExample) Method_for_S3() {
	pageContent, err := alib.GetPageContent()
	if err != nil {
		fmt.Println(err)
	}
	err = blib.RenderWebPage(pageContent)
	if err != nil {
		fmt.Println(err)
	}

	//...
}

//-----------------------------
//Every solution must have an interface that can be injected.
type DependencyOfMethod interface {
	GetPageContent() (string, error)
	RenderWebPage(string) error
}

type DependencyOfMethodImp struct{}

func (dom DependencyOfMethodImp) GetPageContent() (string, error) {
	return alib.GetPageContent()
}

func (dom DependencyOfMethodImp) RenderWebPage(pageContent string) error {
	return blib.RenderWebPage(pageContent)
}

//-----------------------------
//Solution No. 1 - Global Variable

var DoM DependencyOfMethod

func Method_S1() {
	pageConent, err := DoM.GetPageContent()
	if err != nil {
		fmt.Println(err)
	}
	err = DoM.RenderWebPage(pageConent)
	if err != nil {
		fmt.Println(err)
	}

	//...
}

//-----------------------------
//Solution No. 2 - Method Paramter

func Method_S2(dom DependencyOfMethod) {
	pageConent, err := dom.GetPageContent()
	if err != nil {
		fmt.Println(err)
	}
	err = dom.RenderWebPage(pageConent)
	if err != nil {
		fmt.Println(err)
	}
}

//-----------------------------
//Solution No. 3 - Struct Field

type StructExampleD struct {
	DoM DependencyOfMethod
}

func (se StructExampleD) Method_S3() {
	pageConent, err := se.DoM.GetPageContent()
	if err != nil {
		fmt.Println(err)
	}
	err = se.DoM.RenderWebPage(pageConent)
	if err != nil {
		fmt.Println(err)
	}
}
