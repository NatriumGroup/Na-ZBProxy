package minecraft

import (
	"fmt"
	"time"

	"github.com/layou233/ZBProxy/common/mcprotocol"
	"github.com/layou233/ZBProxy/config"
)

func generateKickMessage(s *config.ConfigProxyService, name string) mcprotocol.Message {
	return mcprotocol.Message{
		Color: mcprotocol.White,
		Extra: []mcprotocol.Message{
			{Bold: true, Color: mcprotocol.Green, Text: "Natrium "},
			{Bold: true, Text: "Boost"},
			{Text: " - "},
			{Bold: true, Color: mcprotocol.Red, Text: "连接终止\n"},

			{Text: "你对Natrium Free Project的连接请求已被VeltGop Studio拦截.\n"},
			{Text: "原因: "},
			{Color: mcprotocol.LightPurple, Text: "你没有有效的Natrium Boost许可证(通常是因为已过期).\n"},
			{Text: fmt.Sprintf("请加群%s以获取更多帮助.\n\n", "s.Group"),},

			{
				Color: mcprotocol.Gray,
				Text: fmt.Sprintf("时间戳: %d | 玩家名称: %s | 加速节点: %s\n",
					time.Now().UnixMilli(), name, s.Name),
			},
		},
	}
}

func generatePlayerNumberLimitExceededMessage(s *config.ConfigProxyService, name string) mcprotocol.Message {
	return mcprotocol.Message{
		Color: mcprotocol.White,
		Extra: []mcprotocol.Message{
			{Bold: true, Color: mcprotocol.Green, Text: "Natrium "},
			{Bold: true, Text: "Boost"},
			{Text: " - "},
			{Bold: true, Color: mcprotocol.Red, Text: "连接终止\n"},

			{Text: "你对Natrium Free Project的连接请求已被VeltGop Studio拦截.\n"},
			{Text: "原因: "},
			{Color: mcprotocol.LightPurple, Text: "节点已满员.\n"},
			{Text: fmt.Sprintf("请加群%s以获取更多帮助.\n\n", "s.Group"),},

			{
				Color: mcprotocol.Gray,
				Text: fmt.Sprintf("时间戳: %d | 玩家名称: %s | 加速节点: %s\n",
					time.Now().UnixMilli(), name, s.Name),
			},
		},
	}
}
