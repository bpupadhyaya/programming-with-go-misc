// //////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2022 Bhim Upadhyaya
//
// Licensed under MIT License;
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://opensource.org/licenses/MIT
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
// //////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"strconv"
)

type Item interface{}
type Bag []Item

type Student struct {
	name string
	age  int
	univ string
}

func (s Student) String() string {
	return "(name: " + s.name + " age: " + strconv.Itoa(s.age) + " yrs. old " + "univ: " + s.univ + ")"
}

func main() {
	myBag := make(Bag, 4)
	myBag[0] = "Test"
	myBag[1] = 5
	myBag[2] = Student{"Ken Thompson", 79, "UC Berkeley"}
	myBag[3] = 9.5

	for i, element := range myBag {
		switch value := element.(type) {
		case string:
			fmt.Printf("bag[%d] is a string with value: %s\n", i, value)
		case int:
			fmt.Printf("bag[%d] is a string with value: %d\n", i, value)
		case Student:
			fmt.Printf("bag[%d] is a string with value: %s\n", i, value)
		default:
			fmt.Printf("bag[%d] is of unchecked type\n", i)
		}
	}

}
