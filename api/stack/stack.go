package stack

import (
	"fmt"
	"github.com/gin-gonic/gin"
	stack2 "lag/alg/stack"
	"lag/api/api_code"
)

func JudgeStrIsValid(ctx *gin.Context) {
	needJudgeStr := ctx.Query("needJudgeStr")

	stk := stack2.NewStack()
	for i := 0; i < len(needJudgeStr); i ++ {
		if string(needJudgeStr[i]) == "(" {
			stk.Push(stk)
		}

		if string(needJudgeStr[i]) == ")" {
			i, _ := stk.Pop()
			if i < 0 {
				ctx.JSON(api_code.Success, fmt.Sprintf("第%d个字符，%s", i, string(needJudgeStr[i])))
				return
			}
		}

	}

	ctx.JSON(api_code.Success, "conform to")
	return
}