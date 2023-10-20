package staringos_test

import (
	dgctx "github.com/darwinOrg/go-common/context"
	dglogger "github.com/darwinOrg/go-logger"
	"github.com/darwinOrg/go-openai"
	"github.com/darwinOrg/go-openai/staringos"
	"os"
	"testing"
)

func TestCreateChatCompletion(t *testing.T) {
	ctx := &dgctx.DgContext{TraceId: "123"}
	response, err := staringos.CreateChatCompletion(
		ctx,
		&staringos.ChatByCorpusRequest{
			Model: "hoy-3-16k",
			Prompt: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "我是一名猎头，我将给你上传一份访谈对话文本，访谈对话内容主要是猎头向企业雇主了解招聘需求，请你记住这份对话文本",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "好的",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "A：大家好，我这边看上线会议人齐了，咱们要不正式开始吧\nB：好的，辛苦了\nA：今天主要针对贵公司招聘职位做一些了解，以便更好做候选人招聘。请问您公司业务经营模式的形式和状态情况是怎么样的呢？\nB：我们公司是设计+生产+销售型经营模式，年营收累增。\nA：那公司行业排名情况如何呢？我了解到咱们现在是行业头部品牌\nB：对的，处于行业领先位置，连续三年行业第一了\nA：咱们公司的核心产品或者服务有哪些呢？\nB：我们主要提供劳务派遣、人才招聘、社保代缴等服务，拥有先进的saas系统。\nA：接下来您能详细说下岗位职责吗？\nB：主要就是回答有关公司产品或服务的问题，处理订单和交易，解决问题和排除技术问题，提供有关公司产品的信息，提供主动的客户服务等，还有就是要收集和分析客户反馈，大概就是这些。\nA：好的没问题。下一个问题，请问您对候选人的硬技能有什么要求吗？\nB：我们希望他有三年以上相关行业经验，要有pmp项目管理证书，其他要求暂时没有了，主要是这两点。还有就是候选人工作职责还需要能接受远程出差，刚才忘记说了\nA：嗯嗯好的，请问您对候选人的软技能有什么要求吗?比如个性特征、价值观、学习能力等。\nB：候选人最好是能有自驱力，沟通表达跨部门协作能力要强，还需要有抗压能力，有主人翁精神",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "我记住了",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "请帮我提取这份访谈对话文本的主要内容",
				},
			},
		},
		os.Getenv("STARINGOS_APP_TOKEN"),
	)

	if err != nil {
		dglogger.Error(ctx, err)
		return
	}

	dglogger.Infof(ctx, "%+v", response)
}
