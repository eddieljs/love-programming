package learn

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"untitled/tools"

	"github.com/gin-gonic/gin"
)

type CompilerCon struct{}

func (CompilerCon) Compile(ctx *gin.Context) {
	codeInfo := struct {
		Key  string `form:"key" json:"key"`
		Code string `form:"code" json:"code"`
		Lang string `form:"lang" lang:"lang"`
	}{}
	if err := ctx.ShouldBind(&codeInfo); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	if codeInfo.Key != "lava" {
		ctx.String(http.StatusOK, "aaaaaaaa!!!!!!!")
	}
	comMap := map[string]compiler{"py": pythonComp{}, "java": javaComp{}, "c": cComp{}, "cpp": cppComp{}}
	result, err := comp(comMap[codeInfo.Lang], codeInfo.Code)
	// result := PYTHON()
	if err != "" {
		// bt, berr := base64.RawStdEncoding.DecodeString(string(err))
		// if berr != nil {
		// 	fmt.Printf("berr: %v\n", berr)
		// }
		// fmt.Println("aaaa", err)
		// strings.Split(err.Error(), "")
		tools.Fail(ctx, gin.H{
			"error": err,
		}, "运行失败")
		return
	}
	tools.Success(ctx, gin.H{
		"result": result,
	}, "运行成功")
}

func comp(cer compiler, code string) (string, string) {
	return cer.run(code)
}

type compiler interface {
	run(string) (string, string)
}

type pythonComp struct{}

type javaComp struct{}

type cComp struct{}

type cppComp struct{}

func (pythonComp) run(code string) (string, string) {
	os.WriteFile("./code/test.py", []byte(code), 0664)
	cmd := exec.Command("cmd", "/c", "python ./code/test.py") // 替换成你想要执行的cmd命令
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := stdout.String(), stderr.String()
	// fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	// output, err := cmd.Output()
	if err != nil {
		// fmt.Println("命令执行失败:", err)
		return "Error!", errStr
	}
	// result := string(output)
	result := outStr
	// fmt.Printf("result: %v\n", result)
	return result, ""
}

func (javaComp) run(code string) (string, string) {
	os.WriteFile("./code/test.java", []byte(code), 0664)
	cmd := exec.Command("bash", "-c", "javac -encoding utf8 ./code/test.java && cd ./code && java test") // 替换成你想要执行的cmd命令
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := stdout.String(), stderr.String()
	// output, err := cmd.Output()
	if err != nil {
		// fmt.Println("命令执行失败:", err)
		return "Error!", errStr
	}
	result := outStr
	// fmt.Printf("result: %v\n", result)
	return result, ""
}

func (cComp) run(code string) (string, string) {
	os.WriteFile("./code/test.c", []byte(code), 0664)
	cmd := exec.Command("bash", "-c", "cc ./code/test.c -o ./code/test && ./code/test") // 替换成你想要执行的cmd命令
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := stdout.String(), stderr.String()
	if err != nil {
		// fmt.Println("命令执行失败:", err)
		return "Error!", errStr
	}
	result := string(outStr)
	// fmt.Printf("result: %v\n", result)
	return result, ""
}

func (cppComp) run(code string) (string, string) {
	os.WriteFile("./code/test.cpp", []byte(code), 0664)
	cmd := exec.Command("bash", "-c", "g++ ./code/test.cpp -o ./code/test && ./code/test") // 替换成你想要执行的cmd命令
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := stdout.String(), stderr.String()
	if err != nil {
		// fmt.Println("命令执行失败:", err)
		return "Error!", errStr
	}
	result := string(outStr)
	// fmt.Printf("result: %v\n", result)
	return result, ""
}
