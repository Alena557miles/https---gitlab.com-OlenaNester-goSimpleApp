package collect

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
)

type Collect struct {
	mu            sync.RWMutex
	_arrayOfLinks []string
}

func New() *Collect {
	res := &Collect{}
	return res
}

func (c *Collect) FindLinks(fileContent string) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	re, err := regexp.Compile(`(http|ftp|https):\/\/([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:\/~+#-]*[\w@?^=%&\/~+#-])`)
	if err != nil {
		fmt.Println(err)
	}
	c._arrayOfLinks = re.FindAllString(string(fileContent), -1)
	data := strings.Join(c._arrayOfLinks, " ")
	return data
}

func CreateFileandWriteData(filename, data string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	_, err = file.WriteString(data)
}

func OpenAndPrintFileData(s string) {
	contentLinks, err := os.ReadFile(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(contentLinks))
}
