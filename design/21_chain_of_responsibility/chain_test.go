package chain

import "testing"

func TestExampleChain(t *testing.T) {

	// RequestChain->Manager: &ProjectManager{},
	c1 := NewProjectManagerChain()
	// RequestChain->Manager: &DepManager{},
	c2 := NewDepManagerChain()
	// RequestChain->Manager: &GeneralManager{},
	c3 := NewGeneralManagerChain()

	// RequestChain->successor = RequestChain->Manager: &DepManager{},
	c1.SetSuccessor(c2)
	// RequestChain->successor = RequestChain->Manager: &GeneralManager{},
	c2.SetSuccessor(c3)

	var c Manager = c1

	c.HandleFeeRequest("bob", 400)
	//c.HandleFeeRequest("tom", 1400)
	//c.HandleFeeRequest("ada", 10000)
	//c.HandleFeeRequest("floar", 400)
	// Output:
	// Project manager permit bob 400 fee request
	// Dep manager permit tom 1400 fee request
	// General manager permit ada 10000 fee request
	// Project manager don't permit floar 400 fee request

}