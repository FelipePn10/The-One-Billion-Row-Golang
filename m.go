package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Measurement struct {
	// Não adicionamos o nome, pois ele já está salvo na chave do nosso mapa. Então não precisamos gastar memória à toa.
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	measurements, err := os.Open("measurements.txt")
	//Tratamento de erro caso o arquivo não abra
	if err != nil {
		fmt.Println("Arquivo não conseguiu ser aberto")
		return
	}
	defer measurements.Close()
	// Um nome para uma estrutura. Aqui no caso estamos criando o mapa, então usamos a função make.
	data := make(map[string]Measurement)

	scanner := bufio.NewScanner(measurements)
	//vamos inteirar enquanto tiver linhas para serem lidas.
	for scanner.Scan() {
		//A linha que o scanner leu:
		rawData := scanner.Text()
		//Aqui a função recebe duas strings, a primeira é a string que será procurada e a segunda é a string A SER procurada. E retorna o index onde aquela segunda string foi encontrada na string inicial. Se a string não for encontrada, ela retorna o valor −1.
		semicolon := strings.Index(rawData, ";")
		//Tudo oque está naquela string até o ponto e vírgula
		location := rawData[:semicolon]
		//Agora para pegarmos a temperatura. Esse +1 indica que queremos a partir do primeiro byte depois do ponto e vírgula!
		rawTemp := rawData[semicolon+1:]
		//criando a temperatura de fato, convertendo a string para um float
		temp, _ := strconv.ParseFloat(rawTemp, 64)
		//Vamos ler o nosso mapa, passando o nome da localização atual.
		measurement, ok := data[location]
		if !ok {
			measurement = Measurement{
				Min:   temp,
				Max:   temp,
				Sum:   temp,
				Count: 1,
			}
		} else {
			//Estamos pegando o menor valor passado, comparando com o valor que já está salvo na nossa estrutura e salvando ao final
			measurement.Min = min(measurement.Min, temp)
			measurement.Max = max(measurement.Max, temp)
			measurement.Sum += temp
			measurement.Count++
		}
		//salvando no nosso mapa:
		data[location] = measurement
	}
	//sort das nossas localizações, e para um sort precisamos de um slice
	locations := make([]string, 0, len(data))
	for name := range data {
		locations = append(locations, name)
	}
	sort.Strings(locations)
	fmt.Printf("{")
	for _, name := range locations {
		measurement := data[name]
		//o f imprime o float, e colocando.1f dizemos que queremos imprimir casas decimais
		fmt.Printf(
			"%s=%.1f/%.1f/%.1f, ",
			name,
			measurement.Min,
			measurement.Sum/float64(measurement.Count),
			measurement.Max,
		)
	}
	fmt.Printf("}\n")
}
