package main
import("os"
      "fmt"
      "encoding/csv"
//      "io"
      "strings"
      "log"
      "time")

func main() {
  //Read csv
  fName := "problems.csv"
  records := readCsv(fName)
  //fmt.Println(records)
  questions := make([]question, len(records))
  for i,rec:=range records {
    questions[i]=question{rec[0],strings.TrimSpace(rec[1])}
  }
  //fmt.Println(questions)
  //Create a timer
  timer := time.NewTimer(15*time.Second)
  correct := play(questions, timer)
  fmt.Printf("Number of correct:%d\n", correct)
}

type question struct {
  qs string
  ans string
}

func play(questions []question, timer *time.Timer) int {
  correct :=0
  for i, ques := range questions {
    fmt.Printf("%d:%s = ",i+1,ques.qs)
    ansCh := make(chan string)
    go func() {
      var ans string
      fmt.Scanf("%s",&ans)
      ansCh <- ans
  }()
    select {
    case <-timer.C:
      fmt.Println("\nTime up!")
      return correct
    case ans := <- ansCh:
      if ans==ques.ans {
        fmt.Println("Correct!")
        correct++
        } else {
          fmt.Println("Not correct!")
        }
      }
  }
  return correct
}

func readCsv(fName string) [][]string {
  fp, err := os.Open(fName)
  if err!=nil {
    log.Fatal(err)
  }
  //buffio.NewReader takes in io.Reader that implements Read
  //csv wraps the buffio.NewReader with default comma separation
  r := csv.NewReader(fp)
  //r.ReadAll() returns a 2d slice
  records, err := r.ReadAll()
  if err!= nil {
    log.Fatal(err)
  }
  return records
  /*for {
    record, err := r.Read()
    if err == io.EOF {
      break
    }
    if err != nil {
      log.Fatal(err)
    }
    fmt.Printf("%s %s\n", record[0], record[1])
    return record
  }*/
}
