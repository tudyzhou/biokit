package unique

import (
	"bufio"
	"fmt"
	"github.com/tudyzhou/biokit/utils"
	"os"
	"sort"
)

const (
	Version = "1.0"
	Auther  = "tudyzhb@gmail.com"
)

func Main(fileList []string) {
	// Usage
	Usage := fmt.Sprintf(`
	USAGE:
		... <file1>[ file2, ...]
		
		Sorted file
	`)

	if len(fileList) == 0 {
		fmt.Println(Usage)
	} else {
		fileList = utils.SetString(fileList) // unique values
		chs := make([]chan int, len(fileList))
		for i, f := range fileList {
			chs[i] = make(chan int)
			go Unique(f, chs[i])
		}

		// Drain the channel
		for _, ch := range chs {
			<-ch
		}

		// All done
		return
	}
}

func keySort(m map[string]int) (keys []string) {
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return
}

func Unique(fileName string, ch chan int) {
	// ouput path
	uni_out_p := fileName + ".uni"
	dup_out_p := fileName + ".dup"
	sta_out_p := fileName + ".sta"

	// Files Open
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		ch <- 1
		return
	}
	uni_out, err := os.Create(uni_out_p)
	if err != nil {
		fmt.Println(err)
		ch <- 1
		return
	}
	dup_out, err := os.Create(dup_out_p)
	if err != nil {
		fmt.Println(err)
		ch <- 1
		return
	}
	sta_out, err := os.Create(sta_out_p)
	if err != nil {
		fmt.Println(err)
		ch <- 1
		return
	}

	// Files Close
	defer func() {
		// for i := 0; i < 1000000000000; i++ {} // Debug
		_ = f.Close()
		_ = uni_out.Close()
		_ = dup_out.Close()
		_ = sta_out.Close()
		os.Stdout.WriteString(fmt.Sprintf("output: %s\n", uni_out_p))
		os.Stdout.WriteString(fmt.Sprintf("output: %s\n", dup_out_p))
		os.Stdout.WriteString(fmt.Sprintf("output: %s\n", sta_out_p))
		ch <- 1 // Send finish signal
		close(ch)
	}()

	// main
	var (
		//err   error
		r     = bufio.NewReader(f)
		line  string
		found = make(map[string]int)
		l_num = 0
	)
	for ; err == nil; line, err = utils.Readline(r) {
		if _, ok := found[line]; ok {
			found[line] += 1
		} else {
			found[line] = 1
		}
		l_num += 1
		// fmt.Println(len(line), line)
	}

	// output
	keys := keySort(found)
	sta_out.WriteString("#key\tnum\tpercent%\n") // head

	for _, v := range keys {
		match := found[v]
		percent := float32(found[v]) / float32(l_num) * 100

		// statistics
		sta_out.WriteString(fmt.Sprintf("%s\t%d\t%.2f\n", v, match, percent))

		// unique output, dunplication output
		if len(v) == 0 {
			// debug
			// uni_out.WriteString("Null\n")
		} else {
			uni_out.WriteString(v + "\n")
			// dunplication
			if match > 1 {
				dup_out.WriteString(v + "\n")
			}
		}
	}
}
