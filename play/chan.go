package play

import (
	"fmt"
	"sync"
)

type Stringable interface {
	Repr() string
}

type StringableConstructor func(int) Stringable
type StringableReader func(Stringable)

type Person struct {
	Name string
	Age  int
}

func (m Person) Repr() string {
	return fmt.Sprintf("My name is %s and I am %d years old", m.Name, m.Age)
}

func NewPerson(age int) Stringable {
	return &Person{
		Name: fmt.Sprintf("Name-%d", age),
		Age:  30 + age,
	}
}

type Car struct {
	Make  string
	Model string
	Year  int
}

func (m Car) Repr() string {
	return fmt.Sprintf("My %d %s %s goes brrrrr", m.Year, m.Make, m.Model)
}

func NewCar(year int) Stringable {
	return &Car{
		Make:  fmt.Sprintf("Make-%d", year),
		Model: fmt.Sprintf("Model-%d", year),
		Year:  2000 + year,
	}
}

func Printer(s Stringable) {
	fmt.Println(s.Repr())
}

func SelectorChan(
	create1, create2 StringableConstructor,
	printer1, printer2 StringableReader,
) {

	ch1 := make(chan Stringable, 1)
	ch2 := make(chan Stringable, 1)
	var wg sync.WaitGroup

	// Writer 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch1)

		for i := range 10 {
			ch1 <- create1(i)
		}
	}()

	// Writer 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch2)

		for i := 10; i < 20; i++ {
			ch2 <- create2(i)
		}
	}()

	// Reader
	wg.Add(1)
	go func() {
		defer wg.Done()
		for ch1 != nil || ch2 != nil {
			select {
			case val, ok := <-ch1:
				if ok {
					printer1(val)
				} else {
					ch1 = nil
				}

			case val, ok := <-ch2:
				if ok {
					printer2(val)
				} else {
					ch2 = nil
				}
			}
		}
	}()

	wg.Wait()
	fmt.Println("Done")
}
