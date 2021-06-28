package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"rest_api/app/models"
	"rest_api/app/utils"
	"strconv"
	"strings"
	"sync"
)

var ch_trained_data chan utils.TrainedData

const (
	service_knn_port    = 8002 // servicio de entrenamiento del KNN
	service_listen_port = 8003 // servicio para recibir el KNN entrenado
)

func TrainingService(parameters models.Parameters) (best_K int, best_accuracy float64) {
	ch_trained_data = make(chan utils.TrainedData)
	var wg sync.WaitGroup
	wg = sync.WaitGroup{}
	lines, err := utils.GetFileByUrl(utils.Url_data)

	if err != nil {
		print(err)
	}

	personas := CleanData(lines)

	training_k := utils.TrainingK{
		Epochs:          parameters.Epochs,
		Current_K:       3,
		Parallel_procs:  parameters.ParallelProcs,
		Accuracy_tuples: []utils.Tuple{},
		Personas:        personas,
	}

	SendDataToTraining(training_k)

	var trained_data utils.TrainedData
	wg.Add(1)
	go func() {
		defer wg.Done()
		go GetTrainedData()
	}()
	trained_data = <-ch_trained_data
	best_K = trained_data.Best_k
	best_accuracy = trained_data.Best_accuracy
	wg.Wait()
	close(ch_trained_data)
	return best_K, best_accuracy

}

func CleanData(lines []string) (data []utils.PersonaEncuestada) {
	data = make([]utils.PersonaEncuestada, 0)
	for i, line := range lines {
		attributes := strings.Split(line, "|")
		if i != 0 {
			var ingreso_monetario float64
			var condicion_laboral string

			if attributes[67] == "" {
				ingreso_monetario = 0
			} else {
				ingreso_monetario, _ = strconv.ParseFloat(attributes[67], 64)
			}

			if attributes[92] == "1" {
				condicion_laboral = "Empleado"
			} else if attributes[92] == "2" {
				condicion_laboral = "Desempleado Abierto"
			} else {
				condicion_laboral = "Desempleado Oculto"
			}

			sexo, _ := strconv.ParseFloat(attributes[14], 64)
			edad, _ := strconv.ParseFloat(attributes[15], 64)
			etnia, _ := strconv.ParseFloat(attributes[90], 64)
			nivel_educativo, _ := strconv.ParseFloat(attributes[16], 64)
			ultimo_cargo, _ := strconv.ParseFloat(attributes[47], 64)
			frecuencia_pago, _ := strconv.ParseFloat(attributes[66], 64)
			seguro_salud, _ := strconv.ParseFloat(attributes[82], 64)

			persona := utils.PersonaEncuestada{
				Data:  []float64{sexo, edad, etnia, nivel_educativo, ultimo_cargo, frecuencia_pago, ingreso_monetario, seguro_salud},
				Class: condicion_laboral,
			}
			data = append(data, persona)
		}
	}

	return data
}

func SendDataToTraining(data utils.TrainingK) {
	ip_remote := "192.168.0.2"
	port := strconv.Itoa(service_knn_port)
	hostremote := ip_remote + ":" + port
	conn, _ := net.Dial("tcp", hostremote)

	defer conn.Close()

	bytes, _ := json.Marshal(data)
	fmt.Fprintf(conn, "%s\n", string(bytes))
}

func GetTrainedData() {
	ip_remote := "192.168.0.2"
	port := strconv.Itoa(service_listen_port)
	hostname := ip_remote + ":" + port
	listen, _ := net.Listen("tcp", hostname)

	defer listen.Close()

	for {
		conn, _ := listen.Accept()
		go HandleGetData(conn)
	}
}

func HandleGetData(conn net.Conn) {
	defer conn.Close()

	bufferIn := bufio.NewReader(conn)
	obj, _ := bufferIn.ReadString('\n')
	var trained_data utils.TrainedData
	json.Unmarshal([]byte(obj), &trained_data)
	ch_trained_data <- trained_data
}
