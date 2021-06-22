## 支持SonarQube 8.9 API

> 目前只支持drone-sonar需要的API，如有其他接口自行添加

- [x] api/user_tokens
   - [x] api/user_tokens/generate 生成token
   - [x] api/user_tokens/revoke 销毁token
   - [x] api/user_tokens/search 搜索token
- [x] api/webhooks
    - [x] api/webhooks/create
    - [x] api/webhooks/delete
    - [x] api/webhooks/list
    - [x] api/webhooks/update
- [x] api/system
    - [ ] api/system/health
    - [x] api/system/status
- [x] api/measures
    - [x] api/measures/search
- [x] api/project_branches
    - [x] api/project_branches/list
- [x] api/projects
    - [x] api/projects/create
    - [x] api/projects/delete
    - [x] api/projects/search

  
## 致谢

- [magicsong/sonargo](https://github.com/magicsong/sonargo)
- [sonarapi docs](http://127.0.0.1:9000/web_api)


