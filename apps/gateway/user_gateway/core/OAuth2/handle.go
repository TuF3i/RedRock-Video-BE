package OAuth2

import (
	"LiveDanmu/apps/shared/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"golang.org/x/oauth2"
)

// AuthHandle 授权跳转
func (r *OAuthCore) AuthHandle(c *app.RequestContext) {
	// 生成带state授权URL
	authURL := r.oauthConf.AuthCodeURL("random_csrf_state", oauth2.AccessTypeOnline)
	// 用Hertz重定向
	c.Redirect(consts.StatusTemporaryRedirect, []byte(authURL))
}

// CallbackHandler 授权回调接口
func (r *OAuthCore) CallbackHandler(ctx context.Context, c *app.RequestContext) *models.GitHubUser {
	// 初始化数据
	user := new(models.GitHubUser)
	// 获取code
	code := c.Query("code")
	if code == "" {
		// TODO 重定向到错误页面
		fmt.Printf("Error: %v\n", errors.New("empty code"))
	}

	// 用code交换accessToken
	token, err := r.oauthConf.Exchange(ctx, code)
	if err != nil {
		// TODO 重定向到错误页面
		fmt.Printf("Error: %v\n", err)
		return nil
	}

	// 用access_token请求GitHub用户信息接口
	client := r.oauthConf.Client(ctx, token)
	resp, err := client.Get("https://api.github.com/user")
	defer resp.Body.Close()
	if err != nil {
		// TODO 重定向到错误页面
		fmt.Printf("Error: %v\n", err)
		return nil
	}

	// 解析用户信息
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		// TODO 重定向到错误页面
		fmt.Printf("Error: %v\n", err)
		return nil
	}

	return user
}
