package main

type Payaya struct {
	power       float32
	destruction int
	dead        bool
}

func (r *Payaya) level() float32 {
	return float32(r.destruction) / r.power
}

func main() {
	var b [190]Payaya
	for index := range b {
		b[index] = Payaya{power: float32(index * index * 123 / (index + 1)), destruction: index * index * 123 / (index + 1) * 123, dead: false}
	}
	for _, element := range b {
		println(element.power, element.destruction, element.dead, element.level)
	}
	println("HELLO")
}
