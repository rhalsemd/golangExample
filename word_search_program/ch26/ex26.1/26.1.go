package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다.")
		return
	}

	word := os.Args[1] //var 생략
	files := os.Args[2:]
	fmt.Println("찾으려는 단어: ", word)
	PrintAllFiles(files)
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path) // 많은 파일들을 다뤄야 하는 파이썬 프로그램을 작성할 때, 특정한 패턴이나 확장자를 가진 파일들의 경로나 이름이 필요할 때
}

func PrintAllFiles(files []string) {
	for _, path := range files {
		filelist, err := GetFileList(path) //파일 목록 가져오기
		if err != nil {                    // zero value는 명시적인 초기값을 할당하지 않고 변수를 만들었을 때 해당 변수가 갖게 되는 값이다. nil은 포인터, 인터페이스, 맵, 슬라이스, 채널, 함수의 zero value이다.
			fmt.Println("파일 경로가 잘못되었습니다. err: ", err, "path:", path)
			return
		}
		fmt.Println("찾으려는 파일 리스트")
		for _, name := range filelist {
			fmt.Println(name)
		}
	}
}
