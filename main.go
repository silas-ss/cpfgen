package main

import (
  "fmt"
  "math/rand"
  "time"
  "strconv"
  "strings"
)

func main() {
  var cpf [11]int
  var pesoPriDig []int = []int{10,9,8,7,6,5,4,3,2}
  var pesoSecDig []int = []int{11,10,9,8,7,6,5,4,3,2}

  cpf = populateCpf(cpf)
  
  // Primeiro dígito
  cpf[9] = calcDigito(cpf, pesoPriDig)

  // Segundo dígito
  cpf[10] = calcDigito(cpf, pesoSecDig)

  fmt.Println(prettyCpf(cpf))
}

func populateCpf(cpf [11]int) [11]int {
  for i := 0; i < 9; i += 1 {
    cpf[i] = genDigito()
  }

  return cpf
}

func prettyCpf(cpf [11]int) string {
  var sCpf [11]string
  for i := 0; i < len(cpf); i++ {
    sCpf[i] = strconv.Itoa(cpf[i])
  }

  return fmt.Sprintf("%s.%s.%s-%s",
                      strings.Join(sCpf[:3],""),
                      strings.Join(sCpf[3:6],""),
                      strings.Join(sCpf[6:9], ""),
                      strings.Join(sCpf[9:11], ""))
}

func genDigito() int {
  rand.Seed(time.Now().UnixNano())

  return rand.Intn(9-0) + 0
}

func calcDigito(cpfDig [11]int, pesos []int) int {
  var iPeso int = 0
  var iCpf int = 0
  var dig int
  var soma int

  for ; iPeso <= len(pesos)-1; iPeso += 1 {
    soma += (pesos[iPeso] * cpfDig[iCpf])
    
    iCpf += 1
  }
  dig = soma % 11
  if (dig < 2){
    dig = 0
  } else {
    dig = 11 - dig
  }

  return dig
}