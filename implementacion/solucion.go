package recocido

import (
	_"fmt"
	"strconv"
	"math/rand"
	"github.com/andreagonz/recocido/heuristica"
)

var distancias *[][]float64
var ciudades *[]Ciudad
var problema *[]int
var Max float64
var Avg float64
var c int

func SetDistancias(p *[][]float64) {
	distancias = p
}

func SetCiudades(p *[]Ciudad) {
	ciudades = p
}

func SetC(i int) {
	c = i
}

func SetProblema(p *[]int) {
	problema = p
}

type Ciudad struct {
	Id int
	Nombre string
	Pais string
	Poblacion int
	Latitud float64
	Longitud float64	
}

type Ruta struct {
	Ciudades []int
	funObj float64
	fun float64
}

type Conexiones struct {
	Distancias []float64
}

func (r Ruta) Str() string {
	s := ""
	s += "{"
	for i := 0; i < len(r.Ciudades); i++ {
		s += "(" + strconv.Itoa(r.Ciudades[i]) + ": " + (*ciudades)[r.Ciudades[i]].Nombre + ") "
		if i < len(r.Ciudades) - 1 {
			s += strconv.FormatFloat((*distancias)[r.Ciudades[i]][r.Ciudades[i + 1]], 'f', -1, 64) + " "
		}
	}
	s += "}"
	return s
}

/*
func (r Ruta) Str() string {
	s := ""
	s += "{"
	for i := 0; i < len(r.Ciudades); i++ {
		s += strconv.Itoa(i) + ": " + strconv.Itoa(r.Ciudades[i]) + " "
	}
	s += "}"
	return s
}
*/

func (r Ruta) ObtenFun() float64 {
	return r.fun
}

func (r Ruta) ObtenFunObj() float64 {
	return r.funObj
}

func (ruta Ruta) ObtenVecino(rand *rand.Rand) recocido.Solucion {
	//fmt.Println(ruta.Str())
	var nruta Ruta
	nruta.Ciudades = make([]int, len(ruta.Ciudades))
	for i := 0; i < len(ruta.Ciudades); i++ {
		nruta.Ciudades[i] = ruta.Ciudades[i]
	}
	i := rand.Intn(len(ruta.Ciudades))
	j := rand.Intn(len(ruta.Ciudades))
	for j == i {
		j = rand.Intn(len(ruta.Ciudades))
	}
	a := nruta.Ciudades[i]
	nruta.Ciudades[i] = nruta.Ciudades[j]
	nruta.Ciudades[j] = a
	//fmt.Println(ruta.Str())
	return &nruta
}

func(r Ruta) EsFactible() bool {
	bool := true
	for j := 0; j < len(r.Ciudades) - 1; j++ {
		if (*distancias)[r.Ciudades[j]][r.Ciudades[j + 1]] == 0.0 {
			bool = false
		}
	}
	return bool
}

func MaxAvg() {
	n := 0.0
	p := 0.0
	for i := 0; i < len(*problema); i++ {
	    for j := 0; j < len(*problema); j++ {
		if (*distancias)[(*problema)[i]][(*problema)[j]] > 0.0 {
			if (*distancias)[(*problema)[i]][(*problema)[j]] > Max {
				Max = (*distancias)[(*problema)[i]][(*problema)[j]]
			}
			p += (*distancias)[(*problema)[i]][(*problema)[j]]
			n++
		}
		}
	}
	Avg = p / n
}

func (r *Ruta) CalculaFun() {
	f := 0.0
	for i := 1; i < len(r.Ciudades); i++ {
		if (*distancias)[(r.Ciudades)[i - 1]][(r.Ciudades)[i]] > 0.0 {
			f += (*distancias)[r.Ciudades[i - 1]][r.Ciudades[i]]
		} else {
			f += Max * float64(c)
		}
	}
	r.funObj = f
	r.fun = f / (Avg * float64((len(r.Ciudades)) - 1))
}
