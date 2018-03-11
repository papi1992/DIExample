package tests

import (
	"diExample/dep"
	"testing"
)

//Mock

type DependencyOfMethodMock struct {
	renderWebPageContent string
	tester               *testing.T
}

func (domMock DependencyOfMethodMock) GetPageContent() (string, error) {
	return "<head></head><body></body>", nil
}

func (domMock *DependencyOfMethodMock) RenderWebPage(pageConent string) error {
	domMock.renderWebPageContent = pageConent
	return nil
}

func (domMock DependencyOfMethodMock) AssertWebPageContentIsCorrect() {
	if domMock.renderWebPageContent != "<head></head><body></body>" {
		domMock.tester.Errorf("test failed")
	}
}

//Testing part of Solution No. 1
func TestMethod_S1(t *testing.T) {
	domMock := DependencyOfMethodMock{tester: t}
	dep.DoM = &domMock
	dep.Method_S1()
	domMock.AssertWebPageContentIsCorrect()
}

//Testing part of Solution No. 2
func TestMethod_S2(t *testing.T) {
	domMock := DependencyOfMethodMock{tester: t}
	dep.Method_S2(&domMock)
	domMock.AssertWebPageContentIsCorrect()
}

//Testing part of Solution No. 3

func TestMethod_S3(t *testing.T) {
	domMock := DependencyOfMethodMock{tester: t}
	se := dep.StructExampleD{DoM: &domMock}
	se.Method_S3()
	domMock.AssertWebPageContentIsCorrect()
}
