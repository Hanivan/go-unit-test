package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Tests struct {
	Name, Request, Expected string
}

func TestMain(m *testing.M) {
	fmt.Println("Before unit test")
	m.Run()
	fmt.Println("After unit test")
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Haniv")
	}
}

func TestUsingTableTest(t *testing.T) {
	tests := []Tests{
		{
			Name:     "HelloOne",
			Request:  "One",
			Expected: "Hello One",
		},
		{
			Name:     "HelloTwo",
			Request:  "Two",
			Expected: "Hello Two",
		},
		{
			Name:     "HelloThree",
			Request:  "Three",
			Expected: "Hello Three",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := HelloWorld(test.Request)
			require.Equal(t, test.Expected, result)
		})
	}

}

func TestSubTest(t *testing.T) {
	t.Run("One", func(t *testing.T) {
		result := HelloWorld("Haniv")
		assert.Equal(t, "Hello Haniv", result, "Result must be 'Hello Haniv'")
	})
	t.Run("Two", func(t *testing.T) {
		result := HelloWorld("Anip")
		assert.NotEqual(t, "Hello Rizky", result, "Result must be not equal")
	})
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can not run on Linux operatin system")
	}
	result := HelloWorld("Hanivan")
	assert.Equal(t, "Hello Hanivan", result, "Result must be 'Hello Hanivan'")
}

func TestHelloAssert(t *testing.T) {
	result := HelloWorld("Haniv")
	assert.Equal(t, "Hello Hanivan", result, "Result must be 'Hello Hanivan'")
	fmt.Println("TestHelloAssert done")
}

func TestHelloRequire(t *testing.T) {
	result := HelloWorld("Rizkya")
	require.Equal(t, "Hello Rizky", result, "Result muse be 'Hello Rizky'")
	fmt.Println("TestHelloRequire done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Hanivan")

	if result != "Hello Hanivan" {
		t.Error("Result is not 'Hello Hanivan'")
	}

	fmt.Println("TestHelloWorld done")
}

func TestHelloRizky(t *testing.T) {
	result := HelloWorld("Rizkya")

	if result != "Hello Rizky" {
		t.Error("Result is not 'Hello Rizky'")
	}

	fmt.Println("TestHelloRizky done")
}

func TestHelloS(t *testing.T) {
	result := HelloWorld("S")

	if result != "Hello S" {
		panic("Result is not 'Hello S'")
	}
}
