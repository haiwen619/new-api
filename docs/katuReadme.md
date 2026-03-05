
# Katu

### 1 外部 API 鉴权

- 头部：`Authorization: Bearer <global.api_key>`
- 使用接口：`/v1/models`、`/v1/chat/completions`

### 2 管理后台鉴权

- 登录接口：`POST /api/admin/login`
- 登录成功后返回 session token
- 后续后台接口通过 `Authorization: Bearer <admin_session_token>` 调用
- 账号池与代理池接口也复用同一后台 token 鉴权
- 
### 2 模型名称解析问题
Flow2API 渠道  使用 gemini-3.1-flash-image-Flow 模型时，模型名称会被解析为 gemini-3.1-flash-image，其他类似的模型名称也会被解析掉 -Flow 后缀，例如：


gemini-2.5-flash-image-Flow 解析为 gemini-2.5-flash-image 
gemini-3.0-pro-image-Flow 解析为 gemini-3.0-pro-image
gemini-3.1-flash-image-Flow 解析为 gemini-3.1-flash-image

### 图片生成示例（gemini-3.1-flash-image）
测试请求：
  curl -X POST "http:v1/chat/completions" \
  -H "Authorization: Bearer sk-" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gemini-3.1-flash-image-Flow",
    "contents": [
      { "parts": [ { "text": "A realistic food photo, studio light, clean table." } ] }
    ],
    "generationConfig": {
      "responseModalities": ["IMAGE"],
      "imageConfig": { "aspectRatio": "16:9", "imageSize": "2K" }
    },
    "stream": true
  }'


  curl -X POST "http:v1/chat/completions" \
  -H "Authorization: Bearer sk-" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gemini-3.1-flash-image-Flow",
    "contents": [
      { "parts": [ { "text": "A realistic food photo, studio light, clean table." } ] }
    ],
    "generationConfig": {
      "responseModalities": ["IMAGE"],
      "imageConfig": { "aspectRatio": "16:9", "imageSize": "2K" }
    },
    "stream": true
  }'

